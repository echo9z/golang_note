package main

import (
	"cmp"
	"fmt"
)

func main() {
	// 1.变量声明
	var aa int // 声明一个变量，默认为0
	var bb = 10
	cc := 20 // 声明并初始化，且自动推导类型 为int
	fmt.Println(aa, bb, cc)

	// 2.多变量声明
	var a1, b1 string
	fmt.Println("a1,b2:", a1, b1)
	var c1, c2 string = "h1", "h3"
	fmt.Println("c1,c2:", c1, c2)
	h2, h4 := "h2", "h4"
	fmt.Println(h2, h4)

	m3, n3, q3 := 10, "n2", 30 // 自动推导类型，并初始化值
	fmt.Println(m3, n3, q3)

	// 3.var(...)声明多个变量
	var (
		name string = "John"
		flag bool
		pi   float32 = 3.14
		num  int     = 10
	)
	fmt.Println(name, flag, pi, num)

	// 4. := 自动推导声明变量，并初始化值
	m1, m2 := 98, 100
	fmt.Println(m1, m2)
	v1, v2, v3 := 3.14, true, 'c'
	fmt.Println(v1, v2, v3)

	// 5. _ 丢弃变量值
	value, _ := 1, 2 // 使用_丢弃第二个变量的值
	fmt.Println("value:", value)

	var n1, n2 = 22, 25 // 两者
	var temp int
	temp, _ = n1, n2 // 将n1的值赋值给temp，_表示丢弃变量n2
	fmt.Println("temp", temp)

	// := 接收函数
	r1, r2 := result()

	fmt.Println("r1 r2", r1, r2)

	// 6.在golang中直接使用a,b = b,a 将a和b的值进行交换
	a, b := 15, 20
	a, b = b, a
	fmt.Println("a, b值的交换:", a, b)
	// 三个值进行交换
	num1, num2, num3 := 25, 36, 49
	num1, num2, num3 = num3, num1, num2
	fmt.Println(num1, num2, num3) //49 25 36
	// 下面分解下
	// 第一步 num1, num2, num3 = num3, num1, num2 等式右边num3, num1, num2顺序，n3:49，n1:25，n2:36
	// 第二步 等式左边，按照将49，25，36顺序，对num1, num2, num3进行赋值，最终得到 n1:49, n2:25, n3:36

	// 7.比较
	maxVal := max(15.5, 15, 855, 1024.8)
	minVal := min(0, -1, 0.05)
	fmt.Println(maxVal, minVal)

	// 是内置cmp包
	cmp.Compare(15, 456)
	cmp.Less(1, 2) // x 是否小于 y

	// 8.类型别名 使用 type 类型别名
	type MyInt int
	var aVle, bVle MyInt = 10, 20
	cVle := aVle + bVle // MyInt 类型为 int类型的别名
	fmt.Printf("value: %d", cVle)
}
func result() (int, float32) {
	return 98, 'A'
}
