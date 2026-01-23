package main

import (
	"fmt"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

func main() {
	// 字符串分割
	// Split：根据指定的分隔符将字符串分割成多个子串，并返回一个字符串切片。
	// SplitN：根据指定的分隔符将字符串分割成多个子串，但是可以指定分割的次数（返回的子串个数最多为n）。
	// SplitAfter：根据指定的分隔符将字符串分割成多个子串，但是分隔符会包含在子串中。
	// SplitAfterN：类似于SplitAfter，但是可以指定分割的次数。
	// Fields：根据空白字符（空格、制表符、换行符等）分割字符串。
	// FieldsFunc：根据传入的函数来决定分割符。

	// 1.Split 按分隔符分割
	str := "a,b,c,d,e"
	// 按照指定字符进行分割
	var res []string = strings.Split(str, ",")
	// 返回字符串切片
	fmt.Println(res) // [a b c d e]

	// 通过切片遍历
	for idx, val := range res {
		fmt.Printf("index %d：%s\n", idx, val)
	}

	// 2.SplitN 按分隔符分割，指定分割次数
	str2 := "a,b,c,d,e"
	// n 表示每次分割的字串数量
	// n > 0：最多返回 n 个子字符串；最后一个子字符串将是未分割的剩余部分；
	var par []string = strings.SplitN(str2, ",", 2) // 第一次匹配到的","索引时，进行分割成两个字符串部分
	fmt.Println(par) // [a b,c,d,e] a 和 b,c,d,e 两个字符串元素
	rangStr(par)

	// n=0 返回nil 空值 （零个子字符串）
	var par2 []string = strings.SplitN(str2, ",", 0)
	fmt.Println(par2)

	// n=-1 无限制分割（等同于 Split）
	var par3 []string = strings.SplitN(str2, ",", -1)
	fmt.Println(par3)

	// 3.SplitAfter 分割后保留分隔符
	str3 := "apple|banana|orange"
	var ful []string = strings.SplitAfter(str3, "|")
	fmt.Println(ful) // [apple| banana| orange] 每个元素都包含分隔符

	// 4.SplitAfterN() - 限制分割次数并保留分隔符
	str4 := "apple|banana|orange|grape|pull"
	var ful2 []string = strings.SplitAfterN(str4, "|", 2) // [apple| banana|orange|grape|pull]
	fmt.Println(ful2)
	rangStr(ful2)

	// 5.Fields 按照空白字符（空格、制表符、换行符等）分割字符串
	// 自动识别连续的空白字符作为一个分隔符
	// 自动去除字符串首尾的空白字符
	// 返回的切片中不包含空字符串
	str5 := "  apple  banana   orange  \n grape  \t kiwi \n 123 "
	fie := strings.Fields(str5)
	fmt.Println(fie) // [apple banana orange grape kiwi 123]

	// 6.FieldsFunc(str,func) 根据自定义函数来决定分割规则
	// func 传入一个函数，参数为rune字符类型，返回 bool
	// 返回 true: 该字符作为分隔符
	// 返回 false: 该字符不是分隔符
	str6 := "abc, 455;ufg|kill:open"
	// 按多个字符分割
	var fie2 []string = strings.FieldsFunc(str6, func(r rune) bool {
		fmt.Printf("char: %c\n", r)
		// 当前字符为, ; | : 时，进行分割
		return r == ',' || r == ';' || r == '|' || r == ':'
	})
	fmt.Println(fie2) // abc  455 ufg kill open

	str7 := "Hello, World! 123 Go."
	// 只按字母分割（保留非字母字符）
	var fie3 []string = strings.FieldsFunc(str7, func(r rune) bool {
		// 去unicode值不在a-z 或者A-Z之间，即将a-z A-Z的字母全部截取，只保留字符[,  ! 123  .]
		return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
	})
	fmt.Println(fie3)

	// 分割字符串场景
	splitEg()

	// 使用StringBuilder缓冲区进行字符串分割
	var sBuilder []string = SplitWriteBuilder("a,b,,v,cd")
	fmt.Println(sBuilder)

	// strF
	strF()
}
func splitEg() {
	// 案例一：文件分割
	path := "/home/user/documents/file.txt"
	// Split 会在路径中最后一个分隔符之后立即拆分，将其分为目录和文件名部分。如果路径中没有分隔符，Split 会返回空的目录，并将文件设置为路径本身。
	dir, file := filepath.Split(path)
	fmt.Printf("目录：%s 文件名：%s \n", dir, file)

	// 分割文件扩展名
	ext := filepath.Ext(file) // .txt
	// 字符串切片语法：s[start:end] 截取从 start 到 end（不包括 end）的子串
	name := file[:len(file)-len(ext)] // file[:8-4] 截取从索引0到4的子串，即 "file"
	fmt.Println("file name：", name)

	// 案例二：URL 路径分割
	url := "https://example.com/path/to/resource?query=value"
	parts := strings.SplitN(url, "://", 2) // 在://进行分割 分割成两部分
	if len(parts) == 2 {
		protocol := parts[0]
		rest := parts[1]
		fmt.Printf("协议：%s，路径：%s \n", protocol, rest)
		
		// 分割域名 和 路径
		var domainPath []string = strings.SplitN(rest, "/", 2) // example.com 和 /path/to/resource?query=value
		domain := domainPath[0]
		path := "/" + domainPath[1]
		fmt.Printf("域名：%s，路径：%s", domain, path)
	}
}

func rangStr(strSlice []string) {
	for idx, val := range strSlice {
		fmt.Printf("index %d：%s\n", idx, val)
	}
}

// 简易版 高效拼接例子
func SplitWriteBuilder(str string) []string{
	var result []string // 存放空的字符串切片
	var current strings.Builder // 字符串缓冲区

	// 遍历字符串
	for _, char := range str {
		if char == ',' {
			if current.Len() > 0 {
				// 如果字符为 ，时，缓冲区还又字符时，将普通字符追加到result切片中
				result = append(result, current.String())
				current.Reset() // 清空缓冲区
			}
		} else {
			current.WriteRune(char) // 为普通字符时追加到缓冲区中
		}
	}

	// 当遍历完成所有字符时，需要将最后一个字符所在的缓冲区内容，追加到切片中
	if current.Len() > 0 {
		result = append(result, current.String())
	}
	return result
}

// 输入字符串: "a,,b"
// 遍历过程:
// ┌─────┬──────┬─────────────────────────────────┬──────────────┐
// │ 步  │ 字符 │      current 状态               │  result 状态 │
// ├─────┼──────┼─────────────────────────────────┼──────────────┤
// │ 1   │ 'a'  │ 写入 'a' → "a"                 │ []           │
// │ 2   │ ','  │ 发现分隔符，current有内容       │              │
// │     │      │ 将 "a" 加入 result              │ ["a"]        │
// │     │      │ 清空 current → ""               │              │
// │ 3   │ ','  │ 发现分隔符，但current为空       │              │
// │     │      │ 不添加任何内容                  │ ["a"]        │
// │ 4   │ 'b'  │ 写入 'b' → "b"                 │ ["a"]        │
// │ 5   │ 结束 │ 将最后的 "b" 加入 result        │ ["a", "b"]   │
// └─────┴──────┴─────────────────────────────────┴──────────────┘
// 最终结果: ["a", "b"]

// 创建一个用于统计字节和字符（rune）的程序，并对字符串 asSASA ddd dsjkdsjs dk 进行分析，
// 然后再分析 asSASA ddd dsjkdsjsこん dk，最后解释两者不同的原因（提示：使用 unicode/utf8 包）。
func strF() {
	str1 := "asSASA ddd dsjkdsjs dk"
	str2 := "asSASA ddd dsjkdsjsこん dk"

	var byteSize []byte = []byte(str1)
	var byteSize2 []byte = []byte(str2)
	
	fmt.Println(len(str1) == len(byteSize))
	fmt.Println(len(str2) == len(byteSize2))

	analyStr := func (s string){
		fmt.Printf("字符串: %q\n", s)
		fmt.Printf("字节数len: %d\n", len(s))
		fmt.Printf("字符数: %d\n", utf8.RuneCountInString(s))

		for len(s)>0 { // 只要字符串不为空就继续循环
			// utf8.DecodeRuneInString 逐个解析 UTF-8 字符串中的字符，并打印每个字符及其占用的字节数
			r, size := utf8.DecodeRuneInString(s)
			fmt.Printf("字符 %c， size：%d\n", r, size)
			s = s[size:] // 从第size个字节开始截取到末尾，比如第一个为a，size为1，就从1截取到
		}
	}
	analyStr(str1)
	analyStr(str2)
}