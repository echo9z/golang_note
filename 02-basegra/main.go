package main

// import "golang/notes/02-basegra/utils" // 默认导入
// import utils "golang/notes/02-basegra/utils" // 别名导入
import (
	"fmt"
	_ "golang/notes/02-basegra/test" // 下划线表示匿名导入
	u "golang/notes/02-basegra/utils"
	"math"
)
import (
	base "golang/notes/02-basegra/base"
)

func test() (a, b int) { // 函数名
	return 1, 2
}

/**
* 这里启动时的主函数
 */
func main() {
	fmt.Println("多包导入", math.Pi)
	u.GoRun() // 别名调用utils包下GoRun，函数
	u.SayHello()

	base.Sum()          // 算数运算符
	base.Relationship() // 关系运算符
	base.Logic()        // 逻辑运算符
	base.Bitwise()      // 位运算
	base.Ass()          // 复合运算 num <<= 2等驾驭 num=num<<2
	base.Pointerfnc()   // 指针运算符

	// 变量
	base.Variable()
	base.VariableStat()

	aa, bb := 15, 20
	fmt.Println(base.Swap(aa, bb))

	var arr [10]int = base.IsIota()
	fmt.Println(arr)

	//
	//var aa int  // 声明一个变量，默认为0
	//var bb = 10
	//cc := 20    // 声明并初始化，且自动推导类型
	//
	//fmt.Println(aa, bb, cc)
	//
	//// 多变量声明
	//// var a1,b1 string
	//// var c1,c2 string = "h1","h3"
	//// d1,d2 := "h2","h4"
	//// var(e int f bool)
	//m3, n3, q3 := 10, "n2", 30 // 自动推导类型，并初始化值
	//fmt.Println(m3, n3, q3)
	//
	//// 变量值互换
	//var m = 15
	//var n = 20
	//m, n = n, m // 直接将m与n的值互换
	//fmt.Println("m,n", m, n)
	//
	//var n1, m2 = 22, 25
	//var temp int
	//temp, _ = n1, m2 // 将n1的值赋值给n1，_表示丢弃变量 m2
	//fmt.Println("temp", temp)
	//
	//// 丢弃变量
	//_, ee := 35, 36 // 将值36赋值给ee，并同时丢弃35值
	//fmt.Println("ee", ee)

	// := 声明注意事项
	// num1, num2 := test()
}
