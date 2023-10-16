package go2chinese

import (
	"strings"
	"time"
)

const (
	layoutDate           = "2006年1月2日"
	layoutDatetime       = "2006年1月2日 15时4分5秒"
	layoutDatetimeHour12 = "2006年1月2日 3时4分5秒"
)

func Date(t time.Time) string {
	return t.Format(layoutDate)
}

func Datetime(t time.Time) string {
	switch t.Hour() {
	case 0:
		// 处理  00时
		s := t.Format(layoutDatetime)
		return strings.ReplaceAll(s, " 00", " 0")
	case 1, 2, 3, 4, 5, 6, 7, 8, 9:
		return t.Format(layoutDatetimeHour12)
	default:
		return t.Format(layoutDatetime)
	}
}

// week
const (
	// Sunday ... long weekday names
	Sunday    = "星期天"
	Monday    = "星期一"
	Tuesday   = "星期二"
	Wednesday = "星期三"
	Thursday  = "星期四"
	Friday    = "星期五"
	Saturday  = "星期六"

	// SundayShort ... short weekday names
	SundayShort    = "周日"
	MondayShort    = "周一"
	TuesdayShort   = "周二"
	WednesdayShort = "周三"
	ThursdayShort  = "周四"
	FridayShort    = "周五"
	SaturdayShort  = "周六"
)

func Weekday(t time.Time) string {
	return longWeekdayNames[t.Weekday()]
}

func ShortWeekday(t time.Time) string {
	return shortWeekdayNames[t.Weekday()]
}

var longWeekdayNames = []string{
	Sunday,
	Monday,
	Tuesday,
	Wednesday,
	Thursday,
	Friday,
	Saturday,
}

var shortWeekdayNames = []string{
	SundayShort,
	MondayShort,
	TuesdayShort,
	WednesdayShort,
	ThursdayShort,
	FridayShort,
	SaturdayShort,
}
