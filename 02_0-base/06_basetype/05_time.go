package main

import (
	"fmt"
	"time"
)

func main() {
	// 1.获取当前时间
	now := time.Now()
	fmt.Println("当前时间", now) // 当前时间 2026-02-09 13:35:43.6066288 +0800 CST m=+0.000000001

	// 2.获取年月日，时分秒
	year := now.Year()
	month := now.Month() // 获取月。类型time.Month
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	// nanosecond := now.Nanosecond() // ns（nanosecond）：纳秒，时间单位。一秒的十亿分之一，等于10的负9次方秒（1 ns = 10-9 s）。
	weekday := now.Weekday() // 获取星期

	fmt.Printf("%d年%d月%d日 %d时%d分%d秒 星期：%s\n",
		year, month, day, hour, minute, second, weekday) // 2026年2月9日 15时53分50秒 星期：Monday

	// 3.创建指定时间
	// time.Data(year, month, day,hour,min,sec,nsec,loc)
	// year:int类型 年份
	// month：指定类型time.Month类型，比如time.January、time.December
	// day:int类型 日（1-31）
	// hour:int类型 小时（0-23）
	// min:int分钟（0-59）
	// sec:int秒（0-59）
	// nsec:纳秒（0-999999999）
	// loc:时区 time.Local、time.UTC
	specificTime := time.Date(2024, time.December, 14, 15, 10, 0, 0, time.Local)
	fmt.Println("创建自定义时间：", specificTime) // 创建自定义时间： 2024-12-14 15:10:00 +0800 CST

	// 4.通过时间戳创建
	timestamp := int64(1696735800)
	tStamp := time.Unix(timestamp, 0) // 第二个参数为纳秒
	fmt.Println("时间戳创建：", tStamp)     // 时间戳创建： 2023-10-08 11:30:00 +0800 CST

	// 5.时间格式化
	// Format 返回根据参数定义的布局格式化时间值的文本表示。
	// 格式化为字符串（必须使用Go的特定参考时间）
	// 标准 ISO 8601 格式
	fmt.Println(now.Format("2006-01-02T15:04:05Z07:00")) // 2026-02-09T17:29:33+08:00
	fmt.Println(now.Format("2006-01-02 15:04:05"))       // 2026-02-09 17:30:18
	// 注意：如果写出这种格式 时间必须是2006年01月02日 15:04:05这个时间，Go 的时间格式化必须使用固定的参考时间：Mon Jan 2 15:04:05 MST 2006

	// 简短的时间
	fmt.Println(now.Format("2006/01/02")) // 他会参考这个传入的格式，输入出当前时间2026/02/09
	// 只显示时间  这里传入的时间必须是 15:04:05
	fmt.Println(now.Format("15:04:05")) // 输入当前的时间17:41:02

	// 带星期的
	fmt.Println(now.Format("2006 01 02 Monday")) // 2026 02 09 Monday

	// 其他语言的常见方式（在Go中不支持！）
	// fmt.Println(t.Format("YYYY-MM-DD"))  ❌ 错误！
	// fmt.Println(t.Format("hh:mm:ss"))    ❌ 错误！

	// go中预定义常量
	// 标准库已提供常用格式常量：
	// 	time.RFC3339      // "2006-01-02T15:04:05Z07:00"
	// 	time.RFC1123      // "Mon, 02 Jan 2006 15:04:05 MST"
	// 	time.RFC822       // "02 Jan 06 15:04 MST"
	// 	time.Kitchen      // "3:04PM"
	// 	time.Stamp        // "Jan _2 15:04:05"
	// 	time.StampMilli   // "Jan _2 15:04:05.000"
	// 	time.ANSIC        // "Mon Jan _2 15:04:05 2006"
	t2 := time.Date(2025, 12, 25, 14, 15, 45, 0, time.UTC)
	fmt.Println("RFC3339:", t2.Format(time.RFC3339))   // RFC3339: 2025-12-25T14:15:45Z
	fmt.Println("ANSIC:", t2.Format(time.ANSIC))       // ANSIC: Thu Dec 25 14:15:45 2025
	fmt.Println("UnixDate:", t2.Format(time.UnixDate)) // UnixDate: Thu Dec 25 14:15:45 UTC 2025

	// 6.解析时间字符串
	layout := "2006-01-02 15:04:05" // 模板
	timeStr := "2025-12-12 15:10:59"
	// 使用time.parse
	// 第一个参数 Layout 日期时间必须是：Mon Jan 2 15:04:05 MST 2006 这个世界点的任一格式模板
	// 例如：
	// 第二个参数必须使用第一个参数提供的格式字符串（布局）进行解析。
	// 返回为time类型
	parsedTime, err := time.Parse(layout, timeStr)
	if err != nil {
		fmt.Println("解析错误：", err)
	} else {
		fmt.Println("解析结果：", parsedTime)
	}

	// 带时区的格式
	lot := "Jan 2, 2006 at 3:04pm (MST)"
	t3, _ := time.Parse(lot, "Feb 3, 2025 at 7:54pm (UTC)") // 解析的时间结构 要与lot结构一致
	fmt.Println("解析结果t3：", t3)                              //  解析结果t3： 2025-02-03 19:54:00 +0000 UTC

	// 7.时间计算
	now2 := time.Now()
	// 增加时间
	fmt.Println(time.Hour) // 一小时 1h0m0s
	oneHourLater := now2.Add(time.Hour)
	fmt.Println("一小时后：", oneHourLater) //一小时后： 2026-02-10 23:16:07.0540868 +0800 CST m=+3600.018266301

	// 增加指定时常
	twoDaysLater := now2.Add(time.Hour * 48)
	fmt.Println("两天后：", twoDaysLater) // 两天后： 2026-02-12 22:21:20.4059784 +0800 CST m=+172800.019016601

	// 增加指定的日期
	// AddDate(year,month,day)
	oneMonthLater := now2.AddDate(0, 1, 0) // 年, 月, 日
	fmt.Println("一个月后：", oneMonthLater)    // 一个月后： 2026-03-10 22:41:40.4558462 +0800 CST

	// 减少时间
	oneHourBefore := now.Add(-time.Hour)
	fmt.Println("减一小时", oneHourBefore) // 减一小时 2026-02-11 10:35:28.8194763 +0800 CST m=-3599.999461599

	// 时间差
	// 创建一个时间 2024/12/14 15：10：00
	start := time.Now()
	fmt.Println("开始时间：", start) // 开始时间： 2026-02-11 17:28:17.2852332 +0800 CST m=+0.019361901
	fmt.Println(time.Second)    // 1秒
	// Sleep(d time.Duration)
	// Duration 是 int64 的别名，表示两个时间点之间经过的纳秒数：
	fmt.Println(500 * time.Millisecond) // 500 * 1毫秒
	fmt.Println(500 * time.Millisecond) // 500 * 1毫秒
	fmt.Println(time.Minute)            // 1分钟

	time.Sleep(2 * time.Second) // 暂停当前 goroutine 执行指定的时长，然后恢复执行。即time.Sleep睡2秒后，执行后面代码
	// time.Since(t Time) Duration：计算从时间 t 到现在经过的时长
	duration := time.Since(start)
	fmt.Println("耗时：", duration)           // 耗时： 2.0006042s
	fmt.Println("秒数：", duration.Seconds()) //  2.0006042
	fmt.Println("毫秒数：", duration.Milliseconds()) // 毫秒数： 2000
	// Duration 常用方法：
	// Seconds()	float64	转换为秒
	// Milliseconds()	float64	转换为毫秒
	// Microseconds()	int64	转换为微秒
	// Nanoseconds()	int64	转换为纳秒
	// Minutes()	float64	转换为分钟
	// Hours()	float64	转换为小时
	// String()	string	格式化字符串

	// 	time.Until()
	// 与 Since 相反，计算从现在到未来某个时间点的时长：
	future := time.Now().Add(2 * time.Hour)
	durationF := time.Until(future)
	fmt.Println("距离未来时间还有：", durationF) // 距离未来时间还有： 2h0m0s

	// 两个时间的差值
	t11 := time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)
	t21 := time.Date(2023, 10, 8, 0, 0, 0, 0, time.UTC)
	diffTime := t21.Sub(t11)
	fmt.Println("时间差", diffTime) // 时间差 168h0m0s
	fmt.Println("天数差", diffTime.Hours()/24) // 天数差 7

	// 时间比较
	tt1 := time.Now()
	tt2 := tt1.Add(time.Hour)
	fmt.Println("tt1在tt2之前", tt1.Before(tt2)) // tt1在tt2之前 true
	fmt.Println("tt1在tt2之后", tt1.After(tt2)) // tt1在tt2之后 false
	fmt.Println("tt1与tt2相等", tt1.Equal(tt2)) // tt1与tt2相等 false
	
	// 检查是否在某个时间范围内
	startTime := time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(2023, 10, 31, 23, 59, 59, 0, time.UTC)
	checkTime := time.Date(2023, 10, 15, 12, 0, 0, 0, time.UTC)

	if checkTime.After(startTime) && checkTime.Before(endTime) {
		fmt.Println("在startTime时间之后，在endTime时间之前")
	}


	newTime := func ()  {
		// 定时器 - 只执行一次
		// time.NewTimer(d Duration) 创建一个定时器，2秒后定时器会"到期"；到期后会向 timer.C 通道发送当前时间
		timer := time.NewTimer(2 * time.Second)
		// timer.C 是一个 <-chan time.Time 类型的通道。
		<-timer.C // <-timer.C 表示从通道接收数据；在定时器到期前，这行代码会阻塞，程序停在这里等待
		// 2 秒后，定时器向 timer.C 发送时间，阻塞解除，继续执行
		fmt.Println("定时器触发了")
	}
	newTime()

	timer2 := time.NewTimer(5*time.Second) // 5秒中后会到期，到期后向timer.C通道发生消息

	// 创建一个协程，在这个goroutine中取消定时器
	go func ()  {
		time.Sleep(1*time.Second) // 先sleep睡一秒
		timer2.Stop() // 取消定时器
	}()

	select { // select语法是专门用于处理多个channel的发送/接收
	case <-timer2.C: // 如果定时器没有取消，则timer.C通道接收消息
		fmt.Println("timer2定时器被触发了")
	default:
		fmt.Println("timer2定时器被取消了")
	}
	// select {
	// case msg := <-ch1:
	// 		// 从 ch1 接收数据
	// case ch2 <- value:
	// 		// 向 ch2 发送数据
	// case <-time.After(1 * time.Second):
	// 		// 超时处理
	// default:
	// 		// 没有通道就绪时执行
	// }

	// 时区处理
	local := time.Local
	fmt.Println("本地时区", local)

	// 获取UTC时区
	utc := time.UTC
	fmt.Println("UTC时区:", utc)

	// 加载特定时区
	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err) // 使用 panic - 程序无法继续运行，直接抛出异常
	}
	// 转换时区
	now3 := time.Now()
	nyTime := now3.In(location)
	fmt.Println("纽约时间:", nyTime) // 纽约时间: 2026-02-12 05:52:32.3476491 -0500 EST

	// 时间戳操作
	now4 := time.Now()
	// 获取时间戳（秒）
	timestamp2 := now4.Unix()
	fmt.Println("时间戳（秒）", timestamp2)
	// 获取毫秒时间戳
	millis := now4.UnixMilli()
	fmt.Println("时间戳(毫秒):", millis)

	// 获取微秒时间戳
	micros := now4.UnixMicro()
	fmt.Println("时间戳(微秒):", micros)

	// 获取纳秒时间戳
	nanos := now4.UnixNano()
	fmt.Println("时间戳(纳秒):", nanos)

	// 从时间戳恢复时间
	t := time.Unix(timestamp, 0)
	fmt.Println("从时间戳恢复:", t)
}

// package main

// import (
//     "fmt"
//     "time"
// )

// func main() {
//     // 启动一个goroutine
//     go func() {
//         for i := 0; i < 5; i++ {
//             fmt.Println("Goroutine 1:", i)
//             time.Sleep(100 * time.Millisecond)
//         }
//     }()

//     // 启动另一个goroutine
//     go func() {
//         for i := 0; i < 5; i++ {
//             fmt.Println("Goroutine 2:", i)
//             time.Sleep(100 * time.Millisecond)
//         }
//     }()

//     // 主goroutine等待一段时间，让其他goroutine完成
//     time.Sleep(1 * time.Second)
//     fmt.Println("Main goroutine结束")
// }
