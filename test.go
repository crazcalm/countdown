package main

import (
	"time"
	"fmt"
	"math"
	"os"
	"os/exec"
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


func main(){
	loc := time.Location{}
	future := time.Date(2018, 1, 19, 0, 0, 0, 0, &loc)
	var diff time.Duration

	for i:= 0; i<1000000; i++ {

		Clear()
		diff = time.Until(future)

		fmt.Printf("days: %f\nhours: %f\nminutes: %f\nseconds: %f\n\n\n", Days(diff.Hours()), HoursMod24(diff.Hours()), diff.Minutes(), diff.Seconds())
		time.Sleep(2 * time.Second)
	}
}