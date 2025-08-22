package promptx

import "strings"

// String 字符串选项
type String struct {
	Label string // 选项标签
	Value string // 选项值
}

func (String) ActiveTemplate() string {
	return "* {{ .Label | cyan }} "
}

func (String) InactiveTemplate() string {
	return " {{ .Label | white }}"
}

func (String) SelectedTemplate() string {
	return "* {{ .Label | red | faint }}"
}

func (String) DetailsTemplate() string {
	return `
--------- {{ "VALUE" | faint }} ----------
{{ .Value }}`
}

func (s String) SearchMatch(input string) bool {
	if strings.Contains(s.Label, input) {
		return true
	}
	if strings.Contains(s.Value, input) {
		return true
	}
	return false
}
