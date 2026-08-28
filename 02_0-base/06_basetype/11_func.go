package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
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
	// 实参（argument）：调用时实际传入的值，例如 pack1.Add(3, 5) 中的 3 和 5。
  // 形参（parameter）：函数定义时占位的变量，例如 func Add(a int, b int) 中的 a 和 b 
	// 值传递
	x := 10
	y := 20
	swapVal(x, y)
	fmt.Printf("x:%d y:%d\n", x, y) // x:10 y:20
	
	// 引用传递，传递内存地址
	swapPoint(&x, &y)
	fmt.Printf("x:%d y:%d\n", x, y) // x:20 y:10

	// 函数参数的任意类型
	res4 := intSum3(10, 20, "30", 10)
	fmt.Println("args any", res4) // args any 40

	// 三、返回值
	// 多个返回值
	n1, n2 := Calculate1(5, 10)
	fmt.Println("多个返回值", n1, n2) // 多个返回值 15 -5
	// 函数返回值作为调用函数参数
	fn1(fn2(20, 10)) // 30 10 200
	// 返回值命名
	n1, n2 = Calculate2(100, 200)
	fmt.Println("返回值命名", n1, n2) // 返回值命名 300 -100

	// 函数返回一个匿名函数
	multip := makeMultiplier(10)
	resM := multip(9)
	fmt.Printf("mult 9*10：%d\n", resM) // mult 10*9：90

	// 6.命名返回参数允许 defer 延迟调用通过闭包读取和修改。
	addRes := add2(10, 20)
	fmt.Printf("add2 defer：%d\n", addRes) // add2 defer：40
	addRes2 := add3(10, 20)
	fmt.Printf("add3 defer：%d\n", addRes2) // add3 defer：30

	// 三、匿名函数
  // 没有函数名的函数，可以立即执行，或者赋值给变量：
	// var sum1 func(a int, b int) int = func(a int, b int) int  {
	var sum1 = func(a int, b int) int { // 或者直接简写 sum1 := func(a int, b int) int {
		return a + b
	}
	num2 := sum1(19,20)
	fmt.Println("num2", num2)

	// 匿名函数自调用，声明并立即执行
	func (msg string)  {
		fmt.Println(msg)
	}("hello fn")
	// 匿名函数多用于实现回调函数和闭包。

	// 切片中为func类型
	fnSlices := []func(int)int{
		func(i int) int { return i+1},
		func(i int) int { return i+2},
	}
	fmt.Println("fn",fnSlices[0](10))
	
	// 结构体字段为func类型
	objFn := struct { // 匿名结构体
		fn func(string) string
	}{
		fn: func(str string) string {return str},
	}
	fmt.Println(objFn.fn("ok"))

	// channel通道返回func类型
	chFn := make(chan func() string, 2) // 建一个通道，里面装的是 func() string 类型
	chFn <- func() string { return "im ok"} // ch <- value：把匿名函数「送进」通道（发送）
	res5 := (<-chFn)() // <-chFn：从通道里「取出」那个函数，然后 () 调用它
	fmt.Println("res5", res5) // res5 im ok

	// 四、在go中的闭包
	fnCount := makeCount() // fnCount变量函数并且它引用了其外部作用域中的count变量，此时fnCount就是一个闭包。在fnCount的生命周期内，变量x也一直有效。 
	fmt.Printf("fnCount:%d\n", fnCount()) // 1
	fmt.Printf("fnCount:%d\n", fnCount()) // 2
	fmt.Printf("fnCount:%d\n", fnCount()) // 3

	// 闭包 进阶1
	// fnAdder1
	fnAdder2 := adder2(20)
	fmt.Printf("adder2:%d\n", fnAdder2(20)) // 40
	fmt.Printf("adder2:%d\n", fnAdder2(20)) // 60
	fmt.Printf("adder2:%d\n", fnAdder2(40)) // 100

	// 闭包 进阶2
	jpegFn := makeSuffixFn(".jpeg")
	pngFn := makeSuffixFn(".png")
	fmt.Println(jpegFn("image1")) // image1.jpeg
	fmt.Println(pngFn("image2")) // image2.png

	// 进阶3
	add, sub := Calculate3(10)
	fmt.Printf("add:%d, sub:%d \n", add(1), sub(2)) // 闭包每次都对一个局部变量进行加减操作，先加1：add=11，再减2：sub=9
	fmt.Printf("add:%d, sub:%d \n", add(2), sub(3)) // 再以base=9，加2：add=11,再减3：sub=8

	// 五、递归函数
	// 1. 阶乘：5! = 5×4×3×2×1
	fmt.Printf("5的阶乘：%d\n",factorial(5))
	// 2.斐波那契数列 1 1 2 3 5 8 13...
	for i := 0; i < 8; i++ {
		fmt.Printf("%d ", fibonacci(i)) // 0 1 1 2 3 5 8 13 
		if i == 7 {
			fmt.Println()
		}
	}

	// 匿名函数递归（闭包递归）
	// 匿名函数要实现递归必须先声明函数变量，在进行赋值，否则函数体内无法调用本身
	var countDown func(int) (int, error)
	countDown = func (n int) (int, error) {
		fmt.Printf("%d ", n)
		if n <= 0 {
			return 0, errors.New("传入值大于0")
		}
		if n <= 1 {
			return 1, nil
		}
		val, err := countDown(n - 1)
		if err != nil {
			return 0, err
		}
		return n * val, nil
	}
	val, _ := countDown(5)
	fmt.Println("匿名函数递归", val) // 5 4 3 2 1 匿名函数递归 120

	// 递归遍历目录
	if err := walkDir("02_0-base/06_basetype"); err != nil {
		panic(err)
	}
	// 函数与函数之间相互调用，即相互递归
	// 比如调用isEven内调用isOdd，isOdd
	fmt.Println("22 isEvent", isEvent(22))
	fmt.Println("19 isEvent", isEvent(19))
	fmt.Println("21 isOdd", isOdd(21))
	fmt.Println("18 isOdd", isOdd(18))

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
// args是一个slice通过arg[index]依次访问所有参数，通过len(arg)来判断传递参数的个数。
func intSum2(x, y int, args ...int) int {
	fmt.Printf("%T \n", args) // n是一个切片
	sum := x + y
	for idx, arg := range args {
		fmt.Printf("args[%d]:%d \n", idx, arg)
		sum += arg
	}
	return sum
}
// 任意类型的不定参数：interface{}传递任意类型数据。在 Go 1.18+使用any与interface{} 的完全等价别名。
func intSum3(args ...interface{}) any {
	total := 0
	for _, arg := range args {
		// 使用断言进行判断参数的类型
		if num, ok := arg.(int); ok {
			total += num
		}
	}
	return total
}


