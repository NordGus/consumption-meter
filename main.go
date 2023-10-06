package main

import (
	"embed"
	"flag"
	"fmt"
	"github.com/NordGus/consuption-meter/server/timings"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

var (
	environment *string
	port        *int

	//go:embed dist
	bundle embed.FS
)

func init() {
	environment = flag.String("env", "development", "deployment environment")
	port = flag.Int("port", 4269, "port where the application listens")

	flag.Parse()
}

func main() {
	err := timings.Initialize(*environment)
	if err != nil {
		log.Fatalln(err)
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Route("/", timings.Endpoints)

	router.Get("/dist", http.RedirectHandler("/dist/", http.StatusMovedPermanently).ServeHTTP)
	router.Get("/dist/*", http.FileServer(http.FS(bundle)).ServeHTTP)

	err = http.ListenAndServe(fmt.Sprintf(":%v", *port), router)
	if err != nil {
		log.Fatalf("something went wrong initalizing http server: %v\n", err)
	}
}
