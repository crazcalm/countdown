package countdown

import (
	"testing"
	"time"
)

func TestDays(t *testing.T) {
	cases := []struct {
		Hours  float64
		Answer float64
	}{
		{1, 0},
		{23, 0},
		{27, 1},
		{48, 2},
		{55, 2},
	}

	var result float64
	for _, tests := range cases {
		result = Days(tests.Hours)
		if result != tests.Answer {
			t.Errorf("Days(%.0f) = %.0f, but the answer should be %.0f", tests.Hours, result, tests.Answer)
		}
	}
}

func TestTimeModX(t *testing.T) {
	cases := []struct {
		Time   float64
		Mod    float64
		Answer float64
	}{
		{12, 24, 12},
		{36, 24, 12},
		{66, 60, 6},
	}

	var result float64
	for _, tests := range cases {
		result = TimeModX(tests.Time, tests.Mod)
		if result != tests.Answer {
			t.Errorf("TimeModX(%.0f) = %.0f, but the answer should be %.0f", tests.Time, result, tests.Answer)
		}
	}
}

func TestValidYear(t *testing.T) {
	cases := []struct {
		CurrentYear int
		FutureYear  int
		Answer      bool
	}{
		{2017, 2017, true},
		{2017, 2020, true},
		{2017, 2016, false},
	}

	var result bool
	for _, tests := range cases {
		result = ValidYear(tests.CurrentYear, tests.FutureYear)
		if result != tests.Answer {
			t.Errorf("ValidYear(%d, %d) = %v, but the answer should be %v", tests.CurrentYear, tests.FutureYear, result, tests.Answer)
		}
	}
}

func TestValidMonth(t *testing.T) {
	cases := []struct {
		Month  int
		Answer bool
	}{
		{0, false},
		{1, true},
		{5, true},
		{12, true},
		{13, false},
	}

	var result bool
	for _, tests := range cases {
		result = ValidMonth(tests.Month)
		if result != tests.Answer {
			t.Errorf("ValidMonth(%d) = %v, but the answer should be %v", tests.Month, result, tests.Answer)
		}
	}
}

func TestValidDay(t *testing.T) {
	cases := []struct {
		LowerBound int
		UpperBound int
		Day        int
		Answer     bool
	}{
		{1, 30, 0, false},
		{1, 30, 1, true},
		{1, 30, 15, true},
		{1, 30, 30, true},
		{1, 30, 31, false},
	}

	var result bool
	for _, tests := range cases {
		result = ValidDay(tests.LowerBound, tests.UpperBound, tests.Day)
		if result != tests.Answer {
			t.Errorf("ValidDay(%d, %d, %d) = %v, but the answer should be %b", tests.LowerBound, tests.UpperBound, tests.Day, result, tests.Answer)
		}
	}
}

func TestValidHours(t *testing.T) {
	cases := []struct {
		Hours  int
		Answer bool
	}{
		{-1, false},
		{0, true},
		{1, true},
		{5, true},
		{22, true},
		{23, true},
		{24, false},
	}

	var result bool
	for _, test := range cases {
		result = ValidHours(test.Hours)
		if result != test.Answer {
			t.Errorf("ValidHours(%d) = %v, but it should be %v", test.Hours, result, test.Answer)
		}
	}
}

func TestValidMinutes(t *testing.T) {
	cases := []struct {
		Minutes int
		Answer  bool
	}{
		{-1, false},
		{0, true},
		{1, true},
		{30, true},
		{58, true},
		{59, true},
		{60, false},
	}

	var result bool
	for _, test := range cases {
		result = ValidMinutes(test.Minutes)
		if result != test.Answer {
			t.Errorf("ValidMinutes(%d) = %v, but is should be %v", test.Minutes, result, test.Answer)
		}
	}
}

func TestValidDate(t *testing.T) {
	location, err := time.LoadLocation("Local")
	if err != nil {
		t.Errorf("Error in getting local location")
	}
	cases := []struct {
		Future  time.Time
		Current time.Time
		Error   bool
	}{
		{time.Date(2018, time.Month(1), 19, 0, 0, 0, 0, location), time.Date(2017, time.Month(1), 19, 0, 0, 0, 0, location), false},
		{time.Date(2018, time.Month(1), 19, 0, 0, 0, 0, location), time.Date(2018, time.Month(1), 19, 0, 0, 0, 0, location), true},
		{time.Date(2017, time.Month(1), 19, 0, 0, 0, 0, location), time.Date(2018, time.Month(1), 19, 0, 0, 0, 0, location), true},
	}

	for _, test := range cases {
		err = ValidDate(test.Future, test.Current)

		if test.Error {
			if err == nil {
				t.Errorf("ValidDate(%v, %v) did not throw an error when it should have!", test.Future, test.Current)
			}
		} else {
			if err != nil {
				t.Errorf("ValidDate(%v, %v) did thow an error when it was not suppose to!", test.Future, test.Current)
			}
		}
	}
}

func TestFirstAndLastDayOfTheMonth(t *testing.T) {
	location, err := time.LoadLocation("Local")
	if err != nil {
		t.Errorf("Failed to get local location")
	}

	cases := []struct {
		Year     int
		Month    int
		Location *time.Location
		FirstDay int
		LastDay  int
	}{
		{2018, 1, location, 1, 31},
	}

	var resultDayOne int
	var resultDayTwo int

	for _, test := range cases {
		resultDayOne, resultDayTwo = FirstAndLastDayOfTheMonth(test.Year, test.Month, test.Location)

		if resultDayOne != test.FirstDay || resultDayTwo != test.LastDay {
			t.Errorf("FirstAndLastDayOfTheMonth(%d, %d, %v) = %d, %d but it should be %d %d", test.Year, test.Month, test.Location, resultDayOne, resultDayTwo, test.FirstDay, test.LastDay)
		}
	}

}

func TestValidInput(t *testing.T) {
	cases := []struct {
		Year          int
		Month         int
		Day           int
		Hours         int
		Minutes       int
		CurrentYear   int
		LowerDayBound int
		UpperDayBound int
		Error         bool
	}{
		{0, 0, 0, 0, 0, 0, 0, 0, true},
		{2016, 0, 0, 0, 0, 2017, 0, 0, true},
		{2018, 13, 0, 0, 0, 2017, 0, 0, true},
		{2018, 1, 33, 0, 0, 2017, 1, 30, true},
		{2018, 1, 19, 25, 0, 2017, 1, 30, true},
		{2018, 1, 19, 16, 65, 2017, 1, 30, true},
		{2018, 1, 19, 0, 0, 2017, 1, 30, false},
	}

	var err error

	for _, test := range cases {
		err = ValidateInput(test.Year, test.Month, test.Day, test.Hours, test.Minutes, test.CurrentYear, test.LowerDayBound, test.UpperDayBound)

		if err == nil && test.Error == true {
			t.Errorf("ValidateInput(%d, %d, %d, %d, %d, %d, %d, %d) did not throw an error when it should have!", test.Year, test.Month, test.Day, test.Hours, test.Minutes, test.CurrentYear, test.LowerDayBound, test.UpperDayBound)
		}
		if err != nil && test.Error == false {
			t.Errorf("ValidateInput(%d, %d, %d, %d, %d, %d, %d, %d) did throw and error when it should not have!", test.Year, test.Month, test.Day, test.Hours, test.Minutes, test.CurrentYear, test.LowerDayBound, test.UpperDayBound)
		}
	}
}
