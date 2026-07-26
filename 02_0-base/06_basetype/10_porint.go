package main

import "fmt"

func main(){
	// point
	// 在go中的指针是不能指针偏移和指针运行的，go中只保留指针访问和修改内存地址的能力
	// 任何程序数据载入内存后，在内存都有他们的地址，这就是指针。为了保存一个数据在内存中的地址，此时需要指针变量存放内存地址。
	// 用比喻快速理解
	// 内存 = 一家巨大的酒店
	// 程序数据（比如变量 100） = 入住酒店的客人
	// 内存地址 / 指针 = 客人房间的**“房间号”**（比如 301号）
	// 指针变量 = 你手里拿的一张**“房卡”**，上面写着“301号”

	// 指针地址和指针类型
	// *T 指向类型 T 的指针，比如 *int是指向int的指针
	// & 用于获取变量的地址。取地址
	// * 用于从指针解引用。解引用
	var a int = 10
	var p *int = &a // p
	fmt.Printf("a:%d ptr:%p\n",a, &a) // a:10 ptr:0x3b3c078a0128
	fmt.Printf("p:%p ptr:%T\n", p, p) // p:0x3b3c078a0128 ptr:*int
	fmt.Println(&p) // 0x3e55e104050
	fmt.Println(*p) // 解引地址值，获取指针指向的值
	// 通过指针修改指向地址的值为15
	*p = 15
	fmt.Println(*p) // 15

	// *指针取值，根据指针去内存取值
	c := 100
	d := &c // 取变量c的地址，将指针保存到d中
	fmt.Printf("type of d:%T \n", d) // type of d:*int 
	e := *d // 根据d指针去内存取值
	fmt.Printf("type of e:%T, value of e:%d\n", e, e) // type of e:int, value of e:100

	// 空指针，指针的零值是 nil
	// 初始化指针为nil，解引用nil 指针会 panic：恐慌异常
	var p1 *int
	if p1 != nil {
		fmt.Println("非空")
	} else {
		fmt.Println("空")
	}
	fmt.Println(p1) // <nil>
	fmt.Println(p1 == nil) // true，可以比较
	// 为什么会出现panic异常，*p1 = 100把 100 写入 p1 所指向的那块内存地址，但p指针地址为nil，没有指向任何有效的内存地址。运行时检测到后直接 panic
	// *p1 = 100 // panic: runtime error: invalid memory address or nil pointer dereference
	// p1只是声明了一个指针，但未初始未分配内存是一个空指针，无法正常使用
	// 要么从其他变量取地址符，将其变量的地址赋值给该p1指针
	var num2 int = 110
	p1 = &num2
	fmt.Printf("类型%T, %d\n", p1, *p1) // 类型*int, 110
	// 或者内置函数new手动分配，初始化int为0
	p1 = new(int)
	fmt.Printf("类型%T, %d\n", p1, *p1) // 类型*int, 0

	var str *string
	fmt.Printf("str的值是%v\n", str) // str的值是<nil>

	
	// 使用 new 函数创建指针
	// func new(Type) *Type
	// type表示传入参数类型
	// new函数返回一个指向该类型内存地址的指针
	n1 := new(int)
	*n1 = 100
	fmt.Printf("类型%T, %d\n", n1, *n1) // 类型*int, 100
	s1 := new(string)
	*s1 = "ok"
	fmt.Printf("类型%T, %s\n", s1, *s1) // 类型*string, ok

	fmt.Printf("t:%T, v:%s\n", *new(string), *new(string)) //  t:string, v:   空字符串
	fmt.Printf("t:%T, v:%d\n", *new(int),  *new(int)) // t:int, v:0
	fmt.Printf("t:%T, v:%v\n", *new([5]int),  *new([5]int)) // t:[5]int, v:[0 0 0 0 0]
	fmt.Printf("t:%T, v:%v\n", *new([]float64),  *new([]float64)) // t:[]float64, v:[]  空切片

	type User struct{
		name string
		age int
	}
	u1 := new(User)
	*u1 = User{name: "tom", age: 18}
	*&u1.age = 20 // 这样写 .运算符先级高于 *和&，
	// 这么写不易阅读，先go自动解引，再对20进行取地址值，再又把地址解引用回来（不推荐）
	// 1	u1.age				Go 自动解引用 u1，取到结构体的 age 字段	  类型为int
	// 2	&(u1.age)			再对这个 int 字段取地址										类型为*int
	// 3	*(&(u1.age))	又把地址解引用回来												类型为int（左值）
	// 4	= 20					赋值	
	u1.age = 20 // Go 自动解引用，直接写字段
	(*u1).age = 20 // 显式解引用，也对，但啰嗦
	fmt.Printf("类型%T, %v\n", u1, *u1) // 类型*main.User, {tom 20}

	// new函数主要针对基本类型（int、string、bool、byte…）没有字面量初始化语法
	p2 := new(int) // *int，指向值 0，一行搞定
	*p2 = 15
	// 如果不是new函数，得需要写三行
	var x1 int
	p3 := &x1 // *p3=0 多一个临时变量 x
	*p3 = 155
	fmt.Printf("类型%T, %v\n", p3, *p3) // 类型*int, 155

	// &T{}结构体有字面量语法 User{...}，所以 &User{...} 比 new(User)要好
	// 对结构体：&T{} 基本取代了 new(T)
	u2 := &User{name: "jack", age: 19}
	(*u2).age = 20 // 等价于自动解引用 u2.age = 20
	u3 := new(User) // 也行，但只能零值，还得再赋字段
	u3.name = "tom"
	u3.age = 18

	// 指针的应用场景
	// 1指针传递 作为函数参数
	// Go 默认是值传递（拷贝）。如果希望函数内部修改外部变量，必须传指针。
	md1 := func (x int)  { // 传递值时，只是将值复制副本
		x = 100
	}
	md2 := func (x *int)  { // 形成是传递的复制是内存地址
		*x = 100
	}
	num := 10
	md1(num)
	fmt.Println(num) // 10
	md2(&num)
	fmt.Println(num) // 100

	// 2结构体指针接收者
	account1 := Account{balance: 100}
	account1.updateBalance(100)
	account1.getBalance() // balance is 200

	// 3避免大对象或者数组拷贝 性能
	big := BigData{
		Values: [1024*1024]int{0:10, 100: 110},
	}
	big.processVal(1, 100)
	// 调用函数，传递指针，不需要传递对象
	total := processBigDataP(&big)
	fmt.Println("total", total)

	// 4.map中修改结构体字段
	// map 的元素不可寻址，不能直接改字段，只能整体替换：
	// m := map[string]Account{"a": {balance: 100}}
	acc1 := make(map[string]Account)
	acc1["a"] = Account{balance: 100}
	// map 的值是「不可寻址」(not addressable) 的。
	// acc1["a"] 返回的是值的一个拷贝，而不是对 map 内部实际存储元素的引用
	// acc1["a"].balance += 100 // cannot assign to struct field m["a"].balance in map
	// 第一种整体替换，将"a"key的值account结构体复制给a1，将修改后的a1新account对象赋值给acc1["a"]
	a1 := acc1["a"]
	a1.balance += 100
	acc1["a"] = a1
	fmt.Printf("整体替换新的Acc结构体：%v\n", acc1["a"].balance)

	// 第二种 存指针 推荐
	// m2 := map[string]*Account{"a": {balance: 100}}
	acc2 := make(map[string]*Account)   // map值为指针
	acc2["a"] = &Account{balance: 50}
	acc2["a"].balance += 50  // acc2["a"]为指针，指向同一个地址
	fmt.Printf("map值为指针：%v\n", acc2["a"].balance)

	// 5.实现链表 / 树等数据结构
	// 链表
	type ListNode struct {
    Val  int
    Next *ListNode // 必须是指针，否则无法指向下一个节点
	}
	list := &ListNode{Val: 1}
	list.Next = &ListNode{Val: 2}
	list.Next.Next = &ListNode{Val: 3}
	fmt.Printf("list：%+v\n", list) // list：&{Val:1 Next:0x110bf5e24030}
}

