package promptx

import (
	"errors"
	"github.com/manifoldco/promptui"
)

type Option interface {
	Active() string
	Inactive() string
	Selected() string
	Details() string
}

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
			Label:    "---------- {{ .| red }} ----------",
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

func SelectString(label string, opts []string) (string, error) {
	prompt := promptui.Select{
		Label: label,
		Items: opts,
		Templates: &promptui.SelectTemplates{
			Label:    "---------- {{ .| red }} ----------",
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
