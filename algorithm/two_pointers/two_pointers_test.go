package two_pointers

import (
	"reflect"
	"testing"
)

// 测试排序数组的两数之和
func TestTwoSumSorted(t *testing.T) {
	tests := []struct {
		numbers  []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 3, 4}, 6, []int{1, 3}},
		{[]int{-1, 0}, -1, []int{1, 2}},
		{[]int{1, 2, 3, 4, 5}, 10, []int{-1, -1}},
	}

	for i, test := range tests {
		result := TwoSumSorted(test.numbers, test.target)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("测试%d失败: TwoSumSorted(%v, %d) = %v，期望值 %v",
				i, test.numbers, test.target, result, test.expected)
		}
	}
}

// 测试删除排序数组中的重复项
func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		nums           []int
		expectedLen    int
		expectedSubarr []int
	}{
		{[]int{1, 1, 2}, 2, []int{1, 2}},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5, []int{0, 1, 2, 3, 4}},
		{[]int{}, 0, []int{}},
		{[]int{1}, 1, []int{1}},
		{[]int{1, 2, 3}, 3, []int{1, 2, 3}},
	}

	for i, test := range tests {
		nums := make([]int, len(test.nums))
		copy(nums, test.nums)

		result := RemoveDuplicates(nums)

		if result != test.expectedLen {
			t.Errorf("测试%d失败: RemoveDuplicates(%v) 返回长度 = %d，期望值 %d",
				i, test.nums, result, test.expectedLen)
		}

		// 检查原地修改后的数组
		if !reflect.DeepEqual(nums[:result], test.expectedSubarr) {
			t.Errorf("测试%d失败: RemoveDuplicates(%v) 修改后的数组 = %v，期望值 %v",
				i, test.nums, nums[:result], test.expectedSubarr)
		}
	}
}

// 测试接雨水
func TestTrap(t *testing.T) {
	tests := []struct {
		height   []int
		expected int
	}{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{[]int{4, 2, 0, 3, 2, 5}, 9},
		{[]int{}, 0},
		{[]int{1}, 0},
		{[]int{1, 2}, 0},
		{[]int{3, 2, 1, 2, 3}, 4},
	}

	for i, test := range tests {
		result := Trap(test.height)
		if result != test.expected {
			t.Errorf("测试%d失败: Trap(%v) = %d，期望值 %d",
				i, test.height, result, test.expected)
		}
	}
}

// 测试颜色分类
func TestSortColors(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{[]int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		{[]int{2, 0, 1}, []int{0, 1, 2}},
		{[]int{0}, []int{0}},
		{[]int{1}, []int{1}},
		{[]int{1, 0}, []int{0, 1}},
	}

	for i, test := range tests {
		nums := make([]int, len(test.nums))
		copy(nums, test.nums)

		SortColors(nums)

		if !reflect.DeepEqual(nums, test.expected) {
			t.Errorf("测试%d失败: SortColors(%v) = %v，期望值 %v",
				i, test.nums, nums, test.expected)
		}
	}
}

// 测试最接近的三数之和
func TestThreeSumClosest(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected int
	}{
		{[]int{-1, 2, 1, -4}, 1, 2},
		{[]int{0, 0, 0}, 1, 0},
		{[]int{1, 1, 1, 1}, 0, 3},
		{[]int{1, 2, 4, 8, 16, 32, 64, 128}, 82, 82},
	}

	for i, test := range tests {
		result := ThreeSumClosest(test.nums, test.target)
		if result != test.expected {
			t.Errorf("测试%d失败: ThreeSumClosest(%v, %d) = %d，期望值 %d",
				i, test.nums, test.target, result, test.expected)
		}
	}
}

// 测试反转字符串
func TestReverseString(t *testing.T) {
	tests := []struct {
		s        []byte
		expected []byte
	}{
		{[]byte("hello"), []byte("olleh")},
		{[]byte("Hannah"), []byte("hannaH")},
		{[]byte(""), []byte("")},
		{[]byte("a"), []byte("a")},
		{[]byte("ab"), []byte("ba")},
	}

	for i, test := range tests {
		s := make([]byte, len(test.s))
		copy(s, test.s)

		ReverseString(s)

		if !reflect.DeepEqual(s, test.expected) {
			t.Errorf("测试%d失败: ReverseString(%q) = %q，期望值 %q",
				i, test.s, s, test.expected)
		}
	}
}

// 测试回文字符串判断
func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		s        string
		expected bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{"", true},
		{" ", true},
		{".,", true},
		{"a.", true},
		{"0P", false},
	}

	for i, test := range tests {
		result := IsPalindrome(test.s)
		if result != test.expected {
			t.Errorf("测试%d失败: IsPalindrome(%q) = %v，期望值 %v",
				i, test.s, result, test.expected)
		}
	}
}
