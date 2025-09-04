package main

import "fmt"

func main() {
	// 1.没有专门的char字符，只能使用byte或rune类型来表示字符
	// byte是uint8的别名，用于表示 ASCII 字符（单字节字符）
	var c1 byte = 'A' // 65
	var c2 byte = '0' // 48
	fmt.Println("c1:", c1)
	fmt.Println("c2:", c2)

	var c3 byte = 98          // 98
	fmt.Printf("c3 %d\n", c3) // 98
	fmt.Printf("c3 %c\n", c3) // 输出字符b

	//var c3 byte = '中' // ‘中’的unicode码为20013，byte为uint8，取值范围为0-255
	//fmt.Println("c3:", c3)

	//字符可以和整型进行运算
	c4 := 'A' + 1             // 'A'的unicode码为65，65+1=66
	fmt.Printf("c4=%c", c4)   // 输出66
	fmt.Printf("c4=%d\n", c4) // 输出B

	// 2.rune是int32的别名，Unicode 字符（支持中文、表情等）
	var r1 rune = '你' // 20320
	fmt.Printf("r1=%c", r1)
	fmt.Printf("r1=%d\n", r1)

	var f1 rune = '🚀'
	fmt.Printf("f1=%c\n", f1) // 输出火箭

	// 字符串
	var str1 string = "这是一个普通字符串\nabcd\t123\\zyx"
	fmt.Printf("str1=%s\n", str1)

	// 2.原生字符串
	var str string = `这是一个原生字符串，换行
	  tab缩进，\t制表符但是无效,换行
	  "这是一个普通字符串"
	
	  结束
	`
	fmt.Println(str)
}
