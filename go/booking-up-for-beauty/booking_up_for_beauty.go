// Package booking implements work with the schedule of a beauty salon
package booking

import (
	"fmt"
	"log"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	tm, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		log.Println(err)
	}
	return tm
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	tm, err := time.Parse("January 2, 2006 15:04:05", date)
	if err != nil {
		log.Println(err)
	}
	return time.Since(tm) > 0
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	tm, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if err != nil {
		log.Println(err)
	}
	hrs := tm.Hour()
	return (hrs >= 12 && hrs < 18)
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	tm, err := time.Parse("1/2/2006 15:04:05", date)
	strgTm := tm.Format("Monday, January 2, 2006, at 15:04")
	if err != nil {
		log.Println(err)
	}
	return fmt.Sprintf("You have an appointment on %s.", strgTm)
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
