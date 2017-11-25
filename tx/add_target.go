package tx

import (
	"github.com/draganm/immersadb/modifier"
	"github.com/draganm/snitch/executor"
	uuid "github.com/satori/go.uuid"
)

func (t TX) AddTarget(c *executor.Config) (string, error) {
	id := uuid.NewV4().String()
	return id, t.Transaction(func(m modifier.MapWriter) error {
		err := m.ModifyMap("targets", func(m modifier.MapWriter) error {
			return m.CreateMap(id, func(m modifier.MapWriter) error {
				err := m.SetData("config", c.Write)
				if err != nil {
					return err
				}
				return m.CreateArray("log", nil)
			})
		})

		if err != nil {
			return err
		}

		status := executor.StatusList{}

		err = m.ReadData("status", status.Read)

		if err != nil {
			return err
		}

		status = append(status, executor.Status{
			ID:     id,
			Name:   c.Name,
			Status: "unknown",
		})

		err = m.SetData("status", status.Write)

		if err != nil {
			return err
		}
		return nil
	})
}
