零碎go相关小记

### 控制字符

控制字符是**不可显示的字符**，它们用于控制设备或文本格式，而不是显示为可见符号。
最常用的控制字符

|字符|转义序列|ASCII码|说明|
|---|---|---|---|
|**换行**|`\n`|10|光标移到下一行|
|**制表符**|`\t`|9|水平对齐，跳到下一个tab位置|
|**回车**|`\r`|13|光标回到行首|
|**退格**|`\b`|8|删除前一个字符|
|**空字符**|`\x00`|0|字符串结束标志|
|**转义字符**|`\x1B`|27|ANSI转义序列的开始|

示例代码
```go
import "fmt"

func main() {
    // 1. 换行符 \n
    fmt.Println("Hello\nWorld")
    // 输出：
    // Hello
    // World

    // 2. 制表符 \t（对齐）
    fmt.Println("姓名\t年龄\t城市")
    fmt.Println("张三\t25\t北京")
    // 输出：
    // 姓名    年龄    城市
    // 张三    25      北京

    // 3. 回车符 \r（覆盖行首）
    fmt.Println("加载中\r完成！")
    // 输出：完成！中（"加载中"被覆盖）

    // 4. 空字符 \x00
    str := "A\x00B"  // 中间有个空字符
    fmt.Printf("%q\n", str)  // "A\x00B"

    // 5. 完整的控制字符范围：0x00-0x1F 和 0x7F
    for i := 0; i <= 31; i++ {
        fmt.Printf("\\x%02x ", i)
    }
    // 输出：\x00 \x01 \x02 ... \x1f
}
```


### if 的特殊语法
```go
// if 的特殊语法：初始化语句 + 条件
if 初始化语句; 条件 {
    // 条件为 true 时执行
}

if v, err := strconv.Atoi(config["port"]); err != nil {
    // 错误处理分支
} 
```

优点1：变量作用域限制

```go
// ✅ 推荐：v 和 err 只在 if 块内可见
if v, err := strconv.Atoi("8080"); err != nil {
    // 错误处理
} else {
    fmt.Println(v)  // 可以使用 v
}
// fmt.Println(v)  // ❌ 错误：v 在这里不可见

// ❌ 不推荐：变量污染外部作用域
v, err := strconv.Atoi("8080")
if err != nil {
    // 错误处理
}
fmt.Println(v)  // v 在整个函数内都可见
```

优点2：简洁优雅

```go
// ✅ 简洁
if v, err := strconv.Atoi(config["port"]); err != nil {
    return err
}
// 使用 v...

// ❌ 冗长
portStr := config["port"]
v, err := strconv.Atoi(portStr)
if err != nil {
    return err
}
// 使用 v...
```


### interface{}是什么意思？

解释下`params := make(map[string]interface{})`中得`interface{}`是什么意思？
 Go 语言中，`interface{}` 表示**空接口（empty interface）**

- `interface{}` 是一个没有任何方法要求的接口
- 在 Go 中，**所有类型都实现了空接口**
- 因此它可以存储**任何类型的值**

```go
params := make(map[string]interface{})
```
- 是一个 map集合，键是 `string` 类型
- **值可以是任何类型**（因为使用了 `interface{}`）

实际用途示例
```go
params := make(map[string]interface{})

// 可以存储不同类型的值
params["name"] = "张三"         // string
params["age"] = 25            // int
params["height"] = 175.5      // float64
params["isStudent"] = true    // bool
params["scores"] = []int{85, 90, 95}  // slice
params["info"] = struct {     // 结构体
    City string
    Job  string
}{
    City: "北京",
    Job:  "工程师",
}
```
需要类型断言来获取实际值
```go
// 直接使用会报错
// fmt.Println(params["age"] + 1) // 错误！

// 正确方式：类型断言
if age, ok := params["age"].(int); ok {
    fmt.Println(age + 1)  // 26
}

// 类型判断
switch v := params["name"].(type) {
case string:
    fmt.Println("字符串:", v)
case int:
    fmt.Println("整数:", v)
default:
    fmt.Println("其他类型")
}
```

与具体类型接口的区别
```go
// 空接口 - 可以接受任何值
var a interface{} = "hello"
a = 123
a = true

// 具体接口 - 只能接受实现了该接口的类型
type Writer interface {
    Write([]byte) (int, error)
}
// 只能赋值实现了 Write 方法的类型
```

Go 1.18 引入泛型后，某些场景可以用泛型替代 `interface{}`：

```go
// 使用泛型
func Print[T any](value T) {
    fmt.Println(value)
}

// 替代原来的 interface{} 方式
func PrintOld(value interface{}) {
    fmt.Println(value)
}
```


`select` 是 **Go 语言特有**的控制结构，专门用于处理多个 channel 的发送/接收操作。

 基本语法

```go
select {
case msg := <-ch1:
    // 从 ch1 接收数据
case ch2 <- value:
    // 向 ch2 发送数据
case <-time.After(1 * time.Second):
    // 超时处理
default:
    // 没有通道就绪时执行
}
```

 核心特性

|特性|说明|
|---|---|
|**随机选择**|多个 case 同时就绪时，随机选择一个执行|
|**阻塞/非阻塞**|没有 default 时会阻塞；有 default 时非阻塞|
|**仅用于 channel**|case 语句只能操作 channel|
|**只执行一次**|每次调用 select 只执行一个 case|

基本示例

```go
ch1 := make(chan string, 1)
ch2 := make(chan string, 1)

ch1 <- "消息1"

select {
case msg := <-ch1:
    fmt.Println("从 ch1 收到:", msg)  // 会执行这个
case msg := <-ch2:
    fmt.Println("从 ch2 收到:", msg)
default:
    fmt.Println("没有通道就绪")
}
// 输出: 从 ch1 收到: 消息1
```

```go
timer2 := time.NewTimer(5 * time.Second)

go func() {
    time.Sleep(1 * time.Second)
    timer2.Stop()  // 1秒后取消定时器
}()

select {
case <-timer2.C:
    fmt.Println("timer2定时器被触发了")
default:
    fmt.Println("timer2定时器被取消了")
}
```

**执行流程：**
1. 创建一个 5 秒的定时器
2. 启动 goroutine，1 秒后调用 `timer2.Stop()` 取消定时器
3. `select` 立即检查：
    - `timer2.C` 还未收到数据（定时器被取消了）
    - 执行 `default` 分支
4. 输出："timer2定时器被取消了"


### panic 关键字详解
panic 是 Go 语言中的内置函数，用于处理运行时异常/严重错误。

基本用法

panic("出错了！")
当 panic 被调用时：

立即中断当前函数的执行
逐层向上返回（执行 defer 语句）
打印错误信息和堆栈跟踪
程序崩溃退出（除非被 recover 捕获）

```go
location, err := time.LoadLocation("America/New_York")
if err != nil {
    panic(err)  // 如果加载时区失败，直接崩溃
}
```
含义： 如果加载时区失败，打印错误信息并终止程序。

panic vs error（什么时候用哪个）
场景  使用  示例
可预见的错误  error 文件不存在、网络超时、格式解析失败
不可恢复的严重错误 panic 数组越界、空指针、配置文件缺失导致无法启动

```go
// ✅ 使用 error - 可处理的问题
file, err := os.Open("config.json")
if err != nil {
    log.Println("配置文件不存在，使用默认配置")
    return  // 返回 error，让调用者决定怎么处理
}
```

```go
// ❌ 使用 panic - 程序无法继续运行
config := loadConfig()
if config.DatabaseURL == "" {
    panic("数据库 URL 未配置！程序无法启动")
}
```