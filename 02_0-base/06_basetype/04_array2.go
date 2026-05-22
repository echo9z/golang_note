package main

import "fmt"

// 声明接口
type Animal interface{ Speak(s string) string }

// 结构体
type Dog struct{ Name string }
type Cat struct{ Name string }

// 实现接口方法（接收者为指针）
func (d *Dog) Speak(s string) string { return d.Name + "：" + s }
func (c *Cat) Speak(s string) string { return c.Name + "：" + s }

func main() {
	// 写法1：先创建值，再取地址
	// dog := Dog{Name: "旺财"}  // dog 是 Dog 类型（值）
	// d1 := &dog              // d1 是 *Dog 类型（指针）
	
	// 写法2：创建时直接取地址（简洁写法）
	d1 := &Dog{Name: "旺财"}  //  d1 直接就是 *Dog 类型（指针）
	c1 := &Cat{Name: "小黑"}

	// var dd1 Animal = d1
	// var cc1 Animal = c1

	// 接口指针数组：元素是 Animal 接口类型，存储的是实现了该接口的指针
	var animals [2]Animal = [2]Animal{d1, c1}
	// var animals [2]*Animal = [2]*Animal{&dd1, &cc1}

	// 多态：统一调用，各自表现不同行为
	for _, a := range animals {
		fmt.Println(a.Speak("hello"))
		// fmt.Println((*a).Speak("hello"))
	}
	// 旺财：hello
	// 小黑：hello

	// 也可以用切片，更灵活，推荐
	pets := []Animal{d1, c1, &Cat{Name: "喵喵"}}
	for _, p := range pets {
		fmt.Println(p.Speak("嗨"))
	}
	// 旺财：嗨
	// 小黑：嗨
	// 咪咪：嗨

// 错误做法（真正的接口指针数组）：var arr [3]*Animal —— 极度不推荐，没有任何意义，会让你在调用方法时极其痛苦。
// 正确做法（接口数组/切片）：var arr []Animal —— 推荐。你可以在这个数组里放入 *Dog、*Cat（具体结构体的指针）。
}
