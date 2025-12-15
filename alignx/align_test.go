package alignx

import "testing"

func TestAlign(t *testing.T) {
	var texts = []string{
		"一二三四五六七八九",
		"123456789",
	}

	for _, text := range texts {
		println(Right(text, 15), "|")
	}
	for _, text := range texts {
		println(Left(text, 15), "|")
	}
}
