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
`GOSUMDB` 环境变量的作用是管理和配置 **模块校验和数据库（Checksum Database）**。校验和数据库的，目的是用来验证 Go 模块的完整性和来源，确保模块没有被篡改或破坏，从而提供一个安全的依赖管理机制。

当你拉取某个 Go 模块（论是从源站拉取还是通过 Go module proxy 拉取，如第三方库）时，`GOSUMDB` 会从校验和数据库中获取该模块的校验和，并与本地 `go.sum` 文件中的记录进行对比。如果校验和不一致，Go 会抛出错误，提示模块可能被篡改或下载异常，会立即中止。

GOSUMDB 的默认值为：`sum.golang.org`，是官方提供可信的校验和数据库。由于国内网络无法访问的，但是 GOSUMDB 可以被 Go 模块代理所代理（详见：Proxying a Checksum Database）。

另外若对 GOSUMDB 的值有自定义需求，其支持如下格式：
- 格式 1：`<SUMDB_NAME>+<PUBLIC_KEY>`。
- 格式 2：`<SUMDB_NAME>+<PUBLIC_KEY> <SUMDB_URL>`。
也可以将其设置为“off”，也就是禁止 Go 在后续操作中校验模块版本。

**关闭或自定义 GOSUMDB**
1. **禁用校验和数据库**  
	如果不需要校验模块（例如使用私有模块），可以将 `GOSUMDB` 设置为 `off`：
```bash
GOSUMDB=off
```
禁用后，Go 不会对模块的校验和进行验证。
2. **使用自定义校验和数据库**  
	需要使用公司内部的校验和数据库，可以将 `GOSUMDB` 设置为自定义的地址：
```bash
    GOSUMDB=mychecksumdb.example.com
```

### GONOPROXY/GONOSUMDB/GOPRIVATE
这三个环境变量都是用在当前项目依赖了私有模块，例如像是你公司的私有 git 仓库，又或是 github 中的私有库，都是属于私有模块，都是要进行设置的，否则会拉取失败。

**GOPRIVATE**

指定哪些模块是 **私有模块**
- 对于匹配的私有模块：
	- 不会通过 `GOPROXY` 的公共代理下载
	- 不会通过 `GOSUMDB` 校验和数据库进行校验
- 用于处理公司内部或私有仓库的模块依赖。
- 默认值: 空，表示所有模块都不是私有模块。

**示例配置**:
```bash
GOPRIVATE=*.example.com,gitlab.internal.com
```
上述配置表示：
- 所有`*.example.com` 下的模块，以及 `gitlab.internal.com` 的模块是私有模块。
- Go 不会从公共代理拉取这些模块，也不会验证它们的校验和。

**GONOPROXY**
指定那个模块 **不使用代理模式**，而是直接从源头（如 Git 仓库）拉取。
- 常用于公司内部私有模块仓库。
- 默认值: 空，表示所有模块都通过代理下载。

**示例配置**:
```bash
GONOPROXY=*.example.com
```
- 所有 `*.example.com` 的模块不会通过 `GOPROXY` 下载，而是直接从源头（如 Git 仓库）拉取。

**GONOSUMDB**
指定哪些模块 **跳过校验和验证**。
- 常用于私有模块，避免与公共校验和数据库（如 `sum.golang.org`）交互。
- 默认值: 空，表示所有模块都需要校验和验证

**示例配置**:
```bash
GONOSUMDB=*.example.com
```
- 所有 `*.example.com` 的模块不会经过校验和数据库验证，而是直接信任本地的 `go.sum` 文件。

**综合示例**：
假设你的项目中有以下私有模块：
1. `*.example.com` 是公司内部的模块，不走代理，也不进行校验和验证。
2. `github.com/internal/module` 是公司从 GitHub 托管的私有模块，也需要跳过校验。

```bash
GOPRIVATE=*.example.com,github.com/internal/module
GONOPROXY=*.example.com
GONOSUMDB=*.example.com,github.com/internal/module
```

### 初始化项目

