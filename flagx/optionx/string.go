package optionx

import (
	"flag"

	"github.com/go-xuan/typex"
)

func String(name, usage string, def string) Option {
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

func (o *stringOption) GetUsage() string {
	if o.def != "" {
		return o.getUsage(o.def)
	} else {
		return o.usage
	}
}

func (o *stringOption) SetFS(fs *flag.FlagSet) {
	if o.value == nil {
		o.value = fs.String(o.name, o.def, o.usage)
	}
}

func (o *stringOption) GetValue() typex.Value {
	if value := o.value; value != nil {
		return typex.StringValue(*value)
	}
	return typex.ZeroValue()
}
