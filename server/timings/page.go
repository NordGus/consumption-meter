package timings

import (
	"html/template"
	"net/http"
	"sync"
	"time"
)

type Action int

const (
	Consume Action = iota
	Produce
)

type Page struct {
	timings   []Timing
	TotalTime time.Duration

	isRunning bool
	timingID  uint64

	sync.RWMutex
}

func NewPage() *Page {
	p := Page{
		timings:   make([]Timing, 0, 10),
		TotalTime: 0,
		timingID:  1,
	}

	p.CalculateTotalTime()

	return &p
}

func (page *Page) GetTimings() []Timing {
	return page.timings
}

func (page *Page) CalculateTotalTime() {
	for _, timing := range page.timings {
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
