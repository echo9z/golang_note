package main

import (
	"fmt"
	"io/fs"
	"math/rand"
	"os"
)

func main()  {
	// 三.label标签通常与goto，break，continue进行使用。用于简化流程，if、for、switch 或 select 语句中，进行
	// label配合goto实现跳转
	num := 0
Loop:
	if num < 5 {
		fmt.Println(num)
		num++
		goto Loop // 跳转到Loop标签，模拟循环效果
	}

	// 3.1错误的跳过声明变量
	// var aa int = 15
	// goto goLabel
	// // 不能跳过变量声明语句。如果跳转到变量声明之前，而该变量在跳转之后又被使用，编译器会报错。
	// cc := 20 // goto与label之间的内容都会被跳过，
	// goLabel:
	// 	fmt.Println(aa+cc)

	// 3.2标签和跳转必须在 同一个函数 内部。也不能跳转到 for/switch 的内部
	// sum := func (a int)  {
	// 	var str string = "hello go"
	// 	fmt.Println(str)
	// 	goto labelFn // goto只能在同一个函数内跳转，
	// }

	// sum1 := func (b int)  {
	// 	labelFn:
	// 		fmt.Println(b)
	// }

	// 3.3label与break配合使用
	// 嵌套循环中，break 默认只退出最内层循环。配合 label 可以直接退出外层循环。
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break // break只会停止i=1,j=1--i=1,j=2，只会停止最内层
			}
			fmt.Printf("(%d,%d)\t", i, j)
		}
	}
	fmt.Printf("\n")
	// 立即终止标签所关联的 for、switch 或 select 语句块的执行
BreakF:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break BreakF // 直接跳到最外层，如果不加label。break只会停止i=1,j=1--i=1,j=2，只会停止最内层
			}
			fmt.Printf("(%d,%d)\t", i, j)
		}
	}
	fmt.Println("")
	
	// 3.4continue与Label
ContinueF:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 1 { // 内部循环j等于1时，结束本次循环，跳转至标签ContinueF处，回到外层 i 自增。
				continue ContinueF
			}
			fmt.Printf("(%d,%d)\t", i, j) // 输出：(0,0) (1,0) (2,0)
		}
	}
	fmt.Println("")

	// 3.5 break+label 用于 switch / select
	// 在 select 或 switch 内部，若想跳出外层的 for 循环，也必须使用带标签的 break。
	var breakSelect func() = func() {
		ch := make(chan int) // 创建一个int整型通道
		// chan struct{} 就是一个不传任何数据的通道
		done := make(chan struct{}) // struct{} 是 Go 中的空结构体，占用 0 字节内存。

		// 开启一个goroutine 协程
		go func() {
			ch <- 25 // 向ch通道发生25
			// close(done) 就是用关闭来发信号，和发送 struct{}{} 效果一样——接收方都能收到通知。
			close(done) // 关闭done通道，发出完成信号 或发送空实例 done <- struct{}{}
		}()
	Loop:
		for {
			select {
			case val := <-ch: // 接收25数据，输出
				fmt.Printf("val: %d\n", val)
			case <-done: // 收到完成信号，跳出整个循环
				break Loop
			}
		}
		fmt.Println("Exited")
	}
	breakSelect()

	// 生成[0.0到1.0)的随机浮点数
	flag := rand.Float64()
	if flag >= 0.5 {
		goto A
		// goto 之后，永远不会执行，但编译器仍会检查 
		// float64(a) a是无法调用，a 的作用域仅在这个 {} 块内
		// var val float64 = float64(a) + flag
		// fmt.Printf("flag:%f, val:%f", a, val)
	} else {
		goto B
		// var val float64 = float64(b) + flag
		// fmt.Printf("flag:%f, val:%f", b, val)
	}

A:
	{
		a := 10
		fmt.Println("a", a)
		
		var val float64 = float64(a) + flag
		fmt.Printf("flag:%f, val:%f", a, val)
	}
B:
	{
		b := 20
		fmt.Println("b", b)
	}

	// 统一错误处理/资源清理
	readFile := func(path string) ([]byte, error) {
		var (
			// 这样声明处理golsp出现的跳过了声明的变量
			f    *os.File
			err  error
			info fs.FileInfo
			buf  []byte
		)
		f, err = os.Open(path)
		if err != nil {
			// 如果打开指定文件路径出现错误，err不空值，通过goto跳转至统一处理错误
			goto ErrorL
		}

		info, err = f.Stat() // 获取文件信息
		if err != nil {
			goto ErrorL
		}

		buf = make([]byte, info.Size()) //
		_, err = f.Read(buf)            // 将文件内容读取到字节切片中
		if err != nil {
			goto ErrorL
		}

		f.Close()
		return buf, nil
		// 统一处理错误
	ErrorL:
		f.Close() // 关闭文件资源
		return nil, err
	}
	by, _ := readFile("02_0-base/06_basetype/00_format.go")
	fmt.Println("读取文件", string(by))
	// 用 defer 替代 goto，这是 Go 惯用写法
	// readFile := func(path string) ([]byte, error) {
	//   f, err := os.Open(path)
	//   if err != nil {
	//       return nil, err
	//   }
	//   defer f.Close() // 函数返回时自动关闭，无需 goto

	//   info, err := f.Stat()
	//   if err != nil {
	//       return nil, err
	//   }

	//   buf := make([]byte, info.Size())
	//   _, err = f.Read(buf)
	//   if err != nil {
	//       return nil, err
	//   }

	//   return buf, nil
	// }
}
