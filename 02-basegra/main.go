package main

// import "golang/notes/02-basegra/utils" // 默认导入
// import utils "golang/notes/02-basegra/utils" // 别名导入
import (
	base "basegra/base"
	_ "basegra/test" // 下划线表示匿名导入
	u "basegra/utils"
	"fmt"
	"math"
)

/**
* 这里启动时的主函数
 */
func main() {
	fmt.Println("多包导入", math.Pi)
	u.GoRun() // 别名调用utils包下GoRun，函数
	u.SayHello()

	fmt.Println("运算符")
	// 运算符
	base.Sum()          // 算数运算符
	base.Relationship() // 关系运算符
	base.Logic()        // 逻辑运算符
	base.Bitwise()      // 位运算
	base.Ass()          // 复合运算 num <<= 2等驾驭 num=num<<2
	base.Pointerfnc()   // 指针运算符

	fmt.Println("变量")
	// 变量
	base.Variable()
	base.VariableStat()

	aa, bb := 15, 20
	fmt.Println(base.Swap(aa, bb))

	var arr [10]int = base.IsIota()
	fmt.Println(arr)

	// 数据类型
	fmt.Println("数据类型")
	base.NumberDataType()

	// 格式化
	fmt.Println("格式化输出")
	base.Format()

	u.Coding("hello")

}
