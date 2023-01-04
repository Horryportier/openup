package main

import (
	"log"
	app "github.com/app/openup/v1"
)

func main() {
        err := app.Start()
        if err != nil {
                log.Fatal(err)
        }
}


