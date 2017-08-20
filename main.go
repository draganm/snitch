package main

import (
	"fmt"
	"os"

	"github.com/draganm/reciprocus"
	"github.com/draganm/snitch/executor"
)

func main() {
	r, err := reciprocus.New("js", "db", 10*1024*1024)
	if err != nil {
		panic(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	executor.Start(r.DB)
	r.Serve(fmt.Sprintf(":%s", port))
}
