package main

import (
	"bufio"
	"io/fs"
	"io/ioutil"
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
		column  int
		reverse bool
		unique  bool
		byNum   bool
		file    string
		res     []string
	}{
		{
			column:  2,
			reverse: true,
			unique:  true,
			byNum:   true,
			file:    "test.txt",
			res: []string{
				"drwxr-xr-x 12 user user 4096 янв 14 21:49 Documents",
				"drwxr-xr-x 7 user user 4096 янв 13 11:42 Lightworks",
				"drwxr-xr-x 6 user user 4096 дек 6 14:29 Android",
				"drwx------ 5 user user 12288 янв 15 14:59 Downloads",
				"drwx------ 3 user user 4096 янв 14 22:18 Desktop",
			},
		},
	}

	for _, input := range tableTest {

		f, err := os.Open(input.file)
		if err != nil {
			log.Fatalln(err)
		}
		fl := &FlagsSort{column: 2, reverse: true, unique: true, byNum: true}

		fscan = bufio.NewScanner(f)
		sl = readScan(fscan)
		ioutil.WriteFile("test.txt", Sort(sl, fl), fs.ModePerm)

		for i, v := range input.res {

			if sl[i] != v {
				t.Error("result != value")
			}
		}
	}
}
