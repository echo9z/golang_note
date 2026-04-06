# Golang学习代码库 - AI编码指南

## 项目概述

这是一个结构化的Go语言学习资源库，包含渐进式的代码示例、模块演示和详细文档。主要目的是教学和参考，从基础语法到高级模式。

## 代码库架构

### 学习路径结构
- **01-helloworld/**: 基础Hello World示例
- **02_0-base/**: 核心概念的模块化学习（包管理、变量、运算符、常量、数据类型）
  - 每个子目录独立展示特定Go语言特性
  - 遵循渐进复杂性原则
- **02-basegra/**: 高级集成示例（多包协作、数据库集成）
- **datasource/**: 数据源和数据库交互模式

### 关键特性

**包和可见性模式**（见 [02_0-base/01_package/](02_0-base/01_package/) 和 [02_0-base/02_private_pkg/](02_0-base/02_private_pkg/)）
- 使用别名导入：`import m "math"`
- 下划线导入执行init函数：`import _ "01_package/driver"`
- 公开常量约定：`PI` 需要首字母大写
- internal包用于真正的私有实现
- 私有包与公开API的对比

**模块组织**（见各目录的go.mod文件）
- 每个学习单元是独立的Go模块
- 展示正确的go.mod初始化和依赖管理
- 演示别名导入在多包项目中的应用

**测试约定**（见 [02-basegra/utils/](02-basegra/utils/) 的 `*_test.go`文件）
- 测试文件命名：`<filename>_test.go`
- 测试函数签名：`func Test<FunctionName>(t *testing.T)`

## 开发工作流

### 构建和运行示例
```bash
# 直接运行（推荐用于学习示例）
cd <example-directory>
go run main.go

# 完整编译过程
go build
./<program-name>  # Linux/Mac
<program-name>.exe  # Windows
```

### 依赖管理
```bash
# 在修改imports后清理
go mod tidy

# 下载依赖但不编译
go mod download

# 验证校验和
go mod verify
```

### 测试
```bash
# 当前目录测试
go test -v

# 显示覆盖率
go test -cover

# 基准测试（if present）
go test -bench=.
```

## 项目特定的约定和模式

### 命名约定
- Go标准约定：首字母大写=公开，小写=私有
- 项目中使用中英文混合注释以增加可访问性

### 数据库集成
- 使用 `github.com/lib/pq` PostgreSQL驱动
- 数据源模式在 [datasource/](datasource/) 目录中演示
- 示例包含SQL查询和结果处理

### Go版本
- 本库基于Go 1.22.1和1.23.1
- 大部分示例兼容Go 1.20+

### 环境配置（关键）
对于中国用户或需要国内代理的场景：
```bash
export GOPROXY=https://goproxy.cn,direct
# 备选: https://mirrors.aliyun.com/goproxy/ 或 https://goproxy.io
export GO111MODULE=on
```

## 常见任务

### 添加新学习模块
1. 在 `02_0-base/` 中创建新目录（格式：`XX_topic-name/`）
2. 初始化 `go.mod`：`go mod init <module-name>`
3. 创建 `main.go` 或示例文件
4. 在README.md中添加文档

### 修改现有示例
- 修改后运行 `go mod tidy` 确保依赖正确
- 测试文件运行：`go test ./...` 在模块目录
- 遵循现有的包结构和注释风格

### 引用其他模块
使用导入别名避免冲突：
```go
import (
    base "basegra/base"
    utils "basegra/utils"
)
```

## 文档参考

- **CLAUDE.md**: 完整的代码库说明（此文件的源头）
- **README.md**: 详细的Go学习指南和概念说明（71KB+）
- **go mod.md**: 深入的模块系统和依赖管理
- **go mod publish.md**: 模块发布指南

## 代码质量标准

- 所有示例遵循官方 [Effective Go](https://go.dev/doc/effective_go) 约定
- 使用 `go fmt` 自动格式化
- 关键功能包含测试和注释
- 错误处理贯穿所有示例
