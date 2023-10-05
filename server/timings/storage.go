package timings

import (
	"errors"
	"time"
)

type Action int

const (
	Consume Action = iota
	Produce
)

var (
	currentTiming *Timing
	timings              = make([]Timing, 0, 10)
	timingID      uint64 = 1

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

func CreateTiming(action Action) Timing {
	t := Timing{
		ID:    timingID,
		Start: time.Now(),
		Type:  action,
	}

	timings = append(timings, t)

	timingID++

	currentTiming = &t

	return t
}

func CanCreateTiming() bool {
	return currentTiming == nil
}

func CanStopTiming(action Action) bool {
	if currentTiming == nil {
		return false
	}

	return currentTiming.Type == action
}

func StopTiming() error {
	if currentTiming == nil {
		return ErrTimingAlreadyStopped
	}

	currentTiming.Stop = time.Now()
	currentTiming.Duration = time.Since(currentTiming.Start)

	currentTiming = nil

	return nil
}
