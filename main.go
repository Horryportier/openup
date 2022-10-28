package main

import (
	"log"
	app "openup/src/v1"
)

func main() {
        err := app.Start()
        if err != nil {
                log.Fatal(err)
        }
}


