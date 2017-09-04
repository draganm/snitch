package executor

import (
	"encoding/json"
	"io"

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

func (sl *StatusList) Write(w io.Writer) error {
	return json.NewEncoder(w).Encode(sl)
}
