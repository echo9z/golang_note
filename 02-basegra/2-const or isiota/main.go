package main

import "fmt"

func main() {
	// 1.const关键字，常量在声明时就必须初始化一个值，并且常量的类型可以省略，
	const name string = "tom"
	const msg = "ok 123"
	const PI float32 = 3.14159
	const A byte = 'A'
	const numExpression = (1+2+3)/2%100 + A // 常量表达式
	fmt.Println(name, msg, numExpression, A, PI)

	// 2.批量声明常量可以使用()
	const (
		COUNT int    = 5
		NAME  string = "ok"
	)

	const (
		a = 5
		b = 10
	)

	// 在同一个常量分组中，在已经赋值的常量后面的常量可以不用赋值，其值默认就是前一个的值，比如
	const (
		A1 = 100
		B1
		C1
	)
	fmt.Println(A1, B1, C1)

	// 3.赋值表达式中设计的计算过程，必须在编译期间就能获得。
	// 正确：const c1 = 2/3
	// 错误：const c2 = getNumber() // 引发构建错误: getNumber() used as value
	//const Ln2 = 0.693147180559945309417232121458
	//        \176568075500134360255254120680009
	// 反斜杠 \ 可以在常量表达式中作为多行的连接符使用。(此处的反斜杠已经不能作为多行的连接符使用了)
	const Ln2 = 0.365486
	const Log2E = 1 / Ln2 // this is a precise reciprocal
	const Billion = 1e9   // float constant

	// 4.iota 通常用于表示一个常量声明中的无类型整数序数，一般都是在括号中使用。
	const (
		num1 = iota
		num2 // 1
		num3 // 2
		num4 // 3
	)
	// 还可以通过表带是，表达式方式,iota初始值为0
	const (
		value = iota * 2
		v1    // 2
		v2    // 4
	)
	// 还可以
	const (
		AA = iota<<2*3 + 1 // 1
		BB                 // 1 << 2*3+1=13
		CC                 // 2 << 2*3+1=25
		// 重新初始 iota 递增规则
		DD = iota // 3
		EE        // 4
	)
	// iota是递增的，第一个常量使用iota值的表达式，根据序号值的变化会自动的赋值给后续的常量，直到用新的const重置
	const (
		N1 = iota<<2*3 + 1 // 1 第一行
		N2 = iota<<2*3 + 1 // 13 第二行
		_                  // 25 第三行 占位符进行忽略
		N3                 // 37 第四行
		N4 = iota          // 4 第五行
		_                  // 5 第六行
		N6                 // 6 第七行
	)
}
