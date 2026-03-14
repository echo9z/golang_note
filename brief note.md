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

### switch特殊语法
`switch` 不写条件表达式是 Go 语言的一个特性，它可以实现 `if-else if-else` 的效果。

`switch` 不写条件表达式
基本语法
```go
switch {
case condition1:
    // 当 condition1 为 true 时执行
case condition2:
    // 当 condition2 为 true 时执行
default:
    // 当所有条件都不为 true 时执行
}
```

等价于 `if-else if-else`
```go
// switch 写法
switch {
case score >= 90:
    fmt.Println("优秀")
case score >= 80:
    fmt.Println("良好")
case score >= 60:
    fmt.Println("及格")
default:
    fmt.Println("不及格")
}

// 等价的 if-else if 写法
if score >= 90 {
    fmt.Println("优秀")
} else if score >= 80 {
    fmt.Println("良好")
} else if score >= 60 {
    fmt.Println("及格")
} else {
    fmt.Println("不及格")
}
```

**执行逻辑**：每个 `case` 是一个**布尔表达式**，为 `true` 就执行
为什么允许不写条件表达式？
 1. 灵活性
可以处理复杂的条件判断，不仅仅是相等比较
```go
switch {
case x > 0 && y > 0:
    fmt.Println("第一象限")
case x < 0 && y > 0:
    fmt.Println("第二象限")
case x < 0 && y < 0:
    fmt.Println("第三象限")
case x > 0 && y < 0:
    fmt.Println("第四象限")
}
```

2. 可读性
当有多个条件时，比 `if-else if` 更清晰
```go
// switch 写法（更清晰）
switch {
    case unicode.IsLetter(r):
        fmt.Println("字母")
    case unicode.IsDigit(r):
        fmt.Println("数字")
    case unicode.IsSpace(r):
        fmt.Println("空白")
    default:
        fmt.Println("其他")
}

// if-else 写法（较冗长）
if unicode.IsLetter(r) {
    fmt.Println("字母")
} else if unicode.IsDigit(r) {
    fmt.Println("数字")
} else if unicode.IsSpace(r) {
    fmt.Println("空白")
} else {
    fmt.Println("其他")
}
```
 
 Go 语言 ✅ 支持
```go
// Go 可以不写条件表达式
switch {
case score >= 90:
    fmt.Println("优秀")
case score >= 80:
    fmt.Println("良好")
default:
    fmt.Println("其他")
}
```

JavaScript ❌ 不支持
```javascript
// JavaScript 不允许不写条件表达式
switch {  // ❌ 语法错误
    case score >= 90:  // ❌ 语法错误
        console.log("优秀");
        break;
}

// JavaScript 必须写条件表达式
switch (true) {  // ✅ 必须写 true
    case score >= 90:
        console.log("优秀");
        break;
    case score >= 80:
        console.log("良好");
        break;
}
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


### Go 中有几种"抛出错误"的方式：
1. `panic()` - 程序崩溃（类似 throw）

```go
// 直接抛出错误信息
panic("发生严重错误！")

// 抛出错误对象
panic(errors.New("文件不存在"))

// 抛出格式化错误
panic(fmt.Sprintf("无效的值: %d", value))
```

**执行效果：**
- 打印错误信息和堆栈
- 中断程序执行
- 可被 `recover()` 捕获

```go
func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("捕获到 panic:", r)
        }
    }()
    
    panic("出错了！")  // 程序会中断
    fmt.Println("不会执行")
}
```

2. 返回 error - Go 的惯用方式
```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("除数不能为零")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("错误:", err)
        return
    }
    fmt.Println("结果:", result)
}
```

 3. `fmt.Errorf()` - 格式化错误
```go
func checkAge(age int) error {
    if age < 0 {
        return fmt.Errorf("年龄不能为负数: %d", age)
    }
    if age > 150 {
        return fmt.Errorf("年龄不合理: %d", age)
    }
    return nil
}
```

 4. `log.Fatal()` - 打印并退出
```go
import "log"

