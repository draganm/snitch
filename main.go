package main

import (
	"fmt"
	"os"

	"github.com/draganm/reciprocus"
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

	r.Serve(fmt.Sprintf(":%s", port))
}
