package alignx

import "unicode"

// Left 左对齐
func Left(s string, length int) string {
	fill := Spaces(length - VisualLength(s))
	return string(append([]rune(s), fill...))

}

// Right 右对齐
func Right(s string, length int) string {
	fill := Spaces(length - VisualLength(s))
	return string(append(fill, []rune(s)...))
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
func Spaces(length int) []rune {
	var runes []rune
	for i := 0; i < length; i++ {
		runes = append(runes, 32) // 32为空格
	}
	return runes
}

// MaxLength 计算字符串数组中最大的可视化长度
func MaxLength(ss ...string) int {
	var length int
	for _, s := range ss {
		if l := VisualLength(s); l > length {
			length = l
		}
	}
	return length
}
