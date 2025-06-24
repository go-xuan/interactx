package flagx

import (
	"fmt"
	"testing"
)

func TestCommand(t *testing.T) {
	command := NewCommand("test", "测试")
	command.AddOption(
		IntOption("size", "数量", 0),
	)
	command.SetExecutor(
		func() error {
			command.NeedHelp()
			size := command.GetOptionValue("size")
			fmt.Println("size = ", size.Int(1))
			return nil
		},
	)

	// 注册命令
	if err := Execute(); err != nil {
		fmt.Println(err)
	}
}
