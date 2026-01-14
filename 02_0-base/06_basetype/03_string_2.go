package main

import (
	"fmt"
	"strings"
)

func main() {
	// 字符串分割
	// Split：根据指定的分隔符将字符串分割成多个子串，并返回一个字符串切片。
	// SplitN：根据指定的分隔符将字符串分割成多个子串，但是可以指定分割的次数（返回的子串个数最多为n）。
	// SplitAfter：根据指定的分隔符将字符串分割成多个子串，但是分隔符会包含在子串中。
	// SplitAfterN：类似于SplitAfter，但是可以指定分割的次数。
	// Fields：根据空白字符（空格、制表符、换行符等）分割字符串。
	// FieldsFunc：根据传入的函数来决定分割符。

	// Split
	str := "a,b,c,d,e"
	// 按照指定字符进行分割
	var res []string = strings.Split(str, ",")
	// 返回字符串切片
	fmt.Println(res) // [a b c d e]
	
	// 通过切片遍历
	for idx, val := range res {
		fmt.Printf("index %d：%s\n", idx, val)
	}

	
}
