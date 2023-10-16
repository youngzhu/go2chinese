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

func TestShortWeekday(t *testing.T) {
	testcases := []struct {
		date time.Time
		name string
	}{
		{time.Now(), shortWeekdayNames[time.Now().Weekday()]},
		{time.Date(2023, 10, 13, 0, 0, 0, 0, time.Local), FridayShort},
		{time.Date(2023, 10, 14, 0, 0, 0, 0, time.Local), SaturdayShort},
		{time.Date(2023, 10, 15, 0, 0, 0, 0, time.Local), SundayShort},
		{time.Date(2023, 10, 16, 0, 0, 0, 0, time.Local), MondayShort},
		{time.Date(2023, 10, 17, 0, 0, 0, 0, time.Local), TuesdayShort},
		{time.Date(2023, 10, 18, 0, 0, 0, 0, time.Local), WednesdayShort},
		{time.Date(2023, 10, 19, 0, 0, 0, 0, time.Local), ThursdayShort},
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