// 3.函数参数的传递
// 值传递：在调用函数时将实际参数复制一份将副本传递函数中，在函数中修改传递参数值，对其副本参数进行修改不会影响实际传递的变量
func swapVal(x, y int)  {
	temp := x
	x = y
	y = temp
}
// 引用传递（指针传递）：在调用函数时将实际参数的地址传递到函数中，在函数中修改传递参数值，实际上对内存地址进行修改会影响实际传递的变量
func swapPoint(x, y *int)  {
	temp := *x // *x取得的是内存指针所指向值
	*x = *y
	*y = temp
}

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

// 5.函数返回一个匿名函数
// 函数作为返回值
func makeMultiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

// 6.命名返回参数允许 defer 延迟调用通过闭包读取和修改。
func add2(x, y int) (z int) {
	defer func ()  { // 在return返回函数前调用匿名函数
		fmt.Println("闭包中defer",z)
	}()

	z = x + y
	return z + 10 // 执行顺序 z = x + y -> z + 10 -> defer func() -> return返回
}

func add3(x, y int) int{
	var z int
	defer func ()  { // 在return返回函数前调用匿名函数
		fmt.Println("闭包中defer",z)
	}()

	z = x + y
	return z + 10 // 执行顺序 z = x + y -> z + 10没有定义命名返回参数赋值给返回值变量z+10=40 -> defer func()执行z值为30 -> return返回
}

