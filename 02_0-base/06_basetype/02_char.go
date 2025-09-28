package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	// 1.没有专门的char字符，只能使用byte或rune类型来表示字符
	// byte是uint8的别名，用于表示 ASCII 字符（单字节字符）
	var c1 byte = 'A' // 65
	//  ASCII 码表中，A 的值是 65，而使用 16 进制表示则为 41，所以下面的写法是等效的：
	var c2 byte = 0x41
	var c5 byte = '0' // 48

	fmt.Println("c1:", c1)
	fmt.Printf("%c %d %X\n", c1, c1, c1) // 输出字符A 和对应的整数65 和对应的16进制41

	fmt.Printf("c2: %c\n", c2)
	fmt.Println("c4:", c5)

	var c3 byte = 98          // 98
	fmt.Printf("c3 %d\n", c3) // 98
	fmt.Printf("c3 %c\n", c3) // 输出字符b

	//var c3 byte = '中' // ‘中’的unicode码为20013，byte为uint8，取值范围为0-255
	//fmt.Println("c3:", c3)

	//字符可以和整型进行运算
	c4 := 'A' + 1             // 'A'的unicode码为65，65+1=66
	fmt.Printf("c4=%c\t", c4) // 输出66
	fmt.Printf("c4=%d\n", c4) // 输出B

	// 2.rune是int32的别名，Unicode 字符（支持中文、表情等）
	var r1 rune = '你'         // 20320
	fmt.Printf("r1=%c\t", r1) // 输出字符
	fmt.Printf("r1=%d\t", r1) // 输出值
	fmt.Printf("r1=%U\n", r1) // 输出Unicode码

	var f1 rune = '🚀'
	fmt.Printf("f1=%c\n", f1) // 输出火箭

	// 处理字符串中的得字符
	str := "hello,你好 yeah"

	// 第一种：通过索引访问字符串中的每个字节
	for i := 0; i < len(str); i++ {
		fmt.Printf("索引%d %c ", i, str[i]) // 按字节访问，中文会乱码
	}
	fmt.Println("------")
	// 第二种：按照rune进行遍历
	// 用range遍历字符串时，Go会自动将UTF-8编码的字节序列解码为Unicode字符（rune），而不是简单地逐字节遍历。
	for i, r := range str {
		fmt.Printf("索引%d %c\n ", i, r) // 按字符访问，r是字符的索引
	}

	// 第三种：将字符串转换为rune切片，切片与数组类似，但更灵活。切片是不定长的，切片在容量不够时会自行扩容。
	var runeSline []rune = []rune(str) // []rune就是一个rune类型的切片，切片通常使用make()函数来创建
	for i, r := range runeSline {
		fmt.Printf("索引%d %c\n ", i, r) // 按字符访问，r是字符的索引
	}

	// unicode包中，有一些针对测试字符
	fmt.Println(unicode.IsLetter('a')) // 判断是否为字母
	fmt.Println(unicode.IsDigit('1'))  // 判断是否为数字
	fmt.Println(unicode.IsSpace(' '))  // 判断是否为空白字符
	fmt.Println(unicode.IsUpper('a'))  // 判断是否为大写字母
	fmt.Println(unicode.ToLower('A'))  // 将字符转化为小写
	fmt.Println(unicode.ToUpper('a'))  // 将字符转化为大写

	// 字符与字符串的关系
	var str2 string = "hello你好"
	// 获取每个字符
	for i, r := range str2 {
		fmt.Printf("字符索引%d %c\n", i, r) // 按字符访问，r是字符的索引
	}

	// 输出每个字节
	for i := 0; i < len(str2); i++ {
		fmt.Printf("字节[%d]: %d ('%c')\n", i, str2[i], str2[i])
	}

	// 字符串本质是 UTF-8 编码的字节序列
	fmt.Printf("字符串：%s\n", str2)
	fmt.Printf("字节长度：%d\n", len(str2))                   // 字节5个英文+两个中文，每个中文3个字节，共11个字节
	fmt.Printf("字符长度：%d\n", utf8.RuneCountInString(str)) // 字符串中有多少个字符

}
