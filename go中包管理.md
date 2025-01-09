在版本1.11之前，Go 语言的的包依赖管理一直都被大家所诟病。Go官方也在一直在努力为开发者提供更方便易用的包管理方案，从最初的 `GOPATH` 到 `GO VENDOR`，再到最新的 `GO Modules`。
目前Go语言的包依赖管理方式是 `Go Modules`。

## GOPATH

在 Go 1.11 版本之前，Go 并没有模块（`go mod`）的概念，而是使用了一个全局的工作区机制，依赖通过 `GOPATH` 管理，且所有的 Go 项目代码都要保存在 GOPATH/src 目录下。

```shell
$ go env GOPATH
/home/echo9z/go

$ ls -l /home/echo9z/go
drwxr-xr-x 1 echo9z echo9z 50 12月25日 12:33 bin #存放项目的源代码
drwxr-xr-x 1 echo9z echo9z 38 12月25日 12:45 pkg #存放依赖包的编译产物（.a文件）
drwxr-xr-x 1 echo9z echo9z 42 12月26日 10:53 src #存放由go install命令生成的可执行文件
```

GOPATH目录下一共包含了三个子目录，分别是：
- bin：存储所编译生成的二进制文件。
- pkg：存储预编译的目标文件，以加快程序的后续编译速度。
- src：存储所有`.go`文件或源代码。在编写 Go 应用程序，程序包和库时，一般会以`$GOPATH/src/`的路径进行存放。
```
go
├── bin
├── pkg
└── src
    ├── github.com
    ├── golang.org
    ├── google.golang.org
    ├── gopkg.in
    ... ...
```

### 手动管理依赖

开发者需要手动将依赖的代码库克隆到 `GOPATH/src` 目录下。例如，如果需要使用 `github.com/example/package`，开发者需要执行以下命令：
```bash
git clone https://github.com/example/package $GOPATH/src/github.com/example/package
```
通过`go install github.com/example/package` 将package依赖版本文件构建成`.a`依赖文件存放在`~go/src/pkg`中。
在代码中通过 `import "github.com/example/package"` 来引用该依赖。

比如当前目录下结构

```sh
.
├── bin
├── pkg
└── src
    ├── myapp
    │   └── main.go
    ├── mypkg
    │   └── hello.go
    └── mytest
```

`package main` 是一个特殊的包，用于定义github.com/example/package一个独立的可执行程序。

```go
// src/myapp/main.go
package main  
import "fmt"

func main() {  
    fmt.Println("Hello World")  
}
```

 `package main` 构建的程序会生成一个独立的二进制文件，使用`go install`生成可执行文件会放在 `$GOPATH/bin` 下。

```sh
$ pwd
~/go/src
$ go install myapp

$ tree ~/go -L 3
$GOPATH
├── bin
│   └── myapp  #生成的二进制可执行文件myapp
├── pkg
└── src
    ├── myapp
    │   └── main.go
    ├── mypkg
    │   └── hello.go
    └── mytest
```

如果是自定义的package，例`package mypkg`，`mypkg` 可以通过 `import "mypkg"` 引入使用。
```go
package mypkg 
import "fmt"  
  
func Hello() {  
    fmt.Println("hello")  
}
```

使用`go install`会生成 `.a`后缀文件会添加到 `$GOPATH/pkg`。
```sh
$ pwd
~/go/src
$ go install mypkg

$ tree ~/go -L 3
$GOPATH
├── bin
│   └── myapp
├── pkg
│   └── linux_amd64
│       └── mypkg.a # 依赖包的编译产物（`.a` 文件）
└── src
    ├── myapp
    │   └── main.go
    ├── mypkg
    │   └── hello.go
    └── mytest
```

`GOOS`，表示的是目标操作系统，darwin（Mac）、linux、windows、android等
`GOARCH`，表示目标架构，常见的有 arm，amd64 等

```bash
$ go env GOOS GOARCH
linux
amd64
```

使用 `GOPATH`模式，会遇到的问题：
- 无法项目中，使用指定版本的包，因为不同版本的包的导入方法也都一样
- 其他人运行你的开发的程序时，无法保证下载的包版本是你所期望的版本，当对方使用了其他版本，有可能导致程序无法正常运行
- 在本地，一个包只能保留一个版本，意味着你在本地开发的所有项目，都得用同一个版本的包，这几乎是不可能的。

## vendor 模式
为了解决 GOPATH 方案下不同项目下无法使用多个版本库的问题，Go v1.5 开始支持 vendor ，实现同一个包在不同项目中不同版本、以及无相互侵入的开发和管理。

由于所有项目共享同一个 `GOPATH/src`，不同项目可能需要同一个依赖包的不同版本，这会导致冲突。例如：

- 项目 A 需要 `github.com/pkg/errors v0.8.0`。
- 项目 B 需要 `github.com/pkg/errors v0.9.0`。

