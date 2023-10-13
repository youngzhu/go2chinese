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
