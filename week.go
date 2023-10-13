package go2chinese

import "time"

const (
	// Sunday ... long date names
	Sunday    = "星期天"
	Monday    = "星期一"
	Tuesday   = "星期二"
	Wednesday = "星期三"
	Thursday  = "星期四"
	Friday    = "星期五"
	Saturday  = "星期六"

	// SundayShort ... short date names
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
