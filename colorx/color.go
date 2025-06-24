package colorx

import "fmt"

type Color uint8

const (
	black   Color = iota + 30 // 黑色
	red                       // 红色
	green                     // 绿色
	yellow                    // 黄色
	blue                      // 蓝色
	magenta                   // 洋红色
	cyan                      // 青色
	grey                      // 灰色
)

// Dye 染色
func Dye(c Color, v any) string {
	if black <= c && c <= grey {
		return fmt.Sprintf("\x1b[%dm%v\x1b[0m", uint8(c), v)
	}
	return fmt.Sprint(v)
}

func Black(v any) string {
	return Dye(black, v)
}

func Red(v any) string {
	return Dye(red, v)
}

func Green(v any) string {
	return Dye(green, v)
}

func Yellow(v any) string {
	return Dye(yellow, v)
}

func Blue(v any) string {
	return Dye(blue, v)
}

func Magenta(v any) string {
	return Dye(magenta, v)
}

func Cyan(v any) string {
	return Dye(cyan, v)
}

func Grey(v any) string {
	return Dye(grey, v)
}
