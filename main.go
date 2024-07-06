package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	api "github.com/steven-harris/github-monitor/api"
)

func main() {
	handleSigTerms()
	api.AddRoutes()
	fmt.Printf("Server starting on port 8888")
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func handleSigTerms() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("received SIGTERM, exiting")
		os.Exit(1)
	}()
}
