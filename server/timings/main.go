// Package timings behaves like a singleton ones initialized it can't be prevents the user from reinitializing it.
// The way is design is like a sub-system that handles interactions via a http API that renders a UI for the user to
// interact with it.
package timings

import (
	"embed"
	"errors"
	"sync"
	"time"
)

var (
	ErrPackageAlreadyInitialized = errors.New("timings: package already initialized")

	environment string

	//go:embed templates
	templates embed.FS

	timings   []Timing
	totalTime time.Duration
	isRunning bool
	timingID  uint64 = 1

	lock        sync.RWMutex
	initialized bool
)

func Initialize(env string) error {
	if initialized {
		return ErrPackageAlreadyInitialized
	}

	environment = env
	timings = make([]Timing, 0, 20)
	totalTime = 0
	isRunning = false
	initialized = true

	return nil
}

func calculateTotalTime() {
	totalTime = 0

	for _, timing := range timings {
		if timing.Type == consume {
			totalTime -= timing.Duration
		} else {
			totalTime += timing.Duration
		}
	}
}

func statemachine(action action) func() (interaction, error) {
	switch {
	case len(timings) == 0:
		return func() (interaction, error) {
			return started, startTiming(action)
		}
	case timings[len(timings)-1].Type != action && isRunning:
		return func() (interaction, error) {
			return started, ErrDifferentActionRunning
		}
	case isRunning:
		return func() (interaction, error) {
			return stopped, stopTiming()
		}
	default:
		return func() (interaction, error) {
			return started, startTiming(action)
		}
	}
}
