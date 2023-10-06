package timings

import (
	"fmt"
	"html/template"
	"slices"
	"strings"
	"time"
)

var (
	helpers = template.FuncMap{
		"humanizeDuration": humanizeDuration,
		"environment":      func() string { return environment },
		"creates":          func(t Timing) bool { return t.Type == create },
		"formatTime":       func(t time.Time) string { return t.Format("02/01/06 - 15:04:05") },
	}
)

func humanizeDuration(d time.Duration) string {
	var (
		ms  = d.Milliseconds()
		mil = ms % 1000
		out = make([]string, 0, 3)
	)

	if ms == 0 {
		return "0.000"
	}

	if ms < 0 {
		ms = -ms
		mil = -mil
	}

	ms = ms / 1000
	out = append(out, fmt.Sprintf("%02d.%03d", ms%60, mil))

	ms = ms / 60
	if ms > 0 {
		out = append(out, fmt.Sprintf("%02d", ms%60))
	}

	ms = ms / 60
	if ms > 0 {
		out = append(out, fmt.Sprintf("%d", ms))
	}

	slices.Reverse(out)

	return strings.Join(out, ":")
}
