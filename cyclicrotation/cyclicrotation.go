package main

import "fmt"

func main() {
	A := []int{3, 8, 9, 7, 6}
	K := 3
	fmt.Print(Solution(A, K))
}

// Solution for the Cyclic Rotation
func Solution(A []int, K int) []int {

	l := len(A)
	R := make([]int, l)

	// Get the exact rotate times t.
	t := 0
	if l != 0 {
		t = K % l
	}

	for i := 0; i < l; i++ {
		if (i + t) >= l {
			R[i+t-l] = A[i]
		} else {
			R[i+t] = A[i]
		}

	}
	return R
}
