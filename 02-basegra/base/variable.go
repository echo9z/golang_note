package base

import "fmt"

//ab := 12

func Variable() {
	var a = 5 // 声明一个 a 变量
	a += 1    // a + 1 = 6
	a /= 2    // a / 2 = 3
	a &^= 2   // a &^ 2 = 1
	fmt.Println("a", a)
	var c int // 声明一个变量，默认值为0
	fmt.Println(c)

	b := 20 // 声明并初始化，且自动推导类型，只能在函数内部使用，var定义全局变量
	fmt.Println(b)
}

// VariableStat 变量声明
func VariableStat() {
	// 多变量声明
	var a1, b1 string              // 声明a1，b1类型为string
	var c1, c2 string = "c1", "c2" // 声明变量并赋值
	fmt.Println(a1, b1)
	fmt.Println(c1, c2)

	// 也可以这样声明多个变量
	var (
		e int
		f bool
	)
	fmt.Println("e f", e, f) // 0 false

	// := 自动推导声明的类型
	d1, d2 := "d1", "d2" // 根据值类型 自动推导变量类型
	fmt.Println(d1, d2)
	m3, n3, q3 := 10, "n2", 30 // 自动推导类型，并初始化值
	fmt.Println(m3, n3, q3)
}

func IsIota() [10]int {
	// 关键字iota声明初始值为0，每行递增1
	const (
		a = iota // 0
		b = iota // 1
		c = iota // 2
	)

	// 单独声明iota，在同一常量中的值，每行递增1
	const (
		d = iota // 0
		e        // 1
		f        // 2
	)
	// 定义常量在用一行，则值都一样
	const (
		g       = iota
		h, i, j = iota, iota, iota
		//k  此处不能定义缺省常量，编译错误
	)

	// 两种方式：第一种是显示声明，第二种隐式自动推到类型为[10]int
	//var arr [10]int = [10]int{a, b, c, d, e, f, g, h, i, j}
	arr := [10]int{a, b, c, d, e, f, g, h, i, j}
	return arr
}

//func num() (i, j, k int) {
//	return i, j, k
//}
