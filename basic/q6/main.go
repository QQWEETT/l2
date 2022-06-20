package main

/*
=== Утилита cut ===
Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные
Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/


import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"
)

func init() {
	testing.Init()
	flag.IntVar(&column, "f", 0, "")
	flag.BoolVar(&byDel, "d", false, "")
	flag.BoolVar(&bySep, "s", false, "")
	flag.Parse()
}

func readScan(scan *bufio.Scanner) []string {
	s := make([]string, 0)

	for scan.Scan() {
		s = append(s, scan.Text())
	}

	return s
}

var d string

func fCut(lines []string, k int, byDel bool, bySep bool) []string {
	s := make([][]string, 0)

	k = k - 1
	if k < 0 {
		k = 0
	}

	if byDel {
		fmt.Println("Введите разделитель: ")
		fmt.Scan(&d)

		for _, line := range lines {
			s = append(s, strings.Split(line, d))

		}

	} else {
		for _, line := range lines {
			s = append(s, strings.Split(line, " "))
		}
	}

	if bySep {
		for i := 0; i < len(s); i++ {
			if len(s[i]) == 1 {
				s = append(s[:i], s[i+1:]...)
			}

		}
	}
	sl = make([]string, 0)

	for i := 0; i < len(s); i++ {
		sl = append(sl, s[i][k])

	}

	return sl
}

var fscan *bufio.Scanner
var column int
var sl []string
var byDel bool
var bySep bool

type FlagsSort struct {
	column int
	f      bool
	byDel  bool
	bySep  bool
}

func cut(sl []string, flags *FlagsSort) {
	sl = fCut(sl, flags.column, flags.byDel, flags.bySep)

	fmt.Println((strings.Join(sl, "\n")))
}
func main() {
	f, err := os.Open("test.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	fl := &FlagsSort{column: 2, f: true, byDel: false, bySep: true}
	fscan = bufio.NewScanner(f)
	sl = readScan(fscan)
	cut(sl, fl)
}
