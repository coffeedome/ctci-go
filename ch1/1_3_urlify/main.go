package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "Foo Bar   "
	out := Urlify(input)
	fmt.Println(out)

}

func Urlify(s string) string {
	space_count := 0
	for _, v := range s {
		if v == []rune(" ")[0] {
			space_count++
		}
	}

	if space_count == 0 {
		return s
	}

	r := strings.Split(s, "")

	for i, v := range r {
		if v == " " {
			r[i] = "%20"
		}
	}

	return strings.Join(r, "")
}
