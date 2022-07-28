package utils

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode"
)

//kValidate валидация флага (k)
func kValidate(k int) {
	if k <= 0 {
		if k == 0 {
			_, _ = fmt.Fprintln(os.Stderr, "sort: field number is zero: invalid field specification ‘0’")
		} else {
			_, _ = fmt.Fprintf(os.Stderr, "sort: invalid number at field start: invalid count at start of ‘%d’\n", k)
		}
		os.Exit(0)
	}
}

//ParseFlags парсинг флажков
func ParseFlags() (*map[byte]bool, int) {
	var n, r, u bool
	var k int
	flag.IntVar(&k, "k", 1, "указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)")
	flag.BoolVar(&n, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&r, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&u, "u", false, "не выводить повторяющиеся строки")
	flag.Parse()
	kValidate(k)
	return &map[byte]bool{'n': n, 'r': r, 'u': u}, k - 1
}

//XOR исключающее "или"
func XOR(a, b bool) bool {
	return (a && !b) || (!a && b)
}

//StringComparator сравнение по символьно
func StringComparator(a, b []rune) bool {
	al, bl := len(a), len(b)
	ai, bi := 0, 0

	for a[ai] == '-' && ai < al-1 {
		ai++
	}
	for b[bi] == '-' && bi < bl-1 {
		bi++
	}

	for ; ai < al && bi < bl; ai, bi = ai+1, bi+1 {
		if a[ai] == b[bi] {
			continue
		}
		for (unicode.IsSpace(a[ai]) || a[ai] == '-') && ai < al-1 {
			ai++
		}
		for (unicode.IsSpace(b[bi]) || b[bi] == '-') && bi < bl-1 {
			bi++
		}
		if XOR(ai == al, bi == bl) {
			return ai == al
		}
		if unicode.IsNumber(a[ai]) && unicode.IsNumber(b[bi]) {
			return a[ai] > b[bi]
		}
		return a[ai] > b[bi]
	}
	return al > bl
}

//GetIndex получение индекса символа по номеру колонки с учетом разделителя
func GetIndex(data []rune, delim rune, column int) (index int) {
	var last = 'k'
	for i, v := range data {
		if v == delim {
			if v != last {
				column--
			}
		} else if column == 0 {
			return i
		}
		last = v
	}
	return strings.LastIndexFunc(string(data), func(r rune) bool { return r != delim })
}
