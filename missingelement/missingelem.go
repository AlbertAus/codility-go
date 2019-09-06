package main

import "fmt"

func main() {
	A := []int{2, 3, 1, 5}
	fmt.Print(Solution(A))
}

// Solution for Missing Element in a given permutation
// Calculate the permutation sum first, then sum the array value.
// (perSum - arraySum) will calculate the different between two values, the different is the missing integer.
func Solution(A []int) int {
	l := len(A)
	if l == 0 {
		return 1
	}

	arraySum := 0
	perSum := ((l + 1) + 1) * (l + 1) / 2

	for _, v := range A {
		arraySum += v
	}

	return perSum - arraySum

}
