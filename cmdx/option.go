package cmdx

import (
	"flag"
	"fmt"

	"github.com/go-xuan/typex"
)

type Option interface {
	GetName() string        // 获取选项名
	GetUsage() string       // 获取选项用法
	SetFS(fs *flag.FlagSet) // 设置选项值
	GetValue() typex.Value  // 获取选项值
}

// baseOption 基础选项
type baseOption struct {
	name  string // 选项名
	usage string // 选项用法
}

// GetName 获取选项名
func (o *baseOption) GetName() string {
	return o.name
}

// 生成选项用法
func (o *baseOption) getUsage(def interface{}) string {
	return fmt.Sprintf("%s | default: %v", o.usage, def)
}
