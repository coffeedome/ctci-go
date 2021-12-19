package main

import (
	"errors"
	"fmt"
)

func main() {
	s1 := "hello"
	s2 := "olleh"

	fmt.Print(IsSubstring(s1, s2))
}

func IsSubstring(s1 string, s2 string) (bool, error) {

	if len(s1) == 0 || len(s2) == 0 {
		return false, errors.New("input string empty is invalid")
	}

	if len(s1) != len(s2) {
		return false, nil
	}

	for k := range s2 {
		rk := len(s2) - 1 - k
		if s2[rk] != s1[k] {
			return false, nil
		}
	}
	return true, nil
}
