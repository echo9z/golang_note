package utils

import "fmt"
import pri "02_private_pkg/pri"
import ser "02_private_pkg/pri/ser" // utils可以访问pri包下的共有 ser包

//import priv "02_private_pkg/pri/internal/ser"  // utils无法访问pri包下的私有 internal包

func GoRun() {
	fmt.Println("go run")
	pri.Pri_Test()
	ser.Ser()
	// priv.Ser() 无法引用 pri下文件夹名为 internal的私有函数或类型
}
