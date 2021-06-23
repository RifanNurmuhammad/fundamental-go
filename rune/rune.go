package main

import (
	"fmt"
	"sort"
)

func main() {

	/// when you send "ABCDE" rune will convert to int32 as [65, 66, 67, 68, 69]
	fmt.Println(convertAnyStringToRune("ABCDE"))
	/// when you send "EABCD" rune will convert to int32 as [69, 65, 66, 67, 68]
	fmt.Println(convertAnyStringToRune("EABCD"))
	/// when you send "ECABD" rune will convert to int32 as [69, 67, 65, 66, 68]
	fmt.Println(convertAnyStringToRune("ECABD"))

	// combine rune convert with sorting
	beforeSort := "DBCAFE"
	// sort string using convert rune
	convertedString := convertAnyStringToRune(beforeSort)
	afterSort := sortValueString(convertedString)
	// print "ABCDEF"
	fmt.Println(afterSort)
}

func sortValueString(val []rune) string {
	sort.Slice(val, func(k, v int) bool {
		return val[k] < val[v]
	})
	return string(val)
}

func convertAnyStringToRune(val string) []rune {
	runeValue := []rune{}
	for _, v := range val {
		runeValue = append(runeValue, v)
	}
	return runeValue
}
