package main

import (
	"fmt"
)

// 一、全局变量与局部变量
// 1.全局变量是定义在函数外部的变量，它在程序整个运行周期内都有效。 在函数中可以访问到全局变量。
var num int64 = 10

func GlobalVal() {
	// 函数中访问全局变量
	fmt.Printf("num:%d\n", num)
}
func LocalVal()  {
	// 2.局部变量x，仅在该函数内生效，外部无法直接访问该变量
	var x int = 100
	fmt.Printf("num:%d\n", x)
}

// 如果局部变量和全局变量重名，优先访问局部变量。
var total int64 = 20 // 全局
func testLocal()  {
	total := 100
	fmt.Printf("total:%d\n", total) // 函数中优先使用局部变量
}

func main() {
	// 1
	GlobalVal() // num:10
	LocalVal() 
	// fmt.Println(x) 无法使用局部变量x

	// 函数中优先使用局部变量
	testLocal() // total:100

	// 2
	numbers := []int{1, 5, 8, 12, 3, 15}
	limit := 10
	res := Filter(numbers, GreaterThan(limit))
	fmt.Printf("大于%d\n res:%v\n", limit, res)

	// 3 特殊机制：defer、panic 与 recover
	// defer关键字可以使得一个函数延迟一段时间调用，被 defer 的调用不会立即执行，而是在函数返回之前逐个执行被defer修饰的函数或语句。
	onWork()
	
	// defer的执行顺序是逆序执行，最先写defer语句最后被执行，最后写defer语句最先被执行。类似栈先进后出，后进先出
	fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	// panic(errors.New("恐慌panic"))
	defer fmt.Println(3)
	fmt.Println("end")

	// defer语句中声明变量，在声明时参数值就已经固定的
	defer fmt.Println(onDef())
	fmt.Println(3)
	// 正常认为输出的结果顺序,先输出3再是2,最后1
	// 但结果是2 3 1，在声明defer fmt.Println(onDef()),println输出参数是被固定，会先调用onDef()函数输出2,在确定固定参数defer fmt.Println(1)，再输出3,函数返回之前执行defer输出1,最后得到顺序2，3，1
	
	// 如果想要执行值3 2 1，用闭包
	defer func ()  {
		fmt.Println(onDef())
	}()
	fmt.Println(3)

	// defer在for循环中使用，虽然没有明令禁止，一般建议不要在 for 循环中使用 defer
	for i := 0; i < 5; i++ {
		defer fmt.Println("for未使用闭包",i) // i的值输出顺序 4 3 2 1 0
	}
	// 在go中，每创建一个defer，在当前协程及函数中申请一个片内存空间。当上述for循环较复杂的数据处理流程，外部请求数激增，短时间内会创建大量的defer，循环次数大或者次数不确定时，可能会导致内存暴增导致内存泄漏

	// 四、defer可以在return之后修改函数的最终返回值
	// 配合命名返回值，通过闭包修改函数的最终返回值
	fmt.Println("声明返回值",double()) // 声明返回值 20
	fmt.Println("匿名返回值",badDouble()) // 匿名返回值 10
	
}

// 三、defer延迟调用
func onWork()  {
	defer func ()  {
		fmt.Println("hello")
	}()
	fmt.Println("world")
}
func onDef() int {
	fmt.Println(2)
	return 1
}
// 配合命名返回值，通过闭包修改函数的最终返回值
func double() (val int) {
	defer func ()  {
		val *= 2 // 通过闭包引用了命名返回值变量，修改的是返回值val
	}()
	return 10 // 返回值val被赋值 10 -> 在defer执行 val*=2 ->最后真正返回 20
}
// 匿名返回值修改不良
func badDouble() int {
	var val int = 10
	defer func ()  {
		val *= 2 // 这里的defer修改是局部变量，
	}()
	return val
}

// 二、高阶函数：作为参数与返回值（通用过滤器与工厂）
// 1.定义断言/过滤函数类型
type Predicate func(int) bool

// 2.泛用切片过滤函数（接收函数类型作为参数）
func Filter(num []int, p Predicate) []int {
	var result []int
	for _, v := range num {
		if p(v) { // 通过过滤函数进行条件处理判断
			result = append(result, v)
		}
	}
	return result
}

// 3. 工厂函数（返回函数类型）
func GreaterThan(limit int) Predicate {
	return func(n int) bool {
		return n > limit // 取的切片值大于limit，追加到res[]中
	}
}

