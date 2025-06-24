package flagx

import (
	"flag"

	"github.com/go-xuan/typex"
)

func FloatOption(name, usage string, def float64) Option {
	return &floatOption{
		baseOption: baseOption{
			name:  name,
			usage: usage,
		},
		def: def,
	}
}

type floatOption struct {
	baseOption
	value *float64
	def   float64
}

func (opt *floatOption) Name() string {
	return opt.name
}

func (opt *floatOption) Usage() string {
	if opt.def != float64(0) {
		return genUsage(opt.usage, opt.def)
	} else {
		return opt.usage
	}
}

func (opt *floatOption) Set(fs *flag.FlagSet) {
	if opt.value == nil {
		opt.value = fs.Float64(opt.name, opt.def, opt.usage)
	}
}

func (opt *floatOption) Get() typex.Value {
	if value := opt.value; value != nil {
		return typex.Float64Value(*value)
	}
	return typex.ZeroValue()
}
