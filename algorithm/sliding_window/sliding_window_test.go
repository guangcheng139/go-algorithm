package sliding_window

import (
	"reflect"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "空数组",
			nums:     []int{},
			k:        3,
			expected: []int{},
		},
		{
			name:     "窗口大小为1",
			nums:     []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:        1,
			expected: []int{1, 3, -1, -3, 5, 3, 6, 7},
		},
		{
			name:     "标准示例",
			nums:     []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:        3,
			expected: []int{3, 3, 5, 5, 6, 7},
		},
		{
			name:     "递减数组",
			nums:     []int{7, 6, 5, 4, 3, 2, 1},
			k:        3,
			expected: []int{7, 6, 5, 4, 3},
		},
		{
			name:     "递增数组",
			nums:     []int{1, 2, 3, 4, 5, 6, 7},
			k:        3,
			expected: []int{3, 4, 5, 6, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxSlidingWindow(tt.nums, tt.k)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("MaxSlidingWindow() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestMinSubArrayLen(t *testing.T) {
	tests := []struct {
		name     string
		target   int
		nums     []int
		expected int
	}{
		{
			name:     "空数组",
			target:   7,
			nums:     []int{},
			expected: 0,
		},
		{
			name:     "标准示例",
			target:   7,
			nums:     []int{2, 3, 1, 2, 4, 3},
			expected: 2, // 子数组 [4,3] 长度为 2
		},
		{
			name:     "单个元素满足条件",
			target:   4,
			nums:     []int{1, 4, 4},
			expected: 1, // 子数组 [4] 长度为 1
		},
		{
			name:     "没有满足条件的子数组",
			target:   11,
			nums:     []int{1, 1, 1, 1, 1, 1, 1, 1},
			expected: 0,
		},
		{
			name:     "全部元素刚好满足条件",
			target:   15,
			nums:     []int{5, 1, 3, 5, 10, 7, 4, 9, 2, 8},
			expected: 2, // 子数组 [10,7] 长度为 2
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MinSubArrayLen(tt.target, tt.nums)
			if result != tt.expected {
				t.Errorf("MinSubArrayLen() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestCharacterReplacement(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		k        int
		expected int
	}{
		{
			name:     "空字符串",
			s:        "",
			k:        2,
			expected: 0,
		},
		{
			name:     "标准示例1",
			s:        "ABAB",
			k:        2,
			expected: 4, // 可以替换两个'A'为'B'，得到"BBBB"
		},
		{
			name:     "标准示例2",
			s:        "AABABBA",
			k:        1,
			expected: 4, // 可以替换一个'A'为'B'，得到"AABBBBA"
		},
		{
			name:     "全部相同字符",
			s:        "AAAA",
			k:        2,
			expected: 4,
		},
		{
			name:     "替换次数为0",
			s:        "ABCDE",
			k:        0,
			expected: 1, // 不能替换任何字符
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CharacterReplacement(tt.s, tt.k)
			if result != tt.expected {
				t.Errorf("CharacterReplacement() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		p        string
		expected []int
	}{
		{
			name:     "s比p短",
			s:        "ab",
			p:        "abc",
			expected: []int{},
		},
		{
			name:     "标准示例1",
			s:        "cbaebabacd",
			p:        "abc",
			expected: []int{0, 6}, // "cba"和"bac"是"abc"的字母异位词
		},
		{
			name:     "标准示例2",
			s:        "abab",
			p:        "ab",
			expected: []int{0, 1, 2}, // "ab", "ba", "ab"都是"ab"的字母异位词
		},
		{
			name:     "s和p完全相同",
			s:        "abc",
			p:        "abc",
			expected: []int{0},
		},
		{
			name:     "没有异位词",
			s:        "hello",
			p:        "world",
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindAnagrams(tt.s, tt.p)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("FindAnagrams() = %v, want %v", result, tt.expected)
			}
		})
	}
}
