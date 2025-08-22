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

// Dyeing 染色
func Dyeing(c Color, v any) string {
	if black <= c && c <= grey {
		return fmt.Sprintf("\x1b[%dm%v\x1b[0m", uint8(c), v)
	}
	return fmt.Sprint(v)
}

func Black(v any) string {
	return Dyeing(black, v)
}

func Red(v any) string {
	return Dyeing(red, v)
}

func Green(v any) string {
	return Dyeing(green, v)
}

func Yellow(v any) string {
	return Dyeing(yellow, v)
}

func Blue(v any) string {
	return Dyeing(blue, v)
}

func Magenta(v any) string {
	return Dyeing(magenta, v)
}

func Cyan(v any) string {
	return Dyeing(cyan, v)
}

func Grey(v any) string {
	return Dyeing(grey, v)
}
