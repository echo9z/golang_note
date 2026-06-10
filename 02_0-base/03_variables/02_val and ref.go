package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	// 值类型与引用类型
	// int、float、bool 和 string array struct这些基本类型都属于值类型，使用这些类型的变量直接指向存在内存中的值：

	// 1.值类型
	var i1 int = 7
	var j1 int = i1 // 实际上是在内存中将 i 的值进行了拷贝
	fmt.Println("value：", &i1, &j1) // value： 0x1d7e33b54120 0x1d7e33b54128 地址值不同
	// i 和 j 的内存地址不同，每个变量都有自己独立的内存空间，j=i 是将 i 的值拷贝了一份给 j
	// i内存值0xc00000a130=7，存放在栈区
	// j=i，j变量开辟内存值0xc00000a138，将值7进行拷贝到j内存中

	// 修改j1 不影响i1
	j1 = 100
	fmt.Printf("i1:%d, j1:%d\n", i1, j1) // i1:7, j1:100

	//栈内存布局：
	//+--------+--------+
	//| 地址   | 值     |
	//+--------+--------+
	//| 0xa130 |   7    |  <- i 的位置
	//+--------+--------+
	//| 0xa138 |   7    |  <- j 的位置
	//+--------+--------+

	// 如果想要相同的内存地址，得让j指向i的地址，此时 j 就是一个引用类型 指针类型
	var i int = 7
	var j *int = &i // j 指向 i 的地址

	fmt.Printf("i 的地址: %p\n", &i) // i 的地址: 0x1d7e33b54140
	fmt.Printf("j 指向的地址: %p\n", j) //j 指向的地址: 0x1d7e33b54140
	fmt.Printf("地址相同吗? %t\n", &i == j) // 地址相同吗? true

	// 值类型：数组
	arr1 := [3]int{1, 2, 3}
	arr2 := arr1 // 复制整个数组
	arr2[0] = 100
	fmt.Println(arr1[0]) // 输出 1 (原始不变)

	// 值类型：结构体
	type Person struct {
			Name string
			Age  int
	}
	pp1 := Person{"Alice", 20}
	pp2 := pp1 // 复制整个结构体
	pp2.Age = 30
	fmt.Println(pp1.Age) // 输出 20

	// 引用类型的变量存储一个引用（即内存地址），实际数据位于底层共享的数据结构中。
	// 赋值或传参会复制引用（指针），因此多个变量指向同一份数据。
	// 2.引用类型
	var sli1 []int = []int{1, 2, 3}
	var sli2 []int = sli1 // 将 sli1 的引用赋值给了 sli2

	fmt.Printf("sli1 切片头地址: %p\n", &sli1) // 切片变量本身的地址
	fmt.Printf("sli2 切片头地址: %p\n", &sli2) // 切片变量本身的地址
	fmt.Printf("sli1 底层数组地址: %p\n", sli1) // 底层数组的地址
	fmt.Printf("sli2 底层数组地址: %p\n", sli2) // 底层数组的地址

	// sli1 切片头地址: 0xc000008030
	// sli2 切片头地址: 0xc000008048 <- 不同地址
	// sli1 底层数组地址: 0xc000014108
	// sli2 底层数组地址: 0xc000014108 <- 相同地址

	// 栈上的切片头：
	// +--------+------------+-----+-----+
	// | 地址   | ptr        | len | cap |
	// +--------+------------+-----+-----+
	// | &sli1  | 0x...12120 |  3  |  3  |  <- sli1 切片头
	// +--------+------------+-----+-----+
	// | &sli2  | 0x...12120 |  3  |  3  |  <- sli2 切片头
	// +--------+------------+-----+-----+
	
	// 	堆上的底层数组：
	// +--------+---+---+---+
	// | 地址   | 1 | 2 | 3 |
	// +--------+---+---+---+
	// |12120   | 底层数组  |  <- 两个切片都指向这里
	// +--------+-----------+

	type User struct {
		name string
		age  int
	}

	var u1 User = User{name: "tom", age: 15}
	var u2 User = User{name: "jack", age: 18}
	fmt.Printf("user1 底层对象地址: %p\n", &u1) // user1 底层对象地址: 0x204482200a8
	fmt.Printf("user2 底层对象地址: %p\n", &u2) // user2 底层对象地址: 0x204482200c0

	fmt.Printf("age %d\n", &u1.age) // age 2217413312696
	// string 字段的特殊性
	fmt.Printf("u1.name 地址: %p\n", &u1.name) // u1.name 地址: 0x204482200a8
	fmt.Printf("\n=== string 内部结构 ===\n")
	// string 内部是 {ptr, len} 结构  ai解释
	nameHeader := (*reflect.StringHeader)(unsafe.Pointer(&u1.name))
	fmt.Printf("name 字符串数据地址: %p\n", unsafe.Pointer(nameHeader.Data)) // name 字符串数据地址: 0x4c6b8b
	fmt.Printf("name 字符串长度: %d\n", nameHeader.Len) // name 字符串长度: 3

	// u1 结构体内存:
	// +--------+----------------+--------+--------+
	// | 地址   | name (16字节)   | age    | 填充   |
	// +--------+----------------+--------+--------+
	// | &u1    | ptr | len      |   15   |        |
	// +--------+-----+----------+--------+--------+
	// 			|     |
	// 			|     +-> 字符串长度 (8字节)
	// 			+-> 指向 "tom" 的指针 (8字节)
	// 实际字符串 "tom" 存储在其他内存位置

	// 结构体赋值
	var u3 User = u1                        // 将 u1 的值拷贝了一份给 u3
	// user3 底层对象地址: 0x204482200d8
	fmt.Printf("user3 底层对象地址: %p\n", &u3)   // u1 和 u3 的内存地址不同
	fmt.Printf("u1 == u3 ? %t\n", u1 == u3) // true，值相等

	// 修改 u1，u2 不受影响
	u1.age = 99
	fmt.Printf("u1.age: %d\n", u1.age) // 99
	fmt.Printf("u3.age: %d\n", u3.age) // 15

	// 指针类型
	var p1 *User = &User{name: "alice", age: 20} // 内存地址值
	var p2 *User = p1
	// p1 和 p2 指向同一个内存地址
	p1.age = 99
	fmt.Printf("p1.age: %d\n", p1.age) // 99
	fmt.Printf("p2.age: %d\n", p2.age) // 99 (受影响)

	structDetailAnalysis()
}

