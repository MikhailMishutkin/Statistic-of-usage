// Пакет служит цели инициализации и наполнения переменных, в которых будут
// храниться уникальные слова, а также вспомогательных переменных для возможности
//  сортированного вывода результата работы программы

package storage

import (
	"fmt"
	"sort"
	"strings"
)

var (
	Word string //переменная для отдельных слов

	Ws []string //переменная для сортировки финального вывода
	S  []string //промежуточная переменная для хранения слайса из слов тела сайта

)

type Words map[string]bool    //промежуточная мапа для сравнения со слайсом строк
type ResultVal map[string]int //конечная мапа

type Repository struct {
	W Words
	R ResultVal
}

func NewStorage(r Repository) *Repository {
	r.W = make(map[string]bool)
	r.R = make(map[string]int)
	return &Repository{W: r.W, R: r.R}
}

func (r *Repository) МapFromStr(s string) {

	S = strings.Fields(s) // делаем слайс из входящей строки для чтения в мапу

	//цикл для складывания элементов слайса в промежуточную мапу - убираются дублирующие значения
	for _, v := range S {
		Word = v
		if len(Word) > 4 { // убираем предлоги и подобное, что меньше 3 русских букв
			// складываем значения ключей мапы в слайс строк для сортированного вывода
			if _, ok := r.W[Word]; !ok {
				Ws = append(Ws, Word)
			}
			r.W[Word] = true
		}
	}
}

// вложенный цикл для сравнения промежуточной мапы и слайса строк с подсчётом повторений
func (r *Repository) Compare() map[string]int {

	r.R = make(map[string]int)
	for k := range r.W {
		coincidence := 0
		for i := 0; i < len(S); i++ {
			if k == S[i] {
				coincidence++
				r.Put(k, coincidence)
			} else {
				continue
			}
		}
	}
	return r.R
}

//метод put складывает результат выборки в мапу
func (r *Repository) Put(s string, q int) {
	r.R[s] = q
}

//метод для вызова содержимого мапы с использованием сортированного слайса
func (r *Repository) Get() {
	sort.Strings(Ws)
	sl := Ws
	for _, v := range sl {
		quant := r.R[v]
		fmt.Println(v, quant)
	}
}
