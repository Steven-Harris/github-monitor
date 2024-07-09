package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"text/template"

	api "github.com/steven-harris/github-monitor/api"
)

func main() {
	handleSigTerms()

	client, err := api.NewGitHubHttpClient()
	if err != nil {
		log.Fatalf("Could not create github http client: %s\n", err)
		os.Exit(1)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		tmpl := template.Must(template.ParseFiles("index.html"))
		data, err := client.GetPullRequests()
		if err != nil {
			log.Fatalf("Could not fetch data: %s\n", err)
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			log.Fatalf("Could not load index.html")
		}
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Healthy")
	})

	fmt.Println("Server starting on port 8888")
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
