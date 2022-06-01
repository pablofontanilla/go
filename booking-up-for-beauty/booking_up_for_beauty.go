package booking

import (
	"fmt"
	"time"
)

var stdFormat string = "1/2/2006 15:04:05"
var hasPassedFormat string = "January 2, 2006 15:04:05"
var isAfternoonFormat string = "Monday, January 2, 2006 15:04:05"
var midDay int = 12
var lateDay int = 18
var anniversaryDay int = 15
var anniversaryMonth time.Month = time.September

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {

	t, _ := time.Parse(stdFormat, date)
	return t
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	dateToCheck, _ := time.Parse(hasPassedFormat, date)
	return dateToCheck.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	t, _ := time.Parse(isAfternoonFormat, date)
	hour, _, _ := t.Clock()

	return hour >= midDay && hour < lateDay
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	dateTime := Schedule(date)
	hour, minute, _ := dateTime.Clock()
	year, month, day := dateTime.Date()
	weekday := dateTime.Weekday()
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.",
		weekday, month, day, year, hour, minute)
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	currentYear, _, _ := time.Now().Date()
	return time.Date(currentYear, anniversaryMonth, anniversaryDay, 0, 0, 0, 0, time.UTC)
}
