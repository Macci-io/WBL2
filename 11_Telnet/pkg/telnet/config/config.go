package config

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Config структура для хранения конфигурации
type Config struct {
	timeOut    time.Duration
	host, port *string
}

// SetHost сеттер для установления хоста
func (c *Config) SetHost(host *string) {
	c.host = host
}

// SetPort сеттер для установления порта
func (c *Config) SetPort(port *string) {
	c.port = port
}

// GetConnectionInfo формирование адресной строки для соединения
func (c *Config) GetConnectionInfo() string {
	return *c.host + ":" + *c.port
}

//GetTimeOut геттер для таймаута
func (c *Config) GetTimeOut() time.Duration {
	return c.timeOut
}

// NewConfig создание нового конфига на основе аргументов
func NewConfig() (conf *Config, ok error) {
	var host, port string
	tm := time.Second * 10

	for i := 1; i < len(os.Args); i++ {
		if strings.HasPrefix(os.Args[i], "--timeout=") {
			tm, ok = parseTime(&os.Args[i])
			if ok != nil {
				return nil, ok
			}
		} else if len(host) == 0 {
			host = os.Args[i]
		} else if len(port) == 0 {
			port = os.Args[i]
		} else {
			break
		}
	}
	return &Config{tm, &host, &port}, nil
}

func parseTime(str *string) (tm time.Duration, ok error) {
	var numIndex []int
	var tmTypeIndex []int
	var num int
	numReg, _ := regexp.Compile("\\d+")
	tmReg, _ := regexp.Compile("(ms|m|s)$")
	if numIndex = numReg.FindIndex([]byte(*str)); numIndex == nil {
		return 0, errors.New("wrong argument near --timeout=")
	}
	if num, ok = strconv.Atoi((*str)[numIndex[0]:numIndex[1]]); ok != nil {
		fmt.Println(ok)
	}
	if tmTypeIndex = tmReg.FindIndex([]byte((*str))); tmTypeIndex == nil {
		tm = time.Second * time.Duration(num)
	} else {
		switch (*str)[tmTypeIndex[0]:tmTypeIndex[1]] {
		case "ms":
			tm = time.Millisecond * time.Duration(num)
		case "m":
			tm = time.Minute * time.Duration(num)
		default:
			tm = time.Second * time.Duration(num)
		}
	}
	return tm, nil
}
