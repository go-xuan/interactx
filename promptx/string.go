package promptx

import "strings"

// SelectString 字符串选择器
func SelectString(label string, opts []string) (string, error) {
	var list []String
	for _, opt := range opts {
		list = append(list, String{
			Label: opt,
			Value: opt,
		})
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
--------- {{ "Details" | faint }} ----------
{{ "Label:" | faint }}	{{ .Label }}
{{ "Value:" | faint }}	{{ .Value }}`
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
