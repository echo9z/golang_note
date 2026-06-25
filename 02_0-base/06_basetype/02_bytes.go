package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	// bytes包
	// 1.使用bytes.Buffer高效字符串拼接
  var buf bytes.Buffer
	buf.WriteString("Hello,") // 添加字符串
	buf.WriteByte('W') // 添加单个字节
	buf.Write([]byte("orld!")) // 传入字符串

	// 通过string() 转化为字符串
	res := buf.String()
	fmt.Println(res) // Hello,World!  比 + 或 strings.Builder 更灵活
	fmt.Println(buf.Len())  // 12 字节数（不含已读走的部分）

	// 2.bytes.Reader从字节切片中读取
	data := []byte("ok golang")
	// 从data字节切片读取，返回一个reader读取对象
	reader := bytes.NewReader(data)
	fmt.Printf("str:%v \n", *reader)

	// 读取前 6个字节
	buff := make([]byte, 6) 
	// 将从byte切片data 读去到buff缓存区中
	n,err := reader.Read(buff) 
	if err != nil && err != io.EOF{ 
		panic(err)
	}
	fmt.Printf("读取%d字节 内容:%s\n", n, buff) // 读取6字节 内容:ok gol

	// 通过 Seek 进行跳跃试查找。 Seek 实现了 io.Seeker 接口。
	reader.Seek(3, io.SeekStart) // 定位读取位置（偏移量）
	// 3 表示从目标位置向后移动 3个字节
	// io.SeekStart 表示以文件或流的开头作为基准点
	rest, _ := io.ReadAll(reader) // 读取当前指针所在位置（定位的第3字节）一直到末尾的所有数据。
	fmt.Println(string(rest)) // golang

	// 常用到的工具函数
	// 3.判断两个字节切片是否equal相等
	fmt.Println(bytes.Equal([]byte("abc"), []byte("123"))) // false
	a := []byte("hello world")
	b := []byte("hello")
	// 判断a是否包含b Contains
	fmt.Println(bytes.Contains(a, b)) // true
	// 字节切片 s 是否以 prefix 开头
	fmt.Println(bytes.HasPrefix(a, b))  // 字节切片 a 是否以 hello 开头 true
	fmt.Println(bytes.HasSuffix(b, []byte("Hello"))) // 字节切片 a 是否以 hello 结尾 false
	fmt.Println(bytes.Count(a, []byte("world"))) // 统计world在 字节切片a中出现次数 1次

	// 4.大小写转换（返回新切片[]byte）
	var upper []byte  = bytes.ToUpper([]byte("abcd"))
	var lower []byte  = bytes.ToLower([]byte("ABCD"))
	fmt.Printf("转大写%s, 转小写%s\n", string(upper), string(lower)) // 转大写ABCD, 转小写abcd

	// 5.分割与拼接
	data2 := []byte("a,b,c,d,e")
	parts := bytes.Split(data2, []byte(",")) // 将切片 data2 按 , 分隔为所有子切片{{a},{b},{c},{d},{e}}
	fmt.Printf("%v\n",parts) // 返回二维[][]byte切片 [[97] [98] [99] [100] [101]]
	for _, byteV := range parts {
		fmt.Println(string(byteV)) // a b c d e
	}

	// 通过bytes.join 将 s 的元素连接起来创建一个新的字节切片。
	joined := bytes.Join([][]byte{{'a'},{'b'},{'c'}, {'d'}}, []byte("-"))
	fmt.Println(string(joined)) // a-b-c-d

	// 6. 替换与修剪
	srcByte := []byte("a-b-c-d-e")
	// n=-1表示全部替换
	resR := bytes.ReplaceAll(srcByte, []byte("-"), []byte("_"))
	fmt.Println(string(resR)) // a_b_c_d_e
	resR = bytes.Replace(srcByte, []byte("a"), []byte("1"), 1) // 将s切片中 old字符，替换为new字符，替换前第n个匹配到字符
	fmt.Println(string(resR)) // 1_b_c_d_e  至替换一次

	// 去除首尾字符
	fmt.Println(string(bytes.TrimSpace([]byte("  h i  "))))   // h i 去除首尾空格，
	fmt.Println(string(bytes.Trim([]byte("##hi##"), "#")))    // hi
	fmt.Println(string(bytes.TrimPrefix([]byte("Mr.Tom"), []byte("Mr.")))) // Tom 去除匹配前缀字符串



}
