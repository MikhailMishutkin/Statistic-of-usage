package usecase

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

type Usecase struct {
	repo Repo
	c    Controller
}

type Repo interface {
	МapFromStr(s string)
	Compare() map[string]int
	Put(s string, q int)
	Get()
}

type Controller interface {
	UserAnsw() (s string)
}

func New(r Repo) *Usecase {
	return &Usecase{repo: r}
}

//запрос к пользователю
func (uc *Usecase) AskUser() string {
	var s string
	fmt.Scan(&s)
	return s
}

// сначала убирает всё кроме слов, слитных с разделителями, потом выделяет только русские буквы, без разделителей
func (uc *Usecase) Separate(s string) {
	rx := regexp.MustCompile(`[^А-Яа-я ,.!?";:[]()\n\r\t]+`)
	s = rx.ReplaceAllString(s, "")
	rx = regexp.MustCompile(`[^А-Яа-я ]+`)
	s = rx.ReplaceAllString(s, " ")
	s1 := strings.TrimSpace(s) // проверка на наличие русских символов
	if len(s1) < 2 {
		log.Fatalln("в источнике нет русских символов")
	}
	s = strings.ToTitle(s)
	uc.repo.МapFromStr(s)

}

//сравнение хранилищ и подсчёт совпадений
func (uc *Usecase) ToCompare() {
	uc.repo.Compare()
}

//вывод результатов работы программы на экран
func (uc *Usecase) ResToSrcn() {
	uc.repo.Get()
}
