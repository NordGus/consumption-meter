package main

import (
	"embed"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"html/template"
	"log"
	"net/http"
)

var (
	//go:embed templates
	templates embed.FS
)

func main() {
	tmpls, err := template.ParseFS(templates, "templates/layout.gohtml", "templates/timer-component.gohtml", "templates/timings.gohtml", "templates/timing.gohtml", "templates/timer.gohtml")
	if err != nil {
		log.Fatalf("something went wrong parsing templates: %v\n", err)
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(devCORSMiddleware)

	router.Get("/", func(writer http.ResponseWriter, _ *http.Request) {
		err := tmpls.ExecuteTemplate(writer, "layout.gohtml", nil)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	})

	err = http.ListenAndServe(":4269", router)
	if err != nil {
		log.Fatalf("something went wrong initalizing http server: %v\n", err)
	}
}

func devCORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")

		next.ServeHTTP(writer, request)
	})
}
