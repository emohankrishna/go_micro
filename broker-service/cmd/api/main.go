package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = "80"

type Config struct{}

func main() {
	app := Config{}
	log.Printf("Starting broker service on port %s\n", PORT)
	svr := &http.Server{
		Addr:    fmt.Sprintf(":%s", PORT),
		Handler: app.routes(),
	}
	err := svr.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
