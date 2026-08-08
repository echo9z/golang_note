package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func main()  {
	// 函数func：函数在go中是一等公民。函数可以像普通变量一样被传递、赋值给变量、作为参数传入或作为返回值返回。
	// 一.声明函数，使用func关键字声明
	sayHi()
	fmt.Println("+10", add(100))

	num := sum(10, 20)
	fmt.Println("num", num)

	// 二、函数参数
	intSum(10, 20)
	// 可变参数
	sum := intSum2(10, 20, 1, 2,3,4,5)
	fmt.Println("可变参数", sum)
	sum = intSum2(1,2, 10, 20)
	fmt.Println("可变参数", sum)

	// 三、返回值
	// 多个返回值
	n1, n2 := Calculate1(5, 10)
	fmt.Println("多个返回值", n1, n2) // 多个返回值 15 -5
	// 函数返回值作为调用函数参数
	fn1(fn2(20, 10)) // 30 10 200
	// 返回值命名
	n1, n2 = Calculate2(100, 200)
	fmt.Println("返回值命名", n1, n2) // 返回值命名 300 -100


	// 三、匿名函数
	// var sum1 func(a int, b int) int = func(a int, b int) int  {
	var sum1 = func(a int, b int) int { // 或者直接简写 sum1 := func(a int, b int) int {
		return a + b
	}
	num2 := sum1(19,20)
	fmt.Println("num2", num2)

	// 四.通过 type 关键字可以定义自定义函数类型
	// type MathOperation func(a, b int) int
	// 可以把subN和subtract、multiply函数赋值给cal 定义的类型函数
	var cal MathOperation // 声明一个MathOperation类型的变量cal
	cal = subtract // 把subtract赋值给cal
	fmt.Printf("type of cal:%T\n", cal)  // type of cal:main.MathOperation
	res1 := cal(20, 10) // 像直接调用subtract函数一样调用cal
	fmt.Printf("res1:%d\n", res1) // res1:10
	
	cal2 := subtract // 直接将函数subtract一个变量
	fmt.Printf("type of cal:%T\n", cal2) // type of cal:func(int, int) int
	fmt.Printf("cal2 res1:%d\n", cal2(30, 10)) // cal2 res1:20
	
	// 高阶函数分为函数作为参数和函数作为返回值两部分
	res2 := calc(10, 20, addN) // 将函数addN做参数进行传递
	fmt.Printf("calc res2:%d\n", res2) // calc res2:30
	// 函数作为返回值
	if add, err := doCalc("add"); err == nil {
		res3 := add(20, 20)
		fmt.Printf("return fn res3:%d\n", res3) // return fn res3:40
	}


	// 函数类型作为参数（回调函数）
	// 方式A直接调用，作为参数传入
	res := calculate(10, 20, addN)
	fmt.Println("add func：", res) // add func： 30

	// 方式B：保存在 map 中（策略表）
	ops := map[string]MathOperation{
		"+": addN,
		"-": subtract,
		"*": multiply,
		"/": func(a, b int) int { return a / b }, // 也支持匿名函数
	}

	operation := "*"
	if op,ok := ops[operation]; ok {
		fmt.Println("10*10 = ", op(10, 10)) // 10*10 =  100
	}
	

	// 五、Go 中，自定义类型（包括函数类型）都可以绑定方法。
	// 1. 普通匿名函数转换为 ProcessFunc 类型，
	// 可以这么理解type Age int，Age(18) 把 18 转成 Age 类型
	// ProcessFn(匿名函数)  把这个匿名函数 转成 ProcessFn 类型
	myFunc := ProcessFn(func(data string) { // myFunc类型为ProcessFn
		fmt.Println("正在处理数据:", data)
	})
	// 因为ProcessFn实现了Processor接口方法，所以可以直接作为 Processor 接口传递
	DoProcess(myFunc, "user_info")
	// 类似多态，可以参考下面例子
	/*
	// 定义一个行为约定：只要能说话就行
	type Speaker interface {
			Speak() string
	}

	type Dog struct{}
	func (Dog) Speak() string { return "汪汪" }

	type Cat struct{}
	func (Cat) Speak() string { return "喵喵" }

	// 参数用接口接收 —— Dog、Cat 都能传进来
	func MakeSound(s Speaker) {
			fmt.Println(s.Speak())
	}

	func main() {
			MakeSound(Dog{}) // 汪汪
			MakeSound(Cat{}) // 喵喵
	}*/

	// 六、中间件 / 装饰器模式（链式处理）
	// 函数类型常用于编写中间件（Middleware），对现有函数进行功能增强（如添加日志、耗时统计、权限校验等）。
	pipeline := WithLogging(WithTiming(SaveData))
	pipeline("user:tom,ageL10")
	// [日志] 开始处理请求，参数: user:tom,ageL10
	// --> 成功保存数据: user:tom,ageL10
	// [耗时统计] 执行完成，用时: 100.33288ms
	// [日志] 请求处理结束
}

// 一、声明函数
// 无返回参数
func sayHi()  {
	fmt.Println("say Hi")
}
// 一个参数
func add(n int) int {
	return n + 10
}

// 两个参数 和 一个返回参数
func sum(a int, b int) int {
	return a + b
}
// 连续多个参数类型相同，可以省略前面的类型 func sum(a, b int) int {}

// 在go中函数不支持重写，下面代码无法通过编译
type Person struct {
	Name string
	Age int
	Salary float64
}

func NewPerson(name string, age int, salary float64) *Person {
	return &Person{Name: name, Age: age,Salary: salary}
}
// func NewPerson(name string) *Person {
// 	return &Person{Name: name, Age: age,Salary: salary}
// }
// 在go中如果函数名不一样那就是完全不同的函数，那么就不应该取一样的名字，函数的重载会让代码变得混淆和难以理解。

