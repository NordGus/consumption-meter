package timings

import (
	"html/template"
	"net/http"
	"time"
)

type view struct {
	Timings   []Timing
	TotalTime time.Duration
}

func createPage() *view {
	p := view{
		Timings:   make([]Timing, len(timings)),
		TotalTime: totalTime,
	}

	copy(p.Timings, timings)

	return &p
}

func (p *view) Render(wr http.ResponseWriter, tmpl *template.Template, name string) error {
	return tmpl.ExecuteTemplate(wr, name, p)
}
