package main

import (
	"fmt"
	"math/rand"
)

func SortInts(s []int) []int {
	if len(s) <= 1 {
		return s
	}

	pivote := s[0]

	lower := make([]int, 0, len(s))
	greater := make([]int, 0, len(s))

	for _, i := range s {
		if i < pivote {
			lower = append(lower, i)
		} else if i > pivote {
			greater = append(greater, i)
		}
	}

	new := []int{}
	new = append(new, SortInts(lower)...)
	new = append(new, pivote)
	new = append(new, SortInts(greater)...)

	return new
}

func TestSortInts() {
	var ints []int

	for i := 0; i < 10000; i++ {
		ints = append(ints, rand.Intn(1000))
	}

	for _, i := range ints {
		fmt.Print(i, " ")
	}

	sorted := SortInts(ints)
	fmt.Println()

	for _, i := range sorted {
		fmt.Print(i, " ")
	}
	fmt.Println()
	fmt.Println("steps : ", steps)
}
