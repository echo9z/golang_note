package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

func main() {

	// 十一、从字符串中读取内容
	// strings.NewReader 创建一个从字符串读取数据的 Reader 对象，
	// 它实现了 io.Reader、io.ReaderAt、io.ByteReader、io.RuneReader、io.Seeker、io.WriterTo 等接口。

	// 1.读取字节 (Read)
	strs := "hello, world!"
	var reader *strings.Reader = strings.NewReader(strs) // 返回字符串 Reader对象，读取字符串信息
	fmt.Printf("Reader 长度：%d\n", reader.Len())           // Reader 长度：13

	// 创建缓冲区
	// make([]byte, 5) 是 Go 语言中用于创建一个字节切片（slice of bytes）的语法，
	// 它会分配一个底层数组，并返回一个包含 5 个元素（都初始化为零值）的切片，
	// 其长度（len）和容量（cap）都是 5。这常用于处理二进制数据、网络请求或文件操作，是 Go 语言中创建切片的标准方式
	buf := make([]byte, 5)

	// 通过读取前5个字节内容，
	n, err := reader.Read(buf) // 将reader字符串去读到buf缓冲区中，n表示读取字节的个数；即读取了5字节即hello
	if err != nil {            // !=nil err存在错误，不为空值
		fmt.Println("读取错误", err)
		return
	}
	fmt.Printf("读取了 %d个字节 %s\n", n, buf)        // 读取了 5个字节 hello
	fmt.Printf("Reader 剩余长度%d\n", reader.Len()) // Reader 剩余长度8

	// 2.读取到缓冲区 (ReadAt)
	str2 := "hello, world!"
	reader2 := strings.NewReader(str2)

	// 创建缓冲区
	buf2 := make([]byte, 5)
	// 从指定位置开始读取
	n2, err := reader2.ReadAt(buf2, 7) // 从索引7开始读取
	if err != nil {
		fmt.Println("读取错误", err)
		return
	}
	fmt.Printf("读取到字节共%d 读取内容%s\n", n2, buf2) // 读取到字节共5 读取内容world

	// 3.读取单个字节 (ReadByte)
	str3 := "hello, world!"
	reader3 := strings.NewReader(str3)

	// 循环变量读取单个字节
	for i := 0; i < 5; i++ { // 循环5次（因为"Hello"有5个字符）
		b, err := reader3.ReadByte() // 读取一个字节
		/**
		EOF = End Of File（文件结束标志）
		io.EOF 是一个预定义的错误变量，表示数据流已读取完毕
		var EOF = errors.New("EOF")
		当 Read() 或 ReadByte() 等方法尝试读取超出可用数据时，表示"没有更多数据可读了"
		因为只循环了5次，而"Hello"正好有5个字节。如果尝试读取第6次，就会得到 EOF 错误
		*/
		if err == io.EOF {           // 检查是否到达结尾
			break
		}
		fmt.Printf("字节 %d: %c\n", i, b)
	}

	// 4. 读取单个字符 (ReadRune)
	str4 := "Hello 世界！"
	reader4 := strings.NewReader(str4)
	for {
		ch, size, err := reader4.ReadRune() // 返回
		if err == io.EOF {           // 检查是否到达结尾
			break
		}
		fmt.Printf("字符：%c, 字节大小：%d\n", ch, size)
	}
	// 输出:
	// 字符: H, 字节大小: 1
	// 字符: e, 字节大小: 1
	// 字符: l, 字节大小: 1
	// 字符: l, 字节大小: 1
	// 字符: o, 字节大小: 1
	// 字符: 世, 字节大小: 3
	// 字符: 界, 字节大小: 3

	// 5.定位读取
	str5 := "Hello, World"
	reader5 := strings.NewReader(str5)

	// 跳到索引位置7开始
	reader5.Seek(7, 0) // 从开头(0)偏移7个字节,移动后指针指向字符 '7' 之前。从第八个索引位置开始
	// Reader.Seek(offset, whence)
	// offset：偏移位置
	// whence：0:表示从头开始，比如Seek(7, 0)，字符串为"Hello, World"，从字符串起始位开始，向右偏移7个字节，即第8个位置”W“开始
	// 1:表示当前位置，比如字符串为"Hello, World"，通过b1, _ := reader.ReadByte(),表示读取一个字节b1=0即H
	// 		再通过b2, _ := reader.ReadByte()，再次读取一个字节b2=1即e，此时已经读取了前2个字节
	// 		当reader.Seek(2, 1) 1表示当前位置，当前位置为2，即从位置2移动到4，即第5个位置
	// 2:表示从尾部开始，比如Seek(-5, 2)，字符串为"Hello, World"，从字符串起始位开始，向左偏移5个字节，索引6即第7个位置” “开始

	buf5 := make([]byte, 5) // 创建5个字节大小的缓冲区切片中
	reader5.Read(buf5) // 将reader字符串对象内容，读取到缓冲区中，大小为缓冲区5个字节
	fmt.Printf("从位置7开始读取：%s\n", buf5) // 从位置7开始读取：World

	// 6.写入到 Writer (WriteTo)
	str := "hello world"
	reader6 := strings.NewReader(str)
	
	// 创建一个buffer
	var buff strings.Builder
	// 将reader字符串内容写入到buffer中
	n6, err := reader6.WriteTo(&buff)
	if err != nil {
		fmt.Println("写入错误:", err)
    return
	}
	fmt.Printf("写入了 %d个字节\n", n6) // 写入了 11个字节
	fmt.Printf("buffer内容 %s\n", buff.String()) // buffer内容 hello world

	ReadExe()
}

