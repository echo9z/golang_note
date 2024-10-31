package base

import (
	"fmt"
	"math"
	"unsafe"
)

// 判断两个浮点数是否相等
func FloatEqual(a, b, epsilon float64) bool {
	return math.Abs(a-b) < epsilon
}

func testEqual() {
	a := 0.1 + 0.2
	b := 0.3
	epsilon := 1e-9 // 误差范围
	if FloatEqual(a, b, epsilon) {
		fmt.Println("a and b are equal.")
	} else {
		fmt.Println("a and b are not equal.")
	}
}

// NumberDataType 值类型
func NumberDataType() {
	// 整型
	var n1 int8 = 15
	var n2 uint = 12
	fmt.Println("size", unsafe.Sizeof(n1)) // 查看字节数方法
	fmt.Println(n1, n2)
	// var a int = 12
	// var b int32 = 61
	// c := a + b // 不同类型，不允许互相赋值或操作

	// 浮点
	var f32 float32 = 12.6
	var f64 float64 = 3.14159265359
	fmt.Println(f32, f64)
	// 精度损失
	var f1 float32 = -123.0000803
	var f2 float64 = -123.0000803
	fmt.Println(f1, f2) // 精度缺失 -123.000084 -123.0000803

	var num1 float64 = 12.45666666
	var num2 float64 = 12.45666667
	// 浮点数是否相等
	flags := math.Abs(num2-num1) < 0.00000001
	fmt.Println("浮点型判断：", num2-num1 == 0.00000001)
	fmt.Println("判断：", flags)
	fmt.Println("判断：", math.Abs(num2-num1))

	// 布尔
	var available bool // 一般声明
	valid := false     // 简短声明
	available = true   // 赋值操作
	fmt.Println(available, valid)

	// 字符串
	var str string = "string"
	fmt.Println(str)

	// 复数，形如a+bi（a、b均为实数）的数为复数，其中，a被称为实部，b被称为虚部，i为虚数单位
	var c complex64 = 5 + 5i
	fmt.Printf("复数: %v\n", c)   // 5+5i
	fmt.Println("实部：", real(c)) // 5
	fmt.Println("虚部：", imag(c)) // 5

	// 常量
	const Pi float64 = 3.14159265359
	const MaxThread = 1000
	fmt.Println("常量：", Pi, MaxThread)

	// NaN非数
	var n float64
	fmt.Println("NaN：", n, -n, 1/n, -1/n, n/n) // 0 -0 +Inf -Inf NaN
}

// 格式化输出
type User struct { // 定义一个User类
	Name string
	Age  int
}

func Format() {
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
