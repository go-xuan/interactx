package optionx

import (
	"flag"

	"github.com/go-xuan/typex"
)

func Bool(name, usage string, def bool) Option {
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

func (o *boolOption) GetUsage() string {
	if o.def {
		return o.getUsage(o.def)
	} else {
		return o.usage
	}
}

func (o *boolOption) SetFS(fs *flag.FlagSet) {
	if o.value == nil {
		o.value = fs.Bool(o.name, o.def, o.usage)
	}
}

func (o *boolOption) GetValue() typex.Value {
	if value := o.value; value != nil {
		return typex.NewBool(*value)
	}
	return typex.NewZero()
}
