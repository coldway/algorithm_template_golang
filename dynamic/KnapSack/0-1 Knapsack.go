package main

import (
	"fmt"
	"math"
)

// deepseek-ai 0-1背包问题

// =================== 基础实现（二维数组）===================
// 时间复杂度：O(n*c) 空间复杂度：O(n*c)
func knapSackBasic(weights []int, values []int, capacity int) int {
	n := len(weights)
	// 创建二维DP数组 dp[i][j]表示前i件物品放入容量为j的背包的最大价值
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	// 动态规划过程
	for i := 1; i <= n; i++ {
		w, v := weights[i-1], values[i-1]
		for j := 0; j <= capacity; j++ {
			// 1. 不选当前物品
			dp[i][j] = dp[i-1][j]
			
			// 2. 选当前物品（需满足容量条件）
			if j >= w {
				dp[i][j] = max(dp[i][j], dp[i-1][j-w]+v)
			}
		}
	}
	
	return dp[n][capacity]
}

// =================== 空间优化（一维数组）===================
// 时间复杂度：O(n*c) 空间复杂度：O(c)
func knapSackOptimized(weights []int, values []int, capacity int) int {
	n := len(weights)
	// 创建一维DP数组 dp[j]表示容量为j的背包的最大价值
	dp := make([]int, capacity+1)

	// 动态规划过程（逆序更新）
	for i := 0; i < n; i++ {
		w, v := weights[i], values[i]
		// 必须逆序遍历容量（避免覆盖上一轮结果）
		for j := capacity; j >= w; j-- {
			// 状态转移：max(不选当前物品, 选当前物品)
			dp[j] = max(dp[j], dp[j-w]+v)
		}
	}
	
	return dp[capacity]
}

// =================== 恰好装满背包版本 ===================
func knapSackExact(weights []int, values []int, capacity int) int {
	dp := make([]int, capacity+1)
	
	// 初始化：只有容量0可达（价值0），其它容量初始化为负无穷
	dp[0] = 0
	for j := 1; j <= capacity; j++ {
		dp[j] = math.MinInt32
	}

	// 动态规划
	for i, w := range weights {
		v := values[i]
		for j := capacity; j >= w; j-- {
			if dp[j-w] != math.MinInt32 { // 确保前一个状态可达
				dp[j] = max(dp[j], dp[j-w]+v)
			}
		}
	}
	
	// 返回最大值（若为负无穷表示无法恰好装满）
	return max(0, dp[capacity]) // 处理无法装满的情况
}

// =================== 辅助函数 ===================
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	weights := []int{2, 3, 4, 5}  // 物品重量
	values := []int{3, 4, 5, 6}   // 物品价值
	capacity := 8                 // 背包容量

	fmt.Println("基础版(二维数组):", knapSackBasic(weights, values, capacity)) // 输出: 10
	fmt.Println("优化版(一维数组):", knapSackOptimized(weights, values, capacity)) // 输出: 10
	fmt.Println("恰好装满版本:", knapSackExact(weights, values, capacity)) // 输出: 10
	
	// 无法恰好装满的测试
	weights2 := []int{3, 5}
	values2 := []int{4, 6}
	fmt.Println("恰好装满(不可能):", knapSackExact(weights2, values2, 7)) // 输出: 0
}