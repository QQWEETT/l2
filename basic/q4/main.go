package main

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
