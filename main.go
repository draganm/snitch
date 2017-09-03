package main

import (
	"fmt"
	"os"

	"github.com/draganm/immersadb"
	"github.com/draganm/kickback"
	"github.com/draganm/snitch/executor"

	_ "github.com/draganm/snitch/ui"
)

func main() {

	db, err := immersadb.New("db", 10*1024*1024)
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	err = executor.Start(db)
	if err != nil {
		panic(err)
	}

	kickback.Run(fmt.Sprintf(":%s", port), db, nil)

}
