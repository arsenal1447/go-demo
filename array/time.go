package main

import (
	"fmt"
	"time"
)

func main() {
	// if len(os.Args) > 0 {
	// 	for index, arg := range os.Args {
	// 		fmt.Printf("args[%d]=%v\n", index, arg)
	// 	}
	// }
	timeDemo()
}

func timeDemo() {
	now := time.Now()
	fmt.Printf("current time :%v\n", now) //获取当前时间

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}
