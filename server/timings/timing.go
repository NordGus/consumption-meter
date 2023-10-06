package timings

import (
	"errors"
	"time"
)

var (
	ErrTimingNotStopped       = errors.New("timings: there's another timing running")
	ErrTimingAlreadyStopped   = errors.New("timings: timing already stopped")
	ErrDifferentActionRunning = errors.New("timings: another action's timing is running")
)

type Timing struct {
	ID       uint64
	Start    time.Time
	Stop     time.Time
	Duration time.Duration
	Type     action
}

func startTiming(action action) error {
	if isRunning {
		return ErrTimingNotStopped
	}

	t := Timing{
		ID:    timingID,
		Start: time.Now(),
		Type:  action,
	}

	isRunning = true
	timings = append(timings, t)
	timingID++

	return nil
}

func stopTiming() error {
	if !isRunning {
		return ErrTimingAlreadyStopped
	}

	var (
		idx = len(timings) - 1
	)

	timings[idx].Stop = time.Now()
	timings[idx].Duration = timings[idx].Stop.Sub(timings[idx].Start)

	calculateTotalTime()

	isRunning = false

	return nil
}
