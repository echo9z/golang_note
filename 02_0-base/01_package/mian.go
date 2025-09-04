package main // 程序主入口
import "fmt"
import (
	_ "01_package/driver" // 使用_导入包，表示只执行包的init函数，不使用包中的其他内容
	"01_package/utils"
	m "math"
)

func main() { // 该文件必须包含main函数
	fmt.Println("这是主函数main")
	res := utils.Add(3, 5) // 调用utils包下的add函数
	fmt.Println(res)
	utils.GoRun()
	var max float64 = m.Max(10, 20)
	fmt.Println("max=", max)

	// 使用可见常量
	fmt.Println("PI=", utils.PI)
}
