package alignx

import "unicode"

// Align 对齐字符串
func Align(s string, length int) string {
	switch l := len(s); {
	case length > 0 && length > l:
		return s + Spaces(length-l)
	case length < 0 && -length > l:
		return Spaces(-length-l) + s
	default:
		return s
	}
}

// VisualLength 计算字符串的可视化长度（1个中文占5/3个字符宽度）
func VisualLength(s string) int {
	var runes, cn = []rune(s), 0
	for _, r := range runes {
		if unicode.Is(unicode.Han, r) {
			cn++
		}
	}
	length := len(runes) + cn*2/3
	if y := cn % 3; y == 1 {
		length++
	}
	return length
}

// Spaces 生成指定长度的空格字符串
func Spaces(length int) string {
	var runes []rune
	for i := 0; i < length; i++ {
		runes = append(runes, 32) // 32为空格
	}
	return string(runes)
}

// MaxLength 计算字符串数组中最大的可视化长度
func MaxLength(ss []string) int {
	var length int
	for _, s := range ss {
		if l := VisualLength(s); l > length {
			length = l
		}
	}
	return length
}
