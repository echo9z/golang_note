package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	// 通过 runtime 包在运行时获取所在的操作系统类型，以及如何通过 os 包中的函数 os.Getenv() 来获取环境变量中的值
	var goos string = runtime.GOOS
	fmt.Println("系统os对象：%s\n", goos)
	var path string = os.Getenv("PATH")
	fmt.Println("环境变量PATH：%s\n", path)
}
