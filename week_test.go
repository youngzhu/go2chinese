package go2chinese

import (
	"testing"
	"time"
)

func TestWeekday(t *testing.T) {
	testcases := []struct {
		date time.Time
		name string
	}{
		{time.Now(), longWeekdayNames[time.Now().Weekday()]},
		{time.Date(2023, 10, 13, 0, 0, 0, 0, time.Local), Friday},
		{time.Date(2023, 10, 14, 0, 0, 0, 0, time.Local), Saturday},
		{time.Date(2023, 10, 15, 0, 0, 0, 0, time.Local), Sunday},
		{time.Date(2023, 10, 16, 0, 0, 0, 0, time.Local), Monday},
		{time.Date(2023, 10, 17, 0, 0, 0, 0, time.Local), Tuesday},
		{time.Date(2023, 10, 18, 0, 0, 0, 0, time.Local), Wednesday},
		{time.Date(2023, 10, 19, 0, 0, 0, 0, time.Local), Thursday},
	}
	for _, testcase := range testcases {
		t.Run("", func(t *testing.T) {
			got := Weekday(testcase.date)
			if got != testcase.name {
				t.Errorf("want: %s, but got: %s", testcase.name, got)
			}
		})
	}
}
