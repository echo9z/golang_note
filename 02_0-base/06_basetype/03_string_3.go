package main

import (
	"fmt"
	"strings"
)

func main() {
	// 一.判断字符串
	// 判断字符串s 是否以 prefix开头
	fmt.Println(strings.HasPrefix("hello ok", "He")) // false
	// 判断字符串s 是否以 suffix 结尾
	fmt.Println(strings.HasSuffix("hello ok", "ok")) // true

	// 二.包含判断
	fmt.Println(strings.Contains("hello 世界", "世界"))     // true  是否包含子串
	fmt.Println(strings.ContainsAny("hello", "abceal")) // true 是否包含任意一个字符
	fmt.Println(strings.ContainsRune("hello", 'e'))     // 是否包含指定 rune

	// 三.查找字符位置
	fmt.Println(strings.Index("hello a你好", "a你"))    // 7 查找字符串第一次出现的位置，未找到返回 -1
	fmt.Println(strings.LastIndex("hello a你好", "好")) // 10 查找字符串，最后一次出现的位置，未找到返回 -1
	// 返回 s 中字符中任意 Unicode 码点的第一个字符索引
	fmt.Println(strings.IndexAny("hello a你好", "abced")) // 索引1 取到e字符
	// IndexByte只能查找 ASCII 字符（0-255 范围内的单个字符）
	fmt.Println(strings.IndexByte("hello a你好", 'a')) // 索引6 获取字节第一次出现的位置
	fmt.Println(strings.IndexRune("hello a你好", '你')) // 索引7 rune 第一次出现的位置

	// 四.字符串替换
	// Replace(s, old, new, n) 将s字符串中 old字符，替换为new字符，替换前第n个匹配到字符
	fmt.Println(strings.Replace("hello hello ok", "ll", "LL", 1)) // heLLo hello ok
	fmt.Println(strings.ReplaceAll("hello hello ok", "ll", "LL")) // 全部替换匹配到字符串 heLLo heLLo ok

	// NewReplacer(old1, new1, old2, new2):从旧的、新的字符串对列表中返回一个新的 Replacer函数。
	// 替换按照它们在目标字符串中出现的顺序执行，不会重叠匹配。旧的字符串比较是按参数顺序完成的。
	r := strings.NewReplacer("<", "&lt;", ">", "&gt;") // 返回新的replace函数，替换顺序为<替换为&lt; >替换为&gt;
	fmt.Println(r.Replace("This is <b>HTML</b>!"))     // This is &lt;b&gt;HTML&lt;/b&gt;!

	// 五.字符串分割与连接
	// 分割
	// SplitN 按分隔符分割，指定分割次数
	str2 := "a,b,c,d,e"
	// n 表示每次分割的字串数量
	// n > 0：最多返回 n 个子字符串；最后一个子字符串将是未分割的剩余部分；
	var par []string = strings.SplitN(str2, ",", 2) // 第一次匹配到的","索引时，进行分割成两个字符串部分
	fmt.Println(par)                                // [a b,c,d,e] a 和 b,c,d,e 两个字符串元素

	for idx, c := range par {
		fmt.Printf("index: %d %s \n", idx, c)
	}

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

	// 六、字符串连接
	// 1.运算符
	s1, s2 := "hello", "world"
	res := s1 + " " + s2
	fmt.Println(res)
	optimizedConcat := func() string {
		var builder strings.Builder
		// 如果指定最终字符的大小，可以预分配容量
		builder.Grow(1024) // 1024个字节
		for i := 0; i < 1000; i++ {
			builder.WriteString("a") // 写入字符串
		}
		return builder.String()
	}
	fmt.Println(optimizedConcat())

	// 使用 fmt.Sprintf（格式化字符串）
	name, age := "tom", 100
	result := fmt.Sprintf("姓名：%s 年龄：%d", name, age) // 不适合循环或大量字符串拼接
	fmt.Println(result)

	// 2.使用strings.Join（连接切片）
	var parts []string = []string{"a", "b", "c", "d"}
	res2 := strings.Join(parts, "-") // 将切片中元素连接
	fmt.Println(res2) // a-b-c-d

	// 3.strings.Repeat() 重复字符串
	strRes := "ba" + strings.Repeat("na", 2)
	fmt.Println(strRes) // banana

	// 七、大小写转换
	fmt.Printf("转大写：%s\n", strings.ToUpper("hello")) // 转大写：HELLO
	fmt.Printf("转小写：%s\n", strings.ToUpper("HeLLo")) // 转小写：HELLO
	fmt.Printf("首字母大写：%s\n", strings.ToUpper("HeLLo")) // 首字母大写：HELLO
	fmt.Printf("Unicode格式标题大写：%s\n", strings.ToTitle("HeLLo")) // Unicode格式标题大写：HELLO

	// 八、比较与计数
	fmt.Println(strings.Count("cheese", "e")) // 3 统计字符串出现的次数
	// compare(n1,n2) n1>n2返回1，n1=n2返回0，n1<n2返回-1
	fmt.Println(strings.Compare("a", "b")) // -1
	fmt.Println(strings.EqualFold("HellO", "hEllO")) // true - 忽略大小写比较

	// 九、修剪操作字符
	// 去除首位空格
	fmt.Println(strings.TrimSpace("  hello ok  ")) // hello ok
	fmt.Println(strings.Trim("!!!hello!ok!!", "!")) // hello!ok 去除指定的字符串
	fmt.Println(strings.TrimLeft("!!!hello", "!")) // 去除左边的字符串
	fmt.Println(strings.TrimRight("hello!!!", "!")) // 去除右边的字符串
	fmt.Println(strings.TrimPrefix("hello", "he")) // 去除匹配到的前缀字符
	fmt.Println(strings.TrimSuffix("hello", "ol")) // 去除匹配到的后缀字符

	// 十、高效构建字符串
	var builderStr strings.Builder
	builderStr.WriteString("hello")
	builderStr.WriteString(" ")
	builderStr.WriteString("world")
	fmt.Println(builderStr.String()) // hello world

	
}

func rangStr(strSlice []string) {
	for idx, val := range strSlice {
		fmt.Printf("index %d：%s\n", idx, val)
	}
}
