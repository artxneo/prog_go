package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}
	for i, s := range n {
		wg.Add(1)
		go func(s []int, i int) {
			defer wg.Done()
			fmt.Println("slise ", i, ":", sum(s))
		}(s, i)
	}

	wg.Wait()
}

func sum(s []int) int {
	sum := 0

	for _, v := range s {
		sum += v
	}

	return sum
}
