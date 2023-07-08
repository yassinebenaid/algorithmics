package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func FindInt(n int, t []int) (i int, ok bool) {

	low := 0
	high := len(t) - 1

	for low <= high {
		steps++

		mid := (low + high) / 2

		if n == t[mid] {
			return mid, true
		} else if n < t[mid] {
			high = mid - 1
		} else if n > t[mid] {
			low = mid + 1
		}
	}

	return -1, false
}

func TestFindInt() {
	if len(os.Args) < 2 {
		log.Println("please provide some value")
		os.Exit(1)
	}

	var t = make([]int, 0, 1000)

	for i := 0; i < 1000; i++ {
		t = append(t, i+1)
	}

	needle, _ := strconv.Atoi(os.Args[1])

	i, ok := FindInt(needle, t)

	if ok {
		fmt.Println("exists at : ", i)
	} else {
		fmt.Println("not found")
	}

	fmt.Println("steps : ", steps)
}
