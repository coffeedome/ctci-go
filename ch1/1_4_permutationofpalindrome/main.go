package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	var input string
	fmt.Println("Enter a string: ")
	fmt.Scanln(input)
	out, err := IsPermutationOfPalindrome(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Input string %s is permutation of palindrome: %t", input, out)
}

func IsPermutationOfPalindrome(s string) (bool, error) {

	if len(s) == 1 { // e.g. when s="a"
		return true, nil
	}

	if len(s) == 0 { // when s=""
		return false, errors.New("empty string is invalid")
	}

	tracker := make(map[rune]int)

	var countUneven int

	for _, v := range s {
		if v == []rune(" ")[0] {
			tracker[v] = 0
		} else {
			tracker[v] += 1
		}

	}

	for _, v := range s {
		if tracker[v]%2 != 0 {
			countUneven++
		}
	}

	return countUneven <= 1, nil
} //O(n)
