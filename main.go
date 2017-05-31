package main

import (
	"github.com/crazcalm/countdown/countdown"
	"fmt"
	"flag"
	"time"
	"log"
	"strings"
)

var year = flag.Int("year", 2020, "Year")
var month = flag.Int("month", 1, "Month")
var day = flag.Int("day", 1, "Day")
var hours = flag.Int("hours", 0, "Hours")
var minutes = flag.Int("minutes", 0, "Minutes")

var refreshSeconds = flag.Int("refresh", 1, "Refresh time in seconds")
var msg = flag.String("msg", "The magic day is almost here!", "The message that is displayed with the countdown")

func main (){
	fmt.Println("Hello World")

	flag.Parse()
	loc, err := time.LoadLocation("Local")
	if err != nil {
		log.Fatal("Error with the location")
	}
	
	now := time.Now()
	currentYear, _, _ := now.Date()
	dayLowerBound, dayUpperBound := countdown.FirstAndLastDayOfTheMonth(*year, *month, loc)
	
	err = countdown.ValidateInput(*year, *month, *day, *hours, *minutes, currentYear, dayLowerBound, dayUpperBound)
	if err != nil {
		log.Fatalln(err)
	}
	
	future := time.Date(*year, time.Month(*month), *day, *hours, *minutes, 0, 0, loc)
	
	err = countdown.ValidDate(future)
	if err != nil {
		log.Fatalln(err)
	}

	var diff time.Duration
	
	for {
		countdown.Clear()
	    diff = time.Until(future)
	    if strings.Contains(diff.String(), "-") {
	    	fmt.Printf("The countdown for '%s' is over!\n", *msg)
	        break
	    }
	   fmt.Printf("%d-%d-%d: %s\n\n", *year, *month, *day, *msg)
	   fmt.Printf("%.0fd %.0fh %.0fm %.0fs\n", countdown.Days(diff.Hours()), countdown.TimeModX(diff.Hours(), 24), countdown.TimeModX(diff.Minutes(), 60), countdown.TimeModX(diff.Seconds(), 60))
	   time.Sleep(time.Duration(*refreshSeconds) * time.Second)
	}
}