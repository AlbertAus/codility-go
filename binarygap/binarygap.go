package main

import "fmt"

func main() {
	N := 32
	fmt.Print(Solution(N))
}

// Solution for finding the longest Binary Gap.
func Solution(N int) int {
	maxGap := 0
	tmpGap := 0
	flag := false

	// When meet the first 1, then put the flag as true.
	// Calculate the tmpGap until next 1, if the flag is true, compare the tmpGap and maxGap, get the biggest one.
	for N > 0 {
		if N%2 == 1 {
			if !flag {
				flag = true
			} else {
				if tmpGap > maxGap {
					maxGap = tmpGap
				}
				tmpGap = 0
			}

		} else {
			if flag {
				tmpGap++
			}
		}
		N = N / 2
	}
	return maxGap
}
