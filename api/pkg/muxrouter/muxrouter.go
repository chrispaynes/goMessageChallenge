package muxrouter

import (
	"context"
	"goMessageChallenge/api/pkg/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

// New creates a new router with routes and handlerFuncs
func New() *mux.Router {
	return handle(mux.NewRouter())
}

func handle(r *mux.Router) *mux.Router {
	r.HandleFunc("/healthz", handlers.GetHealth).Methods("GET")

	// r.HandleFunc("/emails", handlers.GetEmails).Name("emails").Methods("GET")

	// r.HandleFunc("/email/{messageId:[0-9]+}", handlers.GetEmail).Name("email").Methods("GET")

	r.HandleFunc("/email", handlers.PostEmail).Name("email").Methods("POST")
	return r
}

// GracefulShutdown shuts the server down gracefully upon syscall notification
func GracefulShutdown(addr string) {
	srv := &http.Server{
		Addr: addr,
	}

	signalChan := make(chan os.Signal, 1)

	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	<-signalChan

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	srv.Shutdown(ctx)

	log.Println("Shutdown signal received, exiting...")
	os.Exit(0)
}
