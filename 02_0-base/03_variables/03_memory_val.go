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
const MaxRetry = 5 // 数值型：编译期直接内联到指令，没有运行时地址
const Greet   = "hello"     // 字符串：字节数据存 .rodata，运行时不可写


// 2.包级 全局变量
var count int      // 未初始化 → .bss段：只占位，程序加载时清零，不增大二进制大小
var timeout = 30   // 已初始化 → .data 段：初值 30 存入二进制，程序启动时直接可用

var name = "go" // 已初始化 → .data 段：string header（ptr+len）在 .data，字节 "go" 本身在 .rodata

func main() {

	const A = 100 // 存在只读数据区（.rodata）
	// fmt.Println(&A) // 常量无法取地址值，编译错误

	fmt.Println(&count)   // 0x5a5d18 可以取地址，与 const 不同
  fmt.Println(&timeout) // 0x57b468 可以取地址


	// 局部变量: 栈
	localVar := 5
	fmt.Printf("局部变量:%p\n", &localVar) // 局部变量:0x19db3d8e0130
}
