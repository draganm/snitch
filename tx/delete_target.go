package tx

import (
	"github.com/draganm/immersadb"
	"github.com/draganm/immersadb/dbpath"
	"github.com/draganm/immersadb/modifier"
	"github.com/draganm/snitch/executor"
)

type TX struct {
	*immersadb.ImmersaDB
}

func (t TX) DeleteTarget(id string) error {
	return t.Transaction(func(ew modifier.EntityWriter) error {
		ew.Delete(dbpath.New("targets", id))
		status := executor.StatusList{}
		err := status.Read(ew.EntityReaderFor(dbpath.New("status")))
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
		err = ew.CreateData(dbpath.New("status"), status.Write)
		if err != nil {
			return err
		}
		return nil
	})
}
