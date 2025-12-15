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
	}); value != "" {
		return typex.NewString(value)
	}
	return typex.NewZero()
}

// InputDefault 获取用户输入, 可以默认值
func InputDefault(label string, def string) typex.Value {
	if value := input(promptui.Prompt{
		Label:     label,
		Default:   def,
		AllowEdit: true,
	}); value != "" {
		return typex.NewString(value)
	}
	return typex.NewString(def)
}

// InputValidate 获取用户输入, 并校验输入
func InputValidate(label string, Validate func(string) error) typex.Value {
	if value := input(promptui.Prompt{
		Label:    label,
		Validate: Validate,
	}); value != "" {
		return typex.NewString(value)
	}
	return typex.NewZero()
}

// InputMust 获取用户输入
func InputMust(label string) typex.Value {
	value := input(promptui.Prompt{
		Label: label,
	}, true)
	return typex.NewString(value)
}

// InputAndDo 根据用户输入执行函数, 直到输入为空
func InputAndDo(label string, do func(string) error) error {
	for text := Input(label).String(); text != ""; text = Input(label).String() {
		if err := do(text); err != nil {
			return err
		}
	}
	return nil
}
