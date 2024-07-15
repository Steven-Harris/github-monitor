package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/steven-harris/github-monitor/api"
)

type dataFetcher func() (interface{}, error)

func main() {
	handleSigTerms()

	client, err := api.NewGitHubHttpClient()
	if err != nil {
		log.Fatalf("Could not create github http client: %s\n", err)
		os.Exit(1)
	}
	http.HandleFunc("/", renderHtmx("templates/index.html", nil))
	http.HandleFunc("/pull-requests", renderHtmx("templates/pull-requests.html", client.GetPullRequests))
	http.HandleFunc("/actions", renderHtmx("templates/actions.html", client.GetActions))
	http.HandleFunc("/jobs", func(w http.ResponseWriter, r *http.Request) {
		repo := r.URL.Query().Get("repo")
		runId := r.URL.Query().Get("runId")
		renderHtmx("templates/jobs.html", func() (interface{}, error) {
			return client.GetJobs(repo, runId)
		})(w, r)
	})
	http.HandleFunc("/reviews", func(w http.ResponseWriter, r *http.Request) {
		repo := r.URL.Query().Get("repo")
		prNumber := r.URL.Query().Get("prNumber")
		renderHtmx("templates/reviews.html", func() (interface{}, error) {
			return client.GetReviews(repo, prNumber)
		})(w, r)
	})

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Healthy")
	})

	http.HandleFunc("/components/*", func(w http.ResponseWriter, r *http.Request) {
		renderHtmx("/templates"+r.URL.Path+".html", nil)(w, r)
	})

	fmt.Println("Server starting on port 8888")
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func renderHtmx(templateFile string, fetchData dataFetcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		tmpl := template.Must(template.ParseFiles(templateFile))

		var data interface{}
		var err error
		if fetchData != nil {
			data, err = fetchData()
			if err != nil {
				_ = fmt.Errorf("could not fetch data: %s", err)
				return
			}
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			_ = fmt.Errorf("could not load template: %s", err)
			return
		}
	}
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
