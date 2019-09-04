package main

import "fmt"

func main() {
	A := []int{9, 3, 9, 3, 9, 7, 9}
	fmt.Print(Solution(A))
}

// Solution for odd Occurrences Array.
// An assignment operation x op= y where op is a binary arithmetic operator is equivalent to x = x op (y)
// but evaluates x only once. The op= construct is a single token. In assignment operations,
// both the left- and right-hand expression lists must contain exactly one single-valued expression,
// and the left-hand expression must not be the blank identifier.
func Solution(A []int) int {
	var oddOccur int
	for i := 0; i < len(A); i++ {
		oddOccur ^= A[i]
	}
	return oddOccur
}
