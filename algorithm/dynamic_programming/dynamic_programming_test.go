package dynamic_programming

import (
	"testing"
)

func TestFibonacciDP(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{
			name:     "n=0",
			n:        0,
			expected: 0,
		},
		{
			name:     "n=1",
			n:        1,
			expected: 1,
		},
		{
			name:     "n=2",
			n:        2,
			expected: 1,
		},
		{
			name:     "n=5",
			n:        5,
			expected: 5,
		},
		{
			name:     "n=10",
			n:        10,
			expected: 55,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FibonacciDP(tt.n)
			if result != tt.expected {
				t.Errorf("FibonacciDP() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestFibonacciOptimized(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{
			name:     "n=0",
			n:        0,
			expected: 0,
		},
		{
			name:     "n=1",
			n:        1,
			expected: 1,
		},
		{
			name:     "n=2",
			n:        2,
			expected: 1,
		},
		{
			name:     "n=5",
			n:        5,
			expected: 5,
		},
		{
			name:     "n=10",
			n:        10,
			expected: 55,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FibonacciOptimized(tt.n)
			if result != tt.expected {
				t.Errorf("FibonacciOptimized() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{
			name:     "n=1",
			n:        1,
			expected: 1,
		},
		{
			name:     "n=2",
			n:        2,
			expected: 2,
		},
		{
			name:     "n=3",
			n:        3,
			expected: 3,
		},
		{
			name:     "n=4",
			n:        4,
			expected: 5,
		},
		{
			name:     "n=5",
			n:        5,
			expected: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ClimbStairs(tt.n)
			if result != tt.expected {
				t.Errorf("ClimbStairs() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestCoinChange(t *testing.T) {
	tests := []struct {
		name     string
		coins    []int
		amount   int
		expected int
	}{
		{
			name:     "无法凑成",
			coins:    []int{2},
			amount:   3,
			expected: -1,
		},
		{
			name:     "标准示例1",
			coins:    []int{1, 2, 5},
			amount:   11,
			expected: 3, // 5 + 5 + 1 = 11
		},
		{
			name:     "标准示例2",
			coins:    []int{2, 5, 10, 1},
			amount:   27,
			expected: 4, // 10 + 10 + 5 + 2 = 27
		},
		{
			name:     "零金额",
			coins:    []int{1, 2, 5},
			amount:   0,
			expected: 0,
		},
		{
			name:     "多种凑法",
			coins:    []int{1, 5, 10, 25},
			amount:   30,
			expected: 2, // 25 + 5 = 30 (最优解需要2个硬币)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CoinChange(tt.coins, tt.amount)
			if result != tt.expected {
				t.Errorf("CoinChange() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestLongestIncreasingSubsequence(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "空数组",
			nums:     []int{},
			expected: 0,
		},
		{
			name:     "单元素",
			nums:     []int{5},
			expected: 1,
		},
		{
			name:     "标准示例1",
			nums:     []int{10, 9, 2, 5, 3, 7, 101, 18},
			expected: 4, // [2, 3, 7, 101] 或 [2, 5, 7, 101]
		},
		{
			name:     "标准示例2",
			nums:     []int{0, 1, 0, 3, 2, 3},
			expected: 4, // [0, 1, 2, 3]
		},
		{
			name:     "单调递增",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 5, // [1, 2, 3, 4, 5]
		},
		{
			name:     "单调递减",
			nums:     []int{5, 4, 3, 2, 1},
			expected: 1, // [1] 或 [2] 或 [3] 或 [4] 或 [5]
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LongestIncreasingSubsequence(tt.nums)
			if result != tt.expected {
				t.Errorf("LongestIncreasingSubsequence() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestKnapsackProblem(t *testing.T) {
	tests := []struct {
		name      string
		weights   []int
		values    []int
		maxWeight int
		expected  int
	}{
		{
			name:      "空背包",
			weights:   []int{},
			values:    []int{},
			maxWeight: 10,
			expected:  0,
		},
		{
			name:      "零容量",
			weights:   []int{1, 2, 3},
			values:    []int{6, 10, 12},
			maxWeight: 0,
			expected:  0,
		},
		{
			name:      "标准示例1",
			weights:   []int{1, 2, 3},
			values:    []int{6, 10, 12},
			maxWeight: 5,
			expected:  22, // 选择物品1和物品3
		},
		{
			name:      "标准示例2",
			weights:   []int{2, 3, 4, 5},
			values:    []int{3, 4, 5, 6},
			maxWeight: 8,
			expected:  10, // 选择物品1和物品3
		},
		{
			name:      "单个物品",
			weights:   []int{5},
			values:    []int{10},
			maxWeight: 4,
			expected:  0, // 无法放入任何物品
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := KnapsackProblem(tt.weights, tt.values, tt.maxWeight)
			if result != tt.expected {
				t.Errorf("KnapsackProblem() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestLongestCommonSubsequence(t *testing.T) {
	tests := []struct {
		name     string
		text1    string
		text2    string
		expected int
	}{
		{
			name:     "空字符串",
			text1:    "",
			text2:    "abc",
			expected: 0,
		},
		{
			name:     "标准示例1",
			text1:    "abcde",
			text2:    "ace",
			expected: 3, // LCS: "ace"
		},
		{
			name:     "标准示例2",
			text1:    "abc",
			text2:    "abc",
			expected: 3, // LCS: "abc"
		},
		{
			name:     "标准示例3",
			text1:    "abc",
			text2:    "def",
			expected: 0, // 没有公共子序列
		},
		{
			name:     "长字符串",
			text1:    "bsbininm",
			text2:    "jmjkbkjkv",
			expected: 1, // LCS: "b" 或 "j" 或 "k"
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LongestCommonSubsequence(tt.text1, tt.text2)
			if result != tt.expected {
				t.Errorf("LongestCommonSubsequence() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestEditDistance(t *testing.T) {
	tests := []struct {
		name     string
		word1    string
		word2    string
		expected int
	}{
		{
			name:     "空字符串到空字符串",
			word1:    "",
			word2:    "",
			expected: 0,
		},
		{
			name:     "空字符串到非空字符串",
			word1:    "",
			word2:    "abc",
			expected: 3, // 插入三个字符
		},
		{
			name:     "非空字符串到空字符串",
			word1:    "abc",
			word2:    "",
			expected: 3, // 删除三个字符
		},
		{
			name:     "标准示例1",
			word1:    "horse",
			word2:    "ros",
			expected: 3, // horse -> rorse (替换h为r) -> rose (删除r) -> ros (删除e)
		},
		{
			name:     "标准示例2",
			word1:    "intention",
			word2:    "execution",
			expected: 5, // intention -> inention (删除t) -> enention (替换i为e) -> exention (替换n为x) -> exection (替换n为c) -> execution (替换t为u)
		},
		{
			name:     "相同字符串",
			word1:    "abc",
			word2:    "abc",
			expected: 0, // 不需要任何操作
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EditDistance(tt.word1, tt.word2)
			if result != tt.expected {
				t.Errorf("EditDistance() = %v, want %v", result, tt.expected)
			}
		})
	}
}
