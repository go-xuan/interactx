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
	Search(string) bool
}

// Select 自定义选择器
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
		Size:              10,
		StartInSearchMode: true,
		Searcher: func(input string, i int) bool {
			return opts[i].Search(input)
		},
	}
	if index, _, err := prompt.Run(); err != nil {
		return opt, err
	} else {
		return opts[index], nil
	}
}

// SelectString 字符串选择器
func SelectString(label string, opts []string) (string, error) {
	var list []String
	for _, opt := range opts {
		list = append(list, String{Label: opt, Value: opt})
	}
	if s, err := Select(label, list); err != nil {
		return "", err
	} else {
		return s.Value, nil
	}
}

// SelectBool 布尔选择器
func SelectBool(label string) bool {
	s, _ := SelectString(label, []string{"TRUE", "FALSE"})
	return s == "TRUE"
}
