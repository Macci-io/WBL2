package grep

import (
	"fmt"
	"grep/pkg/grep/config"
	"grep/pkg/grep/found"
	"grep/pkg/io"
	"log"
	"regexp"
	"strings"
)

//Grep главная структура куда будут записаны изначальные данные для дальнейшей их обработки
type Grep struct {
	cnf     config.Conf
	rawData []string
}

//GetConf геттер конфигурации
func (g Grep) GetConf() config.Conf { return g.cnf }

//GetData геттер данных
func (g Grep) GetData() []string { return g.rawData }

//NewGrep конструктор для grep
func NewGrep() *Grep {
	cnf := config.GetConfig()
	var rawData = io.GetData(cnf)
	return &Grep{cnf, rawData}
}

//Run запуск основной программы
func (g *Grep) Run() string {
	var reg *regexp.Regexp
	var pref, post string
	var err error

	if g.cnf.Keyi {
		pref = "(?i)"
	}
	if g.cnf.KeyF {
		pref = "^" + pref
		post = "$"
	}
	reg, err = regexp.Compile(pref + g.cnf.Request + post)
	if err != nil {
		log.Fatal(err)
	}
	if g.cnf.Keyc {
		count := 0
		for _, v := range g.rawData {
			seg := reg.FindIndex([]byte(v))
			if seg != nil {
				count++
			}
		}
		return fmt.Sprintf("%d\n", count)
	}
	foundGroup := CreateFoundGroup(g, reg)
	sb := strings.Builder{}
	for _, v := range foundGroup {
		sb.WriteString(strings.Join(v.GetData(), ""))
	}
	return sb.String()
}

//CreateFoundGroup TODO
func CreateFoundGroup(g *Grep, reg *regexp.Regexp) []*found.Found {
	var pointIndex = make([]*found.PointIndex, 0, 10)
	for i, v := range g.rawData {
		//seg := reg.FindIndex([]byte(v))
		seg := reg.FindAllIndex([]byte(v), -1)
		if seg != nil {
			pointIndex = append(pointIndex, found.NewPointIndex(i, g.cnf.KeyB, g.cnf.KeyA, seg))
		}
	}
	pointIndex = found.MixPoints(len(g.rawData), pointIndex...)
	var f = make([]*found.Found, 0, len(pointIndex))
	for i := range pointIndex {
		f = append(f, found.NewFound(g.GetConf(), g.rawData, pointIndex[i]))
	}
	return f
}
