package executor

import (
	"encoding/json"
	"io"
)

type Status struct {
	Name   string `json:"name"`
	ID     string `json:"id"`
	Status string `json:"status"`
}

type StatusList []Status

func (sl *StatusList) Read(r io.Reader) error {
	return json.NewDecoder(r).Decode(sl)
}

func (sl *StatusList) Write(w io.Writer) error {
	return json.NewEncoder(w).Encode(sl)
}
