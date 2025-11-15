package main

import (
	"fmt"
	"strings"
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

	// 3.Go 字符串的本质：不可变的字节切片，创建后不能修改
	str2 := "hello世界"
	// str2[0] = 'a' // 编译报错
	// 需要修改字符串，可以转换为[]byte或[]rune切片，修改后再转换为字符串
	byteSlice := []byte(str2)
	byteSlice[0] = 'H'
	str2 = string(byteSlice) // 将byte切片通过 string() 转换为字符串，覆盖原来的字符串
	fmt.Println("1修改后str2:", str2)

	runeSlice := []rune(str2)
	runeSlice[6] = '大'
	str2 = string(runeSlice)
	fmt.Println("2修改后str2:", str2)

	// Ascii 类型可以通过 来直接获取，本质上是使用byte字节类型存放，但是如果是中文、阿拉伯文字等类型，则需要使用rune类型存放即Unicode
	// byte：Go 字符串的基本单位。一个字节8位。
	// rune：Go 语言中的 rune 是 int32 类型的别名。它代表一个单一的 Unicode 码点。
	// 像 'A' 这样的字符可以用一个字节表示，但像 '中' 这样的字符在 UTF-8 中需要三个字节。 rune 可以表示任意一个。

	// len() 返回的是字节数，不是字符数
	fmt.Printf("str2字节长度：%d\n", len(str2)) // 字节长度
	// 方法1：使用 utf8.RuneCountInString (推荐)
	fmt.Printf("str2字符长度：%d\n", utf8.RuneCountInString(str2)) // 使用utf8包，通过Rune统计符串中的 Unicode 字符数量
	// 方法2：转换为 rune 切片
	fmt.Printf("str2字符长度：%d\n", len([]rune(str2))) // 通过切片返回字符长度。将字符串 str2 转换为 rune切片，rune 在 Go 中代表单个 Unicode 字符这个转换会正确处理多字节字符（如中文）， 再获取 rune 切片的长度

	// 4.字符串通过索引进行访问
	var str3 string = "Hello世界"
	fmt.Printf("字节长度：%d \n", len(str3))                    // 11
	fmt.Printf("字节长度：%d \n", utf8.RuneCountInString(str3)) // 7
	// 通过索引访问是字节，
	fmt.Printf("s[0]=%c 字节值：%d\n", str3[0], str3[0])
	fmt.Printf("s[5]=%c 字节值：%d\n", str3[5], str3[5]) // 访问”世” s[5]=ä 字节值：228，访问包含 3 个字节中的 1 个
	// 通过切片方式,截取多个字节访问范文，但存在可能乱码问题
	fmt.Printf("s[0:6]=%s\n", str3[0:6]) // 左包右闭0~5 s[0:6]=Hello�
	fmt.Printf("s[0:8]=%s\n", str3[0:8]) // s[0:8]=Hello世
	fmt.Printf("s[5:8]=%s\n", str3[5:8]) // 4~7 s[5:8]=世

	// 5.字符串遍历
	str5 := "Hello世界مرحبا😀"
	// 1.通过字节遍历 对于非ASCII字符会显示乱码
	for i := 0; i < len(str5); i++ {
		fmt.Printf("索引 %d, 字节 %c, 字符 %c\n", i, str5[i], str5[i])
	}
	
	// 2.range遍历字符串，按字符遍历，不会乱码
	for i, char := range str5 {
		fmt.Printf("索引 %d, 字节 %c, 字符 %c\n", i, char, char)
	}

	// 3.rune 字符串转成 rune[] 切片遍历
	var strRuneSlice []rune = []rune(str5)
	for i, char := range strRuneSlice {
		fmt.Printf("索引 %d, 字符 %c, Unicode %U\n", i, char, char)
	}

	// 6.字符串操作
	// 方式1：字符串拼接 +
	s1 := "hello"
	s2 := "你好"
	s3 := s1 + s2
	fmt.Printf("字符串拼接 %s \n", s3)

	// 方式2：Sprintf方式  Sprint 使用其操作数的默认格式进行格式化，并返回结果字符串。当操作数既不是字符串时，会在操作数之间添加空格
	s4 := fmt.Sprintf("%s %s \n", "yes", "no")
	fmt.Println("Fprintf拼接", s4)

	// 方式3：使用strings包中join函数strings.Join（拼接切片）
	var parts []string = []string{"yellow", "green"} // 定义切片，，不定义长度，，切实长度是可变的
	s5 := strings.Join(parts, ",")                   // Join两个参数 切片,拼接的字符
	fmt.Println("strings.Join:", s5)

	// 方式4：strings.Builder（高性能，循环拼接推荐）
	var builderStr strings.Builder // 声明一个高效的字符串类型
	for i := 0; i < 10; i++ {
		builderStr.WriteString("cc") // 高效追加字符串到Builder中
	}
	var res string = builderStr.String()
	fmt.Println("strings.Builder:", res)

}
