package pri

import (
	priv "02_private_pkg/pri/internal/ser"
	pub "02_private_pkg/pri/ser"
	"fmt"
)

func Pri_Test() {
	fmt.Println("我是pri住测试函数")
	pub.Ser() // 可以访问的ser函数

	// 在pri包中调用可以访问的私有ser函数
	priv.Ser() // 私有的包中的函数
}
