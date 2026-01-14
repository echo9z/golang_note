package main

import (
	"fmt"
	"strings"
)

func main() {
	// 字符串分割
	// Split：根据指定的分隔符将字符串分割成多个子串，并返回一个字符串切片。
	str := "a,b,c,d,e"
	// 按照指定字符进行分割
	var res []string = strings.Split(str, ",")
	// 返回字符串切片
	fmt.Println(res) // [a b c d e]
	
	// 通过切片遍历
	for idx, v := range res {
		
	}
}
