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
	currentTiming *uint64
	timings              = make([]Timing, 0, 10)
	timingID      uint64 = 1
	lock                 = make(chan bool, 1)

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
	lock <- true
	defer func() {
		<-lock
	}()

	t := Timing{
		ID:    timingID,
		Start: time.Now(),
		Type:  action,
	}

	idx := uint64(len(timings))

	currentTiming = &idx

	timings = append(timings, t)

	timingID++

	return t
}

func CanCreateTiming() bool {
	lock <- true
	defer func() {
		<-lock
	}()

	return currentTiming == nil
}

func CanStopTiming(action Action) bool {
	lock <- true
	defer func() {
		<-lock
	}()

	if currentTiming == nil {
		return false
	}

	return timings[*currentTiming].Type == action
}

func StopTiming() error {
	lock <- true
	defer func() {
		<-lock
	}()

	if currentTiming == nil {
		return ErrTimingAlreadyStopped
	}

	timings[*currentTiming].Stop = time.Now()
	timings[*currentTiming].Duration = time.Since(timings[*currentTiming].Start)

	currentTiming = nil

	return nil
}
