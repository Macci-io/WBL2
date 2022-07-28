package sort

import (
	bytes2 "bytes"
	"reflect"
	"strings"
)

/*
	Реализовать поддержку утилитой следующих ключей:

	-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
	-n — сортировать по числовому значению
	-r — сортировать в обратном порядке
	-u — не выводить повторяющиеся строки

	Дополнительно

	Реализовать поддержку утилитой следующих ключей:

	-M — сортировать по названию месяца
	-b — игнорировать хвостовые пробелы
	-c — проверять отсортированы ли данные
	-h — сортировать по числовому значению с учетом суффиксов
*/

//Sort главная структура
type Sort struct {
	//Сырые данные
	data [][]rune
	//Флажки
	keys *map[byte]bool
	//Номер колонки
	k int
}

//IComparator интерфейс с методом сортировки
type IComparator interface {
	Compare(a, b []rune, k int, delim rune) bool
}

//ICommand интерфейс с каким то действием над данными
type ICommand interface {
	Exec([]string) []string
}

//NewSortUtil конструктор
func NewSortUtil(bytes []byte, keys *map[byte]bool, k int) *Sort {
	bytes = bytes2.Trim(bytes, "\n")
	tmp := strings.Split(string(bytes), "\n")
	runes := make([][]rune, len(tmp))
	for i, v := range tmp {
		runes[i] = []rune(v)
	}
	return &Sort{runes, keys, k}
}

//sort сортировки по заданному алгоритму
func (s *Sort) sort(cmp IComparator) *Sort {
	swapper := reflect.Swapper(s.data)
	for i := 0; i < len(s.data); i++ {
		for j := i + 1; j < len(s.data); j++ {
			if cmp.Compare(s.data[i], s.data[j], s.k, ' ') {
				swapper(i, j)
			}
		}
	}
	return s
}

func (s *Sort) toStrings() (result []string) {
	result = make([]string, len(s.data))
	for i, v := range s.data {
		result[i] = string(v)
	}
	return result
}

//Run Точка входа, запуск и выполнение действий
func (s *Sort) Run() string {
	command, iCommands := CommandMaker{}.GetCommands(s)
	toStrings := s.sort(command).toStrings()
	for _, v := range iCommands {
		toStrings = v.Exec(toStrings)
	}
	return strings.Join(toStrings, "\n")
}
