package anagram

import (
	"reflect"
	"strings"
)

// Sort сортирует буквы по алфавиту для дальнейшего использования в качестве ключа
func Sort[T string | rune](data []T) *[]T {
	size := len(data)
	swapper := reflect.Swapper(data)
	for j := 0; j < size; j++ {
		for k := j + 1; k < size; k++ {
			if data[j] > data[k] {
				swapper(j, k)
			}
		}
	}
	return &data
}

// Приводим все в единое состояние
func getSinglegram(raw []string) (*map[string]string, map[string]string) {
	var abcMap = make(map[string]string, len(raw))
	var firstKey = make(map[string]string)
	for i := range raw {
		// Все слова должны быть приведены к нижнему регистру.
		raw[i] = strings.ToLower(raw[i])
		abc := string(*Sort([]rune(raw[i])))
		if _, exist := firstKey[abc]; !exist {
			/*
				Сохраняем в мапу первое слово своей группы,
				из которой потом будем извлекать значения
				которые будут в качестве ключа для основной мапы
			*/
			firstKey[abc] = raw[i]
		} else {
			// Так же сортируем буквы по алфавиту для дальнейшего использования в качестве ключа
			abcMap[raw[i]] = abc
		}
	}
	return &abcMap, firstKey
}

func makeAnagram(rawSingle *map[string]string, first map[string]string) *map[string]*[]string {
	mediaResult := make(map[string]*[]string)
	result := make(map[string]*[]string)
	for k, v := range *rawSingle {
		if _, ok := mediaResult[v]; !ok {
			mediaResult[v] = new([]string)
		}
		*mediaResult[v] = append(*(mediaResult[v]), k)
	}
	for a, v := range mediaResult {
		// Множества из одного элемента не должны попасть в результат.
		if len(*v) > 1 {
			//Массив должен быть отсортирован по возрастанию.
			key, _ := first[a]
			result[key] = Sort(*v)
		}
	}
	return &result
}

// FindAnagram находит все анаграммы и возвращает результат
// Входные данные для функции: ссылка на массив, каждый элемент которого - слово на русском языке в кодировке utf8
// Выходные данные: ссылка на мапу множеств анаграмм
// Ключ - первое встретившееся в словаре слово из множества.
// Значение - ссылка на массив, каждый элемент которого, слово из множества
func FindAnagram(raw *[]string) *map[string]*[]string {
	rawSingle, firstKey := getSinglegram(*raw)
	return makeAnagram(rawSingle, firstKey)
}
