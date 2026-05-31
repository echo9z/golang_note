package main

import (
	"fmt"
	"reflect"
	"slices"
	"strconv"
)

func main()  {
	// 切片
	// 切片是对数组一个连续片段元素引用（可以称之为相关数组），所以切片是一个引用类型
	// 切片本身并不存储任何数据，它只是对底层数组的一个连续片段的引用。

	// 一、切割数组，引出切片
	// 切割数组返回切片 arr[start:end]，分割区域左闭右开，切割完后返回切片类型，前面说的只是对数组一个连续片段的引用
	arr := [5]int{1,2,3,4,5}

	/**
	arr[:] // 子切片范围[0,5) -> [1 2 3 4 5]
	arr[1:] // 子切片范围[1,5) -> [2 3 4 5]
	arr[:5] // 子切片范围[0,5) -> [1 2 3 4 5]
	arr[2:3] // 子切片范围[2,3) -> [3]
	arr[1:3] // 子切片范围[1,3) -> [2 3]
	*/
	var slice1 []int = arr[1:3] // 子切片范围[1～3) 2至3个元素
	fmt.Printf("arr：%v 类型:%T\n",arr, arr) // arr：[1 2 3 4 5] 类型:[5]int
	fmt.Println("slice1:",slice1, "类型:",reflect.TypeOf(slice1)) // slice1: [2 3] 类型: []int。这里通过reflect反射回去变量类型
	
	// 访问切片中引用的元素
	fmt.Println(slice1[0]) // 2

	// 想要将数组转换为切片，直接arr[:]不带索引取值范围转换为切片
	arr2 := [5]int{1,2,3,4,5}
	slice2 := arr2[:] // len(arr2) == 5，cap(arr2) == 5
	// 转换后的切片和数组指向的是同一个内存，当修改切片中元素会影响元数组的数据变化
	slice2[1] = 20
	fmt.Printf("array:%v\n", arr2) // array:[1 20 3 4 5]
	fmt.Printf("slice:%v\n", slice2) // slice:[1 20 3 4 5]
	
	// 若想对原数组不受影响，使用go1.21 引入的slices.Clone()标准库，用于创建切片的浅拷贝
	arr3 := [5]int{1,2,3,4,5}
	// slices.Clone返回一个独立新切片
	slice3 := slices.Clone(arr3[:]) // 克隆arr3[0:5]的切片，返回独立切片
	slice3[0] = 100
	fmt.Printf("array:%v\n", arr3) // array:[1 2 3 4 5]
	fmt.Printf("slice:%v\n", slice3) // slice:[100 2 3 4 5]
	
	// 等价于在go1.21之前版本通过make+copy，实现数组到切片，在通过 copy函数是切片深拷贝
	// make([]T, len, cap)：首先分配一段新的底层数组内存。通过指定长度为 len(originSlice)，确保了拷贝的目标切片拥有足够的空间来容纳所有元素。
	slice4 := make([]int, len(arr3)) // make函数创建一个[]int切片长度与arr3一样
	// func copy(dst, src []Type) int
	// copy 内置函数将源切片中的元素复制到目标切片中。 
	// （作为一种特殊情况，它还将字节从字符串复制到字节切片。）源和目标可能重叠。 
	// Copy 返回复制的元素数量，该数量将是 len(src) 和 len(dst) 中的最小值。
	copy(slice4, arr3[:])
	slice4[4] = 500
	fmt.Printf("array:%v\n", arr3) // array:[1 2 3 4 5]
	fmt.Printf("slice:%v\n", slice4) // slice:[100 2 3 4 5]

	// 这里浅拷贝概念
	// 基本类型 值类型 通过浅拷贝，完全独立的切片
	a := []int{1,2,3}
	b := slices.Clone(a)
	b[0] = 10
	fmt.Printf("a:%v, b:%v\n",a,b) // a:[1 2 3], b:[10 2 3] 不影响原切片

	// 引用类型 切片中是（指针/切片/map）- 只复制了引用，底层对象仍共享
	type User struct {
		Scores []int
	}

	var user1 User = User{
		Scores: []int{100,98,85},
	}
	var user2 User = User{
		Scores: []int{140,98,85},
	}
	// User引用切片
	Users := []User{user1, user2}
	// 通过Clone进行支队第一层了拷贝
	CloneUser := slices.Clone(Users)

	// 修改CloneUser中切片第一个元素，不影响Users对象切片
	CloneUser[0] = User{
		Scores: []int{60,60,60},
	}
	fmt.Printf("users%v\n", Users)  // users[{[100 98 85]} {[140 98 85]}]
	fmt.Printf("CloneUser%v\n", CloneUser) // CloneUser[{[60 60 60]} {[140 98 85]}]

	// 但我对CloneUser[1]中的属性Scores，切片第三个元素进行修改150分，就会影响原切片，切片中Scores还是共享的
	CloneUser[1].Scores[2] = 150
	fmt.Printf("users%v\n", Users) // users[{[100 98 85]} {[140 98 150]}]  150就是被影响了
	fmt.Printf("CloneUser%v\n", CloneUser) // CloneUser[{[60 60 60]} {[140 98 150]}]

	// 需要要深拷贝要自己递归 clone
	for idx := range Users {
		// 将CloneUser对象切片中Scores切片，也进行拷贝，从而达到scores切片为独立的内存
		CloneUser[idx].Scores = slices.Clone(Users[idx].Scores)
	}
	// 通过对scores切片属性进行深拷贝处理，再次对Scores切片第三个元素进行修改120分，就不会影响原切片
	CloneUser[1].Scores[2] = 120
	fmt.Printf("users%v\n", Users)  // users[{[100 98 85]} {[140 98 150]}] 150没有被影响修改
	fmt.Printf("CloneUser%v\n", CloneUser) // CloneUser[{[100 98 85]} {[140 98 120]}]

	// 二、创建切片
	// x := []int{2, 3, 5, 0, 0}[:3]
	//1.字面量初始化切片
	s1 := []int{1,2,3} // 长度和容量都为3
	s2 := []string{"a", "b"} // 长度和容量都为2
	fmt.Printf("slice1:%v(%T), slice2:%v(%T)\n",s1,s1,s2,s2) // slice1:[1 2 3]([]int), slice2:[a b]([]string)

	// 2.使用make函数创建切片
	/* func make(t Type, size ...IntegerType) Type
		make(type, length, capacity) // 仅适用于 slice
		make(type, initialCapacity)  // 适用于 map 和 channel
		type：必须是 slice、map 或 channel 类型。
		length（仅对 slice 有效）：
			切片的初始长度（元素数量）。
			必须指定，否则编译报错。
		capacity（可选，仅对 slice 有效）：
			切片的容量（底层数组的大小）。
			若未指定，默认与 length 相等。

		initialCapacity（对 map 和 channel 有效）：
			map：预分配的哈希表桶数（可选，若未指定则按需动态分配）。
			channel：通道的缓冲区大小（可选，若未指定则为无缓冲通道）。
	*/

	// make 内置函数分配并初始化 slice、map 或 chan（仅）类型的对象。
	// 与 new 一样，第一个参数是类型，而不是值。
	// 与 new 不同，make 的返回类型与其参数的类型相同，而不是指向new函数传入函数参数的指针。

	// make([]T, len)       // size[0] = len，cap = len
	// make([]T, len, cap)  // size[0] = len，size[1] = cap
	// 一个参数les和cap相等
	s3 := make([]int, 5) // 切片长度和容量都为5
	fmt.Printf("%v\n", s3)
	fmt.Printf("len:%d, cap:%d\n", len(s3), cap(s3)) // len:5, cap:5
	// 两个参数：指定len和cap大小
	s4 := make([]string, 2, 5) // len2, cap5   内存中["","",""，""，""]，只有前两个索引位置合法可访问
	fmt.Printf("%v\n", s4)
	fmt.Printf("len:%d, cap:%d\n", len(s4), cap(s4)) // len:2, cap:5

	// 3.通过append()向预分配的切片追加元素
	// func append(slice []Type, elems ...Type) []Type
	// 	slice：目标切片
	// 	elems：可变参数，向目标切片添加1个或多个元素
	// Append 内置函数的作用是将元素添加到切片的末尾。
	// 如果内存容量足够，则会重新分配切片的空间以容纳新添加的元素；
	// 如果内存不足，则会创建一个新的数组来存储这些元素。最后，Append 会返回更新后的切片。
	// 因此，通常需要将 Append 的结果存储在一个变量中，这个变量就应该是保存切片本身的那个变量。
	for i := 0; i < 5; i++ {
		s4 = append(s4, strconv.Itoa(i)+"a")
	}
	fmt.Printf("s4:%v\n",s4) // s4:[  0a 1a 2a 3a 4a]
	// 追加单个元素
	a1 := []int{1,2}
	a1 = append(a1, 3)
	fmt.Println(a1) // [1 2 3]
	// 追加多个元素
	a1 = append(a1, 4,5,6)
	fmt.Println(a1) // [1 2 3 4 5 6]

	// 将一个切片追加到另一个切片中
	b1 := []int{7,8,9}
	// 用...把b切片展开为可变参数
	a1 = append(a1, b1...)
	fmt.Println(a1) // [1 2 3 4 5 6 7 8 9]

	// 追加字符串到[]byte
	var buf []byte
	buf = append(buf, "hello"...) // 将hello分单个字符 buf==hello 
	var str []string
	str = append(str, "hello")
	fmt.Printf("字节%s 字符串切片%s",buf, str) // 字节hello 字符串[hello]

	// append扩容机制

}