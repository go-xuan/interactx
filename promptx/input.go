package promptx

import (
	"strings"

	"github.com/go-xuan/typex"
	"github.com/manifoldco/promptui"
)

// Input 获取用户输入的自定义命令
func Input(label string) typex.Value {
	prompt := promptui.Prompt{
		Label: label,
		Validate: func(s string) error {
			return nil
		},
	}
	if input, err := prompt.Run(); err != nil {
		return typex.ZeroValue()
	} else {
		return typex.StringValue(strings.TrimSpace(input))
	}
}
