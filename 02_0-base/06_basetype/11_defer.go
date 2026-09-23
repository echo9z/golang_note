package main

import (
	"fmt"
	"io"
	"os"
	"sync"
)

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20

	wd, _ := os.Getwd()
	fmt.Println("当前工作目录:", wd)

	written, err := CopyFile("02_0-base/06_basetype/portint2.c", "02_0-base/06_basetype/portint.c", &sync.Mutex{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("写入大小", written)

}

// defer先固定各个被调用的
// 输出结果 A 1 2 3
// B 10 2 12
// BB 10 12 22
// AA 1 3 4

// 打开两个文件，并将一个文件的内容复制到另一个文件中。
// 不使用defer，每次操作，都得处理资源的释放：锁的释放、文件句柄的关闭。
func CopyFile(dstName, srcName string, mu *sync.Mutex) (written int64, err error) {
	mu.Lock() // 互斥锁 mu 锁定，确保同一时间只有一个线程可以执行文件写入操作，避免数据竞争。
	src, err := os.Open(srcName)
	if err != nil {
		mu.Unlock()
		// src.Close()
		return 0, err
	}

	// 创建目标文件
	// 该操作会创建或截断指定的文件。如果文件已经存在，则只会将其截断；如果文件不存在，则会以模式 0o666 创建该文件
	dst, err := os.Create(dstName)
	if err != nil {
		mu.Unlock()
		src.Close() // 关闭占着句柄
		return 0, err
	}

	// 从 src 复制到 dst，直到在 src 处达到文件末尾或发生错误为止。该函数会返回复制的字节数，以及复制发生错误
	written, err = io.Copy(dst, src)
	if err != nil {
		mu.Unlock()
		dst.Close()
		src.Close()
		return 0, err
	}

	mu.Unlock()
	dst.Close()
	src.Close()
	return written, err
}

// 使用defer申请资源之后，马上defer+释放资源。当函数CopyFileDefer函数运行结束时会自动释放锁和关闭文件句柄。
func CopyFileDefer(dstName, srcName string, mu *sync.Mutex) (written int64, err error) {
	mu.Lock()
	defer mu.Unlock()

	src, err := os.Open(srcName)
	if err != nil {
		return 0, err
	}
  defer	src.Close() // 关闭占着句柄

	dst, err := os.Create(dstName)
	if err != nil {
		return 0, err
	}
	defer dst.Close()

	written, err = io.Copy(dst, src)
	if err != nil {
		return 0, err
	}

	return written, err
}
