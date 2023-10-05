package timings

import (
	"errors"
	"sync"
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
	mutex                = new(sync.Mutex)

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
	mutex.Lock()
	defer mutex.Unlock()

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
	return currentTiming == nil
}

func CanStopTiming(action Action) bool {
	if currentTiming == nil {
		return false
	}

	return timings[*currentTiming].Type == action
}

func StopTiming() error {
	mutex.Lock()
	defer mutex.Unlock()

	if currentTiming == nil {
		return ErrTimingAlreadyStopped
	}

	timings[*currentTiming].Stop = time.Now()
	timings[*currentTiming].Duration = time.Since(timings[*currentTiming].Start)

	currentTiming = nil

	return nil
}
