package main

import "fmt"

// 1.全局变量是定义在函数外部的变量，它在程序整个运行周期内都有效。 在函数中可以访问到全局变量。
var num int64 = 10

func GlobalVal() {
	// 函数中访问全局变量
	fmt.Printf("num:%d\n", num)
}
func LocalVal()  {
	// 2.局部变量x，仅在该函数内生效，外部无法直接访问该变量
	var x int = 100
	fmt.Printf("num:%d\n", x)
}

// 如果局部变量和全局变量重名，优先访问局部变量。
var total int64 = 20 // 全局
func testLocal()  {
	total := 100
	fmt.Printf("total:%d\n", total) // 函数中优先使用局部变量
}

func main() {
	GlobalVal() // num:10
	LocalVal() 
	// fmt.Println(x) 无法使用局部变量x

	// 函数中优先使用局部变量
	testLocal() // total:100
}
