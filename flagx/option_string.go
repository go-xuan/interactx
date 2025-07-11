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

func (o *stringOption) Name() string {
	return o.baseOption.Name()
}

func (o *stringOption) Usage() string {
	if o.def != "" {
		return genUsage(o.usage, o.def)
	} else {
		return o.usage
	}
}

func (o *stringOption) Set(fs *flag.FlagSet) {
	if o.value == nil {
		o.value = fs.String(o.name, o.def, o.usage)
	}
}

func (o *stringOption) Get() typex.Value {
	if value := o.value; value != nil {
		return typex.StringValue(*value)
	}
	return typex.ZeroValue()
}
