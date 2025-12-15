package colorx

import "fmt"

type Color uint8

const (
	BLACK   Color = iota + 30 // 黑色
	RED                       // 红色
	GREEN                     // 绿色
	YELLOW                    // 黄色
	BLUE                      // 蓝色
	MAGENTA                   // 洋红色
	CYAN                      // 青色
	GREY                      // 灰色
)

// Dyeing 染色
func Dyeing(v any, c Color) string {
	if BLACK <= c && c <= GREY {
		return fmt.Sprintf("\x1b[%dm%v\x1b[0m", uint8(c), v)
	}
	return fmt.Sprint(v)
}

func Black(v any) string {
	return Dyeing(v, BLACK)
}

func Red(v any) string {
	return Dyeing(v, RED)
}

func Green(v any) string {
	return Dyeing(v, GREEN)
}

func Yellow(v any) string {
	return Dyeing(v, YELLOW)
}

func Blue(v any) string {
	return Dyeing(v, BLUE)
}

func Magenta(v any) string {
	return Dyeing(v, MAGENTA)
}

func Cyan(v any) string {
	return Dyeing(v, CYAN)
}

func Grey(v any) string {
	return Dyeing(v, GREY)
}
