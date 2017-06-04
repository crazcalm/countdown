package countdown

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"time"
)

// Days Converts hours into days
func Days(hours float64) (days float64) {
	days = math.Floor(hours / 24)
	return days
}

// TimeModX is used to convert time into its appropriate range
func TimeModX(time, mod float64) (result float64) {
	result = math.Floor(math.Mod(time, mod))
	return
}

// Clear clears the screen.
func Clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

// ValidYear validates a year
func ValidYear(currentYear, futureYear int) bool {
	return currentYear <= futureYear
}

// ValidateRange is used internally to check the range of items
func ValidateRange(lowerBound, upperBound, item int) bool {
	return lowerBound <= item && item <= upperBound
}

// ValidMonth validates a month
func ValidMonth(month int) bool {
	return ValidateRange(1, 12, month)
}

// ValidDay validates a day
func ValidDay(lowerBound, upperBound, day int) bool {
	return ValidateRange(lowerBound, upperBound, day)
}

// ValidHours validates the hour used
func ValidHours(hours int) bool {
	return ValidateRange(0, 23, hours)
}

// ValidMinutes validates the minutes used
func ValidMinutes(minutes int) bool {
	return ValidateRange(0, 59, minutes)
}

// ValidDate validates ensures that the used date is in the future
func ValidDate(futureDate, currentDate time.Time) (err error) {
	if !currentDate.Before(futureDate) {
		err = fmt.Errorf("Please choose a date in the future")
	}
	return
}

// ValidateInput Checks user input
func ValidateInput(year, month, day, hours, minutes, currentYear, lowerDayBound, upperDayBound int) (err error) {
	if year == 0 && month == 0 && day == 0 {
		err = fmt.Errorf("Please pass a date to be use for the countdown")
		return
	}

	if !ValidYear(currentYear, year) {
		err = fmt.Errorf("year: %d is not in the future", year)
		return
	}

	if !ValidMonth(month) {
		err = fmt.Errorf("%d is not a valid month. Try using the numbers 1 - 12", month)
		return
	}

	if !ValidDay(lowerDayBound, upperDayBound, day) {
		err = fmt.Errorf("%d is not a valid day for the selected month", day)
		return
	}

	if !ValidHours(hours) {
		err = fmt.Errorf("%d is not a valid hour. Select an hour between 0 and 23", hours)
		return
	}

	if !ValidMinutes(minutes) {
		err = fmt.Errorf("%d is not a valid minute. Select an minute between 0 and 59", minutes)
		return
	}
	return
}

// FirstAndLastDayOfTheMonth finds the first and last day of the month
func FirstAndLastDayOfTheMonth(year, month int, location *time.Location) (first, last int) {
	date := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, location)
	first = date.Day()
	last = date.AddDate(0, 1, -1).Day()
	return
}
