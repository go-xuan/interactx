package flagx

import (
	"flag"

	"github.com/go-xuan/typex"
)

func Int64Option(name, usage string, def int64) Option {
	return &int64Option{
		baseOption: baseOption{
			name:  name,
			usage: usage,
		},
		def: def,
	}
}

type int64Option struct {
	baseOption
	value *int64
	def   int64
}

func (opt *int64Option) Name() string {
	return opt.name
}

func (opt *int64Option) Usage() string {
	if opt.def != 0 {
		return genUsage(opt.usage, opt.def)
	} else {
		return opt.usage
	}
}

func (opt *int64Option) Set(fs *flag.FlagSet) {
	if opt.value == nil {
		opt.value = fs.Int64(opt.name, opt.def, opt.usage)
	}
}

func (opt *int64Option) Get() typex.Value {
	if value := opt.value; value != nil {
		return typex.Int64Value(*value)
	}
	return typex.ZeroValue()
}
