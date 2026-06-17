package main

import "fmt"

// +------------------+
// |   Code Segment   | 代码区(.text)
// +------------------+
// |  Read Only Data  | 只读数据段 常量区(.rodata)
// +------------------+
// |   Data Segment   | 数据段 已初始化全局变量(.data)
// +------------------+
// |    BSS Segment   | 未初始化全局变量(.bss)
// +------------------+
// |      Heap        | 堆 通过 make、new 或逃逸分析触发的动态分配都在这里
// |                  |
// +------------------+
// |      Stack       | Goroutine栈 每个 goroutine 独立拥有一个栈，初始大小只有 2-8 KB，按需增长。
// +------------------+

// 1.常量
const MaxRetry = 5    // 数值型：编译期直接内联到指令，没有运行时地址
const Greet = "hello" // 字符串：字节数据存 .rodata，运行时不可写

// 2.包级 全局变量
var count int    // 未初始化 → .bss段：只占位，程序加载时清零，不增大二进制大小
var timeout = 30 // 已初始化 → .data 段：初值 30 存入二进制，程序启动时直接可用

var name = "go" // 已初始化 → .data 段：string header（ptr+len）在 .data，字节 "go" 本身在 .rodata

// go build -gcflags='-m' 03_memory_val.go  生成-> 03_memory_val文件
// go tool nm 03_memory_val|grep -iE "count|timeout
//  5a5d18 D main.count ← D = data 段
// 	57b468 D main.timeout ← D = data 段
// 字母前缀 D(data)、B(bss)、R(rodata)、T(text) 直接告诉你符号在哪个区。

// 3.变量
// x变量未逃逸，留在栈上
func foo() int { // 函数体在 .text，foo(){}编译的机器指令
	x := 42 // 局部变量存放站，栈上
	fmt.Println(x)
	return x
}

// x逃逸：移到堆
func bar() *int {
	x := 42
	return &x // &x 被返回 → 地址逃出了函数作用域
	// 编译器判定：必须分配在堆上，GC 负责回收
}
// 用 -gcflags="-m" 让编译器把逃逸决策打印出来：
// go build -gcflags='-m' 03_memory_val.go
// # command-line-arguments
// ./03_memory_val.go:39:6: can inline bar
// ./03_memory_val.go:59:9: inlining call to bar
// ./03_memory_val.go:34:13: ... argument does not escape
// ./03_memory_val.go:34:14: 42 escapes to heap
// ./03_memory_val.go:40:2: moved to heap: x    # 这里出现x逃逸到堆上


func main() {

	const A = 100 // 存在只读数据区（.rodata）
	// fmt.Println(&A) // 常量无法取地址值，编译错误

	fmt.Println(&count)   // 0x5a5d18 可以取地址，与 const 不同
	fmt.Println(&timeout) // 0x57b468 可以取地址

	// 局部变量: 栈
	localVar := 5
	fmt.Printf("局部变量:%p\n", &localVar) // 局部变量:0x19db3d8e0130

	_ = foo()
	_ = bar()

	// 4. make / new — 堆分配，以及 string 的双区域
	// make，创建切片
	s := make([]int, 3)
	// 栈上：slice header = {ptr=0xc000018060, len=3, cap=3}（24字节）ptr指向底层数组
	// 堆上：实际数组 [0, 0, 0]（3×8=24字节），GC 管理

	// new：对象在堆
	p := new(int)
	// 栈上：指针变量 p（8字节）
	// 堆上：int 对象（8字节），GC 管理

	// string 字面量：头部在栈，字节在 .rodata
	msg := "hello"
	// 栈上：string header = {ptr=0x012b（指向.rodata）, len=5}（16字节）
	// .rodata：77 6f 72 6c 64（只读，运行时不可改）存放字符
	msg = "world" // 合法：修改的是栈上的 header，ptr换了新的指向
	// msg[0] = 'H' ❌ 编译报错：cannot assign to msg[0]
  // 底层字节在 .rodata，写保护，不可修改

	_ = s; _ = p; _ = msg

}
