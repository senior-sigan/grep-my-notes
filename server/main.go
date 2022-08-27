package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"grepmynotes/search"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type SearchResponse struct {
	Entries []search.Entry `json:"entries"`
}

var searcher *search.Searcher

func main() {
	portFlag := flag.Int("port", 3000, "Server port to listen to")
	srcFlag := flag.String("src", ".", "Path to the fodler with file")
	flag.Parse()
	if filesPath, err := filepath.Abs(*srcFlag); err != nil {
		log.Fatal(err)
	} else {
		searcher = &search.Searcher{
			Path: filesPath,
		}
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", Search)

	fmt.Printf("Listening to http://localhost:%d. Content dir is %s", *portFlag, searcher.Path)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *portFlag), r); err != nil {
		log.Fatal(err)
	}
}

func Search(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("query")

	entries := searcher.Find(q, 6)

	res := SearchResponse{
		Entries: entries,
	}

	out, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}
	w.Write(out)
}
