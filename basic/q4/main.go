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

func anagram(words []string) map[string]*[]string {
	anagrams := make(map[string][]string)
	for _, w := range words {
		word := strings.ToLower(w)
		wordSorted := sortString(word)
		anagrams[wordSorted] = append(anagrams[wordSorted], word)
	}

	res := make(map[string]*[]string)
	for _, v := range anagrams {
		v := v
		if len(v) > 1 {
			res[v[0]] = &v

			sort.Slice(v, func(i, j int) bool {
				return v[i] < v[j]
			})
		}
	}
	return res
}

//Функция для сортировки
func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

//Функция для удаления дубликатов
func removeDuplicate(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func check(words []string) map[string][]string {
	words = removeDuplicate(words)
	fmt.Println(words)
	anagrams := anagram(words)
	s := make(map[string][]string)
	fmt.Println("-----------------------------------------")
	for k, v := range anagrams {
		s[k] = *v
	}
	return s
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "Лунь", "нуль", "горечь"}
	s := check(words)
	fmt.Println(s)
}
