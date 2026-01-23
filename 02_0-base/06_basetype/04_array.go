package main

import "fmt"

func main() {
	// 数组 固定长度、相同元素的集合，一旦声明长度不可改变
	// 声明数组
	// 每个元素是一个整型值，当声明数组时所有的元素都会被自动初始化为默认值 0。
	var arr1 [5]int   // 声明了一个长度为5组，默认值都是0
	fmt.Println(arr1) // [0 0 0 0 0]
	
	// 声明并初始化
	var arr2 = [5]int{51, 25, 30,50,63}
	fmt.Println(arr2) // [51 25 30 50 63]

	// 初始化部分元素
	var arr3 = [5]int{51,36,9}
	fmt.Println(arr3) // [51 36 9 0 0]

	// 编译器自己推断数组长度
	var arr4 = [...]int{12, 25, 35, 40}
	fmt.Println(len(arr4)) // 4

	// 指定索引初始化
	var arr5 = [5]int{0: 10, 4: 82} // arr5[0]=10, arr5[4]=82, 其他为0
	fmt.Println(arr5) // [10 0 0 0 82]
	
	// 简短声明
	arr6 := [3]string{"hello", "gg", "ok"}
	fmt.Println(arr6) // [hello gg ok]

	arr7 := [5]int{54, 85, 68, 56, 8}
	// 访问元素
	fmt.Println(arr7[2]) // 68
	// 修改元素
	arr7[0] = 100
	fmt.Println(arr7[0])

	// 第一个元素是 arr1[0]，第三个元素是 arr1[2]；总体来说索引 i 代表的元素是 arr1[i]，最后一个元素是 arr1[len(arr1)-1]。
}
