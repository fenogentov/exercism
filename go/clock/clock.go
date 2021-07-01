// Package clock is a solution to lerning #9 (exercism.io)
package clock

import (
	"fmt"
)

// Clock implement a clock that handles times without dates.
type Clock struct {
	hour, min int
}

// New constructor for Clock
func New(hour, minute int) Clock {
	m := (hour*60 + minute) % 1440
	if m < 0 {
		m = 1440 + m
	}
	return Clock{
		hour: m / 60,
		min:  m % 60,
	}
}

// String converts Clock to string
func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.hour, clock.min)
}

// Add implements the addition of minutes to clock
func (clock Clock) Add(minutes int) Clock {
	return New(clock.hour, clock.min+minutes)
}

// Subtract subtracting minutes from clock
func (clock Clock) Subtract(minutes int) Clock {
	return clock.Add(-minutes)
}
