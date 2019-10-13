package main

import (
	"fmt"
	"github.com/douglasmakey/oauth2-example/handlers"
	"log"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr: fmt.Sprintf(":8000"),
		Handler: handlers.New(),
	}

	log.Printf("Starting HTTP Server. Listening at %q", server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("%v", err)
	} else {
		log.Println("Server closed!")
	}
}
