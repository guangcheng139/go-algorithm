package backtracking

// Subsets 子集
// 给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
// 说明：解集不能包含重复的子集。
// 时间复杂度: O(n * 2^n)
// 空间复杂度: O(n * 2^n)
func Subsets(nums []int) [][]int {
	result := [][]int{}
	backtrackSubsets(nums, 0, []int{}, &result)
	return result
}

// backtrackSubsets 回溯生成子集
func backtrackSubsets(nums []int, start int, current []int, result *[][]int) {
	// 将当前组合添加到结果中
	temp := make([]int, len(current))
	copy(temp, current)
	*result = append(*result, temp)

	// 从start开始尝试，避免重复
	for i := start; i < len(nums); i++ {
		// 选择当前数字
		current = append(current, nums[i])

		// 递归生成子集
		backtrackSubsets(nums, i+1, current, result)

		// 回溯，移除最后一个元素
		current = current[:len(current)-1]
	}
}

// Permute 全排列
// 给定一个不含重复数字的数组 nums ，返回其所有可能的全排列。
// 时间复杂度: O(n!)
// 空间复杂度: O(n * n!)
func Permute(nums []int) [][]int {
	result := [][]int{}
	backtrackPermute(nums, 0, &result)
	return result
}

// backtrackPermute 回溯生成排列
func backtrackPermute(nums []int, index int, result *[][]int) {
	// 达到末尾，添加当前排列
	if index == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, nums)
		*result = append(*result, temp)
		return
	}

	// 从当前位置开始，交换每个数到当前位置
	for i := index; i < len(nums); i++ {
		// 交换
		nums[index], nums[i] = nums[i], nums[index]

		// 递归处理下一个位置
		backtrackPermute(nums, index+1, result)

		// 回溯，还原交换
		nums[index], nums[i] = nums[i], nums[index]
	}
}

// CombinationSum 组合总和
// 给定一个无重复元素的数组 candidates 和一个目标数 target，找出 candidates 中所有可以使数字和为 target 的组合。
// candidates 中的数字可以无限制重复被选取。
// 时间复杂度: O(n^(target/min(candidates)))
// 空间复杂度: O(target/min(candidates))
func CombinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	backtrackCombinationSum(candidates, target, 0, []int{}, &result)
	return result
}

// backtrackCombinationSum 回溯生成组合总和
func backtrackCombinationSum(candidates []int, target int, start int, current []int, result *[][]int) {
	// 找到一个有效组合
	if target == 0 {
		temp := make([]int, len(current))
		copy(temp, current)
		*result = append(*result, temp)
		return
	}

	// 剪枝: 目标值小于0，无解
	if target < 0 {
		return
	}

	// 尝试从start开始的每个候选数
	for i := start; i < len(candidates); i++ {
		// 选择当前数字
		current = append(current, candidates[i])

		// 递归，注意可以重复使用同一个数字，所以下一次起点仍为i
		backtrackCombinationSum(candidates, target-candidates[i], i, current, result)

		// 回溯，移除最后一个元素
		current = current[:len(current)-1]
	}
}

// SolveNQueens N皇后问题
// n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
// 给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。
// 每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
// 时间复杂度: O(n!)
// 空间复杂度: O(n)
func SolveNQueens(n int) [][]string {
	result := [][]string{}
	board := make([][]byte, n)

	// 初始化棋盘
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}

	backtrackNQueens(board, 0, &result)
	return result
}

// backtrackNQueens 回溯解决N皇后问题
func backtrackNQueens(board [][]byte, row int, result *[][]string) {
	// 所有行都放置好了，添加结果
	if row == len(board) {
		solution := make([]string, len(board))
		for i := 0; i < len(board); i++ {
			solution[i] = string(board[i])
		}
		*result = append(*result, solution)
		return
	}

	// 尝试在当前行的每一列放置皇后
	for col := 0; col < len(board); col++ {
		// 检查是否可以在当前位置放置皇后
		if isValid(board, row, col) {
			// 放置皇后
			board[row][col] = 'Q'

			// 递归放置下一行
			backtrackNQueens(board, row+1, result)

			// 回溯，移除皇后
			board[row][col] = '.'
		}
	}
}

// isValid 检查在(row, col)处放置皇后是否有效
func isValid(board [][]byte, row, col int) bool {
	n := len(board)

	// 检查列是否有冲突
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}

	// 检查左上对角线是否有冲突
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	// 检查右上对角线是否有冲突
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	return true
}

// WordBreak 单词拆分
// 给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。
// 说明:
// - 拆分时可以重复使用字典中的单词。
// - 你可以假设字典中没有重复的单词。
// 时间复杂度: O(n^2 * m)，其中 n 是字符串长度，m 是检查单词是否在字典中的复杂度
// 空间复杂度: O(n)
func WordBreak(s string, wordDict []string) bool {
	// 将单词字典转换为哈希表，便于快速查找
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	// 记忆化搜索，避免重复计算
	memo := make(map[int]bool)

	// 从字符串开头开始尝试拆分
	return wordBreakHelper(s, wordSet, 0, memo)
}

// wordBreakHelper 回溯辅助函数
func wordBreakHelper(s string, wordSet map[string]bool, start int, memo map[int]bool) bool {
	// 已经处理到字符串末尾，成功拆分
	if start == len(s) {
		return true
	}

	// 检查当前起点是否已经计算过
	if val, exists := memo[start]; exists {
		return val
	}

	// 尝试从当前位置开始的每个可能的单词
	for end := start + 1; end <= len(s); end++ {
		word := s[start:end]

		// 如果当前子串是字典中的单词，并且剩余部分也可以被拆分
		if wordSet[word] && wordBreakHelper(s, wordSet, end, memo) {
			memo[start] = true
			return true
		}
	}

	// 无法拆分
	memo[start] = false
	return false
}
