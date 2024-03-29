package base

import "fmt"

// Sum 算术运算符
func Sum() {
	var b = 22
	var c = 10
	fmt.Println("b+c", b+c)
	fmt.Println("b-c", b-c)
	fmt.Println("b*c", b*c)
	fmt.Println("b/c", b/c) // 取整数
	fmt.Println("b%c", b%c) // 取余
	b++                     // b变量自增    1
	fmt.Println("b", b)
	c--                 // c变量自减    1
	fmt.Println("c", c) //输出运算结果

	var uInt8Max uint8 = 255              // 声明uInt8Max变量，类型为uint8,值为该类型最大值
	fmt.Println("uInt8Max+1", uInt8Max+1) // uint8 类型最大取值为255
	var int8Max int8 = 127                // 声明int8Max变量，类型为uint8,值为该类型最大值
	fmt.Println(int8Max + 1)              // 输出运算结果
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

func Pointerfnc() {
	var num int = 10
	var pointer *int = &num // 声明pointer变量,类型位指针，值为num变量的内存地址
	fmt.Println(&num)       // num变量的实际内存地址0xc00000a0b8
	fmt.Println(*pointer)   // pointer变量表示的内存地址所存储的变量的值 10
}
