package main

import "fmt"

// lengthOfLIS 计算数组的最长递增子序列长度
// 时间复杂度：O(n^2)，n 为数组的长度
// 空间复杂度：O(n)
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	// 创建一维 DP 数组，dp[i] 表示以 nums[i] 结尾的最长递增子序列长度
	dp := make([]int, n)
	// 初始化每个元素的最长递增子序列长度为 1
	for i := range dp {
		dp[i] = 1
	}

	maxLen := 1
	// 动态规划过程
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				// 更新以 nums[i] 结尾的最长递增子序列长度
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		// 更新全局最长递增子序列长度
		maxLen = max(maxLen, dp[i])
	}

	return maxLen
}

// lengthOfLIS 计算数组的最长递增子序列长度，使用二分查找优化
// 时间复杂度：O(n log n)，n 为数组的长度
// 空间复杂度：O(n)
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	// tail 数组用于存储递增子序列，tail[i] 表示长度为 i+1 的递增子序列的末尾元素的最小值
	tail := make([]int, 0, n)
	tail = append(tail, nums[0])

	for i := 1; i < n; i++ {
		// 如果当前元素大于 tail 数组的最后一个元素，直接添加到 tail 数组末尾
		if nums[i] > tail[len(tail)-1] {
			tail = append(tail, nums[i])
		} else {
			// 二分查找第一个大于等于 nums[i] 的位置
			left, right := 0, len(tail)-1
			for left < right {
				mid := left + (right-left)/2
				if tail[mid] < nums[i] {
					left = mid + 1
				} else {
					right = mid
				}
			}
			// 更新该位置的值为 nums[i]
			tail[left] = nums[i]
		}
	}

	// tail 数组的长度即为最长递增子序列的长度
	return len(tail)
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
