package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Atoi：字符串转int（常用）
	var str string = "123"
	// Atoi 等同于 ParseInt(s, 10, 0)，转换为 int 类型。
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("转换失败", err)
	} else {
		// %T输出值的类型，注意int32和int是两种不同的类型
		fmt.Printf("Atoi: %s -> %d（类型：%T）\n", str, num, num) // Atoi: 123 -> 123（类型：int）
	}

	// 1.Itoa: int转字符串
	var intVal int = 4785
	strVal := strconv.Itoa(intVal)
	fmt.Printf("Itoa：%d -> %s (类型：%T)\n", intVal, strVal, strVal) // Itoa：4785 -> 4785 (类型：string)

	// 2.ParseInt：字符串转int，可指定进制和位数
	// ParseInt(s string, base int, bitSize int)
	// base: 进制（2-36），0表示自动推断；由字符串前缀（如果有符号则在符号之后）决定："0b" 为 2，"0" 或 "0o" 为 8，"0x" 为 16，否则为 10。
	// bitSize: 位数（0, 8, 16, 32, 64） 0、8、16、32 和 64 分别对应 int、int8、int16、int32 和 int64。如果 bitSize 小于 0 或大于 64，则会返回错误。
	// 自动推断类型（0x开头16进制，0开头为8进制）
	// ParseInt(s string, base int, bitSize int) (i int64, err error)
	num64, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println("ParseInt(自动推断)：", num64) // ParseInt(自动推断)： 123

	num0x, _ := strconv.ParseInt("0xFF", 0, 64) // 十六进制字符串0xff 转换为int64输出255
	fmt.Println("ParseInt(十六进制)：", num0x)       // ParseInt(十六进制)： 255

	num8, _ := strconv.ParseInt("0777", 0, 64) // 八进制字符串0777 转换为int64输出511
	fmt.Println("ParseInt(八进制)：", num8)        // ParseInt(八进制)： 511

	// 指定进制：base 设置指定进制进行转换
	num2, _ := strconv.ParseInt("1010", 2, 64)
	fmt.Println("二进制(1010)：", num2) // 二进制(自动推断)： 10

	num16, _ := strconv.ParseInt("ff", 16, 64)
	fmt.Println("十六进制(ff)：", num16) // 十六进制(自动推断)： 255

	num88, _ := strconv.ParseInt("777", 8, 64)
	fmt.Println("八进制(777)：", num88) // 八进制(777)： 511

	// 处理负数字符串
	fuNum, _ := strconv.ParseInt("-123", 10, 64)
	fmt.Printf("负数：(类型：%T) %d\n", fuNum, fuNum) // 负数：(类型：int64) -123

	// 3.formatInt:int64转字符串，可以指定进制
	// base进制2-36
	// 不同进制转换
	fmt.Println("formatInt(十进制)：", strconv.FormatInt(123, 10))  // formatInt(十进制)： 123
	fmt.Println("formatInt(16进制)：", strconv.FormatInt(255, 16)) // formatInt(16进制)： ff
	fmt.Println("formatInt(八进制)：", strconv.FormatInt(511, 8))   // formatInt(八进制)： 777
	fmt.Println("formatInt(2进制)：", strconv.FormatInt(10, 2))    // formatInt(2进制)： 1010

	// 负数处理
	intF := strconv.FormatInt(-123, 10)
	fmt.Printf("负数：(类型：%T) %s\n", intF, intF) // 负数：(类型：string) -123
	intF2 := strconv.FormatInt(-255, 10)
	fmt.Printf("负数：(类型：%T) %s\n", intF2, intF2) // 负数：(类型：string) -255

	// 零值
	zero := strconv.FormatInt(0, 10)
	fmt.Printf("零值(类型：%T) %s\n", zero, zero) // 零值(类型：string) 0

	// 4.strconv.AppendInt(dst []byte, value int64, base int) []byte
	// 向目标dst切片，追加base进制value内容，到dst切片中
	// 初始状态：buf = ['n', 'u', 'm', 'b', 'e', 'r', ':'] 即 "number:"
	// 追加后：buf = ['n', 'u', 'm', 'b', 'e', 'r', ':', '1', '2', '3'] 即 "number:123"
	buf := []byte("number:")               //
	buf = strconv.AppendInt(buf, 123, 10)  // 将整数123以十进制追加到buf字节切片中
	fmt.Println("字节切片", buf)               // 字节切片 [110 117 109 98 101 114 58 49 50 51]
	fmt.Println("AppendInt：", string(buf)) // AppendInt： number:123

	// 不同进制进行追加
	buf1 := []byte("Hex:")
	buf1 = strconv.AppendInt(buf1, 255, 16) // Hex:ff
	buf2 := []byte("Bin:")
	buf2 = strconv.AppendInt(buf2, 10, 2) // Bin:101
	buf3 := []byte("Oct:")
	buf3 = strconv.AppendInt(buf3, 511, 8) // Oct:777
	buf4 := []byte("Dec:")
	buf4 = strconv.AppendInt(buf4, 255, 10) // Dec:255
	fmt.Println(string(buf1), string(buf2), string(buf3), string(buf4))

	// 负数
	buf5 := []byte("负数:")
	buf5 = strconv.AppendInt(buf5, -12, 10) // 负数:-12
	fmt.Println("buf5", string(buf5))

	// 构建网络协议数据包
	buildPacker := func(data string, value int64) []byte {
		var buf []byte = []byte(data) // 创建数据字符串切片
		buf = strconv.AppendInt(buf, value, 10)
		return buf
	}
	packer := buildPacker("Not Found：", 404)
	fmt.Println(string(packer))

	// 日志格式化
	formatLog := func(level string, code int, message string) []byte {
		var buf []byte = []byte(level)
		buf = append(buf, '[')
		buf = strconv.AppendInt(buf, int64(code), 10)
		buf = append(buf, ']')
		buf = append(buf, ' ')
		// append函数的定义需要接收一个切片作为其第一个参数，以及零个或多个与其元素类型相同的值。
		// 通过使用...，你告诉编译器将 message 字符串中的每一个字符（字节）作为单独的参数传递给append函数。
		// 例如，如果 message是"hello", message...相当于传递五个独立的字节参数：'h', 'e', 'l', 'l', 'o'。
		// 在函数声明中的... (例如 func myFunc(args ...int)) 表示该参数是可变的，可以传入任意数量的参数。
		// 在函数调用中的... (例如 append(slice, elements...)) 表示将一个切片（或在这里是字符串）中的元素逐一展开作为单独的参数传入。
		buf = append(buf, message...)
		return buf
	}
	log := formatLog("wring", 302, "wring message")
	fmt.Println(string(log)) // wring[302] wring message

	fmt.Println("---无符号转字符串")

	// 5.无符号整数字符串转换
	unsignedConversion := func() {
		uvar, err := strconv.ParseUint("123", 10, 64)
		if err != nil {
			fmt.Println("转换失败", err)
		} else {
			fmt.Printf("ParseUint: %d(类型%T)\n", uvar, uvar) // ParseUint: 123(类型uint64)
		}

		// 有符号转换解析失败
		_, err = strconv.ParseUint("-123", 10, 64) // 这里的err 变量已经在前面被声明过了，不能再用 :=（短变量声明会导致重新声明错误），只能用 = 来重新赋值
		if err != nil {
			fmt.Println("ParseUint转换错误", err)
		}

		// FormatUint：无符号整数转字符串
		strFUint := strconv.FormatUint(100, 10)
		fmt.Printf("FormatUint：%s %T\n", strFUint, strFUint) // FormatUint：100 string

		// 大数处理
		var bigN uint64 = 18446744073709551615
		strBigN := strconv.FormatUint(bigN, 10)
		fmt.Printf("FormatUint：%s %T\n", strBigN, strBigN) // FormatUint：18446744073709551615 string
		// 将上面长字符串数解析为正常大数
		parsed, _ := strconv.ParseUint(strBigN, 10, 64)     // FormatUint：18446744073709551615 string
		fmt.Printf("ParseUint解析结果：%d %T\n", parsed, parsed) // ParseUint解析结果：18446744073709551615 uint64
	}
	unsignedConversion()

	// 浮点型转换数字
	floatConversion := func() {
		// ParseFloat：字符串转浮点型
		// 32 位用于 float32，或 64 位用于 float64。当 bitSize=32 时，结果仍然是 float64 类型，但它可以在不改变其值的情况下转换为 float32。

		// 1.基本浮点数
		f64, err := strconv.ParseFloat("3.14159", 64)
		if err != nil {
			fmt.Println("转换失败：", err)
		} else {
			fmt.Printf("转换浮点型：%f(类型%T)\n", f64, f64) // 转换浮点型：3.141590(类型float64)
		}

		// 2.科学计数法
		f64e, _ := strconv.ParseFloat("1.23e4", 64) // 1.23 * 10^4
		fmt.Printf("科学计数法：%f(类型%T)\n", f64e, f64e)  // 科学计数法：12300.000000(类型float64)

		// 负数
		f64f, _ := strconv.ParseFloat("-3.14", 64)
		fmt.Printf("负数：%f(类型%T)\n", f64f, f64f) // 负数：-3.140000(类型float64)

		// NaN特殊值
		FNaN, _ := strconv.ParseFloat("NaN", 64)
		fmt.Printf("特殊值：%v(类型%T)\n", FNaN, FNaN) // 特殊值：NaN(类型float64)

		// 正无穷 和 负无穷
		Zinf, _ := strconv.ParseFloat("Inf", 64)
		fmt.Printf("特殊值：%v(类型%T)\n", Zinf, Zinf) // 特殊值：+Inf(类型float64)
		Finf, _ := strconv.ParseFloat("-Inf", 64)
		fmt.Printf("特殊值：%v(类型%T)\n", Finf, Finf) // 特殊值：-Inf(类型float64)

		// FormatFloat：浮点数转字符串
		// FormatFloat(f float64, fmt byte, prec, bitSize int) string
		// 将浮点数 f 转换为字符串，根据格式 fmt 和精度 prec。 bitSize 位（32 位用于 float32，64 位用于 float64
		// 格式 fmt
		// 'b' (-ddddp±ddd，二进制指数), 二进制
		// 'e' (-d.dddde±dd，十进制指数), 科学技术法
		// 'E' (-d.ddddE±dd，十进制指数), 科学技术法
		// 'f' (-ddd.dddd，没有指数), 普通小数格式
		// 'g' (大指数用e，否则用f), 自动选择
		// 'x' (-0xd.ddddp±ddd, 一个十六进制分数和二进制指数), 或
		// 'X' (-0Xd.ddddP±ddd, 一个十六进制分数和二进制指数).
		// 精度 prec 控制由 'e', 'E', 'f', 'g', 'G', 'x', 和 'X' 格式打印的数字位数（不包括指数）。对于 'e', 'E', 'f', 'x', 和 'X'，它是小数点后的位数。
		// 对于 'g' 和 'G'，它是最大有效数字位数（尾部零将被移除）。特殊的精度 -1 使用最少数位的数字，使得 ParseFloat 将精确返回 f。指数以十进制整数形式书写；对于所有除 'b' 以外的格式，它至少是两位数

		const PI float64 = 3.141592653589793
		sPi := strconv.FormatFloat(PI, 'f', 2, 64)  // 保留小数点后2位精度
		sPi2 := strconv.FormatFloat(PI, 'f', 6, 64) // 保留小数点后2位精度
		fmt.Printf("普通格式f：%v（%T类型）\n", sPi, sPi)    // 普通格式：3.14（string类型）
		fmt.Printf("普通格式f：%v（%T类型）\n", sPi2, sPi2)  // 普通格式：3.141593（string类型）

		// 科学计数法
		kPi := strconv.FormatFloat(PI, 'e', 4, 64)
		kPiE := strconv.FormatFloat(PI, 'E', 4, 64)
		fmt.Printf("科学计数法e：%v（%T类型）\n", kPi, kPi)   // 科学计数法e：3.1416e+00（string类型）
		fmt.Printf("科学计数法E：%v（%T类型）\n", kPiE, kPiE) // 科学计数法E：3.1416E+00（string类型）

		// 大数
		// 当数字很大或很小时，自动用科学计数法（'e'）
		// 普通数字用小数格式（'f'）
		bigF := 123456789.123456789                    // 小数点后最多9位
		gBig := strconv.FormatFloat(bigF, 'g', -1, 64) // prec = -1：自动计算最小精度
		fmt.Printf("大数g：%v（%T类型）\n", gBig, gBig)       // 大数g：1.2345678912345679e+08（string类型）
		// prec = 6：保留 6 位有效数字
		gBig = strconv.FormatFloat(bigF, 'g', 6, 64) // "1.23457e+08"
		fmt.Printf("大数g：%v（%T类型）\n", gBig, gBig)
		// prec = 10：保留 10 位有效数字
		gBig = strconv.FormatFloat(bigF, 'g', 10, 64) // "123456789.1"
		fmt.Printf("大数g：%v（%T类型）\n", gBig, gBig)

		// AppendFloat 向浮点字节切片追加内容
		buf := []byte("pi: ")
		buf = strconv.AppendFloat(buf, 3.1415926, 'f', 2, 64)
		fmt.Printf("AppendFloat buf：%s \n", string(buf)) // AppendFloat buf：pi: 3.14
	}
	floatConversion()

	boolConversion := func() {
		// ParseBool：字符串转布尔值
		// ParseBool(str string) (bool, error)
		// 它接受 1、t、T、TRUE、true、True、0、f、F、FALSE、false、False字符串。任何其他值都会返回错误。
		boolArr := []string{
			"1", "t", "T", "true", "True", "TRUE",
			"0", "f", "F", "FALSE", "false", "False",
			"yes", "no", // 最后两个无效
		}

		for _, str := range boolArr {
			b, err := strconv.ParseBool(str)
			if err != nil {
				fmt.Printf("ParsBool 字符串布尔%q 错误: %v \n", str, err)
			} else {
				fmt.Printf("ParsBool %q(%T) ：%v(%T) \n", str, str, b, b)
			}
		}

		// FormatBool：布尔值转字符串
		bool1 := strconv.FormatBool(true)
		bool2 := strconv.FormatBool(false)
		fmt.Printf("FormatBool：%v 类型(%T)\n", bool1, bool1) // FormatBool：true 类型(string)
		fmt.Printf("FormatBool：%v 类型(%T)\n", bool2, bool2) // FormatBool：false 类型(string)

		// AppendBool：向目标字节切片追加bool
		buf := []byte("Status: ")
		buf = strconv.AppendBool(buf, false)
		fmt.Printf("AppendBool: %s\n", string(buf))
	}
	boolConversion()

	// 引号处理
	quoteFunction := func() {
		// Quote：添加双引号并转义特殊字符
		str := "Hello\tWorld\ngo\"语言\""
		fmt.Printf("原始字符串 %s\n", str)
		// 原始字符串 Hello	World
		// go"语言"

		// Quote：添加双引号并转义特殊字符
		fmt.Println("Quote", strconv.Quote(str)) // 输出：Quote "Hello\tWorld\ngo\"语言\""

		// QuoteToASCII：非ASCII字符转Unicode转义序列
		str2 := "Hello 世界"
		fmt.Println("QuoteToASCII：", strconv.QuoteToASCII(str2)) // QuoteToASCII： "Hello \u4e16\u754c"

		// QuoteToGraphic：转义非图形字符
		str3 := "hello 你好\t\n"
		// 会保留中文，对控制字符串转义 控制字符比如\n \t \r \b \x00
		fmt.Println("QuoteToGraphic：", strconv.QuoteToGraphic(str3)) // QuoteToGraphic： "hello 你好\t\n"
		// 对于emoji图标是图形字符，完全保留
		str4 := "hello 🌏"
		fmt.Println("emoji：", strconv.QuoteToGraphic(str4)) // emoji： "hello 🌏"

		// QuoteRune字符加引号
		fmt.Println("QuoteRune:", strconv.QuoteRune('A'))  // QuoteRune: 'A'
		fmt.Println("QuoteRune:", strconv.QuoteRune('\n')) // QuoteRune: '\n'
		fmt.Println("QuoteRune:", strconv.QuoteRune('世'))  // QuoteRune: '世'

		// QuoteRuneToASCII：字符转ASCII表示 只允许传入rune类型
		fmt.Println("QuoteRuneToASCII：", strconv.QuoteRuneToASCII('世')) // QuoteRuneToASCII： '\u4e16'
		
		// Unquote：去除引号（反转义）
		quoted := `"Hello\tWorld\n Go"`
		// %q	该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
		// fmt.Printf("原始字符串：%q", quoted) // 原始字符串："\"Hello\\tWorld\\n Go\""

		unquoted, err := strconv.Unquote(quoted)
		if err != nil {
			fmt.Println("Unquote错误:", err)
		} else {
			fmt.Printf("原字符串 %q	->去除后 %q\n", quoted, unquoted) // 原字符串 "\"Hello\\tWorld\\n Go\""	->去除后 "Hello\tWorld\n Go"
		}
		// 无法取消引用不带引号的字符串
		s1, err := strconv.Unquote("法取消引用不带引号的字符串")
		// 会报错误，报无法取消带引号字符串
		fmt.Println(s1, err) //  invalid syntax

		s2, err := strconv.Unquote("\"字符串必须使用双引号或者单引号或反引号\"")
		fmt.Println(s2, err) // 字符串必须使用双引号或者单引号或反引号 <nil>

		s3, err := strconv.Unquote("`or backquoted.`")
		fmt.Printf("%q, %v\n", s3, err) // "or backquoted.", <nil>

		// 单个字符使用在单引号中
		s4, err := strconv.Unquote("'\u263a'") // 使用Unicode值
		fmt.Println("单个字符使用在单引号中:", s4, err) // 单个字符使用在单引号中: ☺ <nil>

	}
	quoteFunction()
}
