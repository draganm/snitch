package executor

import (
	"encoding/json"

	"github.com/draganm/immersadb/modifier"
)

type Status struct {
	Name   string `json:"name"`
	ID     string `json:"id"`
	Status string `json:"status"`
}

type StatusList []Status

func (sl *StatusList) Read(r modifier.EntityReader) error {
	return json.NewDecoder(r.Data()).Decode(sl)
}
