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

func (o *inputOption) Name() string {
	return o.baseOption.Name()
}

func (o *inputOption) Usage() string {
	return o.usage
}

func (o *inputOption) Set(*flag.FlagSet) {
	return
}

func (o *inputOption) Get() typex.Value {
	return promptx.Input(o.usage)
}
