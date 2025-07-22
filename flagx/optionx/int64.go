package optionx

import (
	"flag"

	"github.com/go-xuan/typex"
)

func Int64(name, usage string, def int64) Option {
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

func (o *int64Option) Name() string {
	return o.baseOption.Name()
}

func (o *int64Option) Usage() string {
	if o.def != 0 {
		return genUsage(o.usage, o.def)
	} else {
		return o.usage
	}
}

func (o *int64Option) Set(fs *flag.FlagSet) {
	if o.value == nil {
		o.value = fs.Int64(o.name, o.def, o.usage)
	}
}

func (o *int64Option) Get() typex.Value {
	if value := o.value; value != nil {
		return typex.Int64Value(*value)
	}
	return typex.ZeroValue()
}
