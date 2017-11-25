package tx

import (
	"github.com/draganm/immersadb"
	"github.com/draganm/immersadb/modifier"
	"github.com/draganm/snitch/executor"
)

type TX struct {
	*immersadb.ImmersaDB
}

func (t TX) DeleteTarget(id string) error {
	return t.Transaction(func(m modifier.MapWriter) error {

		err := m.ModifyMap("targets", func(m modifier.MapWriter) error {
			return m.DeleteKey(id)
		})
		if err != nil {
			return err
		}

		status := executor.StatusList{}
		err = m.ReadData("status", status.Read)
		if err != nil {
			return err
		}

		targetIndex := -1
		for i, s := range status {
			if s.ID == id {
				targetIndex = i
			}
		}
		if targetIndex >= 0 {
			status = append(status[:targetIndex], status[targetIndex+1:]...)
		}
		err = m.SetData("status", status.Write)
		if err != nil {
			return err
		}
		return nil
	})
}
