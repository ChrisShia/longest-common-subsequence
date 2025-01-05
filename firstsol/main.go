package main

import (
	"fmt"
)

func lcs(str1, str2 string) string {
	n := len(str1)
	m := len(str2)

	// Create a table to store lengths of longest common suffixes of substrings
	// Create two dimensional array
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	// Fill dp table
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	// Reconstruct LCS
	lcs := ""
	i, j := n, m
	for i > 0 && j > 0 {
		if str1[i-1] == str2[j-1] {
			lcs = string(str1[i-1]) + lcs
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	return lcs
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

	//todo: check str1=aasdfjhg, str2=asdfkugshdf --- got asdfg....seems right..
	fmt.Print("Enter two strings: ")
	var str1 string
	fmt.Scanln(&str1)
	var str2 string
	fmt.Scanln(&str2)
	lcs := lcs(str1, str2)
	fmt.Println("LCS is:", lcs)
}
