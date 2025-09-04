package main

import "fmt"

// 格式化输出
type User struct { // 定义一个User类
	Name string
	Age  int
}

func main() {
	// 创建一个user实例对象
	var user User = User{
		Name: "tom",
		Age:  15,
	}
	// 通过Printf(输出的format, 原数据)
	fmt.Printf("%%\n")     // 输出 %
	fmt.Printf("%b\n", 16) // 将16以二级制格式进行输出

	fmt.Printf("%c\n", 65)     // A 将数值转化为对应的unicode字符
	fmt.Printf("%c\n", 0x4f60) // 你
	fmt.Printf("%d\n", 'A')    // 将A 转换为十进制数值
	fmt.Printf("%d\n", '你')    // 将你 转换为十进制数值，转换为20320

	fmt.Printf("%x\n", '你') // 将字符转换为十六进制，字母使用小写 4f60
	fmt.Printf("%X\n", '你') // 将字符转换为十六进制，字母使用大写 4F60
	fmt.Printf("%U\n", '你') // 将字符转换为Unicode格式 U+4F60

	fmt.Printf("%t\n", 1 > 2)            // false
	fmt.Printf("%e\n", 4396.7777777)     // 4.396778e+03 默认精度6位
	fmt.Printf("%20.3e\n", 4396.7777777) //            4.397e+03 设置宽度20,精度3,宽度一般用于对齐
	fmt.Printf("%E\n", 4396.7777777)     // 4.396778E+03
	fmt.Printf("%f\n", 4396.7777777)     // 4396.777778
	fmt.Printf("%o\n", 16)               // 20
	fmt.Printf("%p\n", []int{1})         // 0xc000016110
	fmt.Printf("Hello %s\n", "World")    // Hello World
	fmt.Printf("Hello %q\n", "World")    // Hello "World"
	fmt.Printf("%T\n", 3.0)              // 输出值的类型 float64

	fmt.Printf("%v\n", user)  // 值的默认格式输出{tom 15}
	fmt.Printf("%+v\n", user) // 值的默认格式输出{Name:tom Age:15} kv形式的值
	fmt.Printf("%#v\n", user) // 值的默认格式输出base.User{Name:"tom", Age:15}
}
