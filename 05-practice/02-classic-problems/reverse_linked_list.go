package classic

// ListNode 表示链表中的一个节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// ReverseList 反转单链表
// 返回反转后的链表头节点
func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head

	for current != nil {
		// 保存下一个节点
		next := current.Next
		// 反转当前节点的指针
		current.Next = prev
		// 移动prev和current指针
		prev = current
		current = next
	}

	// prev现在指向新的头节点
	return prev
}

// ReverseListRecursive 递归反转单链表
// 返回反转后的链表头节点
func ReverseListRecursive(head *ListNode) *ListNode {
	// 基本情况：空链表或只有一个节点
	if head == nil || head.Next == nil {
		return head
	}

	// 递归反转剩余部分
	newHead := ReverseListRecursive(head.Next)

	// 反转当前节点与下一个节点的关系
	head.Next.Next = head
	head.Next = nil

	// 返回新的头节点
	return newHead
}

// ReverseBetween 反转链表的一部分
// 反转从位置left到位置right的部分（索引从1开始）
// 返回反转后的链表头节点
func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	// 如果链表为空或不需要反转
	if head == nil || left == right {
		return head
	}

	// 创建一个哑节点，简化边界情况处理
	dummy := &ListNode{0, head}
	prev := dummy

	// 移动到反转部分的前一个节点
	for i := 1; i < left; i++ {
		prev = prev.Next
	}

	// 反转部分的开始节点
	start := prev.Next
	// 反转部分的当前节点
	current := start.Next

	// 反转从left+1到right的节点
	for i := left; i < right; i++ {
		start.Next = current.Next
		current.Next = prev.Next
		prev.Next = current
		current = start.Next
	}

	// 返回原始链表的头节点
	return dummy.Next
}

// ReverseKGroup K个一组反转链表
// 返回反转后的链表头节点
func ReverseKGroup(head *ListNode, k int) *ListNode {
	// 如果链表为空或k=1（不需要反转）
	if head == nil || k == 1 {
		return head
	}

	// 创建一个哑节点，简化边界情况处理
	dummy := &ListNode{0, head}
	prev := dummy
	current := head
	count := 0

	// 计算链表长度
	for current != nil {
		count++
		current = current.Next
	}

	// 重置current指针
	current = head

	// 对每组k个节点进行反转
	for count >= k {
		// 当前组的第一个节点
		first := current
		// 前一个节点
		var prevNode *ListNode

		// 反转k个节点
		for i := 0; i < k; i++ {
			next := current.Next
			current.Next = prevNode
			prevNode = current
			current = next
		}

		// 连接反转后的组
		prev.Next = prevNode
		first.Next = current

		// 更新prev指针
		prev = first
		count -= k
	}

	return dummy.Next
}

// HasCycle 检测链表中是否有环
func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	// 使用快慢指针
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next      // 慢指针每次移动一步
		fast = fast.Next.Next // 快指针每次移动两步

		// 如果快慢指针相遇，说明有环
		if slow == fast {
			return true
		}
	}

	// 如果快指针到达链表末尾，说明没有环
	return false
}

// DetectCycle 检测链表中的环，并返回环的入口节点
// 如果没有环，返回nil
func DetectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	// 使用快慢指针检测是否有环
	slow := head
	fast := head

	// 第一阶段：检测是否有环
	hasCycle := false
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			hasCycle = true
			break
		}
	}

	// 如果没有环，返回nil
	if !hasCycle {
		return nil
	}

	// 第二阶段：找到环的入口
	// 将慢指针重置到头节点，快指针保持在相遇点
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	// 当两个指针再次相遇时，相遇点就是环的入口
	return slow
}
