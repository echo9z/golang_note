package main

import (
	"fmt"
	"strconv"
)

func main() {
	// strconv几个实例
	parseConfig := func ()  {
		// 模拟配置文件
		config := map[string]string{ // map集合key是str，value是str
			"port": "3000",
			"timeout": "30",
			"status": "true",
			"pi": "3.1415926",
			"max_connect": "30000",
		}

		// 解析配置
		var (
			port int
			timeout int
			status bool
			pi float64
			max_connect uint64
		)

		// 解析端口号
		// v, err := strconv.Atoi(config["port"])
		// if err != nil {
		// 	port = v
		// }

		// 下面位简写，等价于上面语句
		if v, err := strconv.Atoi(config["port"]); err == nil {
			port = v
		}

		if v, err := strconv.ParseInt(config["timeout"], 10, 64); err == nil { // 解析timeout
			timeout = int(v) // int64转换位int
		}
		if v, err := strconv.ParseBool(config["status"]); err == nil {
			status = v
		}
		if v, err := strconv.ParseFloat(config["pi"], 64); err == nil {
			pi = v
		}
		if v, err := strconv.ParseUint(config["max_connect"], 10, 64); err == nil {
			max_connect = v
		}

		fmt.Printf("配置解析结果:\n")
		fmt.Printf("  端口: %d\n", port)
		fmt.Printf("  超时: %d秒\n", timeout)
		fmt.Printf("  调试模式: %v\n", status)
		fmt.Printf("  PI值: %.5f\n", pi)
		fmt.Printf("  最大连接数: %d\n", max_connect)


		// // ✅ 推荐：v 和 err 只在 if 块内可见
		// if v, err := strconv.Atoi("8080"); err != nil {
		// 		// 错误处理
		// } else {
		// 		fmt.Println(v)  // 可以使用 v
		// }
		// // fmt.Println(v)  // ❌ 错误：v 在这里不可见

		// // ❌ 不推荐：变量污染外部作用域
		// v, err := strconv.Atoi("8080")
		// if err != nil {
		// 		// 错误处理
		// }
		// fmt.Println(v)  // v 在整个函数内都可见
	}
	parseConfig()

	// 解析命令行
	commandLineArgs := func ()  {
		// 模拟命令行
		args := []string{
			"program.exe",
			"--count=100",
			"--ratio=1.5",
			"--verbose=true",
			"--mode=2",
			"--name=Alice",
		}

		// 创建一个接收任何类型的map集合
		params := make(map[string]interface{}) //key为string类型，值为interface空接能存储任何值

		for _, arg := range args[1:] { // 将数组截除第一个，从第二个元素开始
			if len(arg) > 2 && arg[:2] == "--"{ // 过滤参数长度是否大于2，同时从开头开始截取到第二个元素
				// 将count=100 进行分离
				parse := splitKeyValue(arg[2:]) // arg[2:]从--之后开始截取
				if parse != nil {
					key, value := parse[0], parse[1]
					params[key] = parsValue(value) // 将value字符串解析对于类型，存储在map对于的key中
				}
			}
		}

		// 进行解析参数
		fmt.Println("模拟解析运行参数：")
		// 遍历map集合
		for k, v := range params {
			fmt.Printf("参数 %s: %v (类型:%T)\n",k ,v,v)
		}
	}
	commandLineArgs()
}

// 定义一个将key=value的字符串进行分离的函数，返回str切片
func splitKeyValue(s string) []string {
	// 直接遍历字符串
	for idx, ch := range s {
		if ch == '=' {// 当遍历到字符为=时，通过idx进行切片分离
			return []string{s[:idx], s[idx+1:]}
		}
	}
	return nil // 没有=直接返回空
}

// 定义将string解析为其他类型函数
func parsValue(s string) interface{} {
	// 尝试解析整数，如果解析字符串存在错误，走下面解析
	if v, err := strconv.Atoi(s); err == nil {
		return v
	}

	// 解析浮点
	if v, err := strconv.ParseFloat(s, 64); err == nil {
		return v
	}

	// 解析布尔
	if v, err := strconv.ParseBool(s); err == nil {
		return v
	}

	// 默认为字符串
	return s
}
