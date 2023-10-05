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
		"templates/components/timer-app.gohtml",
		"templates/components/timings.gohtml",
		"templates/components/timing.gohtml",
		"templates/components/timer.gohtml",
		"templates/components/total.gohtml",
	}

	helpers = template.FuncMap{
		"produces":   func(timing timings.Timing) bool { return timing.Type == timings.Produce },
		"formatTime": func(t time.Time) string { return t.Format("02/01/06 - 15:04:05") },
		"humanizeDuration": func(d time.Duration) string {
			var (
				ms  = abs(d.Milliseconds())
				mil = ms % 1000
				out = make([]string, 0, 3)
			)

			if ms == 0 {
				return "0.000"
			}

			ms = ms / 1000
			out = append(out, fmt.Sprintf("%02d.%03d", ms%60, mil))

			ms = ms / 60
			if ms > 0 {
				out = append(out, fmt.Sprintf("%02d", ms%60))
			}

			ms = ms / 60
			if ms > 0 {
				out = append(out, fmt.Sprintf("%+02d", ms))
			}

			slices.Reverse(out)

			if d.Milliseconds() > 0 {
				return strings.Join(out, ":")
			}

			return fmt.Sprintf("-%s", strings.Join(out, ":"))
		},
	}

	timingsPage = timings.NewPage()
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(devCORSMiddleware)

	router.Get("/", func(writer http.ResponseWriter, _ *http.Request) {
		tmpl, err := template.New("layout").Funcs(helpers).ParseFS(templates, templatePatterns...)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		err = timingsPage.Render(writer, tmpl, "layout")
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
	timingsPage.Lock()
	defer timingsPage.Unlock()

	if timingsPage.CanCreateTiming() {
		timingsPage.CreateTiming(action)

		tmpl, err := template.New("timings").Funcs(helpers).ParseFS(templates, templatePatterns...)
		if err != nil {
			log.Println(err)
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		err = timingsPage.Render(writer, tmpl, "timings")
		if err != nil {
			log.Println(err)
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

	if timingsPage.CanStopTiming(action) {
		err := timingsPage.StopTiming()
		if err != nil {
			log.Println(err)
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl, err := template.New("timings").Funcs(helpers).ParseFS(templates, templatePatterns...)
		if err != nil {
			log.Println(err)
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		err = timingsPage.Render(writer, tmpl, "timings")
		if err != nil {
			log.Println(err)
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl, err = template.New("total").Funcs(helpers).ParseFS(templates, templatePatterns...)
		if err != nil {
			log.Println(err)
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		err = timingsPage.Render(writer, tmpl, "total")
		if err != nil {
			log.Println(err)
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl, err = template.New("timer").Funcs(helpers).ParseFS(templates, templatePatterns...)
		if err != nil {
			log.Println(err)
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		err = timingsPage.Render(writer, tmpl, "timer")
		if err != nil {
			log.Println(err)
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

	log.Println(timings.ErrInvalidAction)
	http.Error(writer, timings.ErrInvalidAction.Error(), http.StatusInternalServerError)
}

func abs[T int64](num T) T {
	if num < 0 {
		return -num
	}

	return num
}
