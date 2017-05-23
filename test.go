package main

import (
	"time"
	"fmt"
	"math"
	"os"
	"os/exec"
	"flag"
	"log"
)

/*
diff (Duration) seemes to be what I want, but I
will have to format the output...
*/

func Days(hours float64) (days float64) {
	days = math.Floor(hours / 24)
	return days
}

func HoursMod24(allHours float64) (hours float64) {
	hours = math.Floor(math.Mod(allHours, 24))
	return hours
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

func ValidateInput(year, month, day, currentYear, lowerDayBound, upperDayBound int) (err error) {
	if !ValidYear(currentYear, year){
		err = fmt.Errorf("year: %d is not in the future", year)
		return
	}

	if !ValidMonth(month){
		err = fmt.Errorf("%d is not a valid month. Try using the numbers 1 - 12", month)
		return
	}

	if !ValidDay(lowerDayBound, upperDayBound, day){
		err = fmt.Errorf("%d is not a valid day for the selected month", day)
		return
	}
	return
}

var year = flag.Int("year", 0, "Year")
var month = flag.Int("month", 0, "Month")
var day = flag.Int("day", 0, "Day")


func main(){
	flag.Parse()

	loc, err := time.LoadLocation("Local")
	if err!= nil {
		log.Fatal("Error with the location")
	}
	future := time.Date(2018, 1, 19, 0, 0, 0, 0, loc)
	var diff time.Duration

	for i:= 0; i<1000000; i++ {

		Clear()
		fmt.Printf("%d-%d-%d\n", *year, *month, *day)
		diff = time.Until(future)

		fmt.Printf("days: %f\nhours: %f\nminutes: %f\nseconds: %f\n\n\n", Days(diff.Hours()), HoursMod24(diff.Hours()), diff.Minutes(), diff.Seconds())
		time.Sleep(2 * time.Second)
	}
}