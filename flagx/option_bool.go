package flagx

import (
	"flag"

	"github.com/go-xuan/typex"
)

func BoolOption(name, usage string, def bool) Option {
	return &boolOption{
		baseOption: baseOption{
			name:  name,
			usage: usage,
		},
		def: def,
	}
}

type boolOption struct {
	baseOption
	value *bool
	def   bool
}

func (opt *boolOption) Name() string {
	return opt.name
}

func (opt *boolOption) Usage() string {
	if opt.def {
		return genUsage(opt.usage, opt.def)
	} else {
		return opt.usage
	}
}

func (opt *boolOption) Set(fs *flag.FlagSet) {
	if opt.value == nil {
		opt.value = fs.Bool(opt.name, opt.def, opt.usage)
	}
}

func (opt *boolOption) Get() typex.Value {
	if value := opt.value; value != nil {
		return typex.BoolValue(*value)
	}
	return typex.ZeroValue()
}
