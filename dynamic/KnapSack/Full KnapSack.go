//物品可以重复选，无个数限制。

//**问**：关于完全背包，有两种写法，
//一种是外层循环枚举物品，内层循环枚举体积
//另一种是外层循环枚举体积，内层循环枚举物品。如何评价这两种写法的优劣？
//**答**：两种写法都可以，但更推荐前者。
//外层循环枚举物品的写法，只会遍历物品数组一次；
//而内层循环枚举物品的写法，会遍历物品数组多次。
//从 cache 的角度分析，多次遍历数组会导致额外的 cache miss，带来额外的开销。所以虽然这两种写法的时间空间复杂度是一样的，但外层循环枚举物品的写法常数更小。

// 典型题：
// 322. 零钱兑换（https://leetcode.cn/problems/coin-change/description/）
// 518. 零钱兑换 II（https://leetcode.cn/problems/coin-change-ii/description/）
// 279. 完全平方数（https://leetcode.cn/problems/perfect-squares/description/）
package main

import (
	"fmt"
	"math"
)

// =================== 完全背包 ===================
// 基础实现（二维数组）
// 时间复杂度：O(n*c) 空间复杂度：O(n*c)
func fullKnapSackBasic(weights []int, values []int, capacity int) int {
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
			// 不选当前物品
			dp[i][j] = dp[i-1][j]
			// 选当前物品（可重复选，需满足容量条件）
			if j >= w {
				dp[i][j] = max(dp[i][j], dp[i][j-w]+v)
			}
		}
	}

	return dp[n][capacity]
}

// 空间优化（一维数组）
// 时间复杂度：O(n*c) 空间复杂度：O(c)
func fullKnapSackOptimized(weights []int, values []int, capacity int) int {
	n := len(weights)
	// 创建一维DP数组 dp[j]表示容量为j的背包的最大价值
	dp := make([]int, capacity+1)

	// 动态规划过程（正序更新）
	for i := 0; i < n; i++ {
		w, v := weights[i], values[i]
		// 必须正序遍历容量（因为物品可重复选）
		for j := w; j <= capacity; j++ {
			// 状态转移：max(不选当前物品, 选当前物品)
			dp[j] = max(dp[j], dp[j-w]+v)
		}
	}

	return dp[capacity]
}

// 恰好装满背包版本
func fullKnapSackExact(weights []int, values []int, capacity int) int {
	dp := make([]int, capacity+1)

	// 初始化：只有容量0可达（价值0），其它容量初始化为负无穷
	dp[0] = 0
	for j := 1; j <= capacity; j++ {
		dp[j] = math.MinInt32
	}

	// 动态规划
	for i, w := range weights {
		v := values[i]
		for j := w; j <= capacity; j++ {
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
	weights := []int{2, 3, 4, 5} // 物品重量
	values := []int{3, 4, 5, 6}  // 物品价值
	capacity := 8                // 背包容量

	fmt.Println("完全背包 基础版(二维数组):", fullKnapSackBasic(weights, values, capacity))
	fmt.Println("完全背包 优化版(一维数组):", fullKnapSackOptimized(weights, values, capacity))
	fmt.Println("完全背包 恰好装满版本:", fullKnapSackExact(weights, values, capacity))
}

// 322. 零钱兑换
/*
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
你可以认为每种硬币的数量是无限的。

示例 1：
输入：coins = [1, 2, 5], amount = 11
输出：3
解释：11 = 5 + 5 + 1
示例 2：
输入：coins = [2], amount = 3
输出：-1
示例 3：
输入：coins = [1], amount = 0
输出：0
*/
func coinChange(coins []int, amount int) int {
	n := len(coins)
	// 创建二维DP数组 dp[i][j]表示前i件物品放入容量为j的背包的最大价值
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
	}

	for j := range dp[0] {
		dp[0][j] = math.MaxInt / 2 // 除 2 防止下面 + 1 溢出
	}

	dp[0][0] = 0
	// 动态规划过程
	for i := 1; i <= n; i++ {
		v := coins[i-1]
		for j := 0; j <= amount; j++ {
			// 不选当前物品
			dp[i][j] = dp[i-1][j]
			// 选当前物品（可重复选，需满足容量条件）
			if j >= v {
				dp[i][j] = min(dp[i][j], dp[i][j-v]+1)
			}
		}
	}

	ans := dp[n][amount]
	if ans < math.MaxInt/2 {
		return ans
	}
	return -1

}

//518. 零钱兑换 II
/*
给你一个整数数组 coins 表示不同面额的硬币，另给一个整数 amount 表示总金额。

请你计算并返回可以凑成总金额的硬币组合数。如果任何硬币组合都无法凑出总金额，返回 0 。

假设每一种面额的硬币有无限个。

题目数据保证结果符合 32 位带符号整数。

示例 1：
输入：amount = 5, coins = [1, 2, 5]
输出：4
解释：有四种方式可以凑成总金额：
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1

示例 2：
输入：amount = 3, coins = [2]
输出：0
解释：只用面额 2 的硬币不能凑成总金额 3 。

示例 3：
输入：amount = 10, coins = [10]
输出：1
*/

/* way1*/
func change(amount int, coins []int) int {
	n := len(coins)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, amount+1)
	}
	f[0][0] = 1
	for i, x := range coins {
		for c := 0; c <= amount; c++ {
			if c < x {
				f[i+1][c] = f[i][c]
			} else {
				f[i+1][c] = f[i][c] + f[i+1][c-x]
			}
		}
	}
	return f[n][amount]
}

