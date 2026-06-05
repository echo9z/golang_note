package main

import (
	"fmt"
	"reflect"
	"slices"
	"strconv"
)

func main()  {
	// 切片
	// 切片是对数组一个连续片段元素引用（可以称之为相关数组），所以切片是一个引用类型
	// 切片本身并不存储任何数据，它只是对底层数组的一个连续片段的引用。

	// 一、切割数组，引出切片
	// 切割数组返回切片 arr[start:end]，分割区域左闭右开，切割完后返回切片类型，前面说的只是对数组一个连续片段的引用
	arr := [5]int{1,2,3,4,5}

	/**
	arr[:] // 子切片范围[0,5) -> [1 2 3 4 5]
	arr[1:] // 子切片范围[1,5) -> [2 3 4 5]
	arr[:5] // 子切片范围[0,5) -> [1 2 3 4 5]
	arr[2:3] // 子切片范围[2,3) -> [3]
	arr[1:3] // 子切片范围[1,3) -> [2 3]
	*/
	var slice1 []int = arr[1:3] // 子切片范围[1～3) 2至3个元素
	fmt.Printf("arr：%v 类型:%T\n",arr, arr) // arr：[1 2 3 4 5] 类型:[5]int
	fmt.Println("slice1:",slice1, "类型:",reflect.TypeOf(slice1)) // slice1: [2 3] 类型: []int。这里通过reflect反射回去变量类型
	
	// 访问切片中引用的元素
	fmt.Println(slice1[0]) // 2

	// 想要将数组转换为切片，直接arr[:]不带索引取值范围转换为切片
	arr2 := [5]int{1,2,3,4,5}
	slice2 := arr2[:] // len(arr2) == 5，cap(arr2) == 5
	// 转换后的切片和数组指向的是同一个内存，当修改切片中元素会影响元数组的数据变化
	slice2[1] = 20
	fmt.Printf("array:%v\n", arr2) // array:[1 20 3 4 5]
	fmt.Printf("slice:%v\n", slice2) // slice:[1 20 3 4 5]
	
	// 若想对原数组不受影响，使用go1.21 引入的slices.Clone()标准库，用于创建切片的浅拷贝
	arr3 := [5]int{1,2,3,4,5}
	// slices.Clone返回一个独立新切片
	slice3 := slices.Clone(arr3[:]) // 克隆arr3[0:5]的切片，返回独立切片
	slice3[0] = 100
	fmt.Printf("array:%v\n", arr3) // array:[1 2 3 4 5]
	fmt.Printf("slice:%v\n", slice3) // slice:[100 2 3 4 5]
	
	// 等价于在go1.21之前版本通过make+copy，实现数组到切片，在通过 copy函数是切片深拷贝
	// make([]T, len, cap)：首先分配一段新的底层数组内存。通过指定长度为 len(originSlice)，确保了拷贝的目标切片拥有足够的空间来容纳所有元素。
	slice4 := make([]int, len(arr3)) // make函数创建一个[]int切片长度与arr3一样
	// func copy(dst, src []Type) int
	// copy 内置函数将源切片中的元素复制到目标切片中。 
	// （作为一种特殊情况，它还将字节从字符串复制到字节切片。）源和目标可能重叠。 
	// Copy 返回复制的元素数量，该数量将是 len(src) 和 len(dst) 中的最小值。
	copy(slice4, arr3[:])
	slice4[4] = 500
	fmt.Printf("array:%v\n", arr3) // array:[1 2 3 4 5]
	fmt.Printf("slice:%v\n", slice4) // slice:[100 2 3 4 5]

	// 这里浅拷贝概念
	// 基本类型 值类型 通过浅拷贝，完全独立的切片
	a := []int{1,2,3}
	b := slices.Clone(a)
	b[0] = 10
	fmt.Printf("a:%v, b:%v\n",a,b) // a:[1 2 3], b:[10 2 3] 不影响原切片

	// 引用类型 切片中是（指针/切片/map）- 只复制了引用，底层对象仍共享
	type User struct {
		Scores []int
	}

	var user1 User = User{
		Scores: []int{100,98,85},
	}
	var user2 User = User{
		Scores: []int{140,98,85},
	}
	// User引用切片
	Users := []User{user1, user2}
	// 通过Clone进行支队第一层了拷贝
	CloneUser := slices.Clone(Users)

	// 修改CloneUser中切片第一个元素，不影响Users对象切片
	CloneUser[0] = User{
		Scores: []int{60,60,60},
	}
	fmt.Printf("users%v\n", Users)  // users[{[100 98 85]} {[140 98 85]}]
	fmt.Printf("CloneUser%v\n", CloneUser) // CloneUser[{[60 60 60]} {[140 98 85]}]

	// 但我对CloneUser[1]中的属性Scores，切片第三个元素进行修改150分，就会影响原切片，切片中Scores还是共享的
	CloneUser[1].Scores[2] = 150
	fmt.Printf("users%v\n", Users) // users[{[100 98 85]} {[140 98 150]}]  150就是被影响了
	fmt.Printf("CloneUser%v\n", CloneUser) // CloneUser[{[60 60 60]} {[140 98 150]}]

	// 需要要深拷贝要自己递归 clone
	for idx := range Users {
		// 将CloneUser对象切片中Scores切片，也进行拷贝，从而达到scores切片为独立的内存
		CloneUser[idx].Scores = slices.Clone(Users[idx].Scores)
	}
	// 通过对scores切片属性进行深拷贝处理，再次对Scores切片第三个元素进行修改120分，就不会影响原切片
	CloneUser[1].Scores[2] = 120
	fmt.Printf("users%v\n", Users)  // users[{[100 98 85]} {[140 98 150]}] 150没有被影响修改
	fmt.Printf("CloneUser%v\n", CloneUser) // CloneUser[{[100 98 85]} {[140 98 120]}]

	// 二、创建切片
	// x := []int{2, 3, 5, 0, 0}[:3]
	//1.字面量初始化切片
	s1 := []int{1,2,3} // 长度和容量都为3
	s2 := []string{"a", "b"} // 长度和容量都为2
	fmt.Printf("slice1:%v(%T), slice2:%v(%T)\n",s1,s1,s2,s2) // slice1:[1 2 3]([]int), slice2:[a b]([]string)

	// 构建切片指定索引处的元素
	x := []int{2, 3, 5, 8:100} // 长度为9 容量为9
	fmt.Println(len(x), cap(x), x) // 9 9 [2 3 5 0 0 0 0 0 100]


	// 2.使用make函数创建切片
	/* func make(t Type, size ...IntegerType) Type
		make(type, length, capacity) // 仅适用于 slice
		make(type, initialCapacity)  // 适用于 map 和 channel
		type：必须是 slice、map 或 channel 类型。
		length（仅对 slice 有效）：
			切片的初始长度（元素数量）。
			必须指定，否则编译报错。
		capacity（可选，仅对 slice 有效）：
			切片的容量（底层数组的大小）。
			若未指定，默认与 length 相等。

		initialCapacity（对 map 和 channel 有效）：
			map：预分配的哈希表桶数（可选，若未指定则按需动态分配）。
			channel：通道的缓冲区大小（可选，若未指定则为无缓冲通道）。
	*/

	// make 内置函数分配并初始化 slice、map 或 chan（仅）类型的对象。
	// 与 new 一样，第一个参数是类型，而不是值。
	// 与 new 不同，make 的返回类型与其参数的类型相同，而不是指向new函数传入函数参数的指针。

	// make([]T, len)       // size[0] = len，cap = len
	// make([]T, len, cap)  // size[0] = len，size[1] = cap
	// 一个参数les和cap相等
	s3 := make([]int, 5) // 切片长度和容量都为5
	fmt.Printf("%v\n", s3)
	fmt.Printf("len:%d, cap:%d\n", len(s3), cap(s3)) // len:5, cap:5
	// 两个参数：指定len和cap大小
	s4 := make([]string, 2, 5) // len2, cap5   内存中["","",""，""，""]，只有前两个索引位置合法可访问
	fmt.Printf("%v\n", s4)
	fmt.Printf("len:%d, cap:%d\n", len(s4), cap(s4)) // len:2, cap:5
	
	ss := make([]int, 3, 10) // 内存中[0 0 0 _ _ _ _ _ _ _] 前三个元素可访问
	fmt.Printf("len:%d, cap:%d\n", len(ss), cap(ss)) // len:3, cap:10

	// 空切片与 nil 切片
	var nilSlice []int // 声明切片变量，默认为nil，len=0，cap=0，底层array=nil
	emptySli := make([]int, 0) // 非nil空切片，len=0，cap=0 底层array=[]，序列化为JSON会输出 []
	emptySli2 := []int{}       // 同样非nil空切片
	fmt.Printf("nilSlice:%v, emp1:%v, emp2:%v\n", nilSlice, emptySli,emptySli2) // nilSlice:[], emp1:[], emp2:[]

	// 3.通过append()向预分配的切片追加元素
	// func append(slice []Type, elems ...Type) []Type
	// 	slice：目标切片
	// 	elems：可变参数，向目标切片添加1个或多个元素
	// Append 内置函数的作用是将元素添加到切片的末尾。
	// 如果内存容量足够，则会重新分配切片的空间以容纳新添加的元素；
	// 如果内存不足，则会创建一个新的数组来存储这些元素。最后，Append 会返回更新后的切片。
	// 因此，通常需要将 Append 的结果存储在一个变量中，这个变量就应该是保存切片本身的那个变量。
	s5 := make([]string, 2, 5)
	for i := 0; i < 5; i++ {
		s4 = append(s5, strconv.Itoa(i)+"a")
	}
	fmt.Printf("s4:%v\n",s4) // s4:[  0a 1a 2a 3a 4a]
	// 追加单个元素
	a1 := []int{1,2}
	a1 = append(a1, 3)
	fmt.Println(a1) // [1 2 3]
	// 追加多个元素
	a1 = append(a1, 4,5,6)
	fmt.Println(a1) // [1 2 3 4 5 6]

	// 将一个切片追加到另一个切片中
	b1 := []int{7,8,9}
	// 用...把b切片展开为可变参数
	a1 = append(a1, b1...)
	fmt.Println(a1) // [1 2 3 4 5 6 7 8 9]

	// 追加字符串到[]byte
	var buf []byte
	buf = append(buf, "hello"...) // 将hello分单个字符 buf==hello 
	var str []string
	str = append(str, "hello")
	fmt.Printf("字节%s 字符串切片%s\n",buf, str) // 字节hello 字符串切片[hello]

	// append扩容机制
	// cap容量足够时
	s6 := make([]int, 0, 4) // leb:0 cap:4
	s6 = append(s6, 1, 2, 3, 4)
	fmt.Printf("len=%d cap=%d %v\n", len(s6), cap(s6), s6) // len=4 cap=4 [1 2 3 4]
	
	// 触发扩容
	s6 = append(s6, 5) // cap翻一倍
	fmt.Printf("len=%d cap=%d %v\n", len(s6), cap(s6), s6) // len=5 cap=8 [1 2 3 4 5]

	/*
	一、扩容机制三步走
	第一步：计算初始 newcap（Go 1.18+ 源码 runtime/slice.go）
		go// 伪代码，简化自 growslice
		func growslice(oldCap, newLen int) int {
				newCap := oldCap
				doubleCap := newCap + newCap

				if newLen > doubleCap {
						newCap = newLen                    // 一次性追加大量元素
				} else if oldCap < 256 {
						newCap = doubleCap                 // 小切片：直接翻倍
				} else {
						for newCap < newLen {
								// 平滑过渡：从 2× 向 1.25× 收敛
								newCap += (newCap + 3*256) / 4
						}
				}
				return newCap
		}

	第二步：内存对齐（mallocgc size class）
		wcap 后，Go 会把 newcap × elemSize 的字节数向上对齐到内存分配器的 size class（8、16、24、32、48…字节），再用对齐后的字节数除以 elemSize 得到最终 cap。这就是为什么实际 cap 往往比公式结果再大一点。

	第三步：分配 + 拷贝
		go// 分配新底层数组
		newPtr := mallocgc(newcap * elemSize, ...)
		// 把旧数据拷贝过去（O(n) 操作！）
		copy(newPtr, oldPtr, oldLen * elemSize)
		// 返回新切片头部

	二、扩1个元素
	起始状态
	s := []int64{1, 2, 3, 4}   // len=4, cap=4
	s = append(s, 5)            // 触发扩容，因为 len+1=5 > cap=4
	公式算 newcap
	oldCap = 4
	newLen = 5   （需要容下 5 个元素）
	doubleCap = 4 + 4 = 8

	判断：
		newLen(5) > doubleCap(8)?  → 否
		oldCap(4) < 256?           → 是  ✓ int64最大取值范围256
	∴ newCap = doubleCap = 8

	第二步：内存对齐微调
		需要分配的字节 = newCap × elemSize
								= 8 × 8字节(int64)
								= 64 字节

		64 恰好命中 size class，无需向上取整
		∴ 最终 cap = 64字节 ÷ 8 = 8

	第三步：实际发生的事
		旧底层数组（cap=4）         新底层数组（cap=8）
		┌──┬──┬──┬──┐              ┌──┬──┬──┬──┬──┬──┬──┬──┐
		│1 │2 │3 │4 │   copy →     │1 │2 │3 │4 │5 │  │  │  │
		└──┴──┴──┴──┘              └──┴──┴──┴──┴──┴──┴──┴──┘
			被 GC 回收                      ↑
																	5 写入 index[4]

	三、扩5个元素
	起始状态
		s := []int64{1, 2, 3, 4}
		s = append(s, 5, 6, 7, 8, 9)

		oldCap = 4
		newLen = 4 + 5 = 9   ← 需要容下 9 个元素

	第一步：公式算 newcap
		doubleCap = 4 + 4 = 8
		判断：
			newLen(9) > doubleCap(8)?  → 是 ✓  ← 和上次不同！
		∴ newCap = newLen = 9        ← 直接取 newLen，不翻倍
		逻辑是：你要的比翻倍还多，翻倍也不够用，直接给你刚好够的数量，再交给第二步对齐。

	第二步：内存对齐
		需要分配的字节 = newCap × elemSize
								= 9 × 8字节(int64)
								= 72 字节

		size class 表：…64、80、96…
		72 字节 → 向上对齐到 80 字节
		∴ 最终 cap = 80字节 ÷ 8 = 10

		第三步：实际发生的事
		旧数组（cap=4）              新数组（cap=10）
		┌──┬──┬──┬──┐              ┌──┬──┬──┬──┬──┬──┬──┬──┬──┬──┐
		│1 │2 │3 │4 │   copy →     │1 │2 │3 │4 │5 │6 │7 │8 │9 │  │
		└──┴──┴──┴──┘              └──┴──┴──┴──┴──┴──┴──┴──┴──┴──┘
			被 GC 回收                                           ↑
																											预留 1 格
	*/

	// go切片没有内置 delete，所有删除本质都是拼接或覆盖，根据删除位置分三种情况：
	// 删除尾部元素
	ss1 := []int{1,2,3,4,5}
	fmt.Printf("prt：%p\n", &ss1)
	// 删除尾部1个元素
	ss1 = ss1[0:len(ss1) - 1] // {1,2,3,4}
	fmt.Printf("prt：%p\n", &ss1)
	// 删除最后 n个元素，删除最后2个元素
	n := 2
	ss1 = ss1[:len(ss1) - n] // {1,2}
	fmt.Printf("prt：%p\n", &ss1)
	fmt.Printf("拼接后到达删除效果：%v\n",ss1) // 拼接后到达删除效果：[1 2]
	fmt.Printf("len:%d, cap:%d\n", len(ss1), cap(ss1)) // len：2 cap：5
	/* 
	底层数组：[ 1 | 2 | 3 | 4 | 5 ]
	切片头：   ptr → 首地址，len=5，cap=5

	ss1 = ss1[0:len(ss1) - 1] 
	底层数组：[ 1 | 2 | 3 | 4 | 5 ]   ← 没动，5 还在内存里
                            ↑
                    [1,2,3,4] 5 还占着内存，只是看不到了
	切片头：   ptr → 首地址，len=4，cap=5   ← 只改了 len，可见元素1到4

	ss1 = ss1[:len(ss1) - 2]，得到 ss1[0:2] 此时 len=4-2
	底层数组：[ 1 | 2 | 3 | 4 | 5 ]   ← 还是同一个
                    ↑
              [1,2,] 3,4,5 还占着内存
	切片头：   ptr → 首地址，len=2，cap=5   ← len 再次缩小 
	两次修改到印地址完全一样，0x26d82c07c540
  */

	// 删除头部元素
	ss2 := []int{6,7,8,9,10,11}
	// 删除第一个元素，idx：0
	ss2 = ss2[1:] // 1:end [7 8 9 10 11]
	// 删除前n个元素, 删除前3个切片元素。本质是移动指针位置
	n1 := 3
	ss2 = ss2[n1:] // 索引从3开始，第4个元素开始
	fmt.Printf("拼接后到达删除效果：%v\n",ss2) // 拼接后到达删除效果：[10 11]
	fmt.Printf("len:%d, cap:%d\n", len(ss2), cap(ss2)) // len:2, cap:2
	// 同样只移动指针，O(1)。但注意底层数组头部的内存不会释放，如果原切片很大，考虑用 copy 版本。
	// 比如s := make([]int, 10000)   
	// s := s[2:] 删除前两个
	// 底层数组（80KB，无法被 GC）
	// ┌──┬──┬──┬──┬──┬── ... ──┐
	// │x │x │  │  │  │   ...│  │
	// └──┴──┴──┴──┴──┴── ... ──┘
  //    ↑
  //    ptr 移到这里，前两格永远悬空。但整块 80KB 内存都被底层数组占着，GC 无法回收
	
	// 只移动指针，O(1)。但要注意底层数组头部的内存不会释放，如果原切片很大，考虑用 copy 版本。
	// 比如 sBigSlice 底层有一个超大的数组，想跳过前 n 个元素
	sBigSlice := make([]int, 10000)   // 底层数组 10000 个元素，占 80KB
	n3 := 2
	// O(1) 指针移动（有内存隐患）
	smallSlice := sBigSlice[n3:]
	fmt.Println(len(smallSlice), cap(smallSlice)) // 9998 9998
	/*
	底层数组（80KB，无法被 GC）
	┌──┬──┬──┬──┬──┬── ... ──┐
	│x │x │  │  │  │   ...│  │
	└──┴──┴──┴──┴──┴── ... ──┘
				↑
				ptr 移到这里，前两格永远悬空
				但整块 80KB 内存都被底层数组占着，GC 无法回收
	*/
  // copy 版本（安全，O(len(smallSlice))）
	// 先初始化一个切片长度为 len(bigslice) - n的长度
	newSlice := make([]int, len(sBigSlice) - n3)
	copy(newSlice, sBigSlice[n3:]) // 将sBigSlice的从n3处截取切片,截取内容复制给新切片

	sBigSlice = newSlice  // 原sBigSlice指向新数组 旧的 80KB 底层数组没有任何引用了，GC 可以回收
	/*
		旧底层数组（80KB）             新底层数组（79984B）
	┌──┬──┬──┬── ... ──┐            ┌──┬──┬── ... ──┐
	│x │x │v │   ...│v │  copy →    │v │v │   ...│v │
	└──┴──┴──┴── ... ──┘            └──┴──┴── ... ──┘
		没有引用了，等待 GC                  ↑
																		sBigSlice 指向这里
	*/

	// 删除中间元素 第i个元素 最常用的
	ss3 := []int{1,2,3,4,5,6,7,8}
	n4 := 2
	// 通过append方式实现拼接将元素删除
	// ss[:n] = [1,2]   ss3[n+1] = [5,6,7,8]
	ss3 = append(ss3[:n4], ss3[n4+1:]...) // ss3[n+1:]... 将ss3[3:]展开运算，变成切片的多个元素，追加到新的ss3中
	fmt.Printf("%v\n", ss3)

	// 不保序（用尾部元素覆盖）
	ss5 := []string{"A", "B", "C", "D", "E"}
	// 删除第二个元素b，想让b元素索引位置等于ss5切片末尾最后一个元素
	n5 := 1
	ss5[n5] = ss5[len(ss5)-1]  // B替换为E
	ss5 = ss5[:len(ss5)-1] // ["A", "B", "C", "D"] "E" 将最后一个E元素截掉
	fmt.Println(ss5) // 输出: [A E C D] (顺序变了，但 "B" 被成功删除了)
	
	// 通过slices标准库的delete函数
	ss4 := []string{"apple", "banana", "cherry", "date"}
	slices.Delete(ss4, 1, 2) // 删除第二个元素，Delete(目标切片,起始idx,结束idx)  左闭右开 [start, end)
	fmt.Printf("%v\n", ss4) // [apple cherry date ]

	// 批量删除（filter 模式）
	ss6 := []int{1, 2, 3, 4, 5, 6}
	// 删除所有偶数，逆向思维找奇数
  j := 0
	for _, v := range ss6 {
		if v %2 != 0 { // 当元素取余部位0,为奇数
			ss6[j] = v // 将为奇数从索引0开始
			j++ // 记录奇数索引也是个数
		}
	}
	// ss6 [1, 3, 5, 4, 5, 6] j为2，最后截取ss6最有的j索引处位置  
	ss6 = ss6[:j]
	fmt.Println(ss6) // [1 3 5]


}