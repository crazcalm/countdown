package main

import (
	"time"
	"fmt"
	"math"
	"os"
	"os/exec"
	"flag"
	"log"
	"strings"
)


func Days(hours float64) (days float64) {
       days = math.Floor(hours / 24)
       return days
}


func TimeModX(allHours, mod float64) (result float64) {
	result = math.Floor(math.Mod(allHours, mod))
	return
}

func Clear(){
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func ValidYear(currentYear, futureYear int) (bool) {
	return currentYear <= futureYear
}

func ValidMonth(month int) (bool) {
	return 0 <= month && month <= 12
}

func ValidDay(lowerBound, upperBound, day int) bool {
	return lowerBound <= day && day <= upperBound
}

func ValidHours(hours int)(bool){
	return 0 <= hours && hours <= 23
}

func ValidMinutes(minutes int)(bool){
	return 0 <= minutes && minutes <= 59
}

func ValidDate(futureDate time.Time) (err error) {
	now := time.Now()
	if !now.Before(futureDate){
		err = fmt.Errorf("Please choose a date in the future.")
	}
	return
}

func ValidateInput(year, month, day, hours, minutes, currentYear, lowerDayBound, upperDayBound int) (err error) {
	if year == 0 && month == 0 && day == 0 {
		err = fmt.Errorf("Please pass a date to be use for the countdown")
		return
	}

	if !ValidYear(currentYear, year){
		err = fmt.Errorf("year: %d is not in the future.", year)
		return
	}

	if !ValidMonth(month){
		err = fmt.Errorf("%d is not a valid month. Try using the numbers 1 - 12.", month)
		return
	}

	if !ValidDay(lowerDayBound, upperDayBound, day){
		err = fmt.Errorf("%d is not a valid day for the selected month.", day)
		return
	}

	if !ValidHours(hours){
		err = fmt.Errorf("%d is not a valid hour. Select an hour between 0 and 23.", hours)
		return
	}

	if !ValidMinutes(minutes){
		err = fmt.Errorf("%d is not a valid minute. Select an minute between 0 and 59.", minutes)
		return
	}
	return
}

func FirstAndLastDayOfTheMonth(year, month int, location *time.Location)(first, last int){
	date := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, location)
	first = date.Day()
	last =date.AddDate(0, 1, -1).Day()
	return
}

var year = flag.Int("year", 0, "Year")
var month = flag.Int("month", 0, "Month")
var day = flag.Int("day", 0, "Day")
var hours = flag.Int("hours", 0, "Hours")
var minutes = flag.Int("minutes", 0, "Minutes")

var refreshSeconds = flag.Int("refresh", 1, "Refresh time in seconds")
var msg = flag.String("msg", "The magic day is almost here!", "The message that is displayed with the countdown")

func main(){
	flag.Parse()
	loc, err := time.LoadLocation("Local")
	if err != nil {
		log.Fatal("Error with the location")
	}

	now := time.Now()
	currentYear, _, _ := now.Date()
	dayLowerBound, dayUpperBound := FirstAndLastDayOfTheMonth(*year, *month, loc)

	err = ValidateInput(*year, *month, *day, *hours, *minutes, currentYear, dayLowerBound, dayUpperBound)
	if err != nil {
		log.Fatalln(err)
	}

	future := time.Date(*year, time.Month(*month), *day, *hours, *minutes, 0, 0, loc)

	err = ValidDate(future)
	if err != nil {
		log.Fatalln(err)
	}
	
	var diff time.Duration

	for {

		Clear()
		diff = time.Until(future)
		if strings.Contains(diff.String(), "-") {
			fmt.Printf("The countdown for '%s' is over!\n", *msg)
			break
		}
		fmt.Printf("%d-%d-%d: %s\n\n", *year, *month, *day, *msg)

		fmt.Printf("%.0fd %.0fh %.0fm %.0fs\n", Days(diff.Hours()), TimeModX(diff.Hours(), 24), TimeModX(diff.Minutes(), 60), TimeModX(diff.Seconds(), 60))
		time.Sleep(time.Duration(*refreshSeconds) * time.Second)
	}
}