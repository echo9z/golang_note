package utils

import "fmt"
import t "golang/notes/02-basegra/test"
import ser "golang/notes/02-basegra/test/ser"

func GoRun() {
	fmt.Println("go run")
	t.TheTest()
	ser.Ser()
	// ser.IntSer() 无法引用 test下internal的私有函数或类型
}