func main() {
    file, err := os.Open("config.json")
    if err != nil {
        log.Fatal("无法打开配置文件:", err)  // 打印后 os.Exit(1)
    }
    // 后续代码不会执行
}
```

 5. `os.Exit()` - 直接退出
```go
import "os"

func main() {
    os.Exit(1)  // 直接退出，返回状态码 1
    // 不会执行 defer
}
```

 对比总结

|方式|用途|可恢复|适用场景|
|---|---|---|---|
|`panic()`|严重错误|✅ 可 recover|配置错误、不可恢复的故障|
|`return error`|普通错误|-|业务逻辑错误、IO 错误|
|`log.Fatal()`|致命错误|❌|启动失败、无法继续|
|`os.Exit()`|立即退出|❌|需要特定退出码|

### init()函数
`init()` 是 Go 语言中**特殊的保留函数**，用于包的初始化。
基本特性

|特性|说明|
|---|---|
|**自动执行**|程序启动时自动调用，无需手动调用|
|**执行时机**|在 `main()` 之前执行|
|**参数/返回值**|无参数、无返回值|
|**数量**|每个文件/包可以有多个 `init()`|
|**可见性**|必须是小写的 `init`（不能导出）|

基本语法

```go
package main

import "fmt"

// 可以有多个 init()
func init() {
    fmt.Println("init 1 执行")
}

func init() {
    fmt.Println("init 2 执行")
}

func main() {
    fmt.Println("main 执行")
}
```

**输出：**

```
init 1 执行
init 2 执行
main 执行
```

程序执行顺序

```
程序启动
    │
    ▼
┌─────────────────────────────────┐
│  1. 全局变量声明（按声明顺序）      │
└─────────────────────────────────┘
    │
    ▼
┌─────────────────────────────────┐
│  2. 导入包的 init()               │
│     （按导入顺序，每个包的 init）  │
└─────────────────────────────────┘
    │
    ▼
┌─────────────────────────────────┐
│  3. 当前包的 init()               │
│     （按声明顺序）                 │
└─────────────────────────────────┘
    │
    ▼
┌─────────────────────────────────┐
│  4. main()                       │
└─────────────────────────────────┘
```

多文件、多包的 init 执行顺序

```go
// main.go
package main

import (
    "fmt"
    "mypackage"  // 导入自定义包
)

var globalVar = initGlobalVar()

func initGlobalVar() int {
    fmt.Println("main: 全局变量初始化")
    return 100
}

func init() {
    fmt.Println("main: init 1")
}

func init() {
    fmt.Println("main: init 2")
}

func main() {
    fmt.Println("main: main函数")
    mypackage.DoSomething()
}
```

```go
// mypackage/mypackage.go
package mypackage

import "fmt"

var pkgVar = initPkgVar()

func initPkgVar() int {
    fmt.Println("mypackage: 包变量初始化")
    return 200
}

func init() {
    fmt.Println("mypackage: init 1")
}

func init() {
    fmt.Println("mypackage: init 2")
}

func DoSomething() {
    fmt.Println("mypackage: DoSomething")
}
```

**执行顺序：**

```
1. main: 全局变量初始化
2. mypackage: 包变量初始化
3. mypackage: init 1
4. mypackage: init 2
5. main: init 1
6. main: init 2
7. main: main函数
8. mypackage: DoSomething
```

 常见用途

1. 注册驱动/插件

```go
import (
    "database/sql"
    _ "github.com/lib/pq"  // 只执行 init，不直接使用
)

// 在 pq 包的 init 中：
func init() {
    sql.Register("postgres", &Driver{})
}
```

2. 初始化配置

```go
var config Config

func init() {
    // 读取环境变量或配置文件
    config.DBHost = getEnv("DB_HOST", "localhost")
    config.DBPort = getEnv("DB_PORT", "5432")
    config.Debug = getEnv("DEBUG", "false") == "true"
}
```

 3. 验证前置条件

```go
func init() {
    if os.Getenv("API_KEY") == "" {
        log.Fatal("API_KEY 环境变量必须设置")
    }
}
```

4. 初始化单例/缓存

```go
var cache *Cache

