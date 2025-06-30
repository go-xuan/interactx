package promptx

// String 字符串选项
type String struct {
	Name        string
	Description string
}

func (String) Active() string {
	return "* {{ .Name | cyan }} | {{ .Description | cyan }}"
}

func (String) Inactive() string {
	return " {{ .Name | white }} | {{.Description | white }}"
}

func (String) Selected() string {
	return "* {{ .Name | red | faint }}"
}

func (String) Details() string {
	return `
--------- Details ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Description:" | faint }}	{{ .Description }}`
}
