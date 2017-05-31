package countdown

import (
	"testing"
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
