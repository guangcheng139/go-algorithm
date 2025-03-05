package search

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{
			name:     "空数组",
			arr:      []int{},
			target:   5,
			expected: -1,
		},
		{
			name:     "目标值存在（开头）",
			arr:      []int{1, 2, 3, 4, 5},
			target:   1,
			expected: 0,
		},
		{
			name:     "目标值存在（中间）",
			arr:      []int{1, 2, 3, 4, 5},
			target:   3,
			expected: 2,
		},
		{
			name:     "目标值存在（结尾）",
			arr:      []int{1, 2, 3, 4, 5},
			target:   5,
			expected: 4,
		},
		{
			name:     "目标值不存在（小于最小值）",
			arr:      []int{1, 2, 3, 4, 5},
			target:   0,
			expected: -1,
		},
		{
			name:     "目标值不存在（大于最大值）",
			arr:      []int{1, 2, 3, 4, 5},
			target:   6,
			expected: -1,
		},
		{
			name:     "目标值不存在（在范围内）",
			arr:      []int{1, 2, 4, 5},
			target:   3,
			expected: -1,
		},
		{
			name:     "有重复元素",
			arr:      []int{1, 2, 2, 3, 4, 5},
			target:   2,
			expected: 1, // 或 2，取决于实现
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BinarySearch(tt.arr, tt.target)
			// 对于重复元素，我们只检查是否找到了目标值，而不关心返回哪个索引
			if tt.name == "有重复元素" {
				if (result == 1 || result == 2) != (tt.expected == 1 || tt.expected == 2) {
					t.Errorf("BinarySearch() = %v, want %v or %v", result, 1, 2)
				}
			} else if result != tt.expected {
				t.Errorf("BinarySearch() = %v, want %v", result, tt.expected)
			}
		})
	}

	// 测试递归版本
	for _, tt := range tests {
		t.Run("Recursive_"+tt.name, func(t *testing.T) {
			result := BinarySearchRecursive(tt.arr, tt.target)
			// 对于重复元素，我们只检查是否找到了目标值，而不关心返回哪个索引
			if tt.name == "有重复元素" {
				if (result == 1 || result == 2) != (tt.expected == 1 || tt.expected == 2) {
					t.Errorf("BinarySearchRecursive() = %v, want %v or %v", result, 1, 2)
				}
			} else if result != tt.expected {
				t.Errorf("BinarySearchRecursive() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestFirstOccurrence(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{
			name:     "目标值存在（单个）",
			arr:      []int{1, 2, 3, 4, 5},
			target:   3,
			expected: 2,
		},
		{
			name:     "目标值存在（重复）",
			arr:      []int{1, 2, 2, 2, 3, 4, 5},
			target:   2,
			expected: 1,
		},
		{
			name:     "目标值存在（重复）",
			arr:      []int{5, 7, 7, 8, 8, 10},
			target:   8,
			expected: 3,
		},
		{
			name:     "目标值不存在",
			arr:      []int{1, 2, 4, 5},
			target:   3,
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FirstOccurrence(tt.arr, tt.target)
			if result != tt.expected {
				t.Errorf("FirstOccurrence() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestLastOccurrence(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{
			name:     "目标值存在（单个）",
			arr:      []int{1, 2, 3, 4, 5},
			target:   3,
			expected: 2,
		},
		{
			name:     "目标值存在（重复）",
			arr:      []int{1, 2, 2, 2, 3, 4, 5},
			target:   2,
			expected: 3,
		},
		{
			name:     "目标值存在（重复）",
			arr:      []int{5, 7, 7, 8, 8, 10},
			target:   8,
			expected: 4,
		},
		{
			name:     "目标值不存在",
			arr:      []int{1, 2, 4, 5},
			target:   3,
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LastOccurrence(tt.arr, tt.target)
			if result != tt.expected {
				t.Errorf("LastOccurrence() = %v, want %v", result, tt.expected)
			}
		})
	}
}
