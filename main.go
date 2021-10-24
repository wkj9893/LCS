package main

import (
	"fmt"
)

func longestCommonSubsequence(s1, s2 string) []string {
	m, n := len(s1), len(s2)
	dp := [][]int{}
	res := [][][]string{}
	for i := 0; i <= m; i++ {
		dp = append(dp, make([]int, n+1))
		res = append(res, make([][]string, n+1))
	}
	for j := 0; j <= n; j++ {
		res[0][j] = []string{""}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				for _, v := range res[i-1][j-1] {
					res[i][j] = append(res[i][j], v+string(s1[i-1]))
				}
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
				if dp[i-1][j] < dp[i][j-1] {
					res[i][j] = append([]string{}, res[i][j-1]...)
				} else if dp[i-1][j] > dp[i][j-1] {
					res[i][j] = append([]string{}, res[i-1][j]...)
				} else {
					m := map[string]bool{}
					for _, v := range res[i][j-1] {
						m[v] = true
					}
					for _, v := range res[i-1][j] {
						m[v] = true
					}
					for v := range m {
						res[i][j] = append(res[i][j], v)
					}
				}
			}
		}
	}
	return res[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s1 := "ABCD"
	s2 := "ACBAD"
	fmt.Println("s1 =", s1)
	fmt.Println("s2 =", s2)
	fmt.Println("result =", longestCommonSubsequence(s1, s2))
}
