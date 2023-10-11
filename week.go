package go2chinese

import "time"

func Weekday(w time.Weekday) string {
	return longWeekdayNames[w]
}

var longWeekdayNames = []string{
	"星期天",
	"星期一",
	"星期二",
	"星期三",
	"星期四",
	"星期五",
	"星期六",
}

var shortWeekdayNames = []string{
	"周日",
	"周一",
	"周二",
	"周三",
	"周四",
	"周五",
	"周六",
}
