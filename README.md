# golang study notes

## 准备开始go

### 官方文档

地址：[Documentation - The Go Programming Languageopen in new window](https://go.dev/doc/)

文档里有着对于学习Go语言所需要准备的一切东西，包括安装，快速开始，代码示例，风格建议，以及许多在线教程，大多数都是全英文的，少数支持中文，不过并没有什么特别晦涩难懂的词汇，大致意思都比较容易看懂。

### Go之旅

地址：[Go 语言之旅 (go-zh.org)open in new window](https://tour.go-zh.org/welcome/1)

这是由官方编写的一个非常简洁明了的教程，全中文支持，通过互动式的代码教学来帮助你快速了解Go语言的语法与特性，适合想要快速了解Go语言的人，如果将该教程浏览过一遍后，那么本站的基础教程理解起来会轻松很多。

### Effective Go

地址：[Effective Go - The Go Programming Languageopen in new window](https://go.dev/doc/effective_go)

这是由官方编写的一个比较全面的教程，时间最早可以追溯到2009年，内容比较详细，小到变量命名，大到一些设计思想。不过官方也标注了该文档已经很久没有进行大幅度更新，一些内容可能已经过时了，但是大部分教程都仍然适用。

### 参考手册

地址：[The Go Programming Language Specificationopen in new window](https://go.dev/ref/spec)

参考手册的重要性不言而喻，参考手册的内容永远会随着版本的变化而变化，时刻保持最新，其内容有：词法结构，概念定义，语句定义等等，这是一些关于Go语言中最基础的定义，适合有需要的时候查询一些概念，同时里面也有着不少的代码示例。

### 在线代码

地址：[Go Playground - The Go Programming Languageopen in new window](https://go.dev/play/)

由官方搭建的可在线编译并运行Go程序的网站，对于一些代码量不是特别大的Go程序，可以在官方的在线网站直接进行编写，能省去不少时间。

### 更新日志

地址：[Release History - The Go Programming Languageopen in new window](https://go.dev/doc/devel/release)

根据以往的惯例，官方大概每半年发布一个二级版本，每一次更新的变动都可以在更新日志中查看，例如在1.18版本中的最大变动就是增加了泛型，而1.19的更新就相对而言要温和很多，了解一下每一个版本的更新内容也会有所帮助。

### Go安装

推荐使用官方的安装包直接安装，下载地址：https://golang.google.cn/dl/

贴士：本笔记都是基于go1.21

**Win安装Go**：
打开Win安装包下一步下一步即可，默认安装在目录：c:\Go

**Mac安装Go**：
打开Mac安装包下一步下一步即可，需要预装Xcode。安装完毕后需配置环境变量即可使用，但是如果要使用一些`go mod`功能推荐如下配置：

```
vim ~/.bash_profile

export GOROOT=/usr/local/go                 # golang本身的安装位置
export GOPATH=~/go/                         # golang包的本地安装位置
export GOPROXY=https://goproxy.io           # golang包的下载代理
export GO111MODULE=on                       # 开启go mod模式
export PATH=$PATH:$GOROOT/bin               # go本身二进制文件的环境变量
export PATH=$PATH:$GOPATH/bin               # go第三方二进制文件的环境便令

# 重启环境
source ~/.bash_profile
```

测试安装：

```
# 查看go版本
go version

# 查看go环境配置
go env 
```

### 开发工具推荐

笔者推荐的go开发工具：

- goland
- vscode

vscode的相关go插件会出现无法下载情况，解决办法：

```
# 如果开启了go mod，则
    go get -u -v github.com/ramya-rao-a/go-outline
    go get -u -v github.com/acroca/go-symbols
    go get -u -v golang.org/x/tools/cmd/guru
    go get -u -v golang.org/x/tools/cmd/gorename
    go get -u -v github.com/rogpeppe/godef
    go get -u -v github.com/sqs/goreturns
    go get -u -v github.com/cweill/gotests/gotests
    go get -u -v golang.org/x/lint/golint

# 如果未开启go mod，则需要进入cd $GOPATH/src ，使用 git clone 下载上述文件        

# 安装
cd $GOPATH
    go install github.com/ramya-rao-a/go-outline
    go install github.com/acroca/go-symbols
    go install golang.org/x/tools/cmd/guru
    go install golang.org/x/tools/cmd/gorename
    go install github.com/rogpeppe/godef
    go install github.com/sqs/goreturns
    go install github.com/cweill/gotests/gotests
    go install golang.org/x/lint/golint
```

## 基础语法

### HelloWorld

通过一个简单的Hello World示例来进行讲解。

go的项目依赖管理一直饱受诟病，在go1.11后正式引入了`go modules`功能，在go1.13版本中将会默认启用。

1.`go mod init golang/notes`初始化并写入一个新的go.mod至当前目录中，实际上是创建一个以当前目录为根的新模块。可以理解nodejs中package.json

```mod
# go.mod
module golang/notes // 以后自定包，都需要golang/notes/xxx

go 1.21.1
```

2.新建文件`hello.go`，代码如下：

```go
package main                        //每个程序都有且仅有一个main包

import "fmt"    
// 程序的入口文件
func main() {                       //主函数main只有一个
    fmt.Println("Hello World!")     //函数调用：包名.函数名
}
```

`package`关键字代表的是当前go文件属于哪一个包，启动文件通常是`main`包，启动函数是`main`函数，在自定义包和函数时命名应当尽量避免与之重复。

`import`是导入关键字，后面跟着的是被导入的包名。

`func`是函数声明关键字，用于声明一个函数。

`fmt.Println("Hello World!")`是一个语句，调用了`fmt`包下的`Println`函数进行控制台输出。

运行文件：

```
# 执行方式一：先编译，再运行
go build hello.go        # 编译。在同级目录下生成文件`hello`，添加参数`-o 名称` 则可指定生成的文件名 
./hello                  # 运行。贴士：win下生成的是.exe文件，直接双击执行即可 编译二进制文件

# 执行方式二：直接运行
go run hello.go         
```

两种执行流程的区别：

- 先编译方式：可执行文件可以在任意没有go环境的机器上运行，（因为go依赖被打包进了可执行文件）
- 直接执行方式：源码执行时，依赖于机器上的go环境，没有go环境无法直接运行

### package包

在Go中，程序是通过将包链接在一起来构建的，也可以理解为最基本的调用单位是包，而不是go文件。包其实就是一个文件夹，包内共享所有源文件的变量，常量，函数以及其他类型。包的命名风格建议都是小写字母，并且要尽量简短。

#### 包导入

例如创建一个`utils`包，包下有如下函数

utils/say.go

```go
package utils // 注意：文件夹 要与utils一直

import "fmt"

func SayHello() {
	fmt.Println("Hello")
}
```

utils/say.go

```go
package utils

import "fmt"

func GoRun() {
	fmt.Println("go run")
}
```

在`main`函数中调用

```go
package main

import "golang/notes/utils" // "golang/notes" 这里的路径要对应 go.mod 中module名称

func main() {
	utils.GoRun()
}
```

还可以给包起别名

```go
package main

import u "golang/notes/utils"

func main() {
   u.SayHello()
}
```

批量导入时，可以使用括号`()`来表示

```go
package main

import (
	"fmt"
	"math"
	u "golang/notes/02-basegra/utils"
)
func main() {
	fmt.Println("多包导入", math.Pi)
	u.GoRun()
}
```

或者说只导入不调用，通常这么做是为了调用该包下的`init`函数。

```go
package main

import (
   "fmt"
    _ "math" // 下划线表示匿名导入
)

func main() {
   fmt.Println(1)
}
```

> 注意：Go中完全禁止循环导入，不管是直接的还是间接的。例如包A导入了包B，包B也导入了包A，这是直接循环导入，包A导入了包C，包C导入了包B，包B又导入了包A，这就是间接的循环导入，存在循环导入的话将会无法通过编译。

#### 包导出

在Go中，导出和访问控制是通过命名来进行实现的，如果想要对外暴露一个函数或者一个变量，只需要将其名称首字母大写即可，例如`example`包下的`SayHello`函数。

```go
package utils

import "fmt"

// 首字母大写，可以被包外访问
func SayHello() {
   fmt.Println("Hello")
}
```

如果想要不对外暴露的话，只需将名称首字母改为小写即可，例如下方代码

```go
package utils

import "fmt"

// 首字母小写，外界无法访问
func sayHello() {
   fmt.Println("Hello")
}
```

对外暴露的函数和变量可以被包外的调用者导入和访问，如果是不对外暴露的话，那么仅包内的调用者可以访问，外部将无法导入和访问，**该规则适用于整个Go语言**，例如后续会学到的结构体及其字段，方法，自定义类型，接口等等。

#### 私有

go中约定，一个包内名为`internal` 包为私有包，其它的包将无法访问私有包中的任何东西。下面看一个例子。

```bash
 tree ./02-basegra 
./02-basegra
|-- main.go
|-- test
|   |-- internal
|   |   `-- ser
|   |       `-- ser.go
|   |-- ser
|   |   `-- ser.go
|   `-- test.go
`-- utils
    |-- run.go
    `-- say.go
```

文件结构中可知，`utils`包无法访问`ser`包中的类型。