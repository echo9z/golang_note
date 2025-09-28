package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// 1.普通字符串
	var str1 string = "这是一个普通字符串\nabcd\t123\\zyx"
	fmt.Printf("str1=%s\n", str1)

	// 2.原生字符串
	// 由反引号表示，不支持转义，支持多行书写，原生字符串里面所有的字符都会原封不动的输出，包括换行和缩进。
	var str string = `这是一个原生字符串，换行
	  tab缩进，\t制表符但是无效,换行
	  "这是一个普通字符串"
	
	  结束
	`
	fmt.Println(str)

	// Go 字符串的本质：不可变的字节切片，创建后不能修改
	str2 := "hello世界"
	//str2[0] = 'a' // 编译报错
	// Ascii 类型可以通过 来直接获取，本质上是使用byte字节类型存放，但是如果是中文、阿拉伯文字等类型，则需要使用rune类型存放即Unicode
	// byte：Go 字符串的基本单位。一个字节8位。
	// rune：Go 语言中的 rune 是 int32 类型的别名。它代表一个单一的 Unicode 码点。
	// 像 'A' 这样的字符可以用一个字节表示，但像 '中' 这样的字符在 UTF-8 中需要三个字节。 rune 可以表示任意一个。

	fmt.Printf("str2字节长度：%d\n", len(str2)) // 字节长度

	fmt.Printf("str2字符长度：%d\n", utf8.RuneCountInString(str2)) // 使用utf8包，通过Rune统计字符长度
	fmt.Printf("str2字符长度：%d\n", len([]rune(str2)))            // 通过切片返回字符长度
}
