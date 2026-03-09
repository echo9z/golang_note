package main

import (
	"fmt"
)

func main() {
	// for循环
	// 1.for常规循环
	// for init statement; expression; post statement {
	// 	execute statement
	// }
	// 在go中使用 for (i = 0; i < 10; i++) { }，这是无效的代码！不需要括号 ()
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// 初始化循环多个变量
	for i, j := 0, 100; i < 100 && j < 200; i, j = i+1, j+1 {
		fmt.Printf("i: %d j: %d", i, j)
	}
	// 变量二位数组
	var arr [2][3]int = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("访问元素arr二维中第一个元素:", arr[1][0]) // 4
	// 遍历这个二维数组
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j])
			if len(arr[i])-1 == j {
				fmt.Println("")
			}
		}
	}
	// 使用for range遍历
	for i, row := range arr {
		fmt.Println(row)
		for j, column := range row {
			// fmt.Println(column)
			fmt.Printf("arr[%d][%d]:=%d\n", i, j, column)
		}
	}

	// 2.go中类似while语法
	sum := 1
	for sum <= 100 {
		sum += sum
	}
	fmt.Println("sum:", sum)

	// 3.无限死循环
	// 通常通过 break、return 或 panic 退出
	i := 0
	for { // 等价于 for true {}
		if i > 10 {
			break
		}
		if i == 5 {
			fmt.Println("i=5 is continue", i)
			i++ // 一定要对i进行变量添加变化，不然continue跳过就i值永远为5，进入死循环
			continue
		}
		fmt.Println(i)
		i++
	}

	// 九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", i, j, i*j)
		}
		fmt.Println()
	}
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			if i <= j { // i每次要从 i和j相等开始计算
				fmt.Printf("%d*%d=%2d\t", i, j, i*j)
			}
		}
		fmt.Println()
	}

	// 4.for range 语法，用于便捷地遍历数组、切片、字符串、map 或通道（channel）
	// 遍历数组/切片
	nums := []int64{12,13,14,15}
	for i, v := range nums {
		fmt.Printf("索引:%d,值:%d\n",i,v)
	}
	for _, v := range nums { // 忽略索引
    fmt.Println(v)
	}
	for i := range nums { // 只要索引
    fmt.Println(i)
	}

	// 遍历map
	// m := map[string]interface{} { // interface{}表示接收任何值 或者使用在1.18引入的泛型any
	m := map[string]any {
		"name": "tom",
		"age": 15,
	}
	for k,v := range m {
		fmt.Printf("键 %s, 值 %s\n", k, v)
	}
	// 遍历字符串
	str := "hello, 世界！"
	for idx, r := range str {
		fmt.Printf("idx：%d 字符：%c\n",idx, r) // 索引不是连续的，因为中文字符占3个字节
	}
	// 从通道接收
	

	// for i := 0; ; i++ {
  //   fmt.Println("Value of i is now:", i)
	// }
	// for i := 0; i < 3; {
  //   fmt.Println("Value of i:", i)
	// }
	s := ""
	for s != "aaaaa" {
		fmt.Println("Value of s:", s) // 依次输出 "" a aa aaa aaaa
		s = s + "a"
	}

	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j, s = i+1, j+1, s+"a" {
		fmt.Println("Value of i, j, s:", i, j, s)
	}
}
