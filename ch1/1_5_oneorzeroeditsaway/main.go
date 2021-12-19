package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	var firstString string
	var secondString string
	fmt.Println("Enter the first string:")
	fmt.Scanln(&firstString)
	fmt.Println("Enter the second string:")
	fmt.Scanln(&secondString)
	out, err := OneOrZeroEditsAway(firstString, secondString)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("One or less edits away? %t", out)
}

func OneOrZeroEditsAway(s1 string, s2 string) (bool, error) {

	if len(s1) == 0 || len(s2) == 0 {
		return false, errors.New("invalid input")
	}

	if len(s1)-len(s2) > 1 {
		return false, nil
	}

	tracks1 := make(map[rune]int)
	tracks2 := make(map[rune]int)

	for _, v := range s1 {
		tracks1[v] += 1
	}

	for _, v := range s2 {
		tracks2[v] += 1
	}

	var numEdits int

	checker := GetLargerString(s1, s2)

	for _, v := range checker {
		if tracks1[v] != tracks2[v] {
			numEdits++
		}
	}

	return numEdits <= 1, nil
}

func GetLargerString(s1 string, s2 string) string {
	if len(s1) > len(s2) {
		return s1
	}

	return s2
}
