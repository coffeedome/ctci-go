package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {

	input := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("Before Rotation:")
	fmt.Println(input)
	out, err := RotateImageBy90Degrees(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("After Rotation:")
	fmt.Println(out)

}

func RotateImageBy90Degrees(m [][]int) ([][]int, error) {

	mheight := len(m)
	mwidth := len(m[1])

	if mheight != mwidth {
		return m, errors.New("matrix is not square")
	}

	newm := make([][]int, mheight)
	for i := range newm {
		newm[i] = make([]int, mwidth)
	}

	rotatednewm := make([][]int, mheight)
	for i := range rotatednewm {
		rotatednewm[i] = make([]int, mwidth)
	}

	//transpose matrix
	for i := range m {
		for j := range m[i] {
			newm[i][j] = m[j][i]
		}
	}

	//swap each row left to right
	for i := range newm {
		for j := range m[i] {
			h := len(m[i]) - 1 - j
			rotatednewm[i][j] = newm[i][h]
		}
	}

	return rotatednewm, nil
} // O(n^2)
