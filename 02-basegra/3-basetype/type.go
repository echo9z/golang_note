package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	// 1.bool
	var b1 bool = true
	fmt.Println(10 != 100 && b1) // 输出 true

	// 2.int 和 uint 在 32 位操作系统上，它们均使用 32 位（4 个字节），在 64 位操作系统上，它们均使用 64 位（8 个字节）。
	// uintptr 的长度被设定为足够存放一个指针即可。
	var i1 int = 10
	var i2 int32
	//i3 := i1 + i2  int16 也不能够被隐式转换为 int32
	i2 = int32(i1) + 10 // 显式转换 int 到 int32
	fmt.Println(i2)

	// 3.Go 语言中没有 float 类型。只有 float32 和 float64。没有 double 类型。
	// float32 精确到小数点后 7 位，float64 精确到小数点后 15 位。由于精确度的缘故，你在使用 == 或者 != 来比较浮点数时应当非常小心。你最好在正式使用前测试对于精确度要求较高的运算。
	// 你应该尽可能地使用 float64，因为 math 包中所有有关数学运算的函数都会要求接收这个类型
	var f1 float32 = 3.14
	f2 := uint64(f1)
	fmt.Println(f2)

	// 数字值转换
	var i3 int = -10
	value, err := Uint8FromInt(i3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("转换后的值:", value)
	}

	// 4.复数
	// 复数，形如a+bi（a、b均为实数）的数为复数，其中，a被称为实部，b被称为虚部，i为虚数单位
	var c complex64 = 5 + 5i
	fmt.Printf("复数: %v\n", c)   // 5+5i
	fmt.Println("实部：", real(c)) // 5
	fmt.Println("虚部：", imag(c)) // 5

	// 5.NaN非数
	var n float64
	fmt.Println("NaN：", n, -n, 1/n, -1/n, n/n) // 0 -0 +Inf -Inf NaN
	var x, y float64
	x = math.NaN() // 返回一个Nan值
	y = 10.36
	fmt.Println(x == y) // false
	fmt.Println(x == x) // false
	fmt.Println(y == y) // true
	z := x + y          // 进行数值计算，返回NaN
	fmt.Println(z)      // 返回NaN

	// 6.随机数
	RandNum()
}

// 随机函数
func RandNum() {
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Printf("%d / ", a)
	}
	for i := 0; i < 5; i++ {
		r := rand.Intn(8) // 生成一个0-7之间的随机数
		fmt.Printf("%d / ", r)
	}
}

// int => uint8 大取值范围 转换 小取值范围
func Uint8FromInt(n int) (uint8, error) {
	// uint8 的取值范围是 0 到 255
	if n >= 0 && n <= math.MaxUint8 {
		return uint8(n), nil
	}
	return 0, fmt.Errorf("int %d is out of range for uint8", n)
}

// float64 => int32
func IntFromFloat64(x float64) int {
	if math.MinInt32 <= x && x <= math.MaxInt32 { // x lies in the integer range
		whole, fraction := math.Modf(x)
		if fraction >= 0.5 {
			whole++
		}
		return int(whole)
	}
	panic(fmt.Sprintf("%g is out of the int32 range", x))
}
