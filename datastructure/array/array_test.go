package array

import (
	"reflect"
	"testing"
)

// 测试两数之和函数
func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{1, 2, 3, 4, 5}, 10, []int{}},
	}

	for _, tt := range tests {
		got := TwoSum(tt.nums, tt.target)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("TwoSum(%v, %v) = %v, want %v", tt.nums, tt.target, got, tt.want)
		}
	}
}

// 测试三数之和函数
func TestThreeSum(t *testing.T) {
	tests := []struct {
		nums []int
		want [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{[]int{0, 0, 0}, [][]int{{0, 0, 0}}},
		{[]int{1, 2, -2, -1}, [][]int{}},
		{[]int{}, [][]int{}},
		{[]int{0}, [][]int{}},
	}

	for _, tt := range tests {
		got := ThreeSum(tt.nums)
		// 对输入为[]，输出为[]的情况特殊处理
		if len(tt.nums) == 0 && len(got) == 0 && len(tt.want) == 0 {
			continue
		}
		// 对输入为[0]，输出为[]的情况特殊处理
		if len(tt.nums) == 1 && tt.nums[0] == 0 && len(got) == 0 && len(tt.want) == 0 {
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("ThreeSum(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}

// 测试合并两个有序数组函数
func TestMerge(t *testing.T) {
	tests := []struct {
		nums1  []int
		m      int
		nums2  []int
		n      int
		result []int
	}{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{[]int{1}, 1, []int{}, 0, []int{1}},
		{[]int{0}, 0, []int{1}, 1, []int{1}},
	}

	for _, tt := range tests {
		nums1Copy := make([]int, len(tt.nums1))
		copy(nums1Copy, tt.nums1)

		Merge(nums1Copy, tt.m, tt.nums2, tt.n)

		if !reflect.DeepEqual(nums1Copy, tt.result) {
			t.Errorf("Merge(%v, %v, %v, %v) = %v, want %v", tt.nums1, tt.m, tt.nums2, tt.n, nums1Copy, tt.result)
		}
	}
}

// 测试旋转数组函数
func TestRotate(t *testing.T) {
	tests := []struct {
		nums   []int
		k      int
		result []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
		{[]int{1, 2}, 3, []int{2, 1}},
		{[]int{1, 2, 3}, 4, []int{3, 1, 2}},
	}

	for _, tt := range tests {
		numsCopy := make([]int, len(tt.nums))
		copy(numsCopy, tt.nums)

		Rotate(numsCopy, tt.k)

		if !reflect.DeepEqual(numsCopy, tt.result) {
			t.Errorf("Rotate(%v, %v) = %v, want %v", tt.nums, tt.k, numsCopy, tt.result)
		}
	}
}

// 测试移动零函数
func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		nums   []int
		result []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0}, []int{0}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{0, 0, 1}, []int{1, 0, 0}},
	}

	for _, tt := range tests {
		numsCopy := make([]int, len(tt.nums))
		copy(numsCopy, tt.nums)

		MoveZeroes(numsCopy)

		if !reflect.DeepEqual(numsCopy, tt.result) {
			t.Errorf("MoveZeroes(%v) = %v, want %v", tt.nums, numsCopy, tt.result)
		}
	}
}

// 测试盛最多水的容器函数
func TestMaxArea(t *testing.T) {
	tests := []struct {
		height []int
		want   int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
		{[]int{4, 3, 2, 1, 4}, 16},
		{[]int{1, 2, 1}, 2},
	}

	for _, tt := range tests {
		got := MaxArea(tt.height)
		if got != tt.want {
			t.Errorf("MaxArea(%v) = %v, want %v", tt.height, got, tt.want)
		}
	}
}

// 测试最大子数组和函数
func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1}, 1},
		{[]int{5, 4, -1, 7, 8}, 23},
		{[]int{-1}, -1},
	}

	for _, tt := range tests {
		got := MaxSubArray(tt.nums)
		if got != tt.want {
			t.Errorf("MaxSubArray(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}