在完成 Go modules 的开启后，创建一个示例项目来进行演示，执行如下命令：
```bash
mkdir ~/Desktop/arc_web
cd ~/Desktop/arc_web
```

初始化Go modules模块
```bash
go init arc_web
```

该项目根目录下创建 `main.go`文件，如下：
```go
package main  

import (
    "fmt" 
    "github.com/fatih/color"
)
  
func main() {  
    fmt.Println("Hello World!")  
    color.Green("Hello, Go Modules!")
}
```

在项目根目录执行 `go get github.com/fatih/color` 命令，下载`github.com/fatih/color`包模块。
```bash
$ go get github.com/fatih/color  
go: added github.com/fatih/color v1.18.0
go: added github.com/mattn/go-colorable v0.1.13
go: added github.com/mattn/go-isatty v0.0.20
go: added golang.org/x/sys v0.25.0
```

##### go mod 文件

初始项目时，会生成一个 go.mod 文件，是启用了 Go modules 项目所必须的最重要的标识，描述当前项目（也就是当前模块）的元信息。
eg:
```go
module arc_web  
  
go 1.22.1  
  
require (  
    github.com/fatih/color v1.18.0 // indirect  
    github.com/mattn/go-colorable v0.1.13 // indirect  
    github.com/mattn/go-isatty v0.0.20 // indirect  
    golang.org/x/sys v0.25.0 // indirect  
)
```

###### 1.require
`require` 列出当前项目依赖的第三方模块及其版本。
```go
require (     
	github.com/gin-gonic/gin v1.9.0
	github.com/sirupsen/logrus v1.9.0
)
```

- 列出项目依赖的第三方模块及其版本。这里项目依赖于 Gin 框架和 Logrus 日志库。
###### 2.exclude
**`exclude`** 用于在 `go.mod` 文件中排除某些模块的特定版本，防止 Go 工具链下载并使用这些版本。
语法：
```go
exclude <module> <version>
```

```go
module arc_web  
  
go 1.22.1  
  
require (  
    github.com/fatih/color v1.18.0 // indirect  
    github.com/mattn/go-colorable v0.1.13 // indirect  
    github.com/mattn/go-isatty v0.0.20 // indirect  
    golang.org/x/sys v0.25.0 // indirect  
)

exclude github.com/fatih/color v1.17.0
```

- 项目依赖 `color` 的版本 `v1.18.0`。
- 显式排除了 `v1.17.0`，即使有其他依赖模块试图使用 `v1.8.1`，Go 也不会下载或使用该版本。

排除多个版本
``` go
module myproject

go 1.20

require (
    github.com/gin-gonic/gin v1.9.0
    github.com/sirupsen/logrus v1.9.0
)

exclude (
    github.com/gin-gonic/gin v1.7.7
    github.com/gin-gonic/gin v1.8.0
    github.com/sirupsen/logrus v1.8.1
)
```
解析：
- 明确指定`gin`和`logrus`的排除版本范围。
- 任何尝试使用`gin`的`v1.7.7`或`v1.8.0`，以及`logrus`的`v1.8.1`的依赖都会被拒绝。

###### 3.replace
`replace`替换依赖项的路径或版本，将某个依赖模块的路径或版本替换为其他路径或版本。
语法：
```go
replace <module> => <new-module> <new-version>
```

或者：
```go
replace <module> => <local-path>
```
示例：
```go
replace github.com/gin-gonic/gin v1.8.0 => github.com/gin-gonic/gin v1.9.0
```
将 `github.com/example/module` 版本 `v1.8.0` 替换为 `v1.9.0`。

```go
replace github.com/gin-gonic/gin => ../local-gin
```
将 `github.com/gin-gonic/gin` 替换为本地路径 `../local-gin`模块（本地开发常用）

###### 4.注意事项
1. **`replace` 的作用范围**：  
   - 仅对当前模块生效，不会影响其他模块。
   - 如果依赖链中的其他模块也使用 `replace`，可能会覆盖当前配置。