// 二、函数的参数
// 1.函数的参数中如果相邻变量的类型相同，则可以省略类型
func intSum(x, y int) int {
	return x +y
}
// 2.可变参数：Go语言中的可变参数通过在函数参数名最后加...来标识。
func intSum2(x, y int, n ...int) int {
	fmt.Printf("%T \n", n) // n是一个切片
	sum := x + y
	for idx, arg := range n {
		fmt.Printf("args[%d]:%d \n", idx, arg)
		sum += arg
	}
	return sum
}

// 3.函数参数的传递
// 值传递：在调用函数时

// 三、函数返回值
// 1.多个返回值，必须使用()括起来
func Calculate1(x, y int) (int, int) {
	add := x+y
	sub := x-y
	return add, sub
}

// 2.返回值命名 
// 函数定义时可以给返回值命名，并在函数体中直接使用这些变量，最后通过return直接返回。
func Calculate2(x, y int) (add, sub int) {
	add = x+y
	sub = x-y
	return
}
// 3.函数返回值类型为slice时，nil可以看做是一个有效的slice，没必要显示返回一个长度为0的切片。
func someFn(str string) []string {
	if str == "" {
		return nil
	}
	sp := strings.Split(str, ",")
	return sp
}

// 4.函数返回值作为调用函数参数
// 一个函数可以将另一个函数调用作为其参数，被调用函数的返回值个数、返回值类型和返回值顺序与传入函数的参数形参一致。
// 比如fn1(a,b,c int)，fn2返回3个参数：fn2(a,b int) (int, int, int)。就可以fn1(fn2(a,b))调用
func fn1(a, b, c int) {
	fmt.Println(a, b, c)
}
func fn2(a,b int) (int, int, int) {
	add := a+b
	sub := a-b
	mul := a*b
	return add, sub, mul
}


// 三、函数的匿名参数
// go中可以参数只写类型而不写名称（即匿名参数），通常用于接口定义、函数类型声明或未使用的形参。
// 1.在接口使用：定义接口时，方法签名只需要指明参数类型，不需要参数名。
type Reader interface {
	// 只需要知道传入的是 []byte，不需要给它命名
	Read([]byte) (int, error)
}
// 带参数名的接口（辅助阅读）增加可阅读性
type UserServices interface {
	// 加上 userID 和 role，比单纯写 (int, string) 更容易让人理解参数的含义
  AssignRole(userID int, role string) error
}

// 2.在定义函数的类型：声明一个自定义的函数签名或回调类型时。
// 定义一个点击事件的回调函数类型
// 只需要知道参数是 string 和 int，不需要起名字
type OnClick func(string, int)
// 定义一个过滤数据的函数类型
type Filter func(int) bool

// 带参数名的函数类型（增强可读性）
// 加上参数名后，调用者一眼就能看出哪个是旧密码，哪个是新密码
type PasswordValidator func(newPassword string, oldPassword string) bool

// 3.普通函数中忽略未使用的参数：如果函数体不需要某个传入的值，可以用下划线 _ 或直接省略名称
func processData(data string, _ int)  {
	fmt.Println(data)
}

// 四、通过 type 关键字可以定义自定义函数类型（Function Type）。
// 把函数类型作为参数传递，或者放在 map 中作为查找表，可以轻松实现策略模式。
// 2.1 定义函数类型签名
type MathOperation func(a, b int) int
// 2.2 定义符合该签名的具体函数
func addN(a, b int) int {
	return a + b
}
func subtract(a, b int) int {
	return a - b
}
func multiply(a, b int) int {
	return a* b
}
// 2.3. 将函数类型作为参数传入
func calculate(a, b int, op MathOperation) int {
	return op(a, b)
}

// 3.1函数作为参数进行传递
// 三个参数
// x, y int
// op func(int, int) int
func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}
// 3.2函数作为返回值
func doCalc(str string) (func(int, int) int, error) {
	switch str {
	case "add":
		return addN, nil
	case "sub":
		return subtract, nil
	case "mul":
		return multiply, nil
	default:
		err := errors.New("must be add|sub|mul")
		return nil, err
	}
}

// 五、自定义类型（包括函数类型）都可以绑定方法。
// 定义一个接口
type Processor interface {
	Process(data string)
}
// 定义函数类型，可以理解为struct结构，这个函数结构取实现
type ProcessFn func(data string)

// 核心：函数类型实现Processor接口，绑定ProcessFn自定义函数类型
func (f ProcessFn) Process(data string)  {
	// 直接调用自己函数本身
	f(data)
}
// 普通函数参数中接收Processor接口
func DoProcess(p Processor, data string) {
	p.Process(data)
}

// 六、中间件 / 装饰器模式（链式处理）
// 定义业务处理函数类型
type Handler func(payload string)
// 中间件：耗时统计，用函数类型作为参数
func WithTiming(next Handler) Handler {
	return func (payload string)  {
		start := time.Now()
		// 调用下一个处理函数
		next(payload)
		
		fmt.Printf("[耗时统计] 执行完成，用时: %v\n", time.Since(start))
	}
}
// 中间件：日志记录
func WithLogging(next Handler) Handler {
	return func(payload string) {
		fmt.Println("[日志] 开始处理请求，参数:", payload)
		next(payload)
		fmt.Println("[日志] 请求处理结束")
	}
}
// 伪业务逻辑
func SaveData(data string) {
	time.Sleep(100 * time.Millisecond) // 模拟耗时
	fmt.Println("--> 成功保存数据:", data)
}
