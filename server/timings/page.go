package timings

import (
	"html/template"
	"net/http"
	"time"
)

type Page struct {
	Timings   []Timing
	TotalTime time.Duration
}

func NewPage() Page {
	p := Page{
		Timings:   timings,
		TotalTime: 0,
	}

	p.calculateTotalTime()

	return p
}

func (page *Page) calculateTotalTime() {
	for _, timing := range page.Timings {
		if timing.Type == Consume {
			page.TotalTime -= timing.Duration
		} else {
			page.TotalTime += timing.Duration
		}
	}
}

func (page *Page) Render(wr http.ResponseWriter, tmpl *template.Template, name string) error {
	return tmpl.ExecuteTemplate(wr, name, page)
}
