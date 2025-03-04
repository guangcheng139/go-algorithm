package dp

// Knapsack01 解决0-1背包问题
// weights: 物品的重量数组
// values: 物品的价值数组
// capacity: 背包的容量
// 返回能够装入背包的最大价值
func Knapsack01(weights []int, values []int, capacity int) int {
	n := len(weights) // 物品数量

	// 创建一个二维DP表，dp[i][j]表示考虑前i个物品，背包容量为j时的最大价值
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	// 填充DP表
	for i := 1; i <= n; i++ {
		for j := 0; j <= capacity; j++ {
			// 当前物品的索引（因为物品索引从0开始）
			currItem := i - 1

			// 如果当前物品的重量大于当前容量，则不能选择该物品
			if weights[currItem] > j {
				dp[i][j] = dp[i-1][j] // 继承不选择当前物品的最大价值
			} else {
				// 选择当前物品或不选择当前物品，取较大值
				notTake := dp[i-1][j]                                   // 不选择当前物品
				take := dp[i-1][j-weights[currItem]] + values[currItem] // 选择当前物品

				if take > notTake {
					dp[i][j] = take
				} else {
					dp[i][j] = notTake
				}
			}
		}
	}

	return dp[n][capacity]
}

// Knapsack01SpaceOptimized 解决0-1背包问题（空间优化版本）
// weights: 物品的重量数组
// values: 物品的价值数组
// capacity: 背包的容量
// 返回能够装入背包的最大价值
func Knapsack01SpaceOptimized(weights []int, values []int, capacity int) int {
	n := len(weights) // 物品数量

	// 创建一个一维DP数组，dp[j]表示背包容量为j时的最大价值
	dp := make([]int, capacity+1)

	// 填充DP数组
	for i := 0; i < n; i++ {
		// 注意：必须从右向左遍历，以避免重复使用同一物品
		for j := capacity; j >= weights[i]; j-- {
			// 选择当前物品或不选择当前物品，取较大值
			notTake := dp[j]                     // 不选择当前物品
			take := dp[j-weights[i]] + values[i] // 选择当前物品

			if take > notTake {
				dp[j] = take
			}
		}
	}

	return dp[capacity]
}

// Knapsack01WithItems 解决0-1背包问题并返回选择的物品
// weights: 物品的重量数组
// values: 物品的价值数组
// capacity: 背包的容量
// 返回最大价值和选择的物品索引
func Knapsack01WithItems(weights []int, values []int, capacity int) (int, []int) {
	n := len(weights) // 物品数量

	// 创建一个二维DP表，dp[i][j]表示考虑前i个物品，背包容量为j时的最大价值
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	// 填充DP表
	for i := 1; i <= n; i++ {
		for j := 0; j <= capacity; j++ {
			currItem := i - 1 // 当前物品的索引

			if weights[currItem] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				notTake := dp[i-1][j]
				take := dp[i-1][j-weights[currItem]] + values[currItem]

				if take > notTake {
					dp[i][j] = take
				} else {
					dp[i][j] = notTake
				}
			}
		}
	}

	// 回溯找出选择的物品
	selectedItems := []int{}
	totalValue := dp[n][capacity]
	remainingCapacity := capacity

	for i := n; i > 0; i-- {
		currItem := i - 1 // 当前物品的索引

		// 如果当前状态的值不等于前一个状态的值，说明选择了当前物品
		if dp[i][remainingCapacity] != dp[i-1][remainingCapacity] {
			selectedItems = append(selectedItems, currItem)
			remainingCapacity -= weights[currItem]
		}
	}

	// 反转选择的物品列表，使其按照原始顺序排列
	for i, j := 0, len(selectedItems)-1; i < j; i, j = i+1, j-1 {
		selectedItems[i], selectedItems[j] = selectedItems[j], selectedItems[i]
	}

	return totalValue, selectedItems
}
