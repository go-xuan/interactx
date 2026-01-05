package alignx

import (
	"fmt"
	"testing"
)

func TestAlign(t *testing.T) {
	var texts = []string{
		"一二三四五六七八九",
		"123456789",
	}

	for _, text := range texts {
		println(Align(text, 20), "|")
	}

	for _, text := range texts {
		println(Align(text, -20), "|")
	}

	for _, text := range texts {
		println(Align(text, 5), "|")
	}

	for _, text := range texts {
		println(Align(text, -5), "|")
	}

	for _, text := range texts {
		println(Align(text, 0), "|")
	}
}

func TestExtract(t *testing.T) {
	fmt.Println(Extract("1234567", 6))
	fmt.Println(Extract("123456", 6))
	fmt.Println(Extract("12345", 6))

	fmt.Println(Extract("1234567", -6))
	fmt.Println(Extract("123456", -6))
	fmt.Println(Extract("12345", -6))

	fmt.Println(Extract("1234567", -0))
	fmt.Println(Extract("123456", 0))
	fmt.Println(Extract("12345", 0))
}
