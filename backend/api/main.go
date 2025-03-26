package main

import (
	"log"
)

func main() {
	app := &application{
		Addr: ":8080",
	}
	mux := app.mount()
	log.Fatal(app.run(mux))
}
