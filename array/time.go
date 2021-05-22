package main

import (
	"fmt"
	"time"
)

// time.Time类型表示时间。
// 我们可以通过time.Now()函数获取当前的时间对象，然后获取时间对象的年月日时分秒等信息。示例代码如下：

// func timeDemo() {
// 	now := time.Now()
// 	fmt.Printf("current time :%v\n", now) //获取当前时间

// 	year := now.Year()
// 	month := now.Month()
// 	day := now.Day()
// 	hour := now.Hour()
// 	minute := now.Minute()
// 	second := now.Second()
// 	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
// }

// func timeStampDemo() {
// 	now := time.Now()
// 	timestamp1 := now.Unix()     //时间戳
// 	timestamp2 := now.UnixNano() //纳秒时间戳
// 	fmt.Printf("current timestamp1:%v\n", timestamp1)
// 	fmt.Printf("current timestamp2:%v\n", timestamp2)

// }

// // 使用time.Unix()函数可以将时间戳转为时间格式。
// func timestampDemo2(timestamp int64) {
// 	timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式
// 	fmt.Println(timeObj)

// 	year := timeObj.Year()
// 	month := timeObj.Month()
// 	day := timeObj.Day()
// 	hour := timeObj.Hour()
// 	minute := timeObj.Minute()
// 	second := timeObj.Second()
// 	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
// }

// time.Duration是time包定义的一个类型，它代表两个时间点之间经过的时间，以纳秒为单位。
// time.Duration表示一段时间间隔，可表示的最长时间段大约290年。

// time包中定义的时间间隔类型的常量如下：

// const (
// 	Nanosecond  Duration = 1
// 	Microsecond          = 1000 * Nanosecond
// 	Millsecond           = 1000 * Microsecond
// 	Second               = 1000 * Millsecond
// 	Minute               = 60 * Second
// 	Hour                 = 60 * Minute
// )

//例如：time.Duration表示1纳秒，time.Second表示1秒。

func main() {
	// if len(os.Args) > 0 {
	// 	for index, arg := range os.Args {
	// 		fmt.Printf("args[%d]=%v\n", index, arg)
	// 	}
	// }
	// timeDemo()
	// timeStampDemo()
	// timestampDemo2(1621587331)

	now := time.Now()
	later := now.Add(time.Hour) // 当前时间加1小时后的时间
	fmt.Println(later)          //2021-05-21 17:36:37  + 1h = 2021-05-21 18:36:37.455897 +0800 CST m=+3600.000088942

	// tickDemo()
	formatDemo()

}

// Sub
// 求两个时间之间的差值：

// func (t Time) Sub(u Time) Duration
// 返回一个时间段t-u。如果结果超出了Duration可以表示的最大值/最小值，将返回最大值/最小值。要获取时间点t-d（d为Duration），可以使用t.Add(-d)。

// Equal
// func (t Time) Equal(u Time) bool
// 判断两个时间是否相同，会考虑时区的影响，因此不同时区标准的时间也可以正确比较。本方法和用t==u不同，这种方法还会比较地点和时区信息。

// Before
// func (t Time) Before(u Time) bool
// 如果t代表的时间点在u之前，返回真；否则返回假。

// After
// func (t Time) After(u Time) bool
// 如果t代表的时间点在u之后，返回真；否则返回假。

// 使用time.Tick(时间间隔)来设置定时器，定时器的本质上是一个通道（channel）。
func tickDemo() {
	ticker := time.Tick(time.Second) //定义一个1秒间隔的定时器
	for i := range ticker {
		fmt.Println(i) //每秒都会执行的任务
	}
}

//时间类型有一个自带的方法Format进行格式化，
// 需要注意的是Go语言中格式化时间模板不是常见的Y-m-d H:M:S

// 而是使用Go的诞生时间2006年1月2号15点04分（记忆口诀为2006 1 2 3 4）。

// 也许这就是技术人员的浪漫吧。

//补充：如果想格式化为12小时方式，需指定PM

func formatDemo() {
	now := time.Now()

	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan")) //2021-05-21 17:44:19.686 Fri May
	// 12小时制
	fmt.Printf(now.Format("2006-01-02 03:04:05.000 PM Mon Jan")) //2021-05-21 05:44:19.686 PM Fri May
	fmt.Println(now.Format("2006/01/02 15:04"))                  //2021-05-21 05:44:57.177 PM Fri May2021/05/21 17:44
	fmt.Println(now.Format("15:04 2006/01/02"))                  //17:44 2021/05/21
	fmt.Println(now.Format("2006/01/02"))                        //2021/05/21
	fmt.Println(now.Format("2006/01/02 03:04:05"))               //2021/05/22 10:36:02

}
