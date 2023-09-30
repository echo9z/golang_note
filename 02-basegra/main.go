package main

// import "golang/notes/02-basegra/utils" // 默认导入
// import utils "golang/notes/02-basegra/utils" // 别名导入
import (
	"fmt"
	"math"
	u "golang/notes/02-basegra/utils"
)
func main() {
	fmt.Println("多包导入", math.Pi)
	u.GoRun()
}
