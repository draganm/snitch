package executor

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/strslice"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/draganm/immersadb"
	"github.com/draganm/immersadb/dbpath"
	"github.com/draganm/immersadb/modifier"
	shellwords "github.com/mattn/go-shellwords"
)

type startedLog struct {
}

type config struct {
	Name     string `json:"name"`
	Image    string `json:"image"`
	Command  string `json:"command"`
	Interval int    `json:"interval"`
}

const maxLogLength = 100

func Start(db *immersadb.ImmersaDB) error {

	executors := map[string]*executor{}
	client, err := client.NewEnvClient()
	if err != nil {
		return err
	}
	db.AddListenerFunc(dbpath.New("targets"), func(r modifier.EntityReader) {
		err := r.ForEachMapEntry(func(key string, reader modifier.EntityReader) error {

			c := &config{}
			e := json.NewDecoder(reader.EntityReaderFor(dbpath.New("config")).Data()).Decode(c)
			if e != nil {
				log.Println(e)
				return e
			}

			ex, found := executors[key]
			if !found {
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
}

type infoMessage struct {
	Info string `json:"info"`
}

type errorMessage struct {
	Context string `json:"context"`
	Error   string `json:"error"`
}

type successMessage struct {
	Log      string  `json:"log"`
	Duration float64 `json:"duration"`
}

type failureMessage struct {
	Log      string  `json:"log"`
	Duration float64 `json:"duration"`
	ExitCode int     `json:"exitCode"`
}

func (e *executor) start() {
	go func() {
		for {
			err := e.updateStatus("running")
			if err != nil {
				log.Println(err)
			}
			err = e.execute()
			if err != nil {
				err = e.updateStatus("failed")
				if err != nil {
					log.Println(err)
				}
			} else {
				err = e.updateStatus("success")
				if err != nil {
					log.Println(err)
				}
			}
			time.Sleep(time.Duration(e.config.Interval) * time.Second)

		}

	}()
}

func (e *executor) updateStatus(st string) error {
	return e.db.Transaction(func(ew modifier.EntityWriter) error {
		statuses := []*Status{}
		statusPath := dbpath.New("status")
		err := json.NewDecoder(ew.EntityReaderFor(statusPath).Data()).Decode(&statuses)
		if err != nil {
			return err
		}
		for _, s := range statuses {
			log.Println("ID", s.ID, e.id)
			if s.ID == e.id {
				s.Status = st
			}
		}
		return ew.CreateData(statusPath, func(w io.Writer) error {
			return json.NewEncoder(w).Encode(statuses)
		})
	})
}

func (e *executor) execute() error {
	start := time.Now()
	e.log("info", infoMessage{"Executor started"})
	_, _, err := e.dc.ImageInspectWithRaw(context.Background(), e.config.Image)

	if err != nil && !client.IsErrImageNotFound(err) {
		e.log("error", errorMessage{Context: "Pulling image", Error: err.Error()})
		return err
	}

	if client.IsErrImageNotFound(err) {
		e.log("info", infoMessage{fmt.Sprintf("Pulling image %s", e.config.Image)})
		var stream io.ReadCloser
		stream, err = e.dc.ImagePull(context.Background(), e.config.Image, types.ImagePullOptions{})
		if err == nil {
			defer stream.Close()
			io.Copy(os.Stdout, stream)
			e.log("info", infoMessage{"Image pulled successfully"})
		}
		if err != nil {
			e.log("error", errorMessage{Context: "Pulling image", Error: err.Error()})
			return err
		}
	}
	e.log("info", infoMessage{"Starting execution"})
	cmd, err := shellwords.Parse(e.config.Command)

	if err != nil {
		e.log("error", errorMessage{Context: "Parsing command", Error: err.Error()})
		return err
	}

	cc, err := e.dc.ContainerCreate(
		context.Background(),
		&container.Config{
			Image: e.config.Image,
			Cmd:   strslice.StrSlice(cmd),
		},
		nil,
		nil,
		"",
	)
	if err != nil {
		e.log("error", errorMessage{Context: "Creating container", Error: err.Error()})
		return err
	}
	e.log("info", infoMessage{"Created container " + cc.ID})
	err = e.dc.ContainerStart(context.Background(), cc.ID, types.ContainerStartOptions{})
	if err != nil {
		e.log("error", errorMessage{Context: "Starting container", Error: err.Error()})
		return err
	}
	okchan, errchan := e.dc.ContainerWait(context.Background(), cc.ID, container.WaitConditionNotRunning)
	select {
	case <-okchan:
		e.log("info", infoMessage{"Container terminated " + cc.ID})
	case err = <-errchan:
		e.log("error", errorMessage{Context: "Waiting for the container", Error: err.Error()})
		return err
	}
	ci, err := e.dc.ContainerInspect(context.Background(), cc.ID)
	if err != nil {
		e.log("error", errorMessage{Context: "Starting container", Error: err.Error()})
		return err
	}

	exitCode := ci.State.ExitCode

	lr, err := e.dc.ContainerLogs(context.Background(), cc.ID, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Tail:       "20",
	})
	if err != nil {
		e.log("error", errorMessage{Context: "Fetching container logs", Error: err.Error()})
		return err
	}
	defer lr.Close()

	demux := &bytes.Buffer{}

	_, err = stdcopy.StdCopy(demux, demux, lr)
	if err != nil {
		e.log("error", errorMessage{Context: "Demuxing logs", Error: err.Error()})
		return err
	}

	if err != nil {
		e.log("error", errorMessage{Context: "Fetching container logs", Error: err.Error()})
		return err
	}

	err = e.dc.ContainerRemove(context.Background(), cc.ID, types.ContainerRemoveOptions{
		RemoveVolumes: true,
		Force:         true,
	})
	if err != nil {
		e.log("error", errorMessage{Context: "Deleting container", Error: err.Error()})
		return err
	}

	d := time.Now().Sub(start)

	if exitCode != 0 {
		e.log("failure", failureMessage{
			Duration: d.Seconds(),
			Log:      demux.String(),
			ExitCode: exitCode,
		})
		return errors.New("Execution failed")
	}

	e.log("success", successMessage{
		Duration: d.Seconds(),
		Log:      demux.String(),
	})
	return err

}

type logEntry struct {
	Time   time.Time   `json:"time"`
	Fields interface{} `json:"fields"`
	Level  string      `json:"level"`
}

func (e *executor) log(level string, fields interface{}) error {
	json.NewEncoder(os.Stdout).Encode(logEntry{Time: time.Now(), Fields: fields, Level: level})
	return e.db.Transaction(func(ew modifier.EntityWriter) error {
		err := ew.CreateData(dbpath.New("targets", e.id, "log", 0), func(w io.Writer) error {
			return json.NewEncoder(w).Encode(logEntry{Time: time.Now(), Fields: fields, Level: level})
		})
		if err != nil {
			return err
		}
		logReader := ew.EntityReaderFor(dbpath.New("targets", e.id, "log"))
		s := logReader.Size()

		if s <= maxLogLength {
			return nil
		}

		err = ew.Delete(dbpath.New("targets", e.id, "log", int(s-1)))
		if err != nil {
			return nil
		}

		return nil
	})
}

func (e *executor) update(s string) error {
	return e.db.Transaction(func(w modifier.EntityWriter) error {
		statuses := []*Status{}
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
