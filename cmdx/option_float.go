package cmdx

import (
	"flag"

	"github.com/go-xuan/typex"
)

func Float(name, usage string, def float64) Option {
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

func (o *floatOption) GetUsage() string {
	if o.def != float64(0) {
		return o.getUsage(o.def)
	}
	return o.usage
}

func (o *floatOption) SetFS(fs *flag.FlagSet) {
	if o.value == nil {
		o.value = fs.Float64(o.name, o.def, o.usage)
	}
}

func (o *floatOption) GetValue() typex.Value {
	if value := o.value; value != nil {
		return typex.NewFloat64(*value)
	}
	return typex.NewZero()
}