// 四、闭包
// go中闭包函数，本质是一个匿名函数捕获外层变量，被捕获的变量其生命周期超出了定义的作用域，跟着匿名函数存活。
func makeCount() (func() int) {
	count := 0 // 外层局部变量
	return func() int { // 这个匿名函数捕获了 count 外层变量
		count++
		return count
	}
}
// makeCounter()会得到一个独立的 count
// 编译器逃逸分析发现变量被捕获且逃出函数 → 自动搬到堆，闭包存指针

func adder1() (func(int) int) {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}
// 进阶1
func adder2(x int) (func(int) int) { // 调用adder2函数返回闭包函数，闭包函数引用外部的 adder2传入形参x变量，形成闭包
	return func(y int) int {
		x += y
		return x
	}
}
// 进阶2
func makeSuffixFn(suffix string) (func(string) string) {
	return func(s string) string {
		if !strings.HasSuffix(s, suffix) { // s字符串是否以 suffix为结尾
			return s + suffix
		}
		return s
	}
}
// 进阶3 返回两个闭包
func Calculate3(base int) (add, sub func (int) int) {
	add = func (val int) int {
		base += val
		return base
	}
	sub = func (val int) int {
		base -= val
		return base
	}
	return add, sub
}

// 五、递归函数
// 递归函数在执行的过程中，直接或者间接调用自己本身的函数。
// 递归有明确终止条件：无限递归，会导致栈溢出问题
// 每次递归都向终止条件靠近：例如参数递减、问题规模缩小等
// 递归深度可控：Go 的 goroutine 栈可以动态增长，但有上限
// 1.阶乘 5! = 5×4×3×2×1
// 递归终止条件：0! = 1
// n! = n × (n-1)!
func factorial(n int) int {
	if n <= 1 { // 当n <= 1时，递归出口返回
		return 1
	}
	return n * factorial(n - 1)  // 递归调用：规模缩小（n-1）
}

// 2.斐波那契而数列，即前两个数为 1，从第三项开始每个数均为前两项之和
// 0 1 1 2 3 5 8 13...
// F(0)=0,F(1)=1, F(n)=F(n−1)+F(n−2) n表示第几项，比如f6=第5项+第4项
func fibonacci(n int) int {
	// 递归终止条件
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	// 多分支递归调用
	return fibonacci(n - 1) + fibonacci(n - 2)
}

// 3.递归遍历目录
func walkDir(path string) error {
	// 读取指定path目录，按文件名排序的返回切片类型的所有目录条目
	entries, err := os.ReadDir(path)
	if err != nil { // 读取目录发生错误
		return err
	}
	// 遍历读取到的文件切片目录条目
	for _, entry := range entries {
		// 将传入目录和读取文件名进行拼接完成路径
		full := filepath.Join(path, entry.Name())
		if entry.IsDir() {
			// 条目中是目录，递归进入子目录，进行遍历
			if err := walkDir(full); err != nil {
				return nil
			}
		} else {
			fmt.Println(full)
		}
	}
	return nil
}

// go语言中可以互相调用的递归函数：多个函数相互调用形成闭环。
// isEvent判断一个非负数是否为偶数
func isEvent(n int) bool {
	if n == 0 { // 当互相调用传入为0，终止递归循环，偶数返回true
		return true
	}
	return isOdd(n - 1) // 偶数减 1 后应为奇数
}

// isOdd判断一个非负数是否为奇数
func isOdd(n int) bool {
	if n == 0 {
		return false
	}
	return isEvent(n - 1) // 奇数减 1 后应为偶数
}
// 时间复杂度为 O(n)，对于较 n数可能栈溢出，实际中更常用取模运算。

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
