package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	_ "net/http/pprof"

	"github.com/draganm/immersadb"
	"github.com/draganm/immersadb/dbpath"
	"github.com/draganm/immersadb/modifier"
	"github.com/draganm/kickback"
	"github.com/draganm/snitch/executor"
	oauth2 "github.com/goincremental/negroni-oauth2"
	sessions "github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"
	"github.com/urfave/negroni"
	"gopkg.in/urfave/cli.v2"

	_ "github.com/draganm/snitch/ui"
)

func main() {

	go func() {
		http.ListenAndServe("0.0.0.0:6060", nil)
	}()

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "port",
				Aliases: []string{"p"},
				EnvVars: []string{"PORT"},
				Value:   "8000",
			},
			&cli.BoolFlag{
				Name:    "oauth2",
				EnvVars: []string{"OAUTH2"},
			},
			&cli.StringFlag{
				Name:    "oauth2-client-id",
				EnvVars: []string{"OAUTH2_CLIENT_ID"},
			},
			&cli.StringFlag{
				Name:    "oauth2-secret",
				EnvVars: []string{"OAUTH2_SECRET"},
			},
			&cli.StringFlag{
				Name:    "oauth2-redirect-url",
				EnvVars: []string{"OAUTH2_REDIRECT_URL"},
			},
			&cli.StringFlag{
				Name:    "oauth2-scopes",
				EnvVars: []string{"OAUTH2_SCOPES"},
			},
			&cli.StringFlag{
				Name:    "oauth2-auth-url",
				EnvVars: []string{"OAUTH2_AUTH_URL"},
			},
			&cli.StringFlag{
				Name:    "oauth2-token-url",
				EnvVars: []string{"OAUTH2_TOKEN_URL"},
			},
			&cli.StringFlag{
				Name:    "cookie-store-secret",
				EnvVars: []string{"COOKIE_STORE_SECRET"},
			},
		},

		Action: func(c *cli.Context) error {
			db, err := immersadb.New("db", 10*1024*1024)
			if err != nil {
				return err
			}

			err = db.Transaction(func(ew modifier.EntityWriter) error {
				if !ew.Exists(dbpath.New("targets")) {
					err2 := ew.CreateMap(dbpath.New("targets"))
					if err2 != nil {
						return err2
					}
				}

				if !ew.Exists(dbpath.New("status")) {
					status := executor.StatusList{}
					err2 := ew.CreateData(dbpath.New("status"), (&status).Write)
					if err2 != nil {
						return err2
					}
				}
				return nil
			})

			if err != nil {
				return err
			}

			err = db.GC()
			if err != nil {
				return err
			}

			handlers := []negroni.Handler{}

			if c.Bool("oauth2") {

				var cookiestoreSecret []byte

				err := db.Transaction(func(ew modifier.EntityWriter) error {
					cookieStoreSecretPath := dbpath.New("cookieStoreSecret")
					if !ew.Exists(cookieStoreSecretPath) {
						cookiestoreSecret = []byte(c.String("cookie-store-secret"))
						if len(cookiestoreSecret) == 0 {
							cookiestoreSecret = make([]byte, 20)
							_, err2 := rand.Read(cookiestoreSecret)
							if err2 != nil {
								return err
							}
						}
						err = ew.CreateData(cookieStoreSecretPath, func(w io.Writer) error {
							_, err2 := w.Write(cookiestoreSecret)
							return err2
						})
						if err != nil {
							return err
						}
						return nil
					}
					data, err := ioutil.ReadAll(ew.EntityReaderFor(cookieStoreSecretPath).Data())
					if err != nil {
						return err
					}
					cookiestoreSecret = data
					return nil

				})

				if err != nil {
					return err
				}

				handlers = []negroni.Handler{
					sessions.Sessions("snitch_session", cookiestore.New(cookiestoreSecret)),
					oauth2.NewOAuth2Provider(
						&oauth2.Config{
							ClientID:     c.String("oauth2-client-id"),
							ClientSecret: c.String("oauth2-secret"),
							RedirectURL:  c.String("oauth2-redirect-url"),
							Scopes:       strings.Split(c.String("oauth2-scopes"), ","),
						},
						c.String("oauth2-auth-url"),
						c.String("oauth2-token-url"),
					),
					oauth2.LoginRequired(),
				}
			}

			err = executor.Start(db)
			if err != nil {
				return err
			}

			kickback.Run(
				fmt.Sprintf(":%s", c.String("port")),
				db,
				handlers,
			)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Printf("Error: %s\n", err)
	}
	os.Exit(1)

}
