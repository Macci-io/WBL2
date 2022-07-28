package main

import (
	"fmt"
	"reflect"
	"strings"
)

/*
Краткое описание:
	Стратегия — это поведенческий паттерн проектирования, который определяет семейство схожих алгоритмов и помещает
	каждый из них в собственный класс, после чего алгоритмы можно взаимозаменять прямо во время исполнения программы.
Плюсы:
	1. Горячая замена алгоритмов на лету.
	2. Изолирует код и данные алгоритмов от остальных классов.
	3. Уход от наследования к делегированию.
	4. Реализует принцип открытости/закрытости.
Минусы:
	1. Усложняет программу за счёт дополнительных классов.
	2. Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.
*/

func main() {
	data := []string{"Стратегия", "—", "это", "поведенческий", "паттерн", "проектирования,", "который", "определяет", "семейство", "схожих", "алгоритмов", "и", "помещает"}
	ApplyingStrategy(data, Upper{})
	ApplyingStrategy(data, SortLength{})
	ApplyingStrategy(data, СaesarСipher{-12})
	ApplyingStrategy(data, SortLexicographic{})
}

// ApplyingStrategy некая функция реализующая паттерн "стратегия"
func ApplyingStrategy(s []string, operation IOperation) {
	copyData := make([]string, len(s))
	copy(copyData, s)
	fmt.Printf("%-23s", fmt.Sprintf("%v", reflect.TypeOf(operation)))
	fmt.Print("Len:", len(copyData), " ")
	operation.Processing(copyData)
	fmt.Println(copyData)
}

// IOperation интерфейс для алгоритмов
type IOperation interface{ Processing([]string) }

// Upper переводит слайс стрингов в верхний регистр
type Upper struct{}

//SortLength сортирует по длине слова
type SortLength struct{}

//СaesarСipher применяет шифр цезаря
type СaesarСipher struct{ level int }

//SortLexicographic Сортирует лексикографически
type SortLexicographic struct{}

//Processing перевести все буквы в верхний регистр
func (Upper) Processing(s []string) {
	size := len(s)
	for i := 0; i < size; i++ {
		s[i] = strings.ToUpper(s[i])
	}
}

//Processing сортирует слова по длине слова
func (SortLength) Processing(s []string) {
	size, swapper := len(s), reflect.Swapper(s)

	for i := 0; i < size; i++ {
		for x := i + 1; x < size; x++ {
			if len(s[i]) > len(s[x]) {
				swapper(x, i)
			}
		}
	}
}

//Processing обычная лексикографическая сортировка
func (SortLexicographic) Processing(s []string) {
	size, swapper := len(s), reflect.Swapper(s)
	compare := func(a, b string) bool {
		sa, sb := len(a), len(b)
		for i := 0; i < sa && i < sb; i++ {
			if a[i] == b[i] {
				continue
			}
			return a[i] < b[i]
		}
		return sa < sb
	}
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if compare(s[i], s[j]) {
				swapper(i, j)
			}
		}
	}
}

//Processing применить шифр цезаря
func (c СaesarСipher) Processing(s []string) {
	size := len(s)
	up := func(s string) string {
		res := make([]rune, len(s))
		for i, v := range s {
			res[i] = v + rune(c.level)
		}
		return string(res)
	}
	for i := 0; i < size; i++ {
		s[i] = up(s[i])
	}
}
