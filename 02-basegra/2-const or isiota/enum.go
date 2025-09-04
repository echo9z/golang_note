package main

import "fmt"

// 1.枚举基本使用
type Season uint8

const (
	Spring Season = iota
	Summer
	Autumn
	Winter
)

// 枚举实际上就是数字，Go 也不支持直接将其转换为字符串，但我们可以通过给自定义类型添加方法来返回其字符串表现形式，实现Stringer接口即可。
// 实现String()方法让枚举值在打印时更可读
func (s Season) String() string {
	switch s {
	case Spring:
		return "Spring"
	case Summer:
		return "Summer"
	case Autumn:
		return "Autumn"
	case Winter:
		return "Winter"
	default:
		return fmt.Sprintf("Unknown Season: %d", s) // %d十进制输出
	}
}

// IsValid()等方法可以在运行时验证枚举值的有效性
func (s Season) IsValid() bool {
	switch s {
	case Spring, Summer, Autumn, Winter:
		return true
	default:
		return false
	}
}

// 2。字符串枚举
type LogLevel string

const (
	LogLevelDebug   LogLevel = "DEBUG"
	LogLevelInfo    LogLevel = "INFO"
	LogLevelWarning LogLevel = "WARNING"
	LogLevelError   LogLevel = "ERROR"
	LogLevelFatal   LogLevel = "FATAL"
)

// LogLevel 的 IsValid 方法验证日志级别的有效性
func (l LogLevel) IsValid() bool {
	switch l {
	case LogLevelDebug, LogLevelInfo, LogLevelWarning, LogLevelError, LogLevelFatal:
		return true
	default:
		return false
	}
}

/*
(LogLevel)接收者为 LogLevel 类型的 方法，它只能被 LogLevel 类型的变量调用。比如 var l LogLevel; lvl := l.ValidLevels()
*/
func (LogLevel) ValidLevels() []LogLevel {
	return []LogLevel{
		LogLevelDebug,
		LogLevelInfo,
		LogLevelWarning,
		LogLevelError,
		LogLevelFatal,
	}
}

func main() {
	// Go 语言没有为枚举单独设计，而是通过 自定类型+const+iota 实现
	var season Season = Summer
	fmt.Printf("当前季节:%s (值%d)\n", season, int(season)) // 输出: 当前季节: Summer  %s输出字符串, %d输出十进制整数
	fmt.Printf("季节是否有效: %t\n", season.IsValid())

	// 枚举字符串
	var logLevel LogLevel = LogLevelDebug
	fmt.Printf("当前日志级别: %s\n", logLevel) // 输出: 当前日志级别: DEBUG

	fmt.Printf("日志级别是否有效: %t\n", logLevel.IsValid())
	// logLevel.ValidLevels()返回一个包含所有有效日志级别的切片，进行遍历
	for idx, level := range logLevel.ValidLevels() {
		fmt.Print(idx)
		fmt.Printf("有效日志级别: %s\n", level) // 输出: 有效日志级别: DEBUG, INFO, WARNING, ERROR, FATAL
	}
}
