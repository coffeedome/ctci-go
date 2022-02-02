package main

import "fmt"

func main() {

	cache := make(map[int]int)

	fmt.Println(Fibonacci(10, cache))

}

func Fibonacci(n int, cache map[int]int) int {

	if n < 2 {
		cache[n] = n
		return n
	}

	if _, ok := cache[n-1]; !ok {
		cache[n-1] = Fibonacci(n-1, cache)
	}

	if _, ok := cache[n-2]; !ok {
		cache[n-2] = Fibonacci(n-2, cache)
	}

	return cache[n-1] + cache[n-2]

}

type Tile struct {
	Right   *Tile
	Down    *Tile
	Blocked bool
	Visited bool
}

func GridPath(r int, c int, NextTile Tile) bool {

	//cache is list of tiles visited
	if r < 2 {
	}

	//GridPath(r,c) = GridPath(r+1,c) || GridPath(r,c+1), assume no r and d blocked at the same time

	if NextTile.Right == nil && NextTile.Down == nil {
		return false
	}

	if NextTile.Right.Blocked {
		GridPath(r+1, c, *NextTile.Down, visitedcache)
	}

	if NextTile.Down.Blocked {
		GridPath(r, c+1, *NextTile.Right, visitedcache)
	}

	return
}
