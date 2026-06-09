# interactx

终端交互库，提供 CLI 命令框架、交互式提示、颜色输出和文本对齐。

## 安装

```bash
go get github.com/go-xuan/interactx
```

## 快速开始

```go
import (
    "github.com/go-xuan/interactx/cmdx"
    "github.com/go-xuan/interactx/promptx"
    "github.com/go-xuan/interactx/colorx"
)

// 注册命令
cmd := cmdx.NewCommand("hello", "say hello")
cmd.Execute(func() { colorx.Green("Hello, World!") })

// 交互式选择
name, _ := promptx.Choose("请选择:", []string{"Alice", "Bob"})
```

## 子包

| 子包 | 用途 |
|------|------|
| **alignx** | 文本对齐与格式化 |
| **cmdx** | CLI 命令框架（注册、参数解析、选项类型） |
| **colorx** | 终端颜色输出 |
| **promptx** | 交互式提示（输入、选择、确认） |