// 实际应用场景
func ReadExe() {
	// 模拟文件读取
	scan := func ()  {
		// 模拟从字符串"文件"中读取数据，下面时cvs格式数据
		csvData := "name,age,city\nAlice,25,New York\nBob,30,London"
		// 创建reader对象
		reader := strings.NewReader(csvData)

		// 使用bufio追行读取
		// bufio 是 "buffered I/O"（缓冲 I/O）的缩写
		// bufio.Reader	缓冲读取，支持 ReadString()、ReadLine()
		// bufio.Writer	缓冲写入，减少频繁写操作
		// bufio.Scanner 类型提供了按行、按词、按字节扫描数据的能力
		scanner := bufio.NewScanner(reader) // bufio不需要手动处理换行符
		for scanner.Scan() {
			// 自动获取每一行
			fmt.Println(scanner.Text())
		}
	}

	scan()

	// 解析数据
	exeData := func ()  {
		jsonData := `{"name": "Alice", "age": 25}`
    reader := strings.NewReader(jsonData)

		// 解析json
		var result map[string]interface{}
		decoder := json.NewDecoder(reader)
		err := decoder.Decode(&result)
		if err != nil {
        fmt.Println("解析错误:", err)
        return
    }
    fmt.Printf("解析结果: %+v\n", result)
	}
	exeData()

	// 限制读取大小
	str := "Hello, World! This is a long string."
  reader := strings.NewReader(str)

	// 只读取前5个字符
	buf := make([]byte, 5) // 创建5个字节大小的切片
	n, err := reader.Read(buf) // 将reader字符串对象，读取5个字节到buf中缓冲区中
	if err != nil {
		fmt.Println("读取错误", err)
		return
	}
	// buf[:n]
	fmt.Printf("读取字节：%d，读取内容：%s", n, buf[:n]) // 读取字节：5，读取内容：Hello

	// 回退读取
	rockBack := func ()  {
		str := "Hello"
    reader := strings.NewReader(str)
    
    // 读取一个字符
    ch, _, _ := reader.ReadRune()
    fmt.Printf("读取: %c\n", ch)  // H
    
    // 回退一个字符
    reader.UnreadRune()
    
    // 再次读取
    ch, _, _ = reader.ReadRune()
    fmt.Printf("再次读取: %c\n", ch)  // H
	}
	rockBack()

}
