package api

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func AddRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		err := tmpl.Execute(w, data)
		if err != nil {
			log.Fatalf("Could not load index.html")
		}
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Healthy")
	})
}
