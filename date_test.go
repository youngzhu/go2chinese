package go2chinese

import (
	"fmt"
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	testcases := []struct {
		date time.Time
		want string
	}{
		{newDate(2023, 10, 13), "2023年10月13日"},
		{newDate(2023, 9, 1), "2023年9月1日"},
	}
	for _, testcase := range testcases {
		t.Run("", func(t *testing.T) {
			got := Date(testcase.date)
			if got != testcase.want {
				t.Errorf("want: %s, but got: %s", testcase.want, got)
			}
		})
	}
}

func TestDatetime(t *testing.T) {
	fmt.Println(2 << 8)
	testcases := []struct {
		datetime time.Time
		want     string
	}{
		{newDatetime(2023, 10, 13, 0, 0, 0), "2023年10月13日 0时0分0秒"},
		{newDatetime(2023, 10, 13, 1, 59, 59), "2023年10月13日 1时59分59秒"},
		{newDatetime(2023, 10, 13, 1, 9, 59), "2023年10月13日 1时9分59秒"},
		{newDatetime(2023, 10, 13, 23, 59, 59), "2023年10月13日 23时59分59秒"},
		{newDatetime(2006, 1, 2, 0, 0, 0), "2006年1月2日 0时0分0秒"},
		{newDatetime(2006, 1, 2, 15, 4, 5), "2006年1月2日 15时4分5秒"},
	}
	for _, testcase := range testcases {
		t.Run("", func(t *testing.T) {
			got := Datetime(testcase.datetime)
			if got != testcase.want {
				t.Errorf("want: %q, but got: %q", testcase.want, got)
			}
		})
	}
}

func TestWeekday(t *testing.T) {
	testcases := []struct {
		date time.Time
		name string
	}{
		{time.Now(), longWeekdayNames[time.Now().Weekday()]},
		{newDate(2023, 10, 13), Friday},
		{newDate(2023, 10, 14), Saturday},
		{newDate(2023, 10, 15), Sunday},
		{newDate(2023, 10, 16), Monday},
		{newDate(2023, 10, 17), Tuesday},
		{newDate(2023, 10, 18), Wednesday},
		{newDate(2023, 10, 19), Thursday},
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

func TestShortWeekday(t *testing.T) {
	testcases := []struct {
		date time.Time
		name string
	}{
		{time.Now(), shortWeekdayNames[time.Now().Weekday()]},
		{newDate(2023, 10, 13), FridayShort},
		{newDate(2023, 10, 14), SaturdayShort},
		{newDate(2023, 10, 15), SundayShort},
		{newDate(2023, 10, 16), MondayShort},
		{newDate(2023, 10, 17), TuesdayShort},
		{newDate(2023, 10, 18), WednesdayShort},
		{newDate(2023, 10, 19), ThursdayShort},
	}
	for _, testcase := range testcases {
		t.Run("", func(t *testing.T) {
			got := ShortWeekday(testcase.date)
			if got != testcase.name {
				t.Errorf("want: %s, but got: %s", testcase.name, got)
			}
		})
	}
}

func newDate(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}

func newDatetime(year int, month time.Month, day, hour, min, sec int) time.Time {
	return time.Date(year, month, day, hour, min, sec, 0, time.Local)
}