func structDetailAnalysis() {
	type DetailUser struct {
		name    string
		age     int
		hobbies []string
		scores  map[string]int
	}

	u1 := DetailUser{
		name:    "tom",
		age:     15,
		hobbies: []string{"reading", "coding"},
		scores:  map[string]int{"math": 90, "english": 85},
	}

	fmt.Printf("=== 结构体完整内存分析 ===\n")
	fmt.Printf("结构体地址: %p\n", &u1)

	fmt.Printf("\n--- 字段地址 ---\n")
	fmt.Printf("name 字段: %p\n", &u1.name)
	fmt.Printf("age 字段: %p\n", &u1.age)
	fmt.Printf("hobbies 字段: %p\n", &u1.hobbies)
	fmt.Printf("scores 字段: %p\n", &u1.scores)

	fmt.Printf("\n--- 底层数据地址 ---\n")
	fmt.Printf("name 底层: %p\n", unsafe.StringData(u1.name))
	fmt.Printf("hobbies 底层: %p\n", u1.hobbies)
	fmt.Printf("scores 底层: %p\n", u1.scores)

	fmt.Printf("\n--- 切片元素底层 ---\n")
	for i, hobby := range u1.hobbies {
		fmt.Printf("hobbies[%d] 底层: %p\n", i, unsafe.StringData(hobby))
	}
}
