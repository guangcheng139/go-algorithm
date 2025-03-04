package linkedlist

// 反转链表：反转一个单链表。
// 时间复杂度：O(n)，空间复杂度：O(1)
func ReverseList(head *ListNode) *ListNode {
	// 处理空链表的情况
	if head == nil {
		return nil
	}

	var prev *ListNode
	curr := head

	for curr != nil {
		// 保存下一个节点
		next := curr.Next
		// 反转当前节点的指针
		curr.Next = prev
		// 移动prev和curr指针
		prev = curr
		curr = next
	}

	// 返回新的头节点
	return prev
}

// 合并两个有序链表：将两个升序链表合并为一个新的升序链表并返回。
// 时间复杂度：O(n+m)，其中n和m分别是两个链表的长度
// 空间复杂度：O(1)，只使用常数级别的空间
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 处理边界情况
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	// 创建哑节点作为合并链表的头
	dummy := &ListNode{Val: 0}
	curr := dummy

	// 同时遍历两个链表
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}

	// 将剩余部分连接到结果链表
	if l1 != nil {
		curr.Next = l1
	} else {
		curr.Next = l2
	}

	return dummy.Next
}

// 环形链表：给定一个链表，判断链表中是否有环。
// 如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。
// 时间复杂度：O(n)，空间复杂度：O(1)
func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	// 使用快慢指针
	slow, fast := head, head.Next

	for slow != fast {
		// 如果快指针到达链表尾部，说明没有环
		if fast == nil || fast.Next == nil {
			return false
		}

		// 移动指针，慢指针每次移动一步，快指针每次移动两步
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 如果循环结束但没有返回false，说明slow与fast相遇，存在环
	return true
}

// 删除链表的倒数第N个节点：给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
// 时间复杂度：O(n)，空间复杂度：O(1)
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	// 处理空链表情况
	if head == nil {
		return nil
	}

	// 创建哑节点指向头节点
	dummy := &ListNode{Val: 0, Next: head}

	// 使用快慢指针
	first, second := dummy, dummy

	// 先将first指针向前移动n+1步
	for i := 0; i <= n; i++ {
		// 处理n大于链表长度的情况
		if first == nil {
			return head
		}
		first = first.Next
	}

	// 同时移动两个指针，直到first到达链表尾部
	for first != nil {
		first = first.Next
		second = second.Next
	}

	// 此时second指向要删除节点的前一个节点
	second.Next = second.Next.Next

	return dummy.Next
}

// 找到链表的中间节点：给定一个头结点为 head 的非空单链表，返回链表的中间结点。
// 如果有两个中间结点，则返回第二个中间结点。
// 时间复杂度：O(n)，空间复杂度：O(1)
func MiddleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 使用快慢指针
	slow, fast := head, head

	// 当快指针到达尾部，慢指针正好在中间
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

// 相交链表：找到两个单链表相交的起始节点。
// 时间复杂度：O(m+n)，其中m和n分别是两个链表的长度
// 空间复杂度：O(1)
func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	// 创建两个指针
	pA, pB := headA, headB

	// 两个指针分别遍历完两个链表，交替遍历
	// 如果存在相交节点，两个指针最终会在相交节点相遇
	// 如果不存在相交节点，两个指针都会遍历到尾部（都为nil）
	for pA != pB {
		// 当pA到达链表A的尾部时，将其指向链表B的头部
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}

		// 当pB到达链表B的尾部时，将其指向链表A的头部
		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}

	return pA // pA可能为相交节点，也可能为nil（无相交节点）
}

// 回文链表：判断一个链表是否为回文链表。
// 时间复杂度：O(n)，空间复杂度：O(1)
func IsPalindrome(head *ListNode) bool {
	// 处理空链表和单节点链表
	if head == nil || head.Next == nil {
		return true
	}

	// 1. 找到链表的中间节点
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 2. 反转后半部分链表
	secondHalf := reverseListLocal(slow.Next)

	// 3. 比较前半部分和反转后的后半部分
	result := true
	firstHalf := head
	secondHalfPointer := secondHalf

	for secondHalfPointer != nil {
		if firstHalf.Val != secondHalfPointer.Val {
			result = false
			break
		}
		firstHalf = firstHalf.Next
		secondHalfPointer = secondHalfPointer.Next
	}

	// 4. 恢复链表（可选）
	slow.Next = reverseListLocal(secondHalf)

	return result
}

// 本地辅助函数，用于反转链表
func reverseListLocal(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
