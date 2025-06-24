package promptx

import (
	"errors"
	"strings"

	"github.com/go-xuan/typex"

	"github.com/manifoldco/promptui"
)

// Input 获取用户输入的自定义命令
func Input(label string) (typex.Value, error) {
	prompt := promptui.Prompt{
		Label: label,
		Validate: func(s string) error {
			if strings.TrimSpace(s) == "" {
				return errors.New("invalid input")
			}
			return nil
		},
	}
	if input, err := prompt.Run(); err != nil {
		return nil, err
	} else {
		return typex.StringValue(strings.TrimSpace(input)), nil
	}
}
