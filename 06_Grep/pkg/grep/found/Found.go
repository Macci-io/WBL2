package found

import (
	"fmt"
	"grep/pkg/grep/config"
	"sort"
)

//Index найденные индексы и отрезки в них
type Index struct {
	index int
	seg   [][]int
}

//PointIndex слайс найденных индексов
type PointIndex struct {
	startString, endString int
	indexes                []Index
}

//GetIndex получить конкретный индекс с отрезком в ней
func (p *PointIndex) GetIndex(j int) *Index {
	for i := range p.indexes {
		if p.indexes[i].index == j {
			return &p.indexes[i]
		}
	}
	return nil
}

//GetSize получить размер (диапазона)
func (p PointIndex) GetSize() int {
	return p.endString - p.startString
}

//NewPointIndex конструктор
func NewPointIndex(index, startStr, endStr int, seg [][]int) *PointIndex {
	return &PointIndex{
		startString: index - startStr,
		endString:   index + endStr,
		indexes:     []Index{{index, seg}},
	}
}

//MixPoints нахождение всех пересекающихся поинтов и их объединение при необходимости
func MixPoints(length int, pointIndexes ...*PointIndex) []*PointIndex {
	if pointIndexes == nil || len(pointIndexes) == 0 {
		return nil
	}
	var result = make([]*PointIndex, 0, len(pointIndexes))
	sort.SliceIsSorted(pointIndexes, func(i, j int) bool {
		return pointIndexes[i].indexes[0].index > pointIndexes[j].indexes[0].index
	})
	var i int
	result = append(result, pointIndexes[0])
	if result[0].startString < 0 {
		result[0].startString = 0
	}
	for _, v := range pointIndexes[1:] {
		if result[i].endString >= v.startString {
			result[i].endString = v.endString
			result[i].indexes = append(result[i].indexes, v.indexes...)
		} else {
			i++
			result = append(result, v)
		}
	}
	if result[len(result)-1].endString >= length {
		result[len(result)-1].endString = length - 1
	}
	return result
}

//Found найденные данные будут сохранены в этой структуре
type Found struct {
	Conf config.Conf
	data []string

	indexes *PointIndex
}

//NewFound конструктор
func NewFound(conf config.Conf, data []string, indexes *PointIndex) *Found {
	return &Found{conf, data, indexes}
}

//GetData подготовка результата
func (f Found) GetData() []string {
	var result []string
	var row string
	var start = f.indexes.startString
	var end = f.indexes.endString

	for ; start <= end; start++ {
		row = prepareResult(&f, start)
		result = append(result, row)
	}
	return result
}

func prepareResult(f *Found, i int) string {
	var prefix string
	var index *Index
	var row = f.data[i]

	index = f.indexes.GetIndex(i)
	if f.Conf.Keyn {
		if index != nil {
			prefix += fmt.Sprintf("\033[32m%d\033[34m:\033[0m", i+1)
			return prefix + prepareRow([]byte(row), index.seg) + "\n"
		}
		prefix += fmt.Sprintf("\033[32m%d\033[34m-\033[0m", i+1)
	}
	if index != nil {
		return prefix + prepareRow([]byte(row), index.seg) + "\n"
	}
	return prefix + row + "\n"

}

func prepareRow(row []byte, seg [][]int) string {
	var res string
	for i := range seg {
		if i == 0 {
			if seg[i][0] != 0 {
				res += "\033[0m" + string(row[:seg[i][0]])
			}
		} else {
			res += "\033[0m" + string(row[seg[i-1][1]:seg[i][0]])
		}
		res += "\033[31;1m" + string(row[seg[i][0]:seg[i][1]])
	}

	res += "\033[0m" + string(row[seg[len(seg)-1][1]:])
	return res
}
