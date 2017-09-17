package tx

import (
	"github.com/draganm/immersadb/dbpath"
	"github.com/draganm/immersadb/modifier"
	"github.com/draganm/snitch/executor"
	uuid "github.com/satori/go.uuid"
)

func (t TX) AddTarget(c *executor.Config) (string, error) {
	id := uuid.NewV4().String()
	return id, t.Transaction(func(ew modifier.EntityWriter) error {
		err := ew.CreateMap(dbpath.New("targets", id))
		if err != nil {
			return err
		}

		err = ew.CreateData(dbpath.New("targets", id, "config"), c.Write)
		if err != nil {
			return err
		}

		err = ew.CreateArray(dbpath.New("targets", id, "log"))
		if err != nil {
			return err
		}

		status := executor.StatusList{}
		err = status.Read(ew.EntityReaderFor(dbpath.New("status")))
		if err != nil {
			return err
		}

		status = append(status, executor.Status{
			ID:     id,
			Name:   c.Name,
			Status: "unknown",
		})

		err = ew.CreateData(dbpath.New("status"), (&status).Write)
		if err != nil {
			return err
		}
		return nil
	})
}
