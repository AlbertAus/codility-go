package main

import "fmt"

func main() {
	fmt.Print(Solution(10, 85, 30))
}

// Solution for Frog Jump
func Solution(X int, Y int, D int) int {
	var result int
	if (Y-X)%D > 0 {
		result = ((Y - X) / D) + 1
	} else {
		result = (Y - X) / D
	}

	return result
}
