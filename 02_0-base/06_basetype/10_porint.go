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
	var str *string
	fmt.Printf("str的值是%v\n", str) // str的值是<nil>
	
	// 指针传递例子
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




}
