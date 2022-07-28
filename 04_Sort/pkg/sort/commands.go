package sort

import (
	"04_Sort/pkg/utils"
	"reflect"
	"strconv"
	"unicode"
)

/*
	Реализовать поддержку утилитой следующих ключей:

	*-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
	*-n — сортировать по числовому значению
	*-r — сортировать в обратном порядке
	*-u — не выводить повторяющиеся строки

	Дополнительно

	Реализовать поддержку утилитой следующих ключей:

	-M — сортировать по названию месяца
	-b — игнорировать хвостовые пробелы
	-c — проверять отсортированы ли данные
	-h — сортировать по числовому значению с учетом суффиксов
*/

//CommandMaker структура с единственной функцией задача которой описано ниже
type CommandMaker struct{}

//GetCommands определяет поведение сортировки на основе флажков
func (CommandMaker) GetCommands(sortStruct *Sort) (IComparator, []ICommand) {
	var comparators IComparator
	command := make([]ICommand, 0)

	if ok, _ := (*sortStruct.keys)['n']; ok {
		comparators = NumValueSort{}
	} else {
		comparators = DefaultSort{}
	}
	if ok, _ := (*sortStruct.keys)['u']; ok {
		command = append(command, Unique{})
	}
	if ok, _ := (*sortStruct.keys)['r']; ok {
		command = append(command, Reverse{})
	}
	return comparators, command

}

// NumValueSort -n — сортировать по числовому значению
type NumValueSort struct{}

// Compare -n — сортировать по числовому значению
func (NumValueSort) Compare(a, b []rune, k int, delim rune) bool {
	sa, sb := string(a), string(b)
	ai := utils.GetIndex(a, delim, k)
	bi := utils.GetIndex(b, delim, k)
	if ai == -1 || bi == -1 {
		return ai < 0
	}
	if ai >= len(a) || bi >= len(b) {
		return bi >= len(b)
	}
	ssa := string(a[ai])
	ssb := string(b[bi])
	_, _, _, _ = sa, sb, ssa, ssb
	parseInt := func(a []rune) int {
		l := 0
		for l < len(a) && unicode.IsDigit(a[l]) {
			l++
		}
		l64, ok := strconv.ParseInt(string(a[:l]), 10, 32)
		if ok != nil || l == len(a) {
			l64 = -1
		}
		return int(l64)
	}
	massa, massb := parseInt(a[ai:]), parseInt(b[bi:])
	if massa != massb {
		if massa == 0 || massb == 0 {
			return massa != 0
		}
		return massa > massb
	}
	return !utils.StringComparator(a[ai:], b[bi:])
}

// DefaultSort — сортировать по умолчанию
type DefaultSort struct{}

// Compare — сортировать по умолчанию
func (DefaultSort) Compare(a, b []rune, k int, delim rune) bool {
	strA, strB := string(a), string(b)
	_, _ = strA, strB
	ai := utils.GetIndex(a, delim, k)
	bi := utils.GetIndex(b, delim, k)
	lena, lenb := len(a), len(b)
	if lena == 0 || lenb == 0 {
		return lenb == 0
	}
	return utils.StringComparator(a[ai:], b[bi:])
}

// Reverse -r — сортировать в обратном порядке
type Reverse struct{}

// Exec -r — сортировать в обратном порядке
func (Reverse) Exec(data []string) []string {
	swapper := reflect.Swapper(data)
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		swapper(i, j)
	}
	return data
}

// Unique -u — не выводить повторяющиеся строки
type Unique struct{}

// Exec -u — не выводить повторяющиеся строки
func (Unique) Exec(data []string) []string {
	tmp := make(map[string]uint32, len(data))
	for _, v := range data {
		tmp[v] = tmp[v] + 1
	}
	newData := make([]string, 0, len(tmp))
	for _, v := range data {
		tmp[v]--
		if tmp[v] == 0 {
			newData = append(newData, v)
		}
	}
	return newData
}
