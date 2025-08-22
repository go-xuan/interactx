package promptx

import (
	"strings"

	"github.com/go-xuan/typex"
	"github.com/manifoldco/promptui"
)

func input(prompt promptui.Prompt, must ...bool) string {
	value, err := prompt.Run()
	value = strings.TrimSpace(value)
	if len(must) > 0 && must[0] {
		for value == "" || err != nil {
			value, err = prompt.Run()
			value = strings.TrimSpace(value)
		}
	}
	return value
}

// Input 获取用户输入
func Input(label string) typex.Value {
	if value := input(promptui.Prompt{
		Label: label,
	}); value == "" {
		return typex.ZeroValue()
	} else {
		return typex.StringValue(value)
	}
}

// InputDefault 获取用户输入, 可以默认值
func InputDefault(label string, def string) typex.Value {
	if value := input(promptui.Prompt{
		Label:   label,
		Default: def,
	}); value == "" {
		return typex.StringValue(def)
	} else {
		return typex.StringValue(value)
	}
}

// InputMust 获取用户输入
func InputMust(label string) typex.Value {
	value := input(promptui.Prompt{
		Label: label,
	}, true)
	return typex.StringValue(value)
}
