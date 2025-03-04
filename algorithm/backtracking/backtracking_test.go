package backtracking

import (
	"reflect"
	"sort"
	"testing"
)

func TestSubsets(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name:     "空数组",
			nums:     []int{},
			expected: [][]int{{}},
		},
		{
			name:     "单元素",
			nums:     []int{1},
			expected: [][]int{{}, {1}},
		},
		{
			name:     "标准示例",
			nums:     []int{1, 2, 3},
			expected: [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Subsets(tt.nums)

			// 由于子集可能的顺序不同，需要排序后比较
			sortNestedSlices(result)
			sortNestedSlices(tt.expected)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Subsets() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestPermute(t *testing.T) {
	tests := []struct {
		name        string
		nums        []int
		expectedLen int
	}{
		{
			name:        "空数组",
			nums:        []int{},
			expectedLen: 1, // 空数组的排列只有一种
		},
		{
			name:        "单元素",
			nums:        []int{1},
			expectedLen: 1, // 单元素数组的排列只有一种
		},
		{
			name:        "标准示例",
			nums:        []int{1, 2, 3},
			expectedLen: 6, // 3! = 6 种排列
		},
		{
			name:        "四个元素",
			nums:        []int{1, 2, 3, 4},
			expectedLen: 24, // 4! = 24 种排列
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Permute(tt.nums)

			// 检查排列的数量
			if len(result) != tt.expectedLen {
				t.Errorf("Permute() 返回 %d 个排列, 期望 %d 个", len(result), tt.expectedLen)
			}

			// 检查每个排列的长度和元素
			if len(tt.nums) > 0 {
				for _, perm := range result {
					if len(perm) != len(tt.nums) {
						t.Errorf("排列 %v 长度不正确", perm)
					}

					// 检查是否包含所有元素
					sortedPerm := make([]int, len(perm))
					copy(sortedPerm, perm)
					sort.Ints(sortedPerm)

					sortedNums := make([]int, len(tt.nums))
					copy(sortedNums, tt.nums)
					sort.Ints(sortedNums)

					if !reflect.DeepEqual(sortedPerm, sortedNums) {
						t.Errorf("排列 %v 元素不正确", perm)
					}
				}
			}
		})
	}
}

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		expected   [][]int
	}{
		{
			name:       "空数组",
			candidates: []int{},
			target:     8,
			expected:   [][]int{},
		},
		{
			name:       "标准示例1",
			candidates: []int{2, 3, 6, 7},
			target:     7,
			expected:   [][]int{{2, 2, 3}, {7}},
		},
		{
			name:       "标准示例2",
			candidates: []int{2, 3, 5},
			target:     8,
			expected:   [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			name:       "没有解",
			candidates: []int{2, 3, 5},
			target:     1,
			expected:   [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CombinationSum(tt.candidates, tt.target)

			// 由于组合可能的顺序不同，需要排序后比较
			sortNestedSlices(result)
			sortNestedSlices(tt.expected)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("CombinationSum() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestSolveNQueens(t *testing.T) {
	tests := []struct {
		name        string
		n           int
		expectedLen int
	}{
		{
			name:        "n=1",
			n:           1,
			expectedLen: 1, // 只有一种放置方法
		},
		{
			name:        "n=4",
			n:           4,
			expectedLen: 2, // 有两种解决方案
		},
		{
			name:        "n=8",
			n:           8,
			expectedLen: 92, // 有92种解决方案
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SolveNQueens(tt.n)

			// 检查结果数量
			if len(result) != tt.expectedLen {
				t.Errorf("SolveNQueens() 返回 %d 个解决方案, 期望 %d 个", len(result), tt.expectedLen)
			}

			// 验证每个解决方案
			for _, solution := range result {
				// 检查行数
				if len(solution) != tt.n {
					t.Errorf("解决方案行数 %d 不正确, 期望 %d", len(solution), tt.n)
				}

				// 验证每行只有一个皇后
				for _, row := range solution {
					if len(row) != tt.n {
						t.Errorf("解决方案列数 %d 不正确, 期望 %d", len(row), tt.n)
					}

					queenCount := 0
					for _, cell := range row {
						if cell == 'Q' {
							queenCount++
						}
					}

					if queenCount != 1 {
						t.Errorf("每行应该有且只有一个皇后, 但发现 %d 个", queenCount)
					}
				}

				// 验证解决方案是否有效
				if !isValidNQueenSolution(solution) {
					t.Errorf("解决方案无效: %v", solution)
				}
			}
		})
	}
}

// isValidNQueenSolution 验证N皇后解决方案是否有效
func isValidNQueenSolution(board []string) bool {
	n := len(board)

	// 检查列和对角线
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'Q' {
				// 检查同一列
				for row := 0; row < n; row++ {
					if row != i && board[row][j] == 'Q' {
						return false
					}
				}

				// 检查对角线（左上到右下）
				for row, col := i+1, j+1; row < n && col < n; row, col = row+1, col+1 {
					if board[row][col] == 'Q' {
						return false
					}
				}
				for row, col := i-1, j-1; row >= 0 && col >= 0; row, col = row-1, col-1 {
					if board[row][col] == 'Q' {
						return false
					}
				}

				// 检查对角线（右上到左下）
				for row, col := i+1, j-1; row < n && col >= 0; row, col = row+1, col-1 {
					if board[row][col] == 'Q' {
						return false
					}
				}
				for row, col := i-1, j+1; row >= 0 && col < n; row, col = row-1, col+1 {
					if board[row][col] == 'Q' {
						return false
					}
				}
			}
		}
	}

	return true
}

func TestWordBreak(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		wordDict []string
		expected bool
	}{
		{
			name:     "空字符串",
			s:        "",
			wordDict: []string{"leet", "code"},
			expected: true, // 空字符串可以被拆分
		},
		{
			name:     "标准示例1",
			s:        "leetcode",
			wordDict: []string{"leet", "code"},
			expected: true,
		},
		{
			name:     "标准示例2",
			s:        "applepenapple",
			wordDict: []string{"apple", "pen"},
			expected: true,
		},
		{
			name:     "无法拆分",
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			expected: false,
		},
		{
			name:     "多种拆分方式",
			s:        "pineapplepenapple",
			wordDict: []string{"apple", "pen", "applepen", "pine", "pineapple"},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := WordBreak(tt.s, tt.wordDict)
			if result != tt.expected {
				t.Errorf("WordBreak() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// sortNestedSlices 对嵌套切片进行排序，便于比较结果
func sortNestedSlices(slices [][]int) {
	// 首先对每个内部切片排序
	for i := range slices {
		sort.Ints(slices[i])
	}

	// 然后对外部切片按照内部切片的字典序排序
	sort.Slice(slices, func(i, j int) bool {
		// 比较两个切片
		a, b := slices[i], slices[j]

		// 首先比较长度
		if len(a) != len(b) {
			return len(a) < len(b)
		}

		// 长度相同，按元素依次比较
		for k := 0; k < len(a); k++ {
			if a[k] != b[k] {
				return a[k] < b[k]
			}
		}

		return false // 完全相同
	})
}
