package optionx

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

func (o *floatOption) Name() string {
	return o.baseOption.Name()
}

func (o *floatOption) Usage() string {
	if o.def != float64(0) {
		return genUsage(o.usage, o.def)
	} else {
		return o.usage
	}
}

func (o *floatOption) Set(fs *flag.FlagSet) {
	if o.value == nil {
		o.value = fs.Float64(o.name, o.def, o.usage)
	}
}

func (o *floatOption) Get() typex.Value {
	if value := o.value; value != nil {
		return typex.Float64Value(*value)
	}
	return typex.ZeroValue()
}
