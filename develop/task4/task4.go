package main

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"fmt"
	"sort"
	"strings"
)

// findAnagram производит поиск анаграмм
func FindAnagram(arr []string) map[string][]string {
	data := make(map[string][]string)
	keys := make(map[string]string)

	for _, s := range arr {
		s = strings.ToLower(s)
		s1 := []rune(s)

		sort.Slice(s1, func(i, j int) bool {
			return s1[i] < s1[j]
		})

		if _, ok := keys[string(s1)]; !ok {
			data[s] = append(data[s], s)
			keys[string(s1)] = s
			continue
		}

		key := keys[string(s1)]
		data[key] = append(data[key], s)
	}

	for i := range data {
		if len(data[i]) < 2 {
			delete(data, i)
		} else {
			sort.Strings(data[i])
		}
	}

	return data
}

func main() {
	arr := []string{"тяпка", "ПЯТАК", "Пятка", "бетон", "СЛИТОК", "столик", "листок"}

	data := FindAnagram(arr)

	for _, i := range data {
		fmt.Println(i)
	}
}
