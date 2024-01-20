package main

import "fmt"

func min(nums ...int) int {
	minimum := int(^uint(0) >> 1)
	for _, i := range nums {
		if minimum > i {
			minimum = i
		}
	}
	return minimum
}
func print2D[T any](board [][]T) {
	for row := 0; row < len(board); row++ {
		for column := 0; column < len(board[0]); column++ {
			fmt.Print(board[row][column], " ")
		}
		fmt.Print("\n")
	}
}
func getLevenshtein(a, b string) int {
	n, m := len(a), len(b)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
		dp[0][i] = i
		dp[i][0] = i
	}
	if n == 0 {
		return m
	}
	if m == 0 {
		return n
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]) + 1
			}
		}
	}
	print2D(dp)
	return dp[n][m]
}

func main() {
	var a, b string
	fmt.Println("Enter first word:")
	fmt.Scan(&a)
	fmt.Println("Enter second word:")
	fmt.Scan(&b)
	levenshtein := getLevenshtein(a, b)
	fmt.Println(levenshtein)

}
