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
	var aa int // 声明一个变量，默认为0
	var bb = 10
	cc := 20 // 声明并初始化，且自动推导类型

	fmt.Println(aa, bb, cc)

	// 多变量声明
	var a1, b1 string
	fmt.Println("a1,b2:", a1, b1)
	var c1, c2 string = "h1", "h3"
	fmt.Println("c1,c2:", c1, c2)
	h2, h4 := "h2", "h4"
	fmt.Println(h2, h4)

	m3, n3, q3 := 10, "n2", 30 // 自动推导类型，并初始化值
	fmt.Println(m3, n3, q3)

	// 也可以这样声明多个变量
	var (
		e int
		f bool
	)
	fmt.Println("e f", e, f) // 0 false

	// := 自动推导声明的类型
	m1, m2 := "m1", "m2" // 根据值类型 自动推导变量类型
	fmt.Println(m1, m2)
	f1, f2, f3 := 10, "n2", 30 // 自动推导类型，并初始化值
	fmt.Println(f1, f2, f3)

	// 变量值互换
	var m = 15
	var n = 20
	m, n = n, m // 直接将m与n的值互换
	fmt.Println("m,n", m, n)

	var n1, n2 = 22, 25 // 两者
	var temp int
	temp, _ = n1, n2 // 将n1的值赋值给temp，_表示丢弃变量n2
	fmt.Println("temp", temp)

	// 丢弃变量
	_, ee := 35, 36 // 将值36赋值给ee，并同时丢弃35值
	fmt.Println("ee", ee)

	//:= 声明注意事项
	num1, num2 := test()
	fmt.Println(num1, num2)
}

func test() (a, b int) { // 函数名
	return 1, 2
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
