package flagx

import (
	"flag"

	"github.com/go-xuan/typex"
)

func IntOption(name, usage string, def int) Option {
	return &intOption{
		baseOption: baseOption{
			name:  name,
			usage: usage,
		},
		def: def,
	}
}

type intOption struct {
	baseOption
	value *int
	def   int
}

func (opt *intOption) Name() string {
	return opt.name
}

func (opt *intOption) Usage() string {
	if opt.def != 0 {
		return genUsage(opt.usage, opt.def)
	} else {
		return opt.usage
	}
}

func (opt *intOption) Set(fs *flag.FlagSet) {
	if opt.value == nil {
		opt.value = fs.Int(opt.name, opt.def, opt.usage)
	}
}
func (opt *intOption) Get() typex.Value {
	if value := opt.value; value != nil {
		return typex.IntValue(*value)
	}
	return typex.ZeroValue()
}
