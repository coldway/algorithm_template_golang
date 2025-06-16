/*
一般定义 f[i][j] 表示对 (s[:i],t[:j]) 的求解结果。
*/
package main

import "fmt"

// longestCommonSubsequence 计算两个字符串的最长公共子序列长度
// 时间复杂度：O(m*n)，m 和 n 分别为两个字符串的长度
// 空间复杂度：O(m*n)
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	// 创建二维 DP 数组，dp[i][j] 表示 text1 前 i 个字符和 text2 前 j 个字符的最长公共子序列长度
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// 动态规划过程
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				// 当前字符相等，最长公共子序列长度加 1
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// 当前字符不相等，取两种情况的最大值
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}

// max 返回两个整数中的最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// 测试最长公共子序列
	text1 := "abcde"
	text2 := "ace"
	fmt.Println("最长公共子序列长度:", longestCommonSubsequence(text1, text2))

	// 测试最长递增子序列
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println("最长递增子序列长度:", lengthOfLIS(nums))
}
