package go2chinese_test

import (
	"fmt"
	gc "github.com/youngzhu/go2chinese"
	"time"
)

func ExampleWeekday() {
	gc.Weekday(time.Now())

	weekday := gc.Weekday(time.Date(2023, 10, 13, 0, 0, 0, 0, time.Local))
	fmt.Println(weekday)

	// Output:
	// 星期五
}

func ExampleShortWeekday() {
	gc.Weekday(time.Now())

	weekday := gc.ShortWeekday(time.Date(2023, 10, 12, 0, 0, 0, 0, time.Local))
	fmt.Println(weekday)

	// Output:
	// 周四
}

func ExampleDate() {
	gc.Date(time.Now())

	date := gc.Date(time.Date(2023, 10, 16, 0, 0, 0, 0, time.Local))
	fmt.Println(date)

	// Output:
	// 2023年10月16日
}

func ExampleDatetime() {
	gc.Date(time.Now())

	date := gc.Datetime(time.Date(2023, 8, 8, 8, 8, 8, 0, time.Local))
	fmt.Println(date)

	// Output:
	// 2023年8月8日 8时8分8秒
}
