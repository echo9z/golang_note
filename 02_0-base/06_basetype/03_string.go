package main

import (
	"fmt"
	"strings"
	"unicode"
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

	// Go中的字符串默认使用UTF-8编码，UTF-8是一种变长编码方式：
	ss1 := "A"                            // ascii字符 1个字节
	ss2 := "中"                            // 中文：3个字节
	ss3 := "𝄞"                            // 音乐符号：4个字节
	ss4 := "é"                            //带重音的拉丁字符：2字节
	fmt.Printf("ss1所占的字节：%d\n", len(ss1)) // 1
	fmt.Printf("ss2所占的字节：%d\n", len(ss2)) // 3
	fmt.Printf("ss3所占的字节：%d\n", len(ss3)) // 4
	fmt.Printf("ss4所占的字节：%d\n", len(ss4)) // 2

	ss5 := "你好中AAAAi"
	// 字符串转换
	bytes := []byte(ss5) // 字符串转换为字节切片 存放字节码
	runes := []rune(ss5) // 转换为rune字符切片  存放unicode码
	fmt.Println("bytes and runes", bytes, runes)

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
	fmt.Println("------range------")

	// 2.range遍历字符串，按字符遍历，不会乱码
	for i, char := range str5 {
		fmt.Printf("索引 %d, 字节 %c, 字符 %c\n", i, char, char)
	}

	// 索引 0, 字节 H, 字符 H
	// 索引 1, 字节 e, 字符 e
	// 索引 2, 字节 l, 字符 l
	// 索引 3, 字节 l, 字符 l
	// 索引 4, 字节 o, 字符 o
	// 索引 5, 字节 世, 字符 世
	// 索引 8, 字节 界, 字符 界
	// 索引 11, 字节 م, 字符 م
	// 索引 13, 字节 ر, 字符 ر
	// 索引 15, 字节 ح, 字符 ح
	// 索引 17, 字节 ب, 字符 ب
	// 索引 19, 字节 ا, 字符 ا
	// 索引 21, 字节 😀, 字符 😀
	// 注意：索引从5跳到8，中文字符占3个字节

	// 3.rune 字符串转成 rune[] 切片遍历
	var strRuneSlice []rune = []rune(str5) //
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

	// 字符统计案例
	fmt.Println("==字符统计案例==")
	countChar()
	reverseString()
	FindReplace()

	// 7.字符串比较
	// 基础运算符比较
	fruit1 := "apply"
	fruit2 := "banana"
	fruit3 := "apply"

	fmt.Println(fruit1 == fruit2) // false
	fmt.Println(fruit1 == fruit3) // true
	fmt.Println(fruit1 != fruit3) // false

	// 字典顺序比较规则：
	fmt.Println(fruit1 < fruit2)  // true （"apple" < "banana"）
	fmt.Println(fruit1 <= fruit3) // true
	fmt.Println(fruit1 >= fruit3) // true
	// 字典顺序比较规则：
	// 从第一个字符开始逐个比较 Unicode 码点
	// 所有大写字母排在小写字母之前
	// 数字排在字母之前
	// 空格排在可打印字符之前
	fmt.Println("A" < "a")         // true 65<97
	fmt.Println("1" < "A")         // true 49<65
	fmt.Println(" " < "1")         // true 49<65
	fmt.Println("apple" < "Apple") // false97 > 65

	// string包的 Compare(a,b) - 三元比较函数
	fmt.Println(strings.Compare("apply", "banana")) // 第一个参数小于第二个参数，输出为-1
	fmt.Println(strings.Compare("apply", "apply"))  // 两参数相等为0
	fmt.Println(strings.Compare("banana", "apply")) // 第一个参数比第二参数大，输出1
	// strings.Compare 通常不推荐使用，因为直接使用 ==、<、> 运算符更直观且性能更好。

	// strings.EqualFold() - 不区分大小写比较
	fmt.Println("不区分大小写：", strings.EqualFold("gO", "Go")) // true

	// 使用ToLower函数将字符串转换为小写
	fmt.Println(strings.ToLower("gO") == strings.ToLower("Go")) //  但 EqualFold 更高效，因为它不需要分配新字符串

}

// 从性能上看
// == 运算符：最快
// strings.Compare()：稍慢，因为需要额外的函数调用
// strings.EqualFold()：不区分大小写比较的最优选择
// ToLower 后比较：最慢，因为需要分配新字符串

// 统计字符数量
func countChar() {
	text := "hello, 测试! 123"
	// 错误方式：len()返回字节数
	fmt.Printf("字节数：%d\n", len(text))

	// 正确方式1：使用utf8包
	fmt.Printf("字符数：%d\n", utf8.RuneCountInString(text)) // 字符串长度14

	// 正确方式：range遍历一次count++
	count := 0
	// Go 语言中 range 的一个特殊用法，当你不需要使用索引或值时，可以省略它们。
	for range text { // 等价于 for _ range text{count++}
		count++
	}

	// 统计不同的字符
	var letter, digit, spaces, others int
	for _, r := range text {
		switch { // 注意：这里switch的条件表达式是可以省略的，同时再结束case时，自动break，不需要写清楚break
		case unicode.IsLetter(r):
			letter++
		case unicode.IsDigit(r):
			digit++
		case unicode.IsSpace(r):
			spaces++
		default:
			others++
		}
	}
	fmt.Printf("字母:%d, 数字:%d, 空格:%d, 其他:%d\n",
		letter, digit, spaces, others)
}

// 字符串反转
// reverseString
func reverseString() {
	str := "你好，hello ok"

	wrongRevers := func(s string) string {
		// 转换成字节切片
		var byteSlice []byte = []byte(s)
		fmt.Println(byteSlice)
		for i, j := 0, len(byteSlice)-1; i < j; i, j = i+1, j-1 {
			temp := byteSlice[i]
			byteSlice[i] = byteSlice[j]
			byteSlice[j] = temp
		}
		fmt.Println("end", byteSlice)
		return string(byteSlice) // 通过string函数将字节切片转换为字符串
	}
	res := wrongRevers(str)
	fmt.Println("错误反转", res) // 通过将字符串转换字节切片，通过切片数组反转，遇到比如中午出现乱码 ko olleh��･堽�

	// 正确的方式，使用rune字符切片进行数组反转
	// var correctRevers func(string) string = func (s string) string {}
	// correctRevers := func (s string) string {} // 省略方式写法
	correctRevers := func(s string) string {
		var runeSlice []rune = []rune(s)
		for i, j := 0, len(runeSlice)-1; i < j; i, j = i+1, j-1 {
			temp := runeSlice[i]
			runeSlice[i] = runeSlice[j]
			runeSlice[j] = temp
		}
		return string(runeSlice)
	}
	res2 := correctRevers(str)
	fmt.Println("正确反转", res2)
}

// 字符串替换
func FindReplace() {
	s := "Hello 世界, Hello Go!"
	findStr := '世'
	// 查找特定的字符位置
	for i, c := range s {
		if c == findStr {
			fmt.Printf("找到字符 %c 索引位置：%d \n", c, i)
			break
		}
	}

	// 替换特定的字符
	replaceStr := func(str string, findStr, replStr rune) string {
		var runeSlice []rune = []rune(str)
		for i, r := range runeSlice {
			if r == findStr {
				runeSlice[i] = replStr
			}
		}
		return string(runeSlice) // 最后使用string进行自动装箱，转换为字符串
	}
	fmt.Printf("替换后新的字符串 %s\n", replaceStr(s, '世', '时'))
}
