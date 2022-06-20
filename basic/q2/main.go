package main

/*
=== Задача на распаковку ===
Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)
В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.
Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"fmt"
	"strconv"
	"strings"
)

//Проверяем 0й символ слайса. Если он равен цифре, то возвращаем true
func check(a string, numbers []string) bool {
	for _, b := range numbers {
		if a == b {
			return true
		}
	}
	return false
}

//Проверяем цифры в слайсе
func numbersInSlice(a string, numbers []string) bool {
	for _, b := range numbers {
		if b == a {
			return true
		}
	}
	return false
}

func RepeatS(s string) string {
	if len(s) == 0 {
		return ""
	}

	a := strings.Split(s, "")

	numbers := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	result := []string{}

	//Производим проверку. Если true, то выходим из программы
	if check(a[0], numbers) {
		return ""
	}

	for i := 0; i < len(a); i++ {
		if numbersInSlice(a[i], numbers) {
			d, _ := strconv.Atoi(a[i])
			result = append(result, strings.Repeat(a[i-1], d-1))

		} else {
			result = append(result, a[i])
		}
	}
	return strings.Join(result[:], "")
}
func main() {
	a := RepeatS("4a4bc2d5e")
	fmt.Println(a)
}
