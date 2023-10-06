package timings

import (
	"github.com/go-chi/chi/v5"
	"html/template"
	"log"
	"net/http"
)

func Endpoints(router chi.Router) {
	router.Get("/", appHandler)
	router.Post("/create", creationTimingHandler)
	router.Post("/consume", consumptionTimerHandler)
}

func appHandler(w http.ResponseWriter, _ *http.Request) {
	lock.RLock()
	defer lock.RUnlock()

	paths := []string{
		"templates/layout.gohtml",
		"templates/components/timer-app.gohtml",
		"templates/components/timings.gohtml",
		"templates/components/timing.gohtml",
		"templates/components/timer.gohtml",
		"templates/components/total.gohtml",
	}

	tmpl, err := template.New("layout").Funcs(helpers).ParseFS(templates, paths...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	v := createPage()

	err = v.Render(w, tmpl, "layout")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func creationTimingHandler(w http.ResponseWriter, _ *http.Request) {
	lock.Lock()
	defer lock.Unlock()

	command := statemachine(create)

	inter, err := command()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	v := createPage()

	fragments := map[string][]string{
		"timings": {"templates/components/timings.gohtml", "templates/components/timing.gohtml"},
		"total":   {"templates/components/total.gohtml"},
		"timer":   {"templates/components/timer.gohtml"},
	}

	for fragment, paths := range fragments {
		if fragment == "timer" && inter == started {
			continue
		}

		tmpl, err := template.New(fragment).Funcs(helpers).ParseFS(templates, paths...)
		if err != nil {
			log.Printf("timings: api: error occur while parsing fragment[%s]: %v\n", fragment, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = v.Render(w, tmpl, fragment)
		if err != nil {
			log.Printf("timings: api: error occur while rendering fragment[%s]: %v\n", fragment, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func consumptionTimerHandler(w http.ResponseWriter, _ *http.Request) {
	lock.Lock()
	defer lock.Unlock()

	command := statemachine(consume)

	inter, err := command()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	v := createPage()

	fragments := map[string][]string{
		"timings": {"templates/components/timings.gohtml", "templates/components/timing.gohtml"},
		"total":   {"templates/components/total.gohtml"},
		"timer":   {"templates/components/timer.gohtml"},
	}

	for fragment, paths := range fragments {
		if fragment == "timer" && inter == started {
			continue
		}

		tmpl, err := template.New(fragment).Funcs(helpers).ParseFS(templates, paths...)
		if err != nil {
			log.Printf("timings: api: error occur while parsing fragment[%s]: %v\n", fragment, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = v.Render(w, tmpl, fragment)
		if err != nil {
			log.Printf("timings: api: error occur while rendering fragment[%s]: %v\n", fragment, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
