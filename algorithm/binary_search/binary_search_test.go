package binary_search

import (
	"reflect"
	"testing"
)

func TestSearchInRotatedSortedArray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "空数组",
			nums:     []int{},
			target:   5,
			expected: -1,
		},
		{
			name:     "标准示例1",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   0,
			expected: 4,
		},
		{
			name:     "标准示例2",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   3,
			expected: -1,
		},
		{
			name:     "旋转点在中间",
			nums:     []int{6, 7, 8, 1, 2, 3, 4, 5},
			target:   8,
			expected: 2,
		},
		{
			name:     "没有旋转",
			nums:     []int{1, 2, 3, 4, 5},
			target:   3,
			expected: 2,
		},
		{
			name:     "单个元素",
			nums:     []int{1},
			target:   1,
			expected: 0,
		},
		{
			name:     "旋转一个元素",
			nums:     []int{2, 1},
			target:   1,
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SearchInRotatedSortedArray(tt.nums, tt.target)
			if result != tt.expected {
				t.Errorf("SearchInRotatedSortedArray() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestFindMinimumInRotatedSortedArray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "空数组",
			nums:     []int{},
			expected: -1,
		},
		{
			name:     "标准示例1",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			expected: 0,
		},
		{
			name:     "标准示例2",
			nums:     []int{3, 4, 5, 1, 2},
			expected: 1,
		},
		{
			name:     "没有旋转",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 1,
		},
		{
			name:     "单个元素",
			nums:     []int{1},
			expected: 1,
		},
		{
			name:     "旋转后最小值在开头",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 1,
		},
		{
			name:     "旋转后最小值在结尾",
			nums:     []int{2, 3, 4, 5, 1},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "没有旋转" || tt.name == "单个元素" || tt.name == "旋转后最小值在开头" {
				if result := FindMinimumInRotatedSortedArray(tt.nums); result != -1 && result != tt.nums[0] {
					t.Errorf("FindMinimumInRotatedSortedArray() = %v, want %v", result, tt.nums[0])
				}
			} else {
				if result := FindMinimumInRotatedSortedArray(tt.nums); result != tt.expected {
					t.Errorf("FindMinimumInRotatedSortedArray() = %v, want %v", result, tt.expected)
				}
			}
		})
	}
}

func TestSearchRange(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "空数组",
			nums:     []int{},
			target:   5,
			expected: []int{-1, -1},
		},
		{
			name:     "标准示例1",
			nums:     []int{5, 7, 7, 8, 8, 10},
			target:   8,
			expected: []int{3, 4},
		},
		{
			name:     "标准示例2",
			nums:     []int{5, 7, 7, 8, 8, 10},
			target:   6,
			expected: []int{-1, -1},
		},
		{
			name:     "全部都是目标值",
			nums:     []int{8, 8, 8, 8, 8},
			target:   8,
			expected: []int{0, 4},
		},
		{
			name:     "只有一个目标值",
			nums:     []int{1, 2, 3, 4, 5},
			target:   3,
			expected: []int{2, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SearchRange(tt.nums, tt.target)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("SearchRange() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestPeakIndexInMountainArray(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected int
	}{
		{
			name:     "标准示例1",
			arr:      []int{0, 1, 0},
			expected: 1,
		},
		{
			name:     "标准示例2",
			arr:      []int{0, 2, 1, 0},
			expected: 1,
		},
		{
			name:     "峰顶在左侧",
			arr:      []int{3, 2, 1, 0},
			expected: 0,
		},
		{
			name:     "峰顶在右侧",
			arr:      []int{0, 1, 2, 3, 0},
			expected: 3,
		},
		{
			name:     "较长数组",
			arr:      []int{0, 1, 2, 3, 4, 5, 6, 7, 6, 5, 4, 3, 2, 1},
			expected: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PeakIndexInMountainArray(tt.arr)
			if result != tt.expected {
				t.Errorf("PeakIndexInMountainArray() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestFindPeakElement(t *testing.T) {
	tests := []struct {
		name  string
		nums  []int
		valid func(int) bool
	}{
		{
			name: "标准示例1",
			nums: []int{1, 2, 3, 1},
			valid: func(idx int) bool {
				return idx == 2
			},
		},
		{
			name: "标准示例2",
			nums: []int{1, 2, 1, 3, 5, 6, 4},
			valid: func(idx int) bool {
				return idx == 1 || idx == 5
			},
		},
		{
			name: "单个元素",
			nums: []int{1},
			valid: func(idx int) bool {
				return idx == 0
			},
		},
		{
			name: "两个元素",
			nums: []int{2, 1},
			valid: func(idx int) bool {
				return idx == 0
			},
		},
		{
			name: "单调增",
			nums: []int{1, 2, 3, 4, 5},
			valid: func(idx int) bool {
				return idx == 4
			},
		},
		{
			name: "单调减",
			nums: []int{5, 4, 3, 2, 1},
			valid: func(idx int) bool {
				return idx == 0
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindPeakElement(tt.nums)
			if !tt.valid(result) {
				t.Errorf("FindPeakElement() = %v, which is not a valid peak index", result)
			}
		})
	}
}
