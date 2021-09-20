package separate

import (
	"log"
	"regexp"
	"strings"
)

// Пакет последовательно возвращает из сформированной из источника
// строки сначала русские слова с разделителями, затем без них

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
	s1 := strings.TrimSpace(s) // проверка на наличие русских символов
	if len(s1) < 2 {
		log.Fatalln("в источнике нет русских символов")
	}
	return s

}
