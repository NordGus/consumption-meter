package main

import (
	"embed"
	"fmt"
	"github.com/NordGus/consuption-meter/server/timings"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"html/template"
	"log"
	"net/http"
	"slices"
	"strings"
	"time"
)

var (
	//go:embed templates
	templates embed.FS

	templatePatterns = []string{
		"templates/layout.gohtml",
		"templates/components/timer-component.gohtml",
		"templates/components/timings.gohtml",
		"templates/components/timing.gohtml",
		"templates/components/timer.gohtml",
	}

	helpers = template.FuncMap{
		"produces":   func(timing timings.Timing) bool { return timing.Type == timings.Produce },
		"formatTime": func(t time.Time) string { return t.Format("02/01/06 - 15:04:05") },
		"humanizeDuration": func(d time.Duration) string {
			var (
				ms  = d.Milliseconds()
				mil = ms % 1000
				out = make([]string, 0, 3)
			)

			ms = ms / 1000

			if ms > 0 && ms%60 < 10 {
				out = append(out, fmt.Sprintf("0%d.%d", ms%60, mil))
			} else if ms > 0 {
				out = append(out, fmt.Sprintf("%d.%d", ms%60, mil))
			} else {
				out = append(out, fmt.Sprintf("00.%d", mil))
			}

			ms = ms / 60

			if ms > 0 && ms%60 < 10 {
				out = append(out, fmt.Sprintf("0%d", ms%60))
			} else if ms > 0 {
				out = append(out, fmt.Sprintf("%d", ms%60))
			}

			ms = ms / 60

			if ms > 0 && ms < 10 {
				out = append(out, fmt.Sprintf("0%d", ms))
			} else if ms > 0 {
				out = append(out, fmt.Sprintf("%d", ms))
			}

			slices.Reverse(out)

			return strings.Join(out, ":")
		},
	}
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(devCORSMiddleware)

	router.Get("/", func(writer http.ResponseWriter, _ *http.Request) {
		page := timings.NewPage()

		tmpl, err := template.New("layout").Funcs(helpers).ParseFS(templates, templatePatterns...)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		err = page.Render(writer, tmpl, "layout")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	})

	router.Post("/produce", func(writer http.ResponseWriter, _ *http.Request) {
		handleTiming(writer, timings.Produce)
	})

	router.Post("/consume", func(writer http.ResponseWriter, _ *http.Request) {
		handleTiming(writer, timings.Consume)
	})

	err := http.ListenAndServe(":4269", router)
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

func handleTiming(writer http.ResponseWriter, action timings.Action) {
	if timings.CanCreateTiming() {
		_ = timings.CreateTiming(action)
	} else if timings.CanStopTiming(action) {
		err := timings.StopTiming()
		if err != nil {
			log.Println(err)
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		log.Println(timings.ErrInvalidAction)
		http.Error(writer, timings.ErrInvalidAction.Error(), http.StatusInternalServerError)
		return
	}

	page := timings.NewPage()

	tmpl, err := template.New("timings").Funcs(helpers).ParseFS(templates, templatePatterns...)
	if err != nil {
		log.Println(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	err = page.Render(writer, tmpl, "timings")
	if err != nil {
		log.Println(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
