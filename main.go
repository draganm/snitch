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

	"gitlab.netice9.com/dragan/favicon"

	_ "net/http/pprof"

	"github.com/draganm/immersadb"
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
			db, err := immersadb.New("db")
			if err != nil {
				return err
			}

			err = db.Transaction(func(m modifier.MapWriter) error {
				if !m.HasKey("targets") {
					err2 := m.CreateMap("targets", nil)
					if err2 != nil {
						return err2
					}
				}

				if !m.HasKey("status") {
					status := executor.StatusList{}
					err2 := m.SetData("status", (&status).Write)
					if err2 != nil {
						return err2
					}
				}
				return nil
			})

			if err != nil {
				return err
			}

			handlers := []negroni.Handler{}

			if c.Bool("oauth2") {

				var cookiestoreSecret []byte

				err := db.Transaction(func(m modifier.MapWriter) error {

					// cookieStoreSecretPath := dbpath.New("cookieStoreSecret")
					if !m.HasKey("cookieStoreSecret") {
						cookiestoreSecret = []byte(c.String("cookie-store-secret"))
						if len(cookiestoreSecret) == 0 {
							cookiestoreSecret = make([]byte, 20)
							_, err2 := rand.Read(cookiestoreSecret)
							if err2 != nil {
								return err
							}
						}
						err = m.SetData("cookieStoreSecret", func(w io.Writer) error {
							_, err2 := w.Write(cookiestoreSecret)
							return err2
						})
						if err != nil {
							return err
						}
						return nil

					}
					err = m.ReadData("cookieStoreSecret", func(r io.Reader) error {
						data, err := ioutil.ReadAll(r)
						if err != nil {
							return err
						}
						cookiestoreSecret = data
						return nil
					})

					if err != nil {
						return err
					}
					return nil

				})

				if err != nil {
					return err
				}

				handlers = []negroni.Handler{
					favicon.NegroniHandler,
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