2. **`exclude` 的优先级**：  
   - 如果多个模块排除同一版本，最终会选择一个未被排除的版本。
   - 若所有版本均被排除，会导致构建失败。

3. **临时性与长期性**：  
   - `replace` 通常是临时解决方案，长期应推动依赖库修复问题。
   - `exclude` 需谨慎使用，可能破坏依赖关系。

###### 5.完整示例
结合 `replace` 和 `exclude` 的 `go.mod` 文件示例：
```go
module myproject

go 1.20

require (
    github.com/gin-gonic/gin v1.9.0
    github.com/sirupsen/logrus v1.9.0
)

// 替换依赖为本地路径
replace github.com/sirupsen/logrus => ../local-logrus
// 替换依赖为 Fork 后的仓库
replace github.com/buggy/module => github.com/your-fork/module v1.2.3

exclude (
    github.com/gin-gonic/gin v1.7.7
    github.com/gin-gonic/gin v1.8.0
    github.com/sirupsen/logrus v1.8.1
)
```

**require**  
列出项目所依赖的模块及其版本：
- 使用 Gin 框架 `github.com/gin-gonic/gin` 版本 `v1.9.0`
- 使用 Logrus 日志库 `github.com/sirupsen/logrus` 版本 `v1.9.0`
**replace**  
- 将 `github.com/sirupsen/logrus` 的依赖替换为本地路径 `../local-logrus`，方便在本地调试或使用自定义版本。路径相对于 `go.mod` 文件所在目录。
- 将 `replace github.com/buggy/module` 依赖替换为fork后 `github.com/your-fork/module v1.2.3`版本依赖
**exclude**  
- 排除 `github.com/gin-gonic/gin` 模块 `v1.8.0 v1.7.7` 版本。
- 排除 `github.com/sirupsen/logrus`模块 `v1.8.1` 版本
- 即使有间接依赖试图使用该版本，Go 工具链也会拒绝下载和使用它，从而确保项目使用正确的版本。

##### go sum 文件
在第一次拉取模块依赖后，会创建一个 go.sum 文件，其详细罗列了当前项目直接或间接依赖的所有模块版本，并写明了那些模块版本的 SHA-256 哈希值以备 Go 在今后的操作中保证项目所依赖的那些模块版本不会被篡改。

```go
github.com/fatih/color v1.18.0 h1:S8gINlzdQ840/4pfAwic/ZE0djQEH3wM94VfqLTZcOM=  
github.com/fatih/color v1.18.0/go.mod h1:4FelSpRwEGDpQ12mAdzqdOukCy4u8WUtOY6lkT/6HfU=  
github.com/mattn/go-colorable v0.1.13 h1:fFA4WZxdEF4tXPZVKMLwD8oUnCTTo08duU7wxecdEvA=  
github.com/mattn/go-colorable v0.1.13/go.mod h1:7S9/ev0klgBDR4GtXTXX8a3vIGJpMovkB8vQcUbaXHg=  
github.com/mattn/go-isatty v0.0.16/go.mod h1:kYGgaQfpe5nmfYZH+SKPsOc2e4SrIfOl2e/yFXSvRLM=  
github.com/mattn/go-isatty v0.0.20 h1:xfD0iDuEKnDkl03q4limB+vH+GxLEtL/jb4xVJSWWEY=  
github.com/mattn/go-isatty v0.0.20/go.mod h1:W+V8PltTTMOvKvAeJH7IuucS94S2C6jfK/D7dTCTo3Y=  
golang.org/x/sys v0.0.0-20220811171246-fbc7d0a398ab/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=  
golang.org/x/sys v0.6.0/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=  
golang.org/x/sys v0.25.0 h1:r+8e+loiHxRqhXVl6ML1nO3l1+oFoWbnlu2Ehimmi34=  
golang.org/x/sys v0.25.0/go.mod h1:/VUhepiaJMQUp4+oa/7Zr1D23ma6VTLIYjOOTFZPUcA=
```

每一行的格式：
```go
<模块路径> <版本> <校验和类型>:<校验和值>
```

