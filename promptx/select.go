package promptx

import (
	"errors"

	"github.com/manifoldco/promptui"
)

// Option 选项
type Option interface {
	ActiveTemplate() string   // 当前选择项模板
	InactiveTemplate() string // 其他待选项模板
	SelectedTemplate() string // 已选项模板
	DetailsTemplate() string  // 详情模板
	SearchMatch(string) bool  // 搜索匹配
}

// Select 选择器
func Select[OPT Option](label string, opts []OPT) (OPT, error) {
	var opt OPT
	if len(opts) == 0 {
		return opt, errors.New("opts is empty")
	}
	if index, _, err := getSelector(label, opts).Run(); err != nil {
		return opt, err
	} else {
		return opts[index], nil
	}
}

// SelectMust 选择器
func SelectMust[OPT Option](label string, opts []OPT) OPT {
	if index, _, err := getSelector(label, opts).Run(); err != nil {
		return opts[0]
	} else {
		return opts[index]
	}
}

func getSelector[OPT Option](label string, opts []OPT) *promptui.Select {
	opt := opts[0]
	return &promptui.Select{
		Label: label,
		Items: opts,
		Templates: &promptui.SelectTemplates{
			Label:    "---------- {{ . | red }} ----------",
			Active:   opt.ActiveTemplate(),
			Inactive: opt.InactiveTemplate(),
			Selected: opt.SelectedTemplate(),
			Details:  opt.DetailsTemplate(),
		},
		Size:              10,
		StartInSearchMode: true,
		Searcher: func(input string, i int) bool {
			return opts[i].SearchMatch(input)
		},
	}
}
