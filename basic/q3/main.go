package main

/*
=== Утилита sort ===
Отсортировать строки (man sort)
Основное
Поддержать ключи
-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки
Дополнительное
Поддержать ключи
-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func init() {
	testing.Init()
	flag.IntVar(&column, "k", 0, "sort via a key column")
	flag.BoolVar(&byNum, "n", false, "compare according to string numerical value")
	flag.BoolVar(&reverse, "r", false, "reverse the result of comparisons")
	flag.BoolVar(&unique, "u", false, "output only the first of an equal run")
	flag.Parse()
}
func readScan(scan *bufio.Scanner) []string {
	s := make([]string, 0)

	for scan.Scan() {
		s = append(s, scan.Text())
	}

	return s
}

func sortColumn(lines []string, k int, byNum bool) []string {

	s := make([][]string, 0)

	k = k - 1
	if k < 0 {
		k = 0
	}

	for _, line := range lines {
		s = append(s, strings.Split(line, " "))
	}

	if byNum {
		sort.SliceStable(s, func(i, j int) bool {
			if len(s[i]) > k && len(s[j]) > k {
				x, err := strconv.Atoi(s[i][k])
				y, err := strconv.Atoi(s[j][k])
				if err != nil {
					fmt.Println(err)
					return false
				}

				return x < y
			}

			return false
		})
	} else {
		sort.SliceStable(s, func(i, j int) bool {
			if len(s[i]) > k && len(s[j]) > k {
				return strings.ToLower(s[i][k]) < strings.ToLower(s[j][k])
			}
			return false
		})
	}

	var str string
	sl = make([]string, 0)
	// строка файла которая была разделена пробелом, джонится обратно пробелом
	for _, line := range s {
		str = strings.Join(line, " ")
		sl = append(sl, str)
	}
	// возвращаем уже отсортированный слайс
	return sl

}

func Reverse(sl []string) []string {

	for i, j := 0, len(sl)-1; i < j; i, j = i+1, j-1 {
		sl[i], sl[j] = sl[j], sl[i]
	}

	// возвращаем уже отсортированный слайс
	return sl
}

func removeDuplicateStr(strSlice []string) []string {
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

func Sort(sl []string, flags *FlagsSort) []byte {
	if flags.column > -1 {
		sl = sortColumn(sl, flags.column, flags.byNum)
	}

	if flags.reverse {
		sl = Reverse(sl)
	}
	if flags.unique {
		sl = removeDuplicateStr(sl)
	}
	return []byte(strings.Join(sl, "\n"))
}

var fscan *bufio.Scanner
var column int
var byNum bool
var reverse bool
var unique bool

var sl []string

type FlagsSort struct {
	column  int
	reverse bool
	unique  bool
	byNum   bool
}

func main() {
	f, err := os.Open("test.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	fl := &FlagsSort{column: column, reverse: reverse, unique: unique, byNum: byNum}
	fscan = bufio.NewScanner(f)
	sl = readScan(fscan)
	ioutil.WriteFile("test.txt", Sort(sl, fl), fs.ModePerm)
}
