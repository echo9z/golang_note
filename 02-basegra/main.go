package main

// import "golang/notes/02-basegra/utils" // 默认导入
// import utils "golang/notes/02-basegra/utils" // 别名导入
import (
	"fmt"
	"math"
	u "golang/notes/02-basegra/utils"
	_ "golang/notes/02-basegra/test" // 下划线表示匿名导入
)
func main() {
	fmt.Println("多包导入", math.Pi)
	u.GoRun() // 别名调用utils包下GoRun，函数
	u.SayHello()
}
