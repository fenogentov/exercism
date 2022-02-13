package meetup

import (
	"time"
)

type WeekSchedule int

const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Last
	Teenth
)

func Day(week WeekSchedule, weekday time.Weekday, month time.Month, year int) (day int) {
	days := weekdayDays(weekday, month, year)
	switch week {
	case First, Second, Third, Fourth:
		return days[week]
	case Last:
		return days[len(days)-1]
	case Teenth:
		for i, d := range days {
			if d >= 13 && d <= 19 {
				return days[i]
			}
		}
	}

	return 0
}

// weekdayDays returns the list of days which have the specified weekday in the given
// month dan year.
//
// For example:
//   weekdayDays(time.Monday, time.January, 2000) would return all the Monday days in Jan 2000.
//

func weekdayDays(weekday time.Weekday, month time.Month, year int) []int {
	firstOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	offset := int(weekday - firstOfMonth.Weekday())
	if offset < 0 {
		offset = 7 + offset
	}
	days := []int{}
	for d := firstOfMonth.AddDate(0, 0, offset); d.Before(lastOfMonth) || d.Equal(lastOfMonth); d = d.AddDate(0, 0, 7) {
		days = append(days, d.Day())
	}
	return days
}
