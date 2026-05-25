package main

import (
	"fmt"
	"time"
)

func main() {
	// 数组 固定长度、相同元素的集合，一旦声明长度不可改变
	// 一、声明数组
	// 每个元素是一个整型值，当声明数组时所有的元素都会被自动初始化为默认值 0。
	var arr1 [5]int   // 声明了一个长度为5组，默认值都是0
	fmt.Println(arr1) // [0 0 0 0 0]

	// 声明并初始化
	var arr2 = [5]int{51, 25, 30, 50, 63}
	fmt.Println(arr2) // [51 25 30 50 63]

	// 初始化部分元素
	var arr3 = [5]int{51, 36, 9}
	fmt.Println(arr3) // [51 36 9 0 0]

	// 编译器自己推断数组长度
	var arr4 = [...]int{12, 25, 35, 40}
	fmt.Println(len(arr4)) // 4

	// 指定索引初始化
	var arr5 = [5]int{0: 10, 4: 82} // arr5[0]=10, arr5[4]=82, 其他为0
	fmt.Println(arr5)               // [10 0 0 0 82]

	// 简短声明
	arr6 := [3]string{"hello", "gg", "ok"}
	fmt.Println(arr6) // [hello gg ok]

	arr7 := [5]int{54, 85, 68, 56, 8}
	// 第一个元素是 arr1[0]，第三个元素是 arr1[2]；总体来说索引 i 代表的元素是 arr1[i]，最后一个元素是 arr1[len(arr1)-1]。
	// 访问元素
	fmt.Println(arr7[2]) // 68
	// 修改元素
	arr7[0] = 100
	fmt.Println(arr7[0])
	// 最后一个元素
	fmt.Println("end", arr7[len(arr7)-1])

	// [5]int和 [10]int 是属于不同类型的。
	var a1 [5]int = [5]int{0, 1, 2, 3, 4}  // 类型是 [5]int
	var b1 [10]int = [10]int{0: 10, 9: 90} // 类型是 [10]int

	// 在go中数组是值类型，一个 [5]int 变量就是连续 5 个 int 大小的内存块，一个 [10]int 是连续 10 个 int 大小的内存块。它们的尺寸根本不一样，在内存布局、复制开销上都完全不同。
	// 长度是类型的一部分，数组的大小在编译时就是固定的。当你把一个数组赋值给另一个变量，或者传给函数时，Go 会复制整个数组的所有元素。
	// a1 = b1 // 编译错误：cannot use a (type [5]int) as type [10]int in assignment
	fmt.Printf("a1: %d\nb1: %d\n", a1, b1)

	// 切片的类型只有 []int，长度不是类型的一部分。切片的类型不包含长度，所以无论长度是 5 还是 10 的切片，它们的类型都是 []int，可以自由地互相赋值、扩容，或者作为同一个函数的参数。
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("s1:%p val:%d \n", s1, s1) // s1:0xc0000182e8 val:[1 2 3]
	fmt.Printf("s2:%p val:%d \n", s2, s2) // s2:0xc0000200a0 val:[1 2 3 4 5 6 7 8 9 10]
	s1 = s2                               // 都是 []int 类型
	fmt.Printf("s1:%p val:%d \n", s1, s1) // s1:0xc0000200a0 val:[1 2 3 4 5 6 7 8 9 10]
	fmt.Printf("s2:%p val:%d \n", s2, s2) // s2:0xc0000200a0 val:[1 2 3 4 5 6 7 8 9 10]

	// 二、数组指针声明
	// & — 取地址，&a 得到 a 的指针
	// * — 解引用，*p 读写指针指向的值
	numP := [3]int{1, 2, 3}
	// 2.1.数组指针：指向数组的指针
	p := &numP           // go中使用&获取一个变量的指针地址  p *[3]int   &符号：显式取地址
	fmt.Println(p)       // p指针类型类型为：*[3]int  &[1 2 3]
	fmt.Println(*p)      // 解引用地址 得到数组  [1 2 3]
	fmt.Println((*p)[0]) // *p 读写指针指向的值，获取数组中的第一个元素
	fmt.Println(p[0])    // 简写 在 Go 中，不需要手动解引用（(*arrPtr)[0]），可以直接使用 arrPtr[0]
	// 通过指针修改原数组
	p[0] = 100 // 等价于 (*p)[0] = 100
	fmt.Println(p[0])

	// 2.2new 创建数组
	p1 := new([3]int) // *[3]int，默认值[0,0,0]
	// 实际开发中 new 用得很少，&T{...} 更常用，既能拿到指针又能初始化。
	p1[0] = 10
	p1[1] = 20
	p1[2] = 30
	fmt.Println(p1) // &[10 20 30]
	// new([3]int)           // → &[0, 0, 0]，只能得到零值
	// &[3]int{10, 20, 30}  // → &[10, 20, 30]，声明时就能赋值

	// 2.3数组指针可以当切片用
	arrNum := [5]int{15, 35, 55, 75, 95}
	arrP := &arrNum // 取地址，得到arrP指针  arrP类型：*[3]int

	// 直接对指针切片，得到 []int
	sliceP := arrP[1:4] //  直接对指针切片，得到 []int{35,55,75}切片[a,b)
	fmt.Println(sliceP) // 指针切片[]int{35,55,75}

	sliceP[0] = 40      // 修改切片会影响原数组
	fmt.Println(arrNum) // [15 40 55 75 95]

	// 2.4数组指针的使用场景
	// 2.4.1数组是值类型，传参数会复制整个数组，可以使用指针避免拷贝，避免大数组拷贝开销
	// 数组很大时，传值会复制整个数组，传指针只复制一个地址（8字节）：
	// 1e6 = 1 × 10⁶ = 1,000,000（一百万）。
	// 在传递函数参数 p := &arrM，p类型*[1e6]int
	var sumM func(p *[1e6]int) int = func(p *[1e6]int) int { // 通过指针解引用，得到数组地址，传指针，避免复制 100 万个 int
		// fmt.Println(p) 会输出整个100万数组
		total := 0
		for _, v := range p {
			total += v
		}
		return total
	}
	arrM := [1e6]int{}
	arrM[0] = 10
	arrM[99] = 20
	fmt.Println("传指针，零拷贝", sumM(&arrM)) //  传地址，不复制整个数组，零拷贝 30

	// 2.4.2封装函数内修改原数组
	// Go 数组是值类型，传入函数默认是拷贝，用指针才能修改原数组：
	var double = func(p *[3]int) {
		for idx := range p {
			(*p)[idx] *= 2 // 简写为p[idx]  *p解引用取得指针指向的值
		}
	}
	arrM2 := [3]int{1, 2, 3}
	double(&arrM2)
	fmt.Println("arrM2", arrM2) // arrM2 [2 4 6]

	// 对比函数不用数组指针
	arrM3 := [3]int{1, 2, 3}
	var doubleErr = func(p [3]int) {
		for idx := range p {
			p[idx] *= 2 // 修改的是副本p，原数组不变
		}
		fmt.Println("副本p", p) // 副本p[2 4 6]
	}
	doubleErr(arrM3)
	fmt.Println("arrM3", arrM3) // 原数组没有修改arrM3 [1 2 3]

	// 3.指针数组：元素是指针的数组
	a, b, c := 1, 2, 3
	ptrArr := [...]*int{&a, &b, &c} // &获取变量指针地址，类型 [3]*int
	fmt.Println(ptrArr)             // [0xc0000122a0 0xc0000122a8 0xc0000122b0]
	fmt.Println(*ptrArr[0])         // 取第一个指针指向的值

	// 通过 *解引用地址进行修改底层变量的值
	*ptrArr[1] = 20                 // 修改 b 的地址里存的值
	fmt.Println("b val change:", b) // b变量的值被改为: 20

	// 3.1指针数组的使用场景
	// 多个变量需要统一管理、批量处理
	ptrArrFn := func() {
		x, y, z := 10, 20, 30
		// 将变量收集到一个指针*[T]数组中
		ptrs := [...]*int{&x, &y, &z}

		// 循环批量修改变量
		for _, p := range ptrs {
			// 这里的p是复制的副本是指针指向的变量
			*p = *p * 2 //简写 *p *= 2
		}
	}
	ptrArrFn()

	// 3.2多个变量共享同一个数据
	shareVar := func() {
		type Config struct{ Value int }    // 创建一个结构体
		var cfg Config = Config{Value: 10} // 引用cfg指针实例和这个value值

		// 多处持有同一个 Config 的指针
		var handlers [3]*Config = [...]*Config{&cfg, &cfg, &cfg}
		// 修改源数据，所有引用cfg.value同步更新
		cfg.Value = 200

		for i, v := range handlers {
			fmt.Printf("h[%d].Value=%d\n", i, v.Value)
		}
	}
	shareVar()
	// 3.3 当数组中很多元素可能是 nil 时，从而判断哪些是有值的对象。
	// 使用指针数组可以节省大量内存。
	scores := [5]int{0, 98, 0, 76, 0} // 哪些是真的 0 分？哪些是没填？
	fmt.Println(scores)

	// 指针数组：nil 表示未设置
	ptrScores := [5]*int{}
	ss1, ss2 := 0, 98
	ptrScores[1] = &ss1 // 明确填了 0 分
	ptrScores[3] = &ss2 // 填了 98 分
	// ptrScores[2] 是 nil，表示未填写

	for idx, ptrS := range ptrScores {
		if ptrS == nil {
			fmt.Printf("第%d题没有填写\n", idx+1)
		} else {
			fmt.Printf("第%d题：%d 分\n", idx+1, *ptrS)
		}
	}

	// 3.4结构体对象引用
	type User struct {
		Name string
		Age  int
	}

	// 指针数组存储多个用户的引用（可能来自不同数据源）
	printUsers := func(users [3]*User) {
		for _, u := range users {
			if u != nil {
				fmt.Printf("%s (%d)\n", u.Name, u.Age)
			}
		}
	}
	var users [3]*User
	u1 := User{"Alice", 30}
	u2 := User{"Bob", 25}
	users[0] = &u1
	users[1] = &u2
	// users[2] 为 nil
	printUsers(users)

	// // 3.5接口指针数组实现多态
	// type Animal interface { Speak(s string) string } // 声明接口
	// // 创建结构体
	// type Dog struct { Name string }
	// type Cat struct { Name string }

	// // 每个结构体都实现了接口中方法
	// func (d *Dog) Speak(s string) string{ return d.Name+":"+s }
	// func (c *Cat) Speak(s string) string{ return c.Name+":"+s}

	// d1 := &Dog{Name:"柯基"}
	// c1 := &Cat{Name:"喵喵"}
	// animals := [2]Animal{d1, c1}
	// for _, a := range animals {
	// 	fmt.Println(a.Speak("嗨嗨"))
	// }

	// 区分两个概念：
	// *[3]int    // 数组指针 — 一个指针，指向一个数组
	// [3]*int    // 指针数组 — 一个数组，里面存了3个指针
	// 数组指针 *[3]int：
	// p ──→ [10, 20, 30]
	// 			一个指针指向整个数组

	// 指针数组 [3]*int：
	// [0xc01, 0xc02, 0xc03]
	// 	↓       ↓       ↓
	// 	a=1    b=2    c=3
	// 	数组里存的是3个指针

	//三、多维数组
	// 二维数组
	var arrR [2][3]int = [2][3]int{
		{25, 15, 26},
		{15, 35, 36},
	}
	fmt.Printf("[1][2]:%d\n", arrR[1][2])

	// 简短声明
	arrR2 := [2][3]int{
		{25, 15, 26},
		{15, 35, 36},
	}
	fmt.Printf("arrR2:%d\n", arrR2)

	// 自动推导
	arrR3 := [...][3]int{ // [...][...]int❌（只能最外层用...)
		{25, 15, 26},
		{15, 35, 36},
	}
	fmt.Printf("arrR3:%d\n", arrR3)

	// 读取多维数组元素
	fmt.Println(arrR3[1])    // [15 35 36]  第二行
	fmt.Println(arrR3[0][2]) // 26   第1行，第3个元素

	// 修改第二行，第一个元素
	arrR3[1][0] = 266
	fmt.Println("第二行元素", arrR3[1]) // 第二行元素 [266 35 36]

	// 遍历二维数组
	for i := 0; i < len(arrR3); i++ { // len(arrR3)先获取长度有几行
		for j := 0; j < len(arrR3[i]); j++ { // len(arrR3[i])再获取每行有几个元素
			fmt.Printf("arr[%d][%d]:%d \t", i, j, arrR3[i][j])
		}
		fmt.Println()
	}

	// 三维数组
	// [x][y][z]int：x个二维数组，每个二维数组中包含y个一维数组，每个一维长度z个
	arr3d := [2][2][4]int{
		{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
		},
		{
			{10, 20, 30, 40},
			{50, 60, 70, 80},
		},
	}
	// 获取[1][0][3]:40
	fmt.Println("[1][0][3]:", arr3d[1][0][3])
	// 遍历三维数组 [2][2][4]int
	for i := 0; i < len(arr3d); i++ { // 先获x个二维数组
		for j := 0; j < len(arr3d[i]); j++ { // 再获取每个二数组，有len(arr3d[i])个一维数组
			for k := 0; k < len(arr3d[i][j]); k++ { // 最后每个一维数组长度len(arr3d[i][j])
				// fmt.Println(len(arr3d[i][j]))
				fmt.Printf("arr[%d][%d][%d]:%d \t", i, j, k, arr3d[i][j][k])
			}
		}
	}
	fmt.Println()
	// 四、数组元素类型为任意类型
	// any即interface{} 作为元素类型(any是Go 1.18+引入）
	// var a [3]interface{}  // 旧写法
	// var b [3]any          // 新写法，推荐
	anyArr := [5]any{1, true, "hellp", 3.14, []int{1, 2, 3}}
	for i, v := range anyArr {
		fmt.Printf("anyArr[%d]=%v, 类型：%T\n", i, v, v)
	}
	// anyArr[0]=1, 类型：int
	// anyArr[1]=true, 类型：bool
	// anyArr[2]=hellp, 类型：string
	// anyArr[3]=3.14, 类型：float64
	// anyArr[4]=[1 2 3], 类型：[]int

	// 使用时，类型断言取回具体的类型
	arrT := [3]any{51, true, "ok"}

	// 安全断言推荐，用 ok 避免 panic
	if num, ok := arrT[0].(int); ok { // 判断是否为int类型
		fmt.Println(num + 10) // 61
	}

	// 不安全断言（类型不匹配会 panic）
	status := arrT[2].(string) //不对会 panic 不可恢复的致命错误
	fmt.Println(status)
	// 类型不匹配 → panic
	// n := arrT[1].(int) ❌
	// fmt.Println(n) panic: interface conversion: interface {} is bool, not int

	// 使用type switch处理多种类型
	process := func(arr [6]any) {
		for idx, v := range arr { // 遍历每个元素，通过v.(type)
			switch val := v.(type) { // num, ok := element.(type)
			case int:
				fmt.Printf("arr[%d] 整数(%T): %d\n", idx, val, val)
			case string:
				fmt.Printf("arr[%d] 字符串(%T): %d, 长度%d\n", idx, val, val, len(val))
			case bool:
				fmt.Printf("arr[%d] 布尔(%T): %v\n", idx, val, val)
			case float64:
				fmt.Printf("arr[%d] 浮点(%T): %v\n", idx, val, val)
			case []int:
				fmt.Printf("arr[%d] int切片(%T): %v\n", idx, val, val)
			default:
				fmt.Printf("arr[%d] 未知类型(%T): %v\n", idx, val, val)
			}
		}
	}
	anyP := [6]any{1, true, "hello", 3.14, []int{1, 2, 3}, time.Now()}
	process(anyP)

	// 元素类型有约束范围，用泛型比 any 更安全
	var arrF [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(sumArr(arrF))
	
	var arrF2 [5]float64 = [5]float64{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println(sumArr(arrF2))
	// sumArr([5]bool{false,true,false,true,false}) // ❌泛型已经约束参数类型

	// 混合类型配置
	type ConfigItem struct {
		Key string
		Value any // 值为任意值
	}

	configs := [5]ConfigItem{
		{"host",     "localhost"},
		{"port",     8080},
		{"debug",    true},
		{"timeout",  30.5},
		{"tags",     []string{"prod", "v2"}},
	}

	for _, cfgItem := range configs {
		switch val := cfgItem.Value.(type) {
		case string:
			fmt.Printf("%-10s = %q\n", cfgItem.Key, val)
		case int:
			fmt.Printf("%-10s = %d\n", cfgItem.Key, val)
		case bool:
			fmt.Printf("%-10s = %v\n", cfgItem.Key, val)
		case float64:
			fmt.Printf("%-10s = %.1f\n", cfgItem.Key, val)
		case []string:
			fmt.Printf("%-10s = %v\n", cfgItem.Key, val)
		}
	}

	
	// 五、遍历数组 ￼￼[1](https://www.runoob.com/go/go-fmt-sprintf.html)
	iterateArr()
}

// 如果元素类型有约束范围，用泛型比 any 更安全
type Number interface {
	int | int64 | float32 | float64
}

// 泛型函数的参数作为接收类型，Number
func sumArr[T Number](arr [5]T) T {
	var total T
	for _, v := range arr {
		total += v
	}
	return total
}

func iterateArr() {
	// 1.普通for遍历（最基础）
	nums := [5]int{10, 20, 30, 40, 50}
	for i := 0; i < len(nums); i++ {
		fmt.Printf("num[%d]:%d\n", i, nums[i])
	}

	// 2.range 遍历（Go 最常用）
	nums2 := [5]int{10, 20, 30, 40, 50}
	// 同时获取索引和值。idx为数组下标，val数组元素
	for idx, val := range nums2 {
		fmt.Printf("num2[%d]:%d\n", idx, val)
	}

	// 只保留数组值，使用_空白标识符丢弃索引
	for _, val := range nums2 {
		fmt.Printf("value:%d\n", val)
	}

	// 只保留索引idx
	for idx := range nums2 {
		fmt.Printf("index:%d\n", idx)
	}

	// 遍历时修改数组元素
	// 使用普通 基础for 可以修改
	nums3 := [5]int{10, 20, 30, 40, 50}
	for i := 0; i < len(nums3); i++ {
		// 对每个元素都/10
		nums3[i] /= 10
		nums3[i] *= 2
	}
	fmt.Println(nums3) // [2 4 6 8 10]

	// range 修改元素存在坑（重要）
	nums4 := [5]int{10, 20, 30, 40, 50}
	for _, val := range nums4 {
		val /= 10 // 但不会修改原数组
	}
	fmt.Println(nums4) // [10 20 30 40 50] 通过range赋值val变量是元素副本，不是原元素。
	// for range 的 value 是数组元素的副本，修改它不会影响原数组。要修改原数组，必须通过索引操作。
	// 通过数组index索引修改
	for idx := range nums4 {
		nums4[idx] /= 10 // 但不会修改原数组
	}
	fmt.Println(nums4) // [1 2 3 4 5]

}
