package main

import "fmt"

func main()  {
	// panic/recover
	// go在每没有try/catch异常处理机制，go设计哲学把异常错误当作普通的value值处理（Errors is values）
	// 一般在go中通常使用if err != nil return 进行错误处理，但对于那些严重影响程序继续运行的“意外事故”或“运行时致命错误”，Go 提供了内置机制：panic（引发恐慌）与 recover（捕获恐慌/恢复）。

	// 一、触发时机
	// 隐士触发：如空指针解引用（nil pointer dereference）、切片越界访问、被 0 整除等
	// 显式触发：开发者主动调用 panic("严重错误信息")。

	// 隐士触发
	// var slices []int
	// slices[0] = 1 // panic: runtime error: index out of range [0] with length 0

	// num := 10 / 0 // 整数除零(浮点除零不 panic,返回 Inf)

	// var msg interface{} = "hello"
	// num := msg.(int) // 断言失败 panic: interface conversion: interface {} is string, not int
	// fmt.Println(num)

	// 手动触发
	// panic("触发恐慌了") // 直接终止运行程序 panic: 触发恐慌了
	// panic(fmt.Errorf("bad is %d", 404))

	// 二、执行顺序
	// base()
	// 三、发生panic时，一定会执行 defer
	panicDef()

	// 三、panic执行链
	defer fmt.Println("defer in main")
	panicA()
}

// 基本panic例子
func base() {
	fmt.Println("start")
	panic("error 500")
	fmt.Println("end") // 后面的panic
	/* start
	panic: error 500

	goroutine 1 [running]:
	main.base() */
}

// panic发生异常奔溃，当前goroutine 函数中已经注册的defer语句，会执行 defer，顺序是 LIFO，后进先出。
func panicDef()  {
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")

	fmt.Println("panic defer")
	panic("error panic")
}

// panic执行链
func panicF()  {
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	panic("发生恐慌错误！")
}
/*
main() -> 
base() -> 
遇到panic() -> 
停止当前函数后续代码 -> 
执行defer修饰语句。顺序是 LIFO，后进先出。->
向调用栈上一层返回 ->
上一层继续执行自己的 defer ->
一直向上展开 ->
没有 recover ->
程序崩溃 exit status 2
*/
// 调用 panic 后：
// 当前函数停止正常执行，panic 之后的代码不会执行。
// 开始执行当前 goroutine 中已经注册的 defer，顺序是 LIFO，后进先出。
// 如果某个 defer 中调用了 recover 并成功捕获，panic 停止，当前函数返回，调用者继续执行。
// 如果一直没有被 recover，panic 会继续向上层调用栈传播。
// 最终传播到 goroutine 顶层仍未被捕获，整个程序崩溃，打印 panic 信息和堆栈，通常退出码为 2。

// panic 的传播机制
func panicA()  {
	defer fmt.Println("defer in A")
	panicB()
	fmt.Println("continues A") // 下层函数发生panic，不会执行
}
func panicB()  {
	defer fmt.Println("defer in B")
	panicC()
	fmt.Println("continues B") // 不会执行
}
func panicC()  {
	defer fmt.Println("defer in C")
	panic("bad panic in C")
}
/*
执行顺序
main() 
  ↓
a()
  ↓
b()
  ↓
c()
  ↓
panic()

进入panic后，向上返回调用链：
c() 停止，执行当前defer，输出defer in C
  ↓
b() 停止，执行当前defer，输出defer in B
  ↓
a() 停止，执行当前defer，输出defer in A
  ↓
main() 停止，执行当前defer，输出defer in main
  ↓
程序崩溃，输出 bad panic in C

1.panic 传播路径上，函数体 panic 之后的代码、调用点之后的代码全部跳过。只执行 panic直接注册的 defer，遵循后进先出 LIFO
2.只有 panic 发生之前已经注册的 defer 才会执行
3.如果中间defer中没有 recover()函数捕获，程序最终会退出

*/