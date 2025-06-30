package promptx

import (
	"errors"

	"github.com/manifoldco/promptui"
)

// Option 选项
type Option interface {
	Active() string
	Inactive() string
	Selected() string
	Details() string
}

// Select 自定义选项
func Select[OPT Option](label string, opts []OPT) (OPT, error) {
	var opt OPT
	if len(opts) == 0 {
		return opt, errors.New("opts is empty")
	}
	opt = opts[0]
	prompt := promptui.Select{
		Label: label,
		Items: opts,
		Templates: &promptui.SelectTemplates{
			Label:    "---------- {{ . | red }} ----------",
			Active:   opt.Active(),
			Inactive: opt.Inactive(),
			Selected: opt.Selected(),
			Details:  opt.Details(),
		},
		Size: 10,
	}
	if index, _, err := prompt.Run(); err != nil {
		return opt, err
	} else {
		return opts[index], nil
	}
}

// SelectString 选择字符串
func SelectString(label string, opts []string) (string, error) {
	prompt := promptui.Select{
		Label: label,
		Items: opts,
		Templates: &promptui.SelectTemplates{
			Label:    "---------- {{ . | red }} ----------",
			Active:   "* {{ . | cyan }}",
			Inactive: " {{ . | white }}",
			Selected: "* {{. | red | faint }}",
		},
		Size: 10,
	}
	if index, _, err := prompt.Run(); err != nil {
		return "", err
	} else {
		return opts[index], nil
	}
}

// SelectBool 选择布尔值
func SelectBool(label string) (bool, error) {
	opts := []string{"true", "false"}
	prompt := promptui.Select{
		Label: label,
		Items: opts,
		Templates: &promptui.SelectTemplates{
			Label:    "---------- {{ . | red }} ----------",
			Active:   "* {{ . | cyan }}",
			Inactive: " {{ . | white }}",
			Selected: "* {{. | red | faint }}",
		},
		Size: 2,
	}
	if index, _, err := prompt.Run(); err != nil {
		return false, err
	} else {
		return opts[index] == "true", nil
	}
}
