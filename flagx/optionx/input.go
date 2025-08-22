package optionx

import (
	"flag"

	"github.com/go-xuan/typex"

	"github.com/go-xuan/interactx/promptx"
)

func Input(name, usage string) Option {
	return &inputOption{
		baseOption: baseOption{
			name:  name,
			usage: usage,
		},
	}
}

type inputOption struct {
	baseOption
}

func (o *inputOption) GetUsage() string {
	return o.usage
}

func (o *inputOption) SetFS(*flag.FlagSet) {
	return
}

func (o *inputOption) GetValue() typex.Value {
	return promptx.Input(o.usage)
}
