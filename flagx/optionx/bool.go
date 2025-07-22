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

func (o *boolOption) Name() string {
	return o.baseOption.Name()
}

func (o *boolOption) Usage() string {
	if o.def {
		return genUsage(o.usage, o.def)
	} else {
		return o.usage
	}
}

func (o *boolOption) Set(fs *flag.FlagSet) {
	if o.value == nil {
		o.value = fs.Bool(o.name, o.def, o.usage)
	}
}

func (o *boolOption) Get() typex.Value {
	if value := o.value; value != nil {
		return typex.BoolValue(*value)
	}
	return typex.ZeroValue()
}
