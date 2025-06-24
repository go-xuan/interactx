package flagx

import (
	"flag"

	"github.com/go-xuan/typex"
)

func StringOption(name, usage string, def string) Option {
	return &stringOption{
		baseOption: baseOption{
			name:  name,
			usage: usage,
		},
		def: def,
	}
}

type stringOption struct {
	baseOption
	value *string
	def   string
}

func (opt *stringOption) Name() string {
	return opt.name
}

func (opt *stringOption) Usage() string {
	if opt.def != "" {
		return genUsage(opt.usage, opt.def)
	} else {
		return opt.usage
	}
}

func (opt *stringOption) Set(fs *flag.FlagSet) {
	if opt.value == nil {
		opt.value = fs.String(opt.name, opt.def, opt.usage)
	}
}

func (opt *stringOption) Get() typex.Value {
	if value := opt.value; value != nil {
		return typex.StringValue(*value)
	}
	return typex.ZeroValue()
}
