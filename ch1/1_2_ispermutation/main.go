package main

import (
	"errors"
	"fmt"
	"sort"
)

func main() {
	out, err := IsPermutation("abc", "bac")
	if err != nil {
		errors.New("failed")
	}
	fmt.Println(out)
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

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(runeSlice(r))
	return string(r)
}

func IsPermutation(fstring string, sstring string) (bool, error) {

	if fstring == "" || sstring == "" {
		return false, errors.New("empty string")
	}

	sorted_f_string := SortString(fstring)
	sorted_s_string := SortString(sstring)

	return sorted_f_string == sorted_s_string, nil
}
