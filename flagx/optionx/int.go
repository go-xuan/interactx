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

func (o *intOption) GetUsage() string {
	if o.def != 0 {
		return o.getUsage(o.def)
	} else {
		return o.usage
	}
}

func (o *intOption) SetFS(fs *flag.FlagSet) {
	if o.value == nil {
		o.value = fs.Int(o.name, o.def, o.usage)
	}
}
func (o *intOption) GetValue() typex.Value {
	if value := o.value; value != nil {
		return typex.IntValue(*value)
	}
	return typex.ZeroValue()
}
