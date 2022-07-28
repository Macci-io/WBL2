package parse

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
	Реализовать утилиту аналог консольной команды cut (man cut). Утилита должна принимать строки через STDIN, разбивать по разделителю (TAB) на колонки и выводить запрошенные.

	Реализовать поддержку утилитой следующих ключей:
	-f - "fields" - выбрать поля (колонки)
	-d - "delimiter" - использовать другой разделитель
	-s - "separated" - только строки с разделителем
*/

//Config хранение данных о флагах
type Config struct {
	F    [][2]int
	D    byte
	S    bool
	Read io.Reader
}

// NewConfig создание новой структуры с получением данных из аргументов
func NewConfig() *Config {
	var conf = new(Config)
	conf.D = '\t'
	conf.Read = os.Stdin
	length := len(os.Args)
	for i := 1; i < length; i++ {
		if os.Args[i][0] == '-' && len(os.Args[i]) == 2 {
			switch os.Args[i][1] {
			case 'f':
				if i < length-1 {
					i++
					conf.F = parseF(os.Args[i])
				}
			case 'd':
				if i < length-1 {
					i++
					conf.D = os.Args[i][0]
				}
			case 's':
				conf.S = true
			}
		} else {
			conf.Read = getFile(os.Args[i])
		}
	}
	if conf.F == nil {
		fmt.Println("cut: you must specify a list of bytes, characters, or fields")
		os.Exit(0)
	}
	return postScriptF(conf)
}

func sortF(f [][2]int) {
	sort.Slice(f, func(i, j int) bool {
		if f[i][0] == f[j][0] {
			if f[i][1] == TIRE {
				return false
			} else if f[i][1] == NOTHING {
				return true
			}
			return f[i][1] < f[j][1]
		}
		if f[i][0] == TIRE {
			return true
		}
		return f[i][0] < f[j][0]
	})
}

//NOTHING если не было указано второе число
const NOTHING = -1

//TIRE если до или после '-' не было числа
const TIRE = 0

// isBetween находится ли число в пределах диапазона
func isBetween(pair [2]int, el int) bool {
	if (pair[0] == TIRE || pair[0] <= el) && el != TIRE {
		if (pair[1] == TIRE || pair[1] >= el) && pair[1] != NOTHING {
			return true
		}
	}
	return false
}

// xor исключающий (или нет)
func xor(a, b bool) bool {
	return (a && !b) || (!a && b)
}

// isContains находится ли второй диапазон внутри первого диапазона
func isContains(a, b [2]int) bool {
	return isBetween(a, b[0]) && isBetween(a, b[1])
}

// isTouched касается ли один диапазон другого
func isTouched(a, b [2]int) bool {
	cmp1 := isBetween(a, b[0])
	cmp2 := isBetween(a, b[1])
	return (cmp1 || cmp2) && xor(cmp1, cmp2)
}

// postScriptF убрать дубликаты, срезы внутри среза и объединить смежные диапазоны
func postScriptF(config *Config) *Config {
	f := config.F
	var res [][2]int
	sortF(f)
	tmp := f[0]
	for _, v := range f[1:] {
		// 1-5 2-7 || 1-5 2- ... 1-5,2-4
		if isTouched(tmp, v) {
			tmp[1] = v[1]
			// 1-5 2-2
		} else if !isContains(tmp, v) {
			res = append(res, tmp)
			tmp = v
		}
	}
	config.F = append(res, tmp)
	return config
}

// getFile Читаем данные, (с STDIN или файл, смотря что будет в интерфейсе)
func getFile(filePath string) io.Reader {
	open, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return open
}

func atoi(s string) (res int) {
	var err error
	if res, err = strconv.Atoi(s); err != nil {
		log.Fatal(err)
	}
	return res
}

// parseF распарсить и отсортировать это
func parseF(data string) (resultF [][2]int) {
	var pair []string
	rows := strings.Split(data, ",")
	for _, v := range rows {
		// v = 1-3, -3, 4-, 5
		if len(v) == 0 {
			log.Fatal("cut: fields are numbered from 1")
		}
		if !strings.ContainsRune(v, '-') {
			resultF = append(resultF, [2]int{atoi(v), NOTHING})
			continue
		}
		if pair = strings.Split(v, "-"); len(pair) == 0 {
			log.Fatal("cut: invalid range with no endpoint: -")
		}
		// pair [0]2 - [1]9
		if v[0] == '-' {
			resultF = append(resultF, [2]int{TIRE, atoi(pair[1])})
		} else {
			if pair[len(pair)-1] != "" {
				resultF = append(resultF, [2]int{atoi(pair[0]), atoi(pair[1])})
			} else {
				resultF = append(resultF, [2]int{atoi(pair[0]), TIRE})
			}
		}
	}
	return resultF
}
