package main

import (
	"goMessageChallenge/api/pkg/muxrouter"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/urfave/negroni"
)

func main() {
	r := muxrouter.New()
	log.Println("Starting goMessageChallenge application...")

	addr := "0.0.0.0:3000"
	defer muxrouter.GracefulShutdown(addr)

	n := negroni.Classic()
	n.UseHandler(r)

	go func() {
		log.Println("goMessageChallenge application is listening on Port 3000")
		log.Fatal(http.ListenAndServe(":3000", handlers.CORS()(r)))
	}()

}
