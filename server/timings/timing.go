package timings

import (
	"errors"
	"time"
)

var (
	ErrTimingAlreadyStopped = errors.New("timings: timing already stopped")
	ErrInvalidAction        = errors.New("timings: invalid action")
)

type Timing struct {
	ID       uint64
	Start    time.Time
	Stop     time.Time
	Duration time.Duration
	Type     Action
}

func (page *Page) CreateTiming(action Action) {
	t := Timing{
		ID:    page.timingID,
		Start: time.Now(),
		Type:  action,
	}

	page.isRunning = true
	page.timings = append(page.timings, t)
	page.timingID++
}

func (page *Page) CanCreateTiming() bool {
	return !page.isRunning
}

func (page *Page) CanStopTiming(action Action) bool {
	lastIDX := len(page.timings) - 1

	if lastIDX < 0 || !page.isRunning {
		return false
	}

	return page.timings[lastIDX].Type == action
}

func (page *Page) StopTiming() error {
	if !page.isRunning {
		return ErrTimingAlreadyStopped
	}

	lastIDX := len(page.timings) - 1

	page.timings[lastIDX].Stop = time.Now()
	page.timings[lastIDX].Duration = time.Since(page.timings[lastIDX].Start)
	page.isRunning = false

	page.CalculateTotalTime()

	return nil
}