- **模块路径**：依赖的模块名，例如 `github.com/gin-gonic/gin`。
- **版本**：模块的版本号，例如 `v1.9.0`。
- **校验和类型**：目前 Go 使用 `h1` 表示基于 SHA-256 的校验和（哈希值）。
- **校验和值**：经过 Base64 编码的哈希值，用于验证模块内容的完整性。

``` go
github.com/gin-gonic/gin v1.9.0 h1:9kDxQ37l1Ow8K8zKb1gLjkPrZV1sSd6HnAHO6q/B2zY=
github.com/gin-gonic/gin v1.9.0/go.mod h1:yyZ1oYITZFm4+AGuAGDDeL9T3rqv6JLCcaxDJ2Zv0Wk=
```
h1 hash 是 Go modules 将目标模块版本的 zip 文件打开包后，针对所有包内文件依次进行 hash，然后再把它们的 hash 结果按照固定格式和算法组成总的 hash 值。

第一行：
```go
github.com/gin-gonic/gin v1.9.0 h1:9kDxQ37l1Ow8K8zKb1gLjkPrZV1sSd6HnAHO6q/B2zY=
```
- 表示 Gin 模块版本 v1.9.0 的内容经过 SHA-256 哈希后得到的校验和值。

第二行：
```go
github.com/gin-gonic/gin v1.9.0/go.mod h1:yyZ1oYITZFm4+AGuAGDDeL9T3rqv6JLCcaxDJ2Zv0Wk=
```
- 表示 Gin 模块版本 v1.9.0 的 go.mod 文件的内容校验和值。

而 h1 hash 和 go.mod hash 两者，要么同时都存在，要么只存在 go.mod hash。那什么情况下，会不存在 h1 hash 呢？就当 Go 认为肯定用不到某个模块版本的时候就会省略它的 h1 hash，就会出现不存一个简化的例子说明在 h1 hash，只存在 go.mod hash 的情况。

###### **只存在 go.mod hash 示例**
假设你的项目依赖于模块 `github.com/foo/bar` 的某个版本，但你的代码中没有直接导入该模块的任何包，而只是通过其他依赖间接使用它。也就是说，在你的 `go.mod` 文件中，`github.com/foo/bar` 被标记为 **indirect**（间接依赖）。在这种情况下，Go 工具链为了解析依赖关系，只需要下载该模块的 `go.mod` 文件来确定版本和依赖约束，而不必下载整个模块的源码文件（zip 文件）。

因此，在 `go.sum` 文件中，会记录该模块的 `go.mod` 文件的校验和（带 `/go.mod` 后缀），而不会记录整个模块内容的 h1 哈希。

**1. go.mod 文件**
假设你的项目的 `go.mod` 文件如下：
```go
module github.com/example/myproject

go 1.20

require (
    // 作为间接依赖存在
    github.com/foo/bar v1.2.3 // indirect
)
```

这里，由于 `github.com/foo/bar` 只是间接依赖（项目代码中并未直接使用它），所以 Go 工具链在构建时只需要解析该模块的 `go.mod` 文件，而不需要下载整个模块的源码。

**2. go.sum 文件**
经过 `go mod tidy` 或构建操作后，生成的 `go.sum` 文件中，关于 `github.com/foo/bar` 的记录可能只包含如下条目：

```go
github.com/foo/bar v1.2.3/go.mod h1:XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX=`
```
而不会出现类似下面的条目：

```go
github.com/foo/bar v1.2.3 h1:YYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY=`
```
---
**解释说明**
**存在 only go.mod hash 的原因**：  
当 Go 工具链发现某个模块（例如 `github.com/foo/bar v1.2.3`）在编译过程中不会被直接使用（即没有实际导入模块内的包），它只需要该模块的 `go.mod` 文件来解析版本依赖和模块间关系。  
因此，为了节省下载和处理时间，Go 只会下载并校验该模块的 `go.mod` 文件，从而在 `go.sum` 中只记录了 `github.com/foo/bar v1.2.3/go.mod` 的校验和，而不会记录整个模块内容的 h1 hash。