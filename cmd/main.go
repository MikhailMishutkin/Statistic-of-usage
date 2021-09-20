package main

import (
	"fmt"
	"sort"
	"strings"
	"testtask/pkg/compare"
	"testtask/pkg/request"
	"testtask/pkg/separate"
	"testtask/pkg/storage"
)

func main() {
	var s string
	fmt.Println("Получение статитистики по количеству уникальных русских слов")
	fmt.Println("----------------------------------------------------")
	fmt.Println("Введите(скопируйте полностью) URL источника, по которому необходимо получить статистику:")
	fmt.Scan(&s)
	body := string(request.MakeRequest(s))
	body = separate.Separator(body)
	body = separate.OnlyLetters(body) // очищаем от всех знаков препинания
	body = strings.ToTitle(body)      //переводим всё в верхний регистр

	storage.Words = storage.MapFromStr(body)
	compare.M = make(compare.ResultVal)
	compare.M = compare.Compare(storage.Words, storage.S)

	sort.Strings(storage.Ws) // сортировка слайса для вывода в терминале

	compare.M.Get(storage.Ws) // вызов функции вывода содержимого результирующей карты
}
