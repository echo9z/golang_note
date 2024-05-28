package test

import (
	test "basegra/test/internal/ser"
	ser "basegra/test/ser"
	"fmt"
)

func TheTest() {
	fmt.Println("这里时test函数")
	test.IntSer()
	ser.Ser()
}
