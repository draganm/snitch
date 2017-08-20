package executor

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/draganm/immersadb"
	"github.com/draganm/immersadb/dbpath"
	"github.com/draganm/immersadb/modifier"
)

type startedLog struct {
}

type config struct {
	Name     string `json:"name"`
	Image    string `json:"image"`
	Command  string `json:"command"`
	Interval int    `json:"interval"`
}

func Start(db *immersadb.ImmersaDB) error {

	executors := map[string]*executor{}
	client, err := client.NewEnvClient()
	if err != nil {
		return err
	}
	db.AddListenerFunc(dbpath.New("targets"), func(r modifier.EntityReader) {
		err := r.ForEachMapEntry(func(key string, reader modifier.EntityReader) error {

			log.Println("Targets changed!", key)

			c := &config{}
			e := json.NewDecoder(reader.EntityReaderFor(dbpath.New("config")).Data()).Decode(c)
			if e != nil {
				log.Println(e)
				return e
			}

			ex, found := executors[key]
			if !found {
				log.Println("adding executor", key)
				ex = &executor{
					config: c,
					dc:     client,
					db:     db,
					id:     key,
				}
				executors[key] = ex
				ex.start()
				return nil
			}

			// TODO update/restart executer if needed

			return nil
		})
		log.Println("done with targets")
		if err != nil {
			log.Println(err)
		}
		// TODO remove executors that have not been seen
	})

	return nil

}

type executor struct {
	dc     *client.Client
	config *config
	db     *immersadb.ImmersaDB
	id     string
	// path   dbpath.Path
	// chan config{}
}

type infoMessage struct {
	Info string `json:"info"`
}

type errorMessage struct {
	Context string `json:"context"`
	Error   string `json:"error"`
}

func (e *executor) start() {
	go func() {
		// log.Println("executor for", *e.config, "started")
		e.log("info", infoMessage{"Executor started"})
		_, _, err := e.dc.ImageInspectWithRaw(context.Background(), e.config.Image)

		if client.IsErrImageNotFound(err) {
			e.log("info", infoMessage{fmt.Sprintf("Pulling image %s", e.config.Image)})
			var stream io.ReadCloser
			stream, err = e.dc.ImagePull(context.Background(), e.config.Image, types.ImagePullOptions{})
			if err == nil {
				defer stream.Close()
				io.Copy(os.Stdout, stream)
			}
			if err != nil {
				e.log("error", errorMessage{Context: "Pulling image", Error: err.Error()})
			}

			err = nil
		}

		if err != nil {
			e.log("error", errorMessage{Context: "Pulling image", Error: err.Error()})
		}

	}()
}

type logEntry struct {
	Time   time.Time   `json:"time"`
	Fields interface{} `json:"fields"`
	Level  string      `json:"level"`
	// Type   string      `json:"type"`
}

// func (e *executor) log(level string, fields interface{}) error {
//
// })

func (e *executor) log(level string, fields interface{}) error {
	json.NewEncoder(os.Stdout).Encode(logEntry{Time: time.Now(), Fields: fields, Level: level})
	return e.db.Transaction(func(ew modifier.EntityWriter) error {
		return ew.CreateData(dbpath.New("targets", e.id, "log", 0), func(w io.Writer) error {
			return json.NewEncoder(w).Encode(logEntry{Time: time.Now(), Fields: fields, Level: level})
		})
	})
}

func (e *executor) update(s string) error {
	return e.db.Transaction(func(w modifier.EntityWriter) error {
		statuses := []*status{}
		err := json.NewDecoder(w.EntityReaderFor(dbpath.New("status")).Data()).Decode(statuses)
		if err != nil {
			return err
		}

		for _, st := range statuses {
			if st.ID == e.id {
				st.Status = s
			}
		}
		return w.CreateData(dbpath.New("status"), func(w io.Writer) error {
			return json.NewEncoder(w).Encode(statuses)
		})
	})
}

func (e *executor) stop() {

}
