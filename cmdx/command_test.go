package cmdx

import (
	"fmt"
	"testing"
)

func TestCommand(t *testing.T) {
	var command = NewCommand("test", "测试")
	command.AddOption(
		Int("size", "数量", 0),
	)
	command.SetExecutor(func() error {
		size := command.GetOptionValue("size")
		fmt.Println("size = ", size.Int(1))
		return nil
	})

	// 注册命令
	Register(command)

	// 执行命令
	if err := Execute("test", "-size", "88"); err != nil {
		fmt.Println(err)
	}
}
