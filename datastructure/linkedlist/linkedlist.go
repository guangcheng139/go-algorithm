package linkedlist

// ListNode 链表节点定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// 创建一个链表，从切片构建
func CreateLinkedList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	// 创建头节点
	head := &ListNode{Val: vals[0]}
	curr := head

	// 构建链表
	for i := 1; i < len(vals); i++ {
		curr.Next = &ListNode{Val: vals[i]}
		curr = curr.Next
	}

	return head
}

// 将链表转换为切片
func LinkedListToSlice(head *ListNode) []int {
	// 处理空链表
	if head == nil {
		return []int{}
	}

	var result []int
	curr := head

	for curr != nil {
		result = append(result, curr.Val)
		curr = curr.Next
	}

	return result
}

// 打印链表内容
func PrintLinkedList(head *ListNode) string {
	var result string
	curr := head

	for curr != nil {
		if curr != head {
			result += " -> "
		}
		result += string(rune('0' + curr.Val))
		curr = curr.Next
	}

	return result
}
