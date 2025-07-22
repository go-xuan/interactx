package optionx

import (
	"flag"

	"github.com/go-xuan/typex"
)

func Int(name, usage string, def int) Option {
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

func (o *intOption) Name() string {
	return o.baseOption.Name()
}

func (o *intOption) Usage() string {
	if o.def != 0 {
		return genUsage(o.usage, o.def)
	} else {
		return o.usage
	}
}

func (o *intOption) Set(fs *flag.FlagSet) {
	if o.value == nil {
		o.value = fs.Int(o.name, o.def, o.usage)
	}
}
func (o *intOption) Get() typex.Value {
	if value := o.value; value != nil {
		return typex.IntValue(*value)
	}
	return typex.ZeroValue()
}
