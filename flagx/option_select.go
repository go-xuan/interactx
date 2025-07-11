package flagx

import (
	"flag"
	"strings"

	"github.com/go-xuan/typex"

	"github.com/go-xuan/interactx/promptx"
)

func SelectOption(name, usage string, opts []promptx.String) Option {
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

func (o *selectOption) Name() string {
	return o.baseOption.Name()
}

func (o *selectOption) Usage() string {
	def := strings.Builder{}
	for i, opt := range o.opts {
		if i > 0 {
			def.WriteString("/")
		}
		def.WriteString(opt.Value)
	}
	return genUsage(o.usage, def.String())
}

func (o *selectOption) Set(*flag.FlagSet) {
	return
}

func (o *selectOption) Get() typex.Value {
	if s, err := promptx.Select(o.usage, o.opts); err == nil {
		return typex.StringValue(s.Value)
	}
	return typex.ZeroValue()
}