func init() {
    cache = NewCache(1000)  // 初始化容量为1000的缓存
}
```

5. 按平台初始化

```go
var prompt string

func init() {
    if runtime.GOOS == "windows" {
        prompt = "按 Ctrl+Z 退出"
    } else {
        prompt = "按 Ctrl+D 退出"
    }
}
```

 init vs 全局变量初始化

```go
// 方式1：全局变量直接初始化
var count = 100

// 方式2：使用 init
var count int

func init() {
    count = 100
}
```

**何时使用 init？**

- 初始化逻辑复杂时
- 需要错误处理时（可以在 init 中 panic）
- 需要调用其他函数时

注意事项

|注意点|说明|
|---|---|
|**不能手动调用**|`init()` 只能由 Go 运行时调用|
|**不能有参数/返回值**|必须是 `func init()`|
|**执行顺序不确定**|同一文件内按声明顺序，不同文件间不保证|
|**panic 会终止程序**|init 中 panic 会导致程序启动失败|
|**每个文件可有多个**|但不推荐太多，影响可读性|
 实用示例

```go
package main

import (
    "fmt"
    "os"
    "runtime"
)

var (
    appName    string
    env        string
    maxWorkers int
)

// 初始化应用配置
func init() {
    appName = os.Getenv("APP_NAME")
    if appName == "" {
        appName = "MyApp"
    }
}

// 根据环境设置工作线程数
func init() {
    env = os.Getenv("ENV")
    if env == "production" {
        maxWorkers = runtime.NumCPU() * 2
    } else {
        maxWorkers = 2
    }
}

func main() {
    fmt.Printf("App: %s, Env: %s, Workers: %d\n", 
        appName, env, maxWorkers)
}
```

### select 语句详解

是的，`select` 是 **Go 语言特有**的控制结构，专门用于处理多个 channel 的发送/接收操作。

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

代码分析
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
**这里 `default` 的作用：让 select 变成非阻塞的。**

select 使用模式
模式1：多路复用（监听多个通道）
```go
ch1 := make(chan string)
ch2 := make(chan string)

go func() {
    time.Sleep(100 * time.Millisecond)
    ch1 <- "来自 channel 1"
}()

go func() {
    time.Sleep(200 * time.Millisecond)
    ch2 <- "来自 channel 2"
}()

select {
case msg := <-ch1:
    fmt.Println("收到:", msg)
case msg := <-ch2:
    fmt.Println("收到:", msg)
}
// 输出: 收到: 来自 channel 1 (先到的先执行)
```

模式2：超时控制
```go
result := make(chan string)

go func() {
    // 模拟耗时操作
    time.Sleep(3 * time.Second)
    result <- "完成"
}()

select {
case res := <-result:
    fmt.Println("操作结果:", res)
case <-time.After(2 * time.Second):
    fmt.Println("超时了！")
}
// 输出: 超时了！
```

模式3：非阻塞接收（带 default）
```go
ch := make(chan string)

// 非阻塞尝试接收
select {
case msg := <-ch:
    fmt.Println("收到:", msg)
default:
    fmt.Println("通道为空，没有收到数据")
}
```

模式4：无限循环监听
```go
ch1 := make(chan string)
ch2 := make(chan string)
quit := make(chan bool)

go func() {
    for {
        select {
        case msg := <-ch1:
            fmt.Println("ch1:", msg)
        case msg := <-ch2:
            fmt.Println("ch2:", msg)
        case <-quit:
            fmt.Println("退出")
            return
        }
    }
}()