/* way2*/
func change(amount int, coins []int) int {
	f := make([]int, amount+1)
	f[0] = 1
	for _, x := range coins {
		for c := x; c <= amount; c++ {
			f[c] += f[c-x]
		}
	}
	return f[amount]
}

// 作者：灵茶山艾府
// 链接：https://leetcode.cn/problems/coin-change-ii/solutions/2706227/shi-pin-wan-quan-bei-bao-cong-ji-yi-hua-o3ew0/

//279. 完全平方数
/*
给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。
完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。


示例 1：
输入：n = 12
输出：3
解释：12 = 4 + 4 + 4

示例 2：
输入：n = 13
输出：2
解释：13 = 4 + 9

提示：
1 <= n <= 104
*/

/*way1*/
func numSquares(n int) int {
	// 创建二维DP数组 dp[i][j]表示前i件物品放入容量为j的背包的最大价值
	var dp [101][10000 + 1]int
	for i := 1; i <= 10000; i++ {
		dp[0][i] = math.MaxInt
	}

	// 动态规划过程
	for i := 1; i*i <= 10000; i++ {
		for j := 0; j <= 10000; j++ {
			// 不选当前物品
			dp[i][j] = dp[i-1][j]
			// 选当前物品（可重复选，需满足容量条件）
			if j >= i*i {
				dp[i][j] = min(dp[i][j], dp[i][j-i*i]+1)
			}
		}
	}

	return dp[int(math.Sqrt(float64(n)))][n] // 也可以写 f[100][n]
}

/*way2*/
const N = 10000

var f [N + 1]int

func init() {
	for i := 1; i <= N; i++ {
		f[i] = math.MaxInt
	}
	for i := 1; i*i <= N; i++ {
		for j := i * i; j <= N; j++ {
			f[j] = min(f[j], f[j-i*i]+1) // 不选 vs 选
		}
	}
}

func numSquares(n int) int {
	return f[n]
}

// 作者：灵茶山艾府
// 链接：https://leetcode.cn/problems/perfect-squares/solutions/2830762/dong-tai-gui-hua-cong-ji-yi-hua-sou-suo-3kz1g/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

// 1449. 数位成本和为目标值的最大数字
/*
给你一个整数数组 cost 和一个整数 target 。请你返回满足如下规则可以得到的 最大 整数：

给当前结果添加一个数位（i + 1）的成本为 cost[i] （cost 数组下标从 0 开始）。
总成本必须恰好等于 target 。
添加的数位中没有数字 0 。
由于答案可能会很大，请你以字符串形式返回。

如果按照上述要求无法得到任何整数，请你返回 "0" 。

示例 1：

输入：cost = [4,3,2,5,6,7,2,5,5], target = 9
输出："7772"
解释：添加数位 '7' 的成本为 2 ，添加数位 '2' 的成本为 3 。所以 "7772" 的代价为 2*3+ 3*1 = 9 。 "977" 也是满足要求的数字，但 "7772" 是较大的数字。
 数字     成本
  1  ->   4
  2  ->   3
  3  ->   2
  4  ->   5
  5  ->   6
  6  ->   7
  7  ->   2
  8  ->   5
  9  ->   5
示例 2：

输入：cost = [7,6,5,5,5,6,8,7,8], target = 12
输出："85"
解释：添加数位 '8' 的成本是 7 ，添加数位 '5' 的成本是 5 。"85" 的成本为 7 + 5 = 12 。
示例 3：

输入：cost = [2,4,6,2,4,6,4,4,4], target = 5
输出："0"
解释：总成本是 target 的条件下，无法生成任何整数。
示例 4：

输入：cost = [6,10,15,40,40,40,40,40,40], target = 47
输出："32211"


提示：

cost.length == 9
1 <= cost[i] <= 5000
1 <= target <= 5000
*/
func largestNumber(cost []int, target int) string {
	dp := make([][]int, 10)
	from := make([][]int, 10)
	for i := range dp {
		dp[i] = make([]int, target+1)
		for j := range dp[i] {
			dp[i][j] = math.MinInt32
		}
		from[i] = make([]int, target+1)
	}
	dp[0][0] = 0
	for i, c := range cost {
		for j := 0; j <= target; j++ {
			if j < c {
				dp[i+1][j] = dp[i][j]
				from[i+1][j] = j
			} else {
				if dp[i][j] > dp[i+1][j-c]+1 {
					dp[i+1][j] = dp[i][j]
					from[i+1][j] = j
				} else {
					dp[i+1][j] = dp[i+1][j-c] + 1
					from[i+1][j] = j - c
				}
			}
		}
	}
	if dp[9][target] < 0 {
		return "0"
	}
	ans := make([]byte, 0, dp[9][target])
	i, j := 9, target
	for i > 0 {
		if j == from[i][j] {
			i--
		} else {
			ans = append(ans, '0'+byte(i))
			j = from[i][j]
		}
	}
	return string(ans)
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/form-largest-integer-with-digits-that-add-up-to-target/solutions/824378/shu-wei-cheng-ben-he-wei-mu-biao-zhi-de-dnh86/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
