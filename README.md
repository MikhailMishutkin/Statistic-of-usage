# Simbir_test_task.git
test task 09/21
<!-- /* Задача
 Приложение, которое позволяет скачивать произвольную HTML-страницу посредством HTTP-запроса
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
ОБЕСПЕЧЕНИЯ - 4 */ -->

Логика программы:
1. Считываем источник (начинается с https://) pkg/request
2. С помощью разделителей извлекаем только слова из русских букв длинной от 3-х символов в слайс строк. pkg/separate
3. Результат п.2 складываем в промежуточную мапу - для удаления Не уникальных значений ("String" - "Bool"). pkg/storage
4. Слайс строк (п.2) сравниваем во вложенном цикле с промежуточной мапой (п. 3): при совпадении значения слайса с уникальным значением из мапы увеличиваем счётчик количества уникальных слов. pkg/compare
5. Уникальное значение слова и счётчик складываем в результирующую мапу вида: "String" - "Int". pkg/compare
6. Для упорядоченного вывода содержимого результирующей мапы (п.5) используем разложение мапы в слайс строк. pkg/compare