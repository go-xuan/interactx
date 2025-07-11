package flagx

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/go-xuan/typex"

	"github.com/go-xuan/interactx/colorx"
)

// Main 根命令
var Main = NewCommand("main", "根命令")

// NewCommand 创建命令
func NewCommand(name, usage string) *Command {
	return &Command{
		name:      name,
		usage:     usage,
		options:   []string{},
		optionMap: make(map[string]Option),
		subs:      []string{},
		subMap:    make(map[string]*Command),
	}
}

// Register 注册子命令
func Register(commands ...*Command) {
	Main.Register()
	for _, command := range commands {
		command.Join(Main)
		command.Register()
	}
}

// Execute 执行入口
func Execute(args ...string) error {
	if Main == nil {
		return errors.New("根命令未初始化")
	} else if Main.status != 1 {
		return errors.New("根命令未注册")
	}
	if len(args) > 0 {
		Main.args = args
	} else {
		Main.args = os.Args[1:]
	}
	return Main.Execute()
}

type Command struct {
	name      string              // 命令名
	usage     string              // 命令用法说明
	parent    *Command            // 父命令
	subs      []string            // 子命令名，有序
	subMap    map[string]*Command // 子命令map
	options   []string            // 选项名，有序
	optionMap map[string]Option   // 选项map
	fs        *flag.FlagSet       // FlagSet
	args      []string            // 当前命令的执行参数
	status    int                 // 状态（0:初始化/1:注册/2:执行）
	executor  func() error        // 命令执行器
}

// Join 添加父命令
func (c *Command) Join(command *Command) *Command {
	if command.parent != nil {
		panic(fmt.Sprintf("[%s]不是根命令", command.name))
	}
	c.parent = command
	return c
}

// AddOption 添加参数
func (c *Command) AddOption(options ...Option) *Command {
	for _, option := range options {
		if name := option.Name(); name != "" {
			if _, ok := c.optionMap[name]; !ok {
				c.options = append(c.options, name)
			}
			c.optionMap[name] = option
		}
	}
	return c
}

// SetExecutor 设置执行器
func (c *Command) SetExecutor(executor func() error) *Command {
	c.executor = executor
	return c
}

func (c *Command) Register() {
	// 已注册则直接跳过
	if c.status == 1 {
		return
	}
	var name = c.name
	if parent := c.parent; parent != nil {
		if c.executor == nil {
			panic(fmt.Sprintf("[%s]子命令未设置执行器！", name))
		}
		if _, ok := parent.subMap[name]; ok {
			panic(fmt.Sprintf("[%s]子命令重复注册！", name))
		}
		parent.subs = append(parent.subs, name)
		parent.subMap[name] = c
	} else if len(c.subs) > 0 {
		for _, sub := range c.subs {
			if subCommand, ok := c.subMap[sub]; ok {
				subCommand.Register()
			}
		}
	}
	c.status = 1
	c.addDefaultOption()
}

// Execute 执行命令
func (c *Command) Execute() error {
	if c.status == 0 {
		return errors.New("请先注册命令")
	} else if c.status == 2 {
		return errors.New("请勿重复执行命令")
	}
	if err := c.execute(); err != nil {
		return errors.New("执行命令失败：" + err.Error())
	}
	c.status = 2
	return nil
}

func (c *Command) execute() error {
	if args := c.args; len(args) > 0 {
		name := strings.ToLower(args[0])
		if sub, ok := c.subMap[name]; ok {
			sub.args = args[1:]
			return sub.execute()
		}
	}

	var name = c.Name()
	if err := c.ParseArgs(); err != nil {
		return err
	}
	if executor := c.executor; executor != nil {
		fmt.Printf("======当前执行命令:[%s]======\n", name)
		if err := executor(); err != nil {
			return err
		}
	} else if c.parent != nil {
		return fmt.Errorf("子命令未设置执行器:[%s]", name)
	} else {
		c.PrintSubs()
	}
	c.status = 2
	return nil
}

func (c *Command) Name() string {
	if c.parent != nil {
		return c.parent.Name() + "." + c.name
	}
	return c.name
}

// ParseArgs 解析参数值到选项中
func (c *Command) ParseArgs() error {
	fs := c.FlagSet()
	// 绑定参数到FlagSet
	for _, option := range c.optionMap {
		option.Set(fs)
	}
	// 解析FlagSet
	if err := fs.Parse(c.args); err != nil {
		return errors.New("FlagSet解析失败：" + err.Error())
	}
	return nil
}

// GetOptionValue 获取参数值
func (c *Command) GetOptionValue(name string) typex.Value {
	if option, ok := c.optionMap[name]; ok {
		if value := option.Get(); value != nil {
			if value.String() == "-h" {
				_ = c.FlagSet().Set("h", "true")
			} else {
				return value
			}
		}
	}
	return typex.ZeroValue()
}

// FlagSet 初始化FlagSet并将参数注册到FlagSet
func (c *Command) FlagSet() *flag.FlagSet {
	if c.fs == nil {
		c.fs = flag.NewFlagSet(c.name, flag.ExitOnError)
	}
	return c.fs
}

func (c *Command) NeedHelp() bool {
	return c.GetOptionValue("h").Bool()
}

func (c *Command) GetArgs() []string {
	return c.args
}

func (c *Command) GetArg(index int) string {
	if index >= 0 && index < len(c.args) {
		return c.args[index]
	}
	return ""
}

func (c *Command) addDefaultOption() {
	c.AddOption(
		BoolOption("h", "帮助说明", false),
	)
}

// PrintSubs 打印子命令
func (c *Command) PrintSubs() {
	if len(c.subs) == 0 {
		return
	}
	fmt.Printf("[%s]子命令：\n", colorx.Cyan(c.name))
	for _, name := range c.subs {
		sub := c.subMap[name]
		fmt.Printf("%-50s %s\n", colorx.Magenta(name), sub.usage)
	}
}

// PrintOptions 打印命令选项
func (c *Command) PrintOptions() {
	if len(c.options) == 0 {
		return
	}
	fmt.Printf("[%s]命令选项：\n", colorx.Cyan(c.name))
	for _, optName := range c.options {
		option := c.optionMap[optName]
		fmt.Printf("%-50s %s\n", colorx.Magenta("-"+option.Name()), option.Usage())
	}
}
