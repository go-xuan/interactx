package flagx

import (
	"flag"
	"fmt"

	"github.com/go-xuan/typex"
)

type Option interface {
	Name() string
	Usage() string
	Set(fs *flag.FlagSet)
	Get() typex.Value
}

type OptionHandler func(*baseOption)

// baseOption 基础选项
type baseOption struct {
	name  string // 选项名
	usage string // 选项用法
}

func (opt *baseOption) Name() string {
	return opt.name
}

// 通用的 Usage 方法生成逻辑
func genUsage(usage string, def interface{}) string {
	return fmt.Sprintf("%s | default: %v", usage, def)
}
