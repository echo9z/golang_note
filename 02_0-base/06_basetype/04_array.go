package main

import "fmt"

func main() {
	// 数组 固定长度、相同元素的集合，一旦声明长度不可改变
	// 一、声明数组
	// 每个元素是一个整型值，当声明数组时所有的元素都会被自动初始化为默认值 0。
	var arr1 [5]int   // 声明了一个长度为5组，默认值都是0
	fmt.Println(arr1) // [0 0 0 0 0]

	// 声明并初始化
	var arr2 = [5]int{51, 25, 30, 50, 63}
	fmt.Println(arr2) // [51 25 30 50 63]

	// 初始化部分元素
	var arr3 = [5]int{51, 36, 9}
	fmt.Println(arr3) // [51 36 9 0 0]

	// 编译器自己推断数组长度
	var arr4 = [...]int{12, 25, 35, 40}
	fmt.Println(len(arr4)) // 4

	// 指定索引初始化
	var arr5 = [5]int{0: 10, 4: 82} // arr5[0]=10, arr5[4]=82, 其他为0
	fmt.Println(arr5)               // [10 0 0 0 82]

	// 简短声明
	arr6 := [3]string{"hello", "gg", "ok"}
	fmt.Println(arr6) // [hello gg ok]

	arr7 := [5]int{54, 85, 68, 56, 8}
	// 第一个元素是 arr1[0]，第三个元素是 arr1[2]；总体来说索引 i 代表的元素是 arr1[i]，最后一个元素是 arr1[len(arr1)-1]。
	// 访问元素
	fmt.Println(arr7[2]) // 68
	// 修改元素
	arr7[0] = 100
	fmt.Println(arr7[0])
	// 最后一个元素
	fmt.Println("end",arr7[len(arr7)-1])

	// [5]int和 [10]int 是属于不同类型的。
	var a1 [5]int = [5]int{0, 1, 2, 3, 4}  // 类型是 [5]int
	var b1 [10]int = [10]int{0: 10, 9: 90} // 类型是 [10]int

	// 在go中数组是值类型，一个 [5]int 变量就是连续 5 个 int 大小的内存块，一个 [10]int 是连续 10 个 int 大小的内存块。它们的尺寸根本不一样，在内存布局、复制开销上都完全不同。
	// 长度是类型的一部分，数组的大小在编译时就是固定的。当你把一个数组赋值给另一个变量，或者传给函数时，Go 会复制整个数组的所有元素。
	// a1 = b1 // 编译错误：cannot use a (type [5]int) as type [10]int in assignment
	fmt.Printf("a1: %d\nb1: %d\n", a1, b1)

	// 切片的类型只有 []int，长度不是类型的一部分。切片的类型不包含长度，所以无论长度是 5 还是 10 的切片，它们的类型都是 []int，可以自由地互相赋值、扩容，或者作为同一个函数的参数。
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("s1:%p val:%d \n",s1, s1) // s1:0xc0000182e8 val:[1 2 3] 
	fmt.Printf("s2:%p val:%d \n",s2, s2) // s2:0xc0000200a0 val:[1 2 3 4 5 6 7 8 9 10] 
	s1 = s2 // 都是 []int 类型
	fmt.Printf("s1:%p val:%d \n",s1, s1) // s1:0xc0000200a0 val:[1 2 3 4 5 6 7 8 9 10] 
	fmt.Printf("s2:%p val:%d \n",s2, s2) // s2:0xc0000200a0 val:[1 2 3 4 5 6 7 8 9 10] 

	// 数组指针声明
	numP := [3]int{1,2,3}

	p := &numP // go中使用&获取一个变量的指针地址
	fmt.Println(p) // p指针类型类型为：*[3]int
	fmt.Println((*p)[0]) // *p 读写指针指向的值，获取数组中的第一个元素
	fmt.Println(p[0]) // 简写

	// new 创建数组
	p1 := new([3]int) // *[3]int，默认值[0,0,0]
	// 实际开发中 new 用得很少，&T{...} 更常用，既能拿到指针又能初始化。
	p1[0] = 10
	p1[1] = 20
	p1[2] = 30
	fmt.Println(p1)  // &[10 20 30]
	// new([3]int)           // → &[0, 0, 0]，只能得到零值
	// &[3]int{10, 20, 30}  // → &[10, 20, 30]，声明时就能赋值

	//二、多维数组
	// 二维数组
	var arrR[2][3]int = [2][3]int{
		{25, 15, 26},
		{15, 35, 36},
	}
	fmt.Printf("[1][2]:%d\n",arrR[1][2])

	// 简短声明
	arrR2 := [2][3]int{
		{25, 15, 26},
		{15, 35, 36},
	}
	fmt.Printf("arrR2:%d\n",arrR2)

	// 自动推导
	arrR3 := [...][3]int{ // [...][...]int❌（只能最外层用...)
		{25, 15, 26},
		{15, 35, 36},
	}
	fmt.Printf("arrR3:%d\n",arrR3)

	// 读取多维数组元素
	fmt.Println(arrR3[1]) // [15 35 36]  第二行
	fmt.Println(arrR3[0][2]) // 26   第1行，第3个元素

}