ch1 <- "消息1"
quit <- true
```


## Go `defer` 详解
 一、基本概念

`defer` 会将函数调用**推迟到当前函数返回前**执行，无论函数是正常返回还是 panic。

```go
func main() {
    fmt.Println("start")
    defer fmt.Println("deferred")
    fmt.Println("end")
}
// 输出：
// start
// end
// deferred
```
 
 二、核心特性
 1. 后进先出（LIFO）—— 多个 defer 倒序执行

```go
func main() {
    defer fmt.Println("1")
    defer fmt.Println("2")
    defer fmt.Println("3")
}
// 输出：
// 3
// 2
// 1
```

> 可以理解为 defer 把函数调用压入一个**栈**，函数返回时从栈顶依次弹出执行。
 
 2. 参数在注册时求值，不是执行时
```go
func main() {
    i := 0
    defer fmt.Println(i) // i 的值在这里就被捕获了，是 0
    i = 100
    fmt.Println(i)
}
// 输出：
// 100
// 0   ← 不是 100！
```

```go
// 对比：闭包捕获的是引用，执行时才读值
func main() {
    i := 0
    defer func() {
        fmt.Println(i) // 执行时才读 i，是 100
    }()
    i = 100
    fmt.Println(i)
}
// 输出：
// 100
// 100
```

3. 可以修改命名返回值

```go
// 普通返回值：defer 修改无效
func noEffect() int {
    result := 0
    defer func() {
        result = 100 // 修改的是局部变量，不影响返回值
    }()
    return result // 返回 0
}

// 命名返回值：defer 可以修改！
func withEffect() (result int) {
    defer func() {
        result = 100 // 直接修改返回值变量
    }()
    return 0 // 实际返回 100
}

func main() {
    fmt.Println(noEffect())  // 0
    fmt.Println(withEffect()) // 100
}
```

**原理**：`return` 语句并不是原子操作，实际拆成了：
 1. 给返回值赋值
 2. 执行 defer
 3. 函数返回
命名返回值就是那个"返回值变量"，defer 可以在第2步修改它。


## 问：这里为啥输出结果是3 3 3
```go
for i := 0; i < 3; i++ {
    i := i // 重新声明一个局部变量
    go func() {
        fmt.Println(i)
    }()
}
```

1. **循环变量的单一性**  
    `for i := 0; i < 3; i++` 中的 `i` 在整个循环过程中只分配一次内存，每次迭代结束会执行 `i++` 改变其值，但变量本身始终是同一个。
2. **goroutine 调度时机**  
    goroutine 的启动和实际执行是并发的，当循环快速执行完毕后，`i` 的值已经变为 `3`（因为最后一次 `i++` 后 `i=3`，不满足 `i<3` 而退出循环）。此时，各个 goroutine 才开始执行 `fmt.Println(i)`，它们读取的都是同一个变量 `i` 的当前值，即 `3`。
3. **闭包捕获的是变量引用**  
    匿名函数（闭包）直接引用外部变量 `i`，而不是在创建时复制其值。所以当函数执行时，它看到的是变量最新的值。

```go
package main
import (
	"fmt"
	"time"
)
func main() {
	for i := 0; i < 3; i++ {
		go func() {
			fmt.Print(i, " ")
		}()
	}
	time.Sleep(time.Second) // 等待 goroutine 执行
}
```

输出几乎总是（可能顺序不同）：
```go
3 3 3
```
如果循环中增加一些延时（比如 `time.Sleep`）让 goroutine 在循环结束前执行，则可能看到 `0 1 2`的混合，但大多数情况下，循环结束很快，goroutine 看到的就是最终值 `3`。
如何避免？
使用局部变量复制值：
```go
for i := 0; i < 3; i++ {
    i := i // 创建局部副本
    go func() {
        fmt.Print(i, " ")
    }()
}
```
或通过参数传递值：
```go
for i := 0; i < 3; i++ {
    go func(i int) {
        fmt.Print(i, " ")
    }(i)
}
```
这样每个 goroutine 捕获的是独立的副本，值就固定为迭代时的值（0、1、2）。
总结：“都是3”的原因是：循环结束后，所有 goroutine 读取的是同一个循环变量，而该变量最后的值为 3。这正是 Go 循环变量捕获的经典陷阱。