在 `GOPATH`模式的时候，所有的项目都共享一个`GOPATH`，需要导入依赖的时候，都在`~/go/src`下找，在 GOPATH 模式下只能有一个版本的第三方库。`vendor`为兼容 `GOPATH`工作模式。
`Dep`官方开发的工具（Go 1.9 和 1.10 推荐使用），`govendor`

`vendor`解决方式，在`~/go/src`的每个项目下都创建一个 vendor 目录，每个项目所需的依赖都只会下载到自己vendor目录下，项目之间的依赖包互不影响。在编译时，v1.5 的 Go 在设置了开启 `GO15VENDOREXPERIMENT=1` （注：这个变量在 v1.6 版本默认为1，但是在 v1.7 版本时，已去掉该环境变量，默认开启 `vendor` 特性，无需你手动设置）后，会提升 vendor 目录的依赖包搜索路径的优先级（相较于 GOPATH）。

`vendor`搜索包的优先级顺序，由高到低：

- 当前包下的 vendor 目录
- 向上级目录查找，直到找到 src 下的 vendor 目录
- 在 GOROOT 目录下查找
- 在 GOPATH 下面查找依赖包

这个方案解决了一些问题，但是解决得并不完美。

- 如果多个项目用到了同一个包的同一个版本，这个包会存在于该机器上的不同目录下，不仅对磁盘空间是一种浪费，而且没法对第三方包进行集中式的管理（分散在各个项目的角落）。
- 并且如果要分享开源你的项目，需要将你项目所有的依赖包悉数上传，别人使用的时候，除了你的项目源码外，还有所有的依赖包全部下载下来，才能保证别人使用的时候，不会因为版本问题导致项目不能如你预期那样正常运行。

关于`vendor`具体使用参考：[govendor](https://shockerli.net/post/go-package-manage-tool-govendor/)文章

## Modules
### mod提供的命令
在 Go modules 中，我们能够使用如下命令进行操作：

| 命令              | 作用                   |
| --------------- | -------------------- |
| go mod init     | 生成 go.mod 文件         |
| go mod download | 下载 go.mod 文件中指明的所有依赖 |
| go mod tidy     | 整理现有的依赖              |
| go mod graph    | 查看现有的依赖结构            |
| go mod edit     | 编辑 go.mod 文件         |
| go mod vendor   | 导出项目所有的依赖到vendor目录   |
| go mod verify   | 校验一个模块是否被篡改过         |
| go mod why      | 查看为什么需要依赖某模块         |

### GO111MODULE
go modules 在 v1.11 版本正式推出，在最新发布的 v1.13 版本中，默认将`GO111MODULE`

从 v1.11 开始，`go env` 多了个环境变量： `GO111MODULE` ，这里的 111，其实就是 v1.11 的象征标志， go 很喜欢这样的命名方式，比如环境使用`vendor`出现的时候，也多了个 `GO15VENDOREXPERIMENT`环境变量，其中 15表示的vendor 是在 v1.5 时才诞生的。

`GO111MODULE` 是一个开关，通过它可以开启或关闭 go mod 模式。
```bash
$ go env -w GO111MODULE="on"
```

1. `GO111MODULE=off`禁用模块支持，编译时会从`GOPATH`和`vendor`文件夹中查找包。
2. `GO111MODULE=on`启用模块支持，编译时会忽略`GOPATH`和`vendor`文件夹，根据 `go.mod`下载依赖。
3. `GO111MODULE=auto`，当项目在`$GOPATH/src`外且项目根目录有`go.mod`文件时，自动开启模块支持。

### GOPROXY
这个环境变量主要是用于设置 Go 模块代理（Go module proxy），其作用是用于使 Go 在后续拉取模块版本时能够脱离传统的 VCS 方式，直接通过镜像站点来快速拉取。

GOPROXY 的默认值是：`https://proxy.golang.org,direct`，`proxy.golang.org` 在国内是无法访问的，因此在下载Go模块包直接卡住，所以你必须在开启 Go modules 的时，同时设置国内的 Go 模块代理：
```bash
$ go env -w GOPROXY=https://goproxy.cn,direct
```
`“direct” `是一个特殊指示符，用于指示 Go 回源到模块版本的源地址去抓取（比如 GitHub 等），场景如下：当值列表中上一个 Go 模块代理返回 404 或 410 错误时，Go 自动尝试列表中的下一个，遇见 `“direct” `时回源，也就是回到源地址去抓取，而遇见 EOF 时终止并抛出类似 “invalid version: unknown revision...” 的错误。

### GOSUMDB
它的值是一个 Go checksum database，用于在拉取模块版本时（无论是从源站拉取还是通过 Go module proxy 拉取）保证拉取到的模块版本数据未经过篡改，若发现不一致，也就是可能存在篡改，将会立即中止。



