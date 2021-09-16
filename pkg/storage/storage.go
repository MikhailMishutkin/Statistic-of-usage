package storage

import "strings"

var (
	Word  string          //переменная для отдельных слов
	Words map[string]bool //промежуточная мапа для сравнения со слайсом строк
	Ws    []string        //переменная для сортировки финального вывода
	S     []string        //промежуточная переменная для хранения слайса из слов тела сайта
)

func MapFromStr(s string) map[string]bool {

	S = strings.Fields(s) // делаем слайс из строки для

	Words := make(map[string]bool) //инициализация мапы

	//цикл для складывания элементов слайса в промежуточную мапу - убираются дублирующие значения
	for _, v := range S {
		Word = v
		if len(Word) > 4 { // убираем предлоги меньше 3 русских букв
			// складываем значения ключей мапы в слайс строк для сортированного вывода
			if _, ok := Words[Word]; !ok {
				Ws = append(Ws, Word)
			}
			Words[Word] = true
		}
	}
	return Words
}
