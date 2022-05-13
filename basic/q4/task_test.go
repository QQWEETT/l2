package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetAnagrams(t *testing.T) {
	anagrams := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "Лунь", "нуль", "горечь"}

	testM := make(map[string][]string)
	testM["листок"] = append(testM["листок"], "листок", "слиток", "столик")
	testM["лунь"] = append(testM["лунь"], "лунь", "нуль")
	testM["пятак"] = append(testM["пятак"], "пятак", "пятка", "тяпка")

	m := check(anagrams)
	for k := range m {
		fmt.Println(m[k])
		if ok := reflect.DeepEqual(testM[k], m[k]); !ok {
			t.Error("result != value")
		}
	}
}
