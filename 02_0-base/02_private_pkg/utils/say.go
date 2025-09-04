package utils

import "fmt"

// 首字母大写，可以被包外访问
func SayHello() {
	fmt.Println("Hello")
}

// 首字母小写，外界无法访问
func sayHello() {
	fmt.Println("Hello")
}
