package dynamic_programming

// FibonacciDP 斐波那契数列 - 动态规划实现
// 返回斐波那契数列的第n个数 F(n)
// F(0) = 0, F(1) = 1
// F(n) = F(n-1) + F(n-2), 当 n > 1
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func FibonacciDP(n int) int {
	if n <= 1 {
		return n
	}

	// 创建DP数组
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1

	// 自底向上计算
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// FibonacciOptimized 斐波那契数列 - 优化空间复杂度
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func FibonacciOptimized(n int) int {
	if n <= 1 {
		return n
	}

	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}

	return b
}

// ClimbStairs 爬楼梯
// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
// 时间复杂度: O(n)
// 空间复杂度: O(n) 可优化到 O(1)
func ClimbStairs(n int) int {
	if n <= 2 {
		return n
	}

	// 创建DP数组
	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 2

	// 自底向上计算
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// CoinChange 零钱兑换
// 给定不同面额的硬币 coins 和一个总金额 amount，计算可以凑成总金额所需的最少的硬币个数。
// 如果没有任何一种硬币组合能组成总金额，返回 -1。
// 时间复杂度: O(amount * len(coins))
// 空间复杂度: O(amount)
func CoinChange(coins []int, amount int) int {
	// 创建DP数组，初始值设为amount+1（一个不可能达到的值）
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0 // 凑成金额0需要0个硬币

	// 自底向上计算
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i {
				// 选择当前硬币 vs 不选择当前硬币，取最小值
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	// 如果dp[amount]仍为初始值，说明无解
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

// LongestIncreasingSubsequence 最长递增子序列
// 给定一个无序的整数数组，找到其中最长上升子序列的长度。
// 时间复杂度: O(n^2)
// 空间复杂度: O(n)
func LongestIncreasingSubsequence(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// dp[i] 表示以nums[i]结尾的最长递增子序列的长度
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1 // 初始化为1，因为最短的递增子序列就是元素本身
	}

	maxLen := 1
	// 自底向上计算
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxLen = max(maxLen, dp[i])
	}

	return maxLen
}

// KnapsackProblem 01背包问题
// 给定n个物品，每个物品有重量weights[i]和价值values[i]。
// 背包的最大承重为maxWeight，求解放入背包的物品的最大价值。
// 每个物品要么完全放入背包，要么不放入（01背包）。
// 时间复杂度: O(n * maxWeight)
// 空间复杂度: O(n * maxWeight) 可优化到 O(maxWeight)
func KnapsackProblem(weights []int, values []int, maxWeight int) int {
	n := len(weights)
	if n == 0 || maxWeight == 0 {
		return 0
	}

	// 创建二维DP数组，dp[i][w]表示考虑前i个物品，背包容量为w时的最大价值
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, maxWeight+1)
	}

	// 自底向上计算
	for i := 1; i <= n; i++ {
		for w := 1; w <= maxWeight; w++ {
			// 当前物品不放入背包
			dp[i][w] = dp[i-1][w]

			// 如果当前物品可以放入背包，则考虑放入
			if weights[i-1] <= w {
				// 选择当前物品 vs 不选择当前物品，取最大值
				dp[i][w] = max(dp[i][w], dp[i-1][w-weights[i-1]]+values[i-1])
			}
		}
	}

	return dp[n][maxWeight]
}

// LongestCommonSubsequence 最长公共子序列
// 给定两个字符串text1和text2，返回这两个字符串的最长公共子序列的长度。
// 一个字符串的子序列是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
// 时间复杂度: O(m * n)
// 空间复杂度: O(m * n)
func LongestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	if m == 0 || n == 0 {
		return 0
	}

	// 创建二维DP数组，dp[i][j]表示text1[0...i-1]和text2[0...j-1]的LCS长度
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// 自底向上计算
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				// 当前字符相同，LCS长度加1
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// 当前字符不同，取前面计算的最大值
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}

// EditDistance 编辑距离
// 给定两个单词word1和word2，计算出将word1转换成word2所使用的最少操作数。
// 你可以对一个单词进行如下三种操作：插入一个字符、删除一个字符、替换一个字符
// 时间复杂度: O(m * n)
// 空间复杂度: O(m * n)
func EditDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)

	// 创建二维DP数组，dp[i][j]表示word1[0...i-1]转换到word2[0...j-1]的最小操作数
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// 边界情况初始化
	for i := 0; i <= m; i++ {
		dp[i][0] = i // 将word1[0...i-1]转换为空串需要删除i个字符
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j // 将空串转换为word2[0...j-1]需要插入j个字符
	}

	// 自底向上计算
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				// 当前字符相同，不需要操作
				dp[i][j] = dp[i-1][j-1]
			} else {
				// 当前字符不同，取三种操作的最小值加1
				// 替换: dp[i-1][j-1] + 1
				// 删除: dp[i-1][j] + 1
				// 插入: dp[i][j-1] + 1
				dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1]))
			}
		}
	}

	return dp[m][n]
}

// min 返回两个整数中的较小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// max 返回两个整数中的较大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
