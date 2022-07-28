package main

import (
	"03_Unpack/unpack"
	"fmt"
)

/*
	Задача на распаковку
	Создать Go-функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны, например:
	"a4bc2d5e" => "aaaabccddddde"
	"abcd" => "abcd"
	"45" => "" (некорректная строка)
	"" => ""
*/

func main() {
	input := ""
	fmt.Println("Enter the word")
	for {
		_, _ = fmt.Scan(&input)
		fmt.Println(input + " => " + unpack.Unpack(&input))
	}

}
