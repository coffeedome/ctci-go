package main

import (
	"fmt"
	"sort"
)

func main() {
	var input string
	fmt.Print("Enter a string:\n")
	fmt.Scanln(&input)

	output := IsUniqueWithMap(input)
	fmt.Printf("String has unique characters: %t\n", output)

	second_output := IsUniqueWithoutMap(input)
	fmt.Printf("String has unique characters: %t\n", second_output)
}

func IsUniqueWithMap(str string) bool {
	if len(str) > 256 {
		return false // assume extended ASCII
	}

	char_set := make(map[rune]int)

	for _, c := range str {
		if char_set[c] > 0 {
			return false
		}
		char_set[c] += 1
	}
	return true
}

type runeSlice []rune

func (r runeSlice) Len() int {
	return len(r)
}

func (r runeSlice) Less(i int, j int) bool {
	return r[i] < r[j]
}

func (r runeSlice) Swap(i int, j int) {
	r[i], r[j] = r[j], r[i]
}

func IsUniqueWithoutMap(s string) bool {

	r := []rune(s)
	sort.Sort(runeSlice(r))

	for i := 1; i < len(r); i++ {
		if r[i-1] == r[i] {
			return false
		}
	}
	return true
}
