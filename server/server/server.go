package server

import (
	"encoding/json"
	"fmt"
	"grepmynotes/search"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type SearchResponse struct {
	Entries []search.Entry `json:"entries"`
}

var searcher *search.Searcher

func Run(port int, filesPath string) error {
	searcher = search.NewSearcher(filesPath)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"chrome-extension://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", Search)

	fmt.Printf("Listening to http://localhost:%d. Content dir is %s", port, searcher.Path)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), r)
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