// 2结构体指针接收者
type Account struct {
		id int
		name string
		balance int
}
// 结构体的字段时，即便拿到的是一枚指针，也可以直接用 . 操作符，Go 会自动帮你解引用。
func (a *Account) updateBalance(amount int)  {
	 a.balance += amount  // 实际上等价于  (*a).balance += balance
}
// 只是读取，不会修改
func (a Account) getBalance()  {
	fmt.Println("balance is", a.balance)
}

// 3避免大对象拷贝
type BigData struct{
	// 对象属性中存在100万数组，如果将整个结构对象的值拷贝传递，会复制整个数组
	// Values [1_000_000]int // [1_000_000] 下划线 _ 是数字分隔符，代表 100 万
	Values [1024*1024]int // 占用 1MB 内存（数组是值类型，复制时会完整拷贝）
}
func (bigData *BigData)processVal(idx, num int) {
	// 比如将Values的100万元素进行相加求和
	bigData.Values[idx] = num
}

// 3.1值传递函数（Go 默认拷贝整个结构体，包含 1MB 的 Data 数组）
func processBigDataV(bigData BigData) int {
	var total int
	for i := 0; i < len(bigData.Values); i++ {
		total += i
	}
	return total
}
// 3.2指针传递函数（只复制 8 字节的内存地址）
func processBigDataP(bigData *BigData) int {  // 直接操作原数据，无需复制整个数组
	// 比如将Values的100万元素进行相加求和，直接复制传递指针只复制 8 字节，不会
	var total int
	for i := 0; i < len(bigData.Values); i++ {
		total += i
	}
	return total
}

