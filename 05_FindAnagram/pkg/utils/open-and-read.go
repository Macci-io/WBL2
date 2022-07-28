package utils

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//ReadFromFile читает файл и возвращает результат в виде строчного слйса
func ReadFromFile(filePath string) []string {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	split := strings.Split(string(file), "\n")
	for i := range split {
		split[i] = strings.TrimRight(split[i], "\r")
	}
	return split
}
