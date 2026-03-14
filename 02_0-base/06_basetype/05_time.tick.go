package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	// time.NewTimer(d Duration) *Timer 创建一个一次性定时器，返回*time.Timer类型：
	// 	2 秒后定时器会"到期"
	// 	到期后会向 timer.C通道发送当前时间
	// time.Timer 结构
	// type Timer struct {
	//     C <-chan Time  // 定时器到期时发送时间的通道
	//     // ... 其他字段
	// }

	// Example:
	// timerA()
	// timerB2()
	// timerC()

	// time.NewTick(d Duration) *Ticker创建一个周期性定时器，每隔指定时间向通道发送当前时间
	// type Ticker struct {
	//   C <-chan Time  // 周期性发送时间的通道
	// }

	// Example：
	// Ticker1()
	heartbeat()
	retryExample()
	dataPolling()
}

func timerA() {
	fmt.Println("start")
	// 1.超时控制
	timer3 := time.NewTimer(3 * time.Second) // 3等待3秒向timer.C通道发送消息
	// 如果将3秒改为2秒，NewTimer中等待2秒的时间快与goroutine的sleep2秒时间，程序退出时，goroutine 还没来得及输出消
	// 通过goroutine 模拟一个可能很慢的操作
	go func() {
		// 模拟耗时操作
		time.Sleep(2 * time.Second)
		fmt.Println("模拟操作操作任务完成")
	}()

	<-timer3.C // 阻塞等待接收timer3的通道消息
	fmt.Println("等待超时或任务完成")
	// 执行过程
	// 创建一个 3 秒的定时器
	// 启动 goroutine，2 秒后调用 输出：模拟操作操作任务完成
	// 定时器到期后向timer3.C 发送消息
	// timer3.C 接收到数据
	// 输出：等待超时或任务完成
}
func timerB2() {
	// 简单的延时
	fmt.Println("开始：", time.Now().Format("2006-01-02 15:04:05"))
	<-time.After(2 * time.Second) // 阻塞 等待2秒
	fmt.Println("2秒之后：", time.Now().Format("2006-01-02 15:04:05"))
}
func timerB() {
	// 可取消的延迟，time.After返回一个 <-chan Time 类型的通道，在指定时间（1秒）后，会自动向该通道发送当前时间。
	timer := time.NewTimer(2 * time.Second)

	select {
	case <-timer.C:
		fmt.Println("定时器到期") // 2秒后执行
	case <-time.After(1 * time.Second): // 在指定时间（1秒）后，会自动timer该通道发送当前时间。
		fmt.Println("1秒后执行其他逻辑")
		timer.Stop() // 取消原定时器
	}
	// 如果2秒内没有从 timer 收到结果，就会执行超时分支
	// 	执行流程：
	// 创建一个2秒的定时器
	// select 同时监听两个通道：
	// 		timer.C：2秒后收到信号
	// 			t, ok := <-ticker.C	t 是 time.Time，ok 是 bool
	// 		time.After(1 * time.Second)：1秒后收到信号
	// 哪个先到就执行哪个：1秒通道先触发
	// 执行 "1秒后执行其他逻辑"，并取消2秒定时器

	// 超时控制	防止操作无限期等待
	// 优先级选择	多个定时任务，先到先执行
	// 资源清理	超时后释放资源或取消操作
	//注意：time.After 会创建一个新定时器，如果频繁使用应考虑用 time.NewTimer + Reset 重用，避免内存泄漏
}

func timerC() {
	// 循环定时器
	timer := time.NewTimer(1 * time.Second)

	for i := 0; i < 3; i++ {
		<-timer.C
		fmt.Println("第", i+1, "次触发定时器")
		// NewTimer定时器只会执行一次，通过timer.Reset重制定时器
		timer.Reset(1 * time.Second) // 重置NewTimer时间，下次继续用
	}
	timer.Stop() // 避免内存泄漏
}

// 超时控制（最常见用法）

func selectEx() {
	// select 语句详解
	// select 是 Go 语言特有的控制结构，专门用于处理多个 channel 的发送/接收操作。
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1) // 创建string的
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "来自 channel 1"
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "来自 channel 2"
	}()

	select {
	case msg := <-ch1:
		fmt.Println("收到:", msg)
	case msg := <-ch2:
		fmt.Println("收到:", msg)
	}
	// 输出从ch1中的消息，因为ch1中goroutine的sleep为100毫秒，先唤醒，select捕获到ch1通道消息输出"来自 channel 1"
}

