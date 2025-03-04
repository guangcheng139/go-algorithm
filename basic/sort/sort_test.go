package sort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "空数组",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "单元素数组",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "已排序数组",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "逆序数组",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "乱序数组",
			input:    []int{3, 1, 5, 2, 4},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "有重复元素数组",
			input:    []int{3, 1, 3, 2, 5},
			expected: []int{1, 2, 3, 3, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BubbleSort(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("BubbleSort() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "空数组",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "单元素数组",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "已排序数组",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "逆序数组",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "乱序数组",
			input:    []int{3, 1, 5, 2, 4},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "有重复元素数组",
			input:    []int{3, 1, 3, 2, 5},
			expected: []int{1, 2, 3, 3, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := QuickSort(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("QuickSort() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "空数组",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "单元素数组",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "已排序数组",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "逆序数组",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "乱序数组",
			input:    []int{3, 1, 5, 2, 4},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "有重复元素数组",
			input:    []int{3, 1, 3, 2, 5},
			expected: []int{1, 2, 3, 3, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MergeSort(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("MergeSort() = %v, want %v", result, tt.expected)
			}
		})
	}
}
