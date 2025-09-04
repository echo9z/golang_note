package driver

import "fmt"

var idCard int

// init函数会在包初始化时自动调用，且每个包可以有多个init函数
func init() {
	fmt.Println("init: driver初始化")
}
