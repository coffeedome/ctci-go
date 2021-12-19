package main

import "fmt"

func main() {

	input := [][]int{
		{1, 0, 3},
		{4, 5, 6},
		{0, 8, 9},
	}

	fmt.Println(Zeroify(input))

}

func Zeroify(m [][]int) [][]int {

	n := make([][]int, len(m))
	for i := range m {
		n[i] = make([]int, len(m[0]))
		copy(n[i], m[i])
	}

	var zcolms []int

	for i, v := range m {
		for j := range v {
			if m[i][j] == 0 {
				n[i] = make([]int, len(m[0]))
				zcolms = append(zcolms, j)
			}
		}
	}

	for _, j := range zcolms {
		for i := 0; i < len(n[0]); i++ {
			n[i][j] = 0
		}
	}

	return n
}
