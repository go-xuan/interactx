package optionx

import (
	"flag"
	"strings"

	"github.com/go-xuan/typex"

	"github.com/go-xuan/interactx/promptx"
)

func Select(name, usage string, opts []promptx.String) Option {
	return &selectOption{
		baseOption: baseOption{
			name:  name,
			usage: usage,
		},
		opts: opts,
	}
}

type selectOption struct {
	baseOption
	opts []promptx.String
}

func (o *selectOption) GetUsage() string {
	def := strings.Builder{}
	for i, opt := range o.opts {
		if i > 0 {
			def.WriteString("/")
		}
		def.WriteString(opt.Value)
	}
	return o.getUsage(def.String())
}

func (o *selectOption) SetFS(*flag.FlagSet) {
	return
}

func (o *selectOption) GetValue() typex.Value {
	if s, err := promptx.Select(o.usage, o.opts); err == nil {
		return typex.StringValue(s.Value)
	}
	return typex.ZeroValue()
}
