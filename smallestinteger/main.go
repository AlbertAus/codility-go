package main

import (
	"fmt"
	"sort"
)

func main() {
	// Initial array for testing purpose
	A := []int{1, 2, 3, 6, 4, 1, 2}
	fmt.Print(Solution(A))
}

// Solution for finding the smallest positive integer (greater than 0).
// Some task call finding the Missing positive integer (greater than 0).
func Solution(A []int) int {
	// Sort the Array from small to large
	sort.Ints(A)

	// Define r as result, then try to compare values in Array A.
	r := 1
	for i := 0; i < len(A); i++ {
		if A[i] <= 0 {
			r = 1
		} else {
			if A[i] == r {
				r++
			}

		}
	}
	return r
}
