package linkedlist

import (
	"reflect"
	"testing"
)

// 测试链表创建和转换函数
func TestLinkedListCreation(t *testing.T) {
	// 测试用例：空链表
	emptyList := CreateLinkedList([]int{})
	if emptyList != nil {
		t.Errorf("空链表创建失败: 期望 nil, 得到 %v", emptyList)
	}

	// 测试用例：单节点链表
	singleNodeList := CreateLinkedList([]int{5})
	if singleNodeList.Val != 5 || singleNodeList.Next != nil {
		t.Errorf("单节点链表创建失败: 期望值 5, 实际 %d", singleNodeList.Val)
	}

	// 测试用例：多节点链表
	vals := []int{1, 2, 3, 4, 5}
	list := CreateLinkedList(vals)
	resultVals := LinkedListToSlice(list)

	if !reflect.DeepEqual(vals, resultVals) {
		t.Errorf("链表创建和转换失败: 期望 %v, 得到 %v", vals, resultVals)
	}
}

// 测试反转链表
func TestReverseList(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
	}

	for i, test := range tests {
		head := CreateLinkedList(test.input)
		reversed := ReverseList(head)
		result := LinkedListToSlice(reversed)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("测试%d失败: ReverseList(%v) = %v, 期望 %v",
				i, test.input, result, test.expected)
		}
	}
}

// 测试合并两个有序链表
func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		list1    []int
		list2    []int
		expected []int
	}{
		{[]int{}, []int{}, []int{}},
		{[]int{}, []int{1, 3, 5}, []int{1, 3, 5}},
		{[]int{2, 4, 6}, []int{}, []int{2, 4, 6}},
		{[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
		{[]int{1, 3, 5, 7}, []int{2, 4, 6, 8}, []int{1, 2, 3, 4, 5, 6, 7, 8}},
	}

	for i, test := range tests {
		l1 := CreateLinkedList(test.list1)
		l2 := CreateLinkedList(test.list2)
		merged := MergeTwoLists(l1, l2)
		result := LinkedListToSlice(merged)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("测试%d失败: MergeTwoLists(%v, %v) = %v, 期望 %v",
				i, test.list1, test.list2, result, test.expected)
		}
	}
}

// 测试删除链表的倒数第N个节点
func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		input    []int
		n        int
		expected []int
	}{
		{[]int{1}, 1, []int{}},
		{[]int{1, 2}, 1, []int{1}},
		{[]int{1, 2}, 2, []int{2}},
		{[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
		{[]int{1, 2, 3, 4, 5}, 5, []int{2, 3, 4, 5}},
	}

	for i, test := range tests {
		head := CreateLinkedList(test.input)
		removed := RemoveNthFromEnd(head, test.n)
		result := LinkedListToSlice(removed)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("测试%d失败: RemoveNthFromEnd(%v, %d) = %v, 期望 %v",
				i, test.input, test.n, result, test.expected)
		}
	}
}

// 测试找到链表的中间节点
func TestMiddleNode(t *testing.T) {
	tests := []struct {
		input    []int
		expected int // 中间节点的值
	}{
		{[]int{1}, 1},
		{[]int{1, 2}, 2},
		{[]int{1, 2, 3}, 2},
		{[]int{1, 2, 3, 4}, 3},
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{1, 2, 3, 4, 5, 6}, 4},
	}

	for i, test := range tests {
		head := CreateLinkedList(test.input)
		middle := MiddleNode(head)

		if middle == nil {
			t.Errorf("测试%d失败: MiddleNode(%v) = nil, 期望 %d",
				i, test.input, test.expected)
			continue
		}

		if middle.Val != test.expected {
			t.Errorf("测试%d失败: MiddleNode(%v).Val = %d, 期望 %d",
				i, test.input, middle.Val, test.expected)
		}
	}
}

// 测试回文链表判断
func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{[]int{}, true},
		{[]int{1}, true},
		{[]int{1, 2}, false},
		{[]int{1, 2, 1}, true},
		{[]int{1, 2, 2, 1}, true},
		{[]int{1, 2, 3, 2, 1}, true},
		{[]int{1, 2, 3, 4, 5}, false},
	}

	for i, test := range tests {
		head := CreateLinkedList(test.input)
		result := IsPalindrome(head)

		if result != test.expected {
			t.Errorf("测试%d失败: IsPalindrome(%v) = %v, 期望 %v",
				i, test.input, result, test.expected)
		}

		// 对于空链表或单节点链表，不检查恢复
		if len(test.input) <= 1 {
			continue
		}

		// 检查链表是否恢复原状
		restored := LinkedListToSlice(head)
		if !reflect.DeepEqual(restored, test.input) {
			t.Errorf("测试%d失败: 链表未恢复原状 %v, 期望 %v",
				i, restored, test.input)
		}
	}
}

// 不易测试的函数:
// 1. HasCycle - 需要创建循环链表
// 2. GetIntersectionNode - 需要创建相交链表
// 这些函数需要特殊的测试设置
