package config

import (
	"flag"
	"io"
	"log"
	"os"
)

//Conf набор флагов
type Conf struct {
	KeyA    int  // "after" печатать +N строк после совпадения
	KeyB    int  // "before" печатать +N строк до совпадения
	KeyC    int  // "context" (A+B) печатать ±N строк вокруг совпадения
	Keyc    bool // "count" (количество строк)
	Keyi    bool // "ignore-case" (игнорировать регистр)
	Keyv    bool // "invert" (вместо совпадения, исключать)
	KeyF    bool // "fixed", точное совпадение со строкой, не паттерн
	Keyn    bool // "line num", напечатать номер строки
	Input   io.Reader
	Request string
}

func prepareFlagsABC(conf *Conf) {
	if conf.KeyC != -1 {
		if conf.KeyA == -1 {
			conf.KeyA = conf.KeyC
		}
		if conf.KeyB == -1 {
			conf.KeyB = conf.KeyC
		}
	} else {
		if conf.KeyA == -1 {
			conf.KeyA = 0
		}
		if conf.KeyB == -1 {
			conf.KeyB = 0
		}
	}
}

//GetConfig парсинг флагов и возвращает конфигурацию флагов
func GetConfig() Conf {
	conf := Conf{}
	flag.IntVar(&conf.KeyA, "A", -1, "\"after\" печатать +N строк после совпадения")
	flag.IntVar(&conf.KeyB, "B", -1, "\"before\" печатать +N строк до совпадения")
	flag.IntVar(&conf.KeyC, "C", -1, "\"context\" (A+B) печатать ±N строк вокруг совпадения")
	flag.BoolVar(&conf.Keyc, "c", false, "\"count\" (количество строк)")
	flag.BoolVar(&conf.Keyi, "i", false, "\"ignore-case\" (игнорировать регистр)")
	flag.BoolVar(&conf.Keyv, "v", false, "\"invert\" (вместо совпадения, исключать)")
	flag.BoolVar(&conf.KeyF, "F", false, "\"fixed\" точное совпадение со строкой, не паттерн")
	flag.BoolVar(&conf.Keyn, "n", false, "\"line num\", напечатать номер строки")
	flag.Parse()

	prepareFlagsABC(&conf)
	var filename string
	conf.Request, filename = findFile()
	conf.Input = GetFile(filename)
	return conf
}

func findFile() (request, filename string) {
	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "-A", "-B", "-C":
			i++
		case "-c", "-i", "-v", "-F", "-n":
		default:
			if request == "" {
				request = os.Args[i]
			} else {
				filename = os.Args[i]
			}
		}
	}
	return request, filename
}

//GetFile определяет откуда будут получены данные
func GetFile(filename string) io.Reader {
	var ok error
	var file *os.File
	if filename != "" {
		if file, ok = os.Open(filename); ok == nil {
			return file
		}
		log.Fatal(ok)
	}
	return os.Stdin
}
