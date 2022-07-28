package io

import (
	"grep/pkg/grep/config"
	"io/ioutil"
	"log"
	"strings"
)

//GetData чтение данных
func GetData(conf config.Conf) []string {
	var (
		raw []byte
		ok  error
	)
	if raw, ok = ioutil.ReadAll(conf.Input); ok != nil {
		log.Fatal(ok)
	}
	return strings.Split(string(raw), "\n")
}
