package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"os"
	"runtime"
	"strconv"
)

func main() {
	// 一。if-else结构
	var a int = 10
	var b int = 20
	if a < b {
		fmt.Println("a小于b")
	} else {
		fmt.Println("a大于b")
	}

	// if-else if-else
	var source float64 = 3.1415926
	if source > 3 {
		fmt.Println("大于3")
	} else if source < 3 {
		fmt.Println("大于3")
	} else {
		fmt.Println("等于3")
	}

	var source2 int64 = 98
	var level string
	// 利用了if语句是从上到下的判断的前提，所以代码要更简洁些
	if source2 > 0 && source2 < 60 {
		level = "E"
		fmt.Println("level", level)
	} else if source2 < 70 {
		level = "D"
		fmt.Println("level", level)
	} else if source2 < 80 {
		level = "C"
		fmt.Println("level", level)
	} else if source2 < 90 {
		level = "B"
		fmt.Println("level", level)
	} else if source2 <= 100 {
		level = "A"
		fmt.Println("level", level)
	} else {
		level = "nil"
	}

	// 判断空字符串
	// if s == "" {}
	// 判断长度 if len(s) == 0
	if str := ""; len(str) == 0 {
		fmt.Println("str字符串为空")
	}
	// Go 程序的操作系统类型，可以通过常量 runtime.GOOS 来判断
	if runtime.GOOS == "windows" {
		fmt.Println("windows")
	} else if runtime.GOOS == "linux" {
		fmt.Println("linux")
	}

	// if用于返回一个整数的绝对值
	var Abs func(x int) int = func (x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	fmt.Println(Abs(-15))

	// if 可以包含一个初始化语句（如：给一个变量赋值）
	if val, err := strconv.Atoi("8080"); err != nil{
		fmt.Println("err for", err)
	} else {
		fmt.Println("端口转换的值：",val)
	}

	// 这是测试 err 变量是否包含一个真正的错误（if err != nil）的习惯用法。如果确实存在错误，则会打印相应的错误信息然后通过 return 提前结束函数的执行。
	// 通过 os.Open 方法打开一个名为 name 的只读文件：
	if file,err := os.Open(`c:\Users\11312\Desktop\golang_note\02_0-base\06_basetype\05_time.go`); err != nil {
		fmt.Println(err)
		// return err 在 main() 函数中，main 不能有返回值，所以 return err 会导致编译错误。
		return
	} else {
		fmt.Println(file)
	}

	// 多返回值的函数
	mySqrt := func (f float64) (v float64, b bool, err error) {
		if f < 0 {
			return 0, false, errors.New("传入值小于0")
		}
		return math.Sqrt(f), true, nil
	}
	if v, b, _ := mySqrt(25.0); b {
		fmt.Println("值的平方根：", v)
	}


	// 二。switch结构
	chr := 'A'
	switch chr {
	case 'A':
		fmt.Println("字符A")
	case 'B':
		fmt.Println("字符B")
	default:
		fmt.Println("其他字符")
	}
	// 1.switch不带表达式
	source3 := 85
	// switch (true) { 等价于 switch {
	switch { // 在其他语言不可以省略
	case source3 > 0 && source3 < 60:
		fmt.Println("level E")
	case source3 < 70:
		fmt.Println("level D")
	case source3 < 80:
		fmt.Println("level C")
	case source3 < 90:
		fmt.Println("level B")
	case source3 <= 100:
		fmt.Println("level A")
	default:
		fmt.Println("level nil")
	}
	// 2.通过fallthrough关键字来继续执行相邻的下一个分支
	i := 26
	switch true {
	case i > 10: fallthrough
	case i > 20:
		fmt.Println("i数值", i)
	}
	
	// 3.switch 语句条件语句中，包含一个初始化语句：
	switch v, err :=strconv.Atoi("12"); true{
	case err != nil:
    fmt.Println("转换错误:", err)
	case v < 0:
		fmt.Println("转换的值小于0")
	case v > 0:
		fmt.Println("转换的值大于0")
	default:
    fmt.Println("值为0")
	}

	// label标签通常与goto，break，continue进行使用
	
	// 生成[0.0到1.0)的随机浮点数
	flag := rand.Float64()
	if flag >= 0.5 {
		goto A
		var val float64 = float64(a) + flag
		fmt.Printf("flag:%f, val:%f", a, val)
	} else {
		goto B
		var val float64 = float64(b) + flag
		fmt.Printf("flag:%f, val:%f", b, val)
	}

	A:{
		a := 10
		fmt.Println("a", a)
	}
	B:{
		b := 20
		fmt.Println("b", b)
	}
}

// 使用var 变量为全局变量直接初始化
var prompt = "Enter a digit, e.g. 3 " + "or %s to quit."

// init()是 Go 语言中特殊的保留函数，用于包的初始化。
// 自动执行：程序启动时自动调用，无需手动调用
// 执行时机：在 main() 之前执行
// 参数/返回值：无参数、无返回值
// 数量：每个文件/包可以有多个 init()
// 可见性：必须是小写的 init（不能导出）
func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
		fmt.Println(prompt)
	} else { // Unix-like
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
		fmt.Println(prompt)
	}
}
