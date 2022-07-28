package main

import (
	"cut/pkg/cut"
	"cut/pkg/cut/parse"
	"fmt"
)

/*
	Реализовать утилиту аналог консольной команды cut (man cut).
	Утилита должна принимать строки через STDIN,
	разбивать по разделителю (TAB) на колонки и выводить запрошенные.

	Реализовать поддержку утилитой следующих ключей:
	-f - "fields" - выбрать поля (колонки)
	-d - "delimiter" - использовать другой разделитель
	-s - "separated" - только строки с разделителем
*/

func main() {
	config := parse.NewConfig()
	cut := cut.NewCut(config, nil)
	fmt.Println(cut.GetResult())
}
