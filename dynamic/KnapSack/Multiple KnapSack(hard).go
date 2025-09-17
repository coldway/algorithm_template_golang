/*
【模板】多重背包求方案数（Python/Java/C++/Go）
https://leetcode.cn/problems/number-of-ways-to-earn-points/solutions/2148313/fen-zu-bei-bao-pythonjavacgo-by-endlessc-ludl/

更快的做法请看 2902. 和带限制的子多重集合的数目
func waysToReachTarget(target int, types [][]int) int {
	const mod = 1_000_000_007
	f := make([]int, target+1)
	f[0] = 1
	for _, p := range types {
		count, marks := p[0], p[1]
		for j := target; j > 0; j-- {
			for k := 1; k <= min(count, j/marks); k++ {
				f[j] += f[j-k*marks]
			}
			f[j] %= mod
		}
	}
	return f[target]
}
*/
/*
https://leetcode.cn/problems/number-of-ways-to-earn-points/description/
https://leetcode.cn/problems/find-the-original-typed-string-ii/description/
https://leetcode.cn/problems/count-of-sub-multisets-with-bounded-sum/description/
https://leetcode.cn/problems/zero-array-transformation-iv/description/
*/