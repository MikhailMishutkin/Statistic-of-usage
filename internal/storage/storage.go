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
	Word      string          //переменная для отдельных слов
	Words     map[string]bool //промежуточная мапа для сравнения со слайсом строк
	Ws        []string        //переменная для сортировки финального вывода
	S         []string        //промежуточная переменная для хранения слайса из слов тела сайта
	ResultVal map[string]int  //конечная мапа
)

type Repository struct {
	uc UseCase
}

type UseCase interface {
	AskUser()
	Separate(sl string) (s string)
	ToCompare()
	ResToSrcn()
}

func NewStorage(s Repository) *Repository {
	Words = make(map[string]bool)
	ResultVal = make(map[string]int)
	return &Repository{uc: s.uc}
}

func (r *Repository) МapFromStr(s string) {

	S = strings.Fields(s) // делаем слайс из входящей строки для чтения в мапу

	//цикл для складывания элементов слайса в промежуточную мапу - убираются дублирующие значения
	for _, v := range S {
		Word = v
		if len(Word) > 4 { // убираем предлоги и подобное, что меньше 3 русских букв
			// складываем значения ключей мапы в слайс строк для сортированного вывода
			if _, ok := Words[Word]; !ok {
				Ws = append(Ws, Word)
			}
			Words[Word] = true
		}
	}
}

// вложенный цикл для сравнения промежуточной мапы и слайса строк с подсчётом повторений
func (r *Repository) Compare() map[string]int {

	ResultVal = make(map[string]int)
	for k := range Words {
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
	return ResultVal
}

//метод put складывает результат выборки в мапу
func (r *Repository) Put(s string, q int) {
	ResultVal[s] = q
}

//метод для вызова содержимого мапы с использованием сортированного слайса
func (r *Repository) Get() {
	sort.Strings(Ws)
	sl := Ws
	for _, v := range sl {
		quant := ResultVal[v]
		fmt.Println(v, quant)
	}
}
