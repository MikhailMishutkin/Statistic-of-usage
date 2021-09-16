package compare

import (
	"fmt"
	"testtask/pkg/storage"
)

// Пакет выполняет сравнение мапы из хранилища (с уникальными значениями)
// с очищенным от разделителей слайсом строк и считает совпадения. Также содержит методы
// для записи в результирующую мапу и извлечения из мапы в промежуточный слайс строк для
// сортированного вывода результата работы программы

type ResultVal map[string]int //задаём тип для результата работы программы

var M ResultVal //конечная мапа

// вложенный цикл для сравнения промежуточной мапы и слайса строк с подсчётом повторений
func Compare(m map[string]bool, s []string) ResultVal {
	for k := range storage.Words {
		coincidence := 0
		for i := 0; i < len(storage.S); i++ {
			if k == storage.S[i] {
				coincidence++
				M.Put(k, coincidence)
			} else {
				continue
			}
		}
	}
	return M
}

//метод put складывает результат выборки в мапу
func (m ResultVal) Put(s string, q int) ResultVal {
	m[s] = q
	return m
}

//метод для вызова содержимого мапы с использованием сортированного слайса
func (m ResultVal) Get(sl []string) {
	for _, v := range sl {
		quant := m[v]
		fmt.Println(v, quant)
	}
}
