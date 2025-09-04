package main

import "fmt"

func main() {
	// 算数运算符
	var a int = 21
	var b int = 10
	// + - * / %
	fmt.Println("a+b", a+b) // 31
	fmt.Println("a-b", a-b)
	fmt.Println("a*b", a*b) // 210
	fmt.Println("a/b", a/b) // 取整
	fmt.Println("a%b", a%b) // 取余
	a++                     // a变量自增    22
	fmt.Println("a", a)
	b-- // b变量自减    11
	fmt.Println("b", b)

	// uint8 类型取值范围 0~255
	// 用8位，2^8 = 256,0~255 共256个数，在二进制中255表示为11111111，11111111上加1结果100000000，
	// uint8 类型只能容纳 8 位。它无法存储那第 9 位（最高位的 1 ）。所以，这一位被舍弃了，剩下的就是 00000000 ，0的二进制表示
	var uInt8Max uint8 = 255              // 声明uInt8Max变量，类型为uint8,值为该类型最大值
	fmt.Println("uInt8Max+1", uInt8Max+1) // 0  256 % 256 = 0

	// int8 类型取值范围 -128~127
	// int8Max 设置为 127，int8是有符号得。二进制表示为 01111111，在 01111111 上加 1。
	// 二补码系统中，二进制模式 10000000 不代表 128 。有符号即-128。
	var int8Max int8 = 127   // 声明int8Max变量，类型为uint8,值为该类型最大值
	fmt.Println(int8Max + 1) // 输出运算结果 -128

	Relationship()
	Bitwise()
}

// Relationship 关系运算符
func Relationship() {
	fmt.Println(100 == (50 + 50))
	fmt.Println((51 + 49) != (50 * 2))
	var str string = "abcde"
	fmt.Println(str[0] == 97)
}

// Logic 逻辑运算符
func Logic() {
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!(1 > 6))
}

// Bitwise 位运算符
func Bitwise() {
	fmt.Println(1561 & 0) // 与运算
	fmt.Println(3 | 5)    // 或运算
	numOne, numTwo := 10, 20
	fmt.Println(numOne ^ numTwo)  // 输出 numOne 和 numTwo 变量的按位异或结果
	fmt.Println(numOne &^ numTwo) // 输出 numOne 和 numTwo 变量的按位清空结果

	// &^ 清除运算
	a := 12                     // 二进制 0000 1100
	b := 10                     // 二进制 0000 1010
	fmt.Println("通过清除运算", a&^b) // 0000 0100 = 4
	// &^ b 保留 a 中的位，但清除 b 中为 1 的对应位
	//当 `b` 的某一位是 `1` 时：无论 `a` 是什么，结果都是 `0`（清除）
	//当 `b` 的某一位是 `0` 时：保持 `a` 的原值（保留）

	// 在go使用^进行取反
	x := 12                                        // 二进制 0000 1100
	y := ^x                                        // 取反运算
	fmt.Printf("x的值 %d (%08b)\n", x, uint8(x))     // x的值 12 (00001100)
	fmt.Printf("x取反，y结果 %d (%08b)\n", y, uint8(y)) // x取反，y结果 -13 (11110011)

	// go中 ^x 是一元取反运算符，但也可以用二元异或 m^x 来达到相同效果，m 的选择：
	// 无符号类型: m 应该是"全部位设置为1"的值
	// 有符号类型: m 应该是 -1
	// 任何数与"全1"异或，结果就是该数的按位取反：
	// x ^ 11111111 = ~x  (按位取反)
	// 无符号类型: m 应该是"全部位设置为1"的值
	var x1 uint8 = 12
	var z1 uint8 = x1 ^ 255 // 11111111 = 255
	fmt.Printf("无符号与255的异或运算 %d (%08b)\n", z1, uint8(z1))
	var z2 uint8 = ^x1
	fmt.Printf("z2取反结果 %d (%08b)\n", z2, uint8(z2))

	// 有符号类型: m 应该是 -1
	var x2 int8 = 12
	var z3 int8 = x2 ^ -1
	// -1 的二进制表示为11111111
	// 1 的原码: 00000001
	// 1 的反码: 11111110
	// 1 的补码: 11111111 (反码+1) 最终-1=11111111
	fmt.Printf("有符号与-1的异或运算 %d (%08b)\n", z3, uint8(z3))
	var z4 int8 = ^x2
	fmt.Printf("z4取反结果 %d (%08b)\n", z4, uint8(z4))

	fmt.Println(174 << 2)
	fmt.Println(174 >> 2)
}

// Swap 定义函数 func 函数名称(参数1,参数2 参数类型) (函数返回值的类型) {}
func Swap(n1, n2 int) (int, int) {
	if n1 != n2 {
		n1 ^= n2
		n2 ^= n1
		n1 ^= n2
	}
	return n1, n2
}

// Ass 赋值运算
func Ass() {
	var num int = 10
	num += 10 // num = num + 10
	fmt.Println(num)
	num -= 9
	num *= 100
	num /= 20
	num %= 4  // num = num%4
	num <<= 2 // num = num << 2
	num >>= 2 // num = num >> 2
	fmt.Println("num", num)
}

// 指针
func Pointerfnc() {
	var num int = 10
	var pointer *int = &num // 声明pointer变量,类型位指针，值为num变量的内存地址
	fmt.Println(&num)       // num变量的实际内存地址0xc00000a0b8
	fmt.Println(*pointer)   // pointer变量表示的内存地址所存储的变量的值 10
}
