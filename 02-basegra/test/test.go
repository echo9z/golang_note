package test

import (
	"fmt"
	test "golang/notes/02-basegra/test/internal/ser"
	ser "golang/notes/02-basegra/test/ser"
)

func TheTest() {
	fmt.Println("这里时test函数")
	test.IntSer()
	ser.Ser()
}