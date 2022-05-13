package main

import (
	"bufio"
	"log"
	"os"
	"testing"
)

type args struct {
	files []string
	flags []string
}

func TestWithoutFlags(t *testing.T) {
	tableTest := []struct {
		column int
		f      bool
		byDel  bool
		bySep  bool
		file   string
		res    []string
	}{
		{
			column: 2,
			f:      true,
			byDel:  true,
			bySep:  true,
			file:   "test.txt",
			res: []string{
				"6,",
				"3,",
				"12,",
				"5,",
				"7,",
			},
		},
	}

	for _, input := range tableTest {

		f, err := os.Open(input.file)
		if err != nil {
			log.Fatalln(err)
		}
		fl := &FlagsSort{column: 2, f: true, byDel: false, bySep: true}

		fscan = bufio.NewScanner(f)
		sl = readScan(fscan)
		cut(sl, fl)

		for i, v := range input.res {

			if sl[i] != v {
				t.Error("result != value")
			}
		}
	}
}
