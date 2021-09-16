package separate

import (
	"regexp"
)

// убирает всё кроме слов, слитных с разделителями
func Separator(s string) string {
	rx := regexp.MustCompile(`[^А-Яа-я ,.!?";:[]()\n\r\t]+`)
	s = rx.ReplaceAllString(s, "")
	return s
}

// функция для выделения только русских букв, без разделителей
func OnlyLetters(s string) string {
	rx := regexp.MustCompile(`[^А-Яа-я ]+`)
	s = rx.ReplaceAllString(s, " ")
	return s
}
