package promptx

import (
	"strings"
)

// SelectFunction 选择函数
func SelectFunction(label string, functions []Function) (func() error, error) {
	function, err := Select(label, functions)
	if err != nil {
		return nil, err
	}
	return function.Function, nil
}

type Function struct {
	Name     string       // 函数名称
	Desc     string       // 函数描述
	Function func() error // 函数
}

func (Function) ActiveTemplate() string {
	return "* {{ .Name | cyan }}({{ .Desc | cyan }})"
}

func (Function) InactiveTemplate() string {
	return " {{ .Name | white }}"
}

func (Function) SelectedTemplate() string {
	return "* {{ .Name | red | faint }}"
}

func (Function) DetailsTemplate() string {
	return `
--------- {{ "Details" | faint }} ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Desc:" | faint }}	{{ .Desc }}`
}

func (s Function) SearchMatch(input string) bool {
	if strings.Contains(s.Name, input) {
		return true
	}
	if strings.Contains(s.Desc, input) {
		return true
	}
	return false
}
