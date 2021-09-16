/* Приложение, которое позволяет скачивать произвольную HTML-страницу посредством HTTP-запроса
на жесткий диск компьютера и выдает статистику по количеству уникальных слов в консоль.
Требования к приложению
1 В качестве входных данных в приложение принимает строку с адресом web-страницы.
Пример входной строки: https://www.simbirsoft.com/
2 Приложение разбивает текст страницы на отдельные слова с помощью списка разделителей.
Пример списка:
{' ', ',', '.', '! ', '?','"', ';', ':', '[', ']', '(', ')', '\n', '\r', '\t'}
3 В качестве результата работы пользователь должен получить статистику по
количеству уникальных слов в тексте.
Пример:
РАЗРАБОТКА -1
ПРОГРАММНОГО - 2
ОБЕСПЕЧЕНИЯ - 4 */

package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
	"testtask/pkg/compare"
	"testtask/pkg/separate"
	"testtask/pkg/storage"
)

//читаем источник в байтах
func MakeRequest() []byte {
	resp, err := http.Get("https://www.simbirsoft.com/")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return body
}

func main() {

	MakeRequest()
	body := string(MakeRequest())
	body = separate.Separator(body)
	body = separate.OnlyLetters(body) // очищаем от всех знаков препинания
	body = strings.ToTitle(body)      //переводим всё в верхний регистр

	storage.Words = storage.MapFromStr(body)
	compare.M = make(compare.ResultVal)
	compare.M = compare.Compare(storage.Words, storage.S)

	sort.Strings(storage.Ws) // сортировка слайса для вывода в терминале

	compare.M.Get(storage.Ws) // вызов функции вывода содержимого результирующей карты
}