func Ticker1() {
	// 创建一个周期性定时器，每隔指定时间向通道发送当前时间，
	ticker := time.NewTicker(1 * time.Second) // 每隔1秒向通道发送一次时间；即每隔1秒，ticker 会自动往这个 channel 里塞入一个时间值。
	defer ticker.Stop()                       // defer 关键字作用，将调用函数推迟到当前函数的返回前执行，无论函数是正常返回还是 panic。
	// 可以理解ticker.stop函数会在Ticker1最后一行执行
	// 	func main() {
	//     fmt.Println("start")
	//     defer fmt.Println("deferred")
	//     fmt.Println("end")
	// }
	// 输出：
	// start
	// end
	// deferred

	// count := 0 // 停止周期定时，flag
	// // for range channel不停地从 channel 里取值，取一次执行一次，直到 channel 关闭。
	// for range ticker.C { // 阻塞等待，每1秒接收一次通道信息
	// 	count++
	// 	fmt.Println("第",count,"次消息")
	// 	if count >= 5 { // 当前count大于5时，停止for循环；停止向通过过去时间
	// 		break
	// 	}
	// }
	// go func ()  {
	// 	time.Sleep(5*time.Second) // 5秒后 关闭ticker通道
	// 	ticker.Stop()
	// }()
	// 等价于
	count2 := 0
	for {
		count2++
		// t, ok := <-ticker.C	t 是 time.Time，ok 是 bool
		t, ok := <-ticker.C // 阻塞等待 channel 时间
		if !ok{ // ok为false，channel 被关闭时退出
			break
		}
		fmt.Println("第", count2, "次消息, 返回时间：", t.Format("15:04:05"))
		if count2 >= 5 {
			break
		}
	}
	// ticker.C 永远不会自动关闭（除非调用 ticker.Stop()），所以 ok 值始终为 true（只要 ticker 还在运行）

	// 或者 直接使用for循环方式，比上述写法都要简洁
	for i:=0; i< 5;i++{
		t := <-ticker.C // 阻塞等待，每1秒接收一次通道信息
		fmt.Println("第", i, "次消息, 返回时间：", t.Format("15:04:05"))
	}

	// ticker2 := time.NewTicker(1 * time.Second)
	// defer ticker2.Stop() // 必须手动 Stop，否则 goroutine 泄漏！

	// for t := range ticker2.C {
	// 		fmt.Println("tick:", t.Format("15:04:05"))
	// }

}

// 1.心跳检测
func heartbeat() {
	ticker := time.NewTicker(2*time.Second)// 每2秒发送一次
	defer ticker.Stop()

	// 创建一个channel通道，无缓冲区的通道
	done := make(chan bool)

	// 模拟10秒后结束停止心跳
	go func ()  {
		time.Sleep(10*time.Second)
		done<-true // 10秒后向缓冲区通道发送true消息
	}()

	beatCount :=0
	for {
		beatCount++
		select {
		case t:=<-ticker.C:
			fmt.Printf("第%d次心跳检测：%s\n", beatCount, t.Format("15:04:05"))
		case <-done:
			fmt.Println("心跳检测停止...")
			// close(done)
			return // 直接返回，停止for循环
		}
	}
}

// 定时任务
func scheduledTask()  {
	ticker := time.NewTicker(2*time.Second)// 每2秒发送一次
	defer ticker.Stop()

	taskCount := 0
	for t := range ticker.C { // 持续从 channel 读取
		taskCount++
		fmt.Printf("执行第 %d 次任务：%s\n", taskCount, t.Format("15:04:05"))
    // 执行定时任务逻辑... ...
		if taskCount >= 5{
			return
		}
	}
}

// 
func retryExample() {
    ticker := time.NewTicker(2 * time.Second)
    defer ticker.Stop()

    for i := 0; i < 3; i++ {
        <-ticker.C
        fmt.Printf("第%d 次重试\n", i+1)
        
        if i == 2 {
            fmt.Println("重试成功")
            return
        }
        fmt.Println("重试失败")
    }
}

// 数据轮询
func dataPolling() error {
	// 循环定时器，每1秒钟循环执行一次
	ticker := time.NewTicker(1  *time.Second)
	defer ticker.Stop() // 以免内存泄漏
	// 创建一次性定时器，超过6秒
	timeout := time.NewTimer(6*time.Second)
	defer timeout.Stop()

	pollCount := 0
	for {
		select {
		case t := <-ticker.C: // 阻塞等待，每1秒接收一次通道信息
			pollCount++
			fmt.Printf("第%d次轮询：%s\n",pollCount, t.Format("15:04:05"))
			if pollCount  == 3{
				fmt.Println("轮询找到数据")
				return nil
			}
		case <-timeout.C: // 超出6秒：轮询超时
			fmt.Errorf("轮询超时")
			return errors.New("轮询超时")
		}
	}
}