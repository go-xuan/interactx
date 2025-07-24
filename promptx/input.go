package promptx

import (
	"strings"

	"github.com/go-xuan/typex"
	"github.com/manifoldco/promptui"
)

// Input 获取用户输入
func Input(label string) typex.Value {
	prompt := promptui.Prompt{Label: label}
	input, err := prompt.Run()
	input = strings.TrimSpace(input)
	if err != nil || input == "" {
		return typex.ZeroValue()
	} else {
		return typex.StringValue(input)
	}
}

// InputMust 获取用户输入
func InputMust(label string) typex.Value {
	prompt := promptui.Prompt{Label: label}
	input, err := prompt.Run()
	for strings.TrimSpace(input) == "" || err != nil {
		input, err = prompt.Run()
	}
	return typex.StringValue(input)
}

// InputValidate 获取用户输入（自定义验证）
func InputValidate(label string, validate func(string) error) typex.Value {
	prompt := promptui.Prompt{
		Label:    label,
		Validate: validate,
	}
	if input, err := prompt.Run(); err != nil {
		return typex.ZeroValue()
	} else {
		return typex.StringValue(strings.TrimSpace(input))
	}
}
