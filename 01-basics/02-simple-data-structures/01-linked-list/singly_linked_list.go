package linkedlist

// Node 表示链表中的一个节点
type Node struct {
	Value int
	Next  *Node
}

// SinglyLinkedList 表示单链表
type SinglyLinkedList struct {
	Head *Node
	Size int
}

// NewSinglyLinkedList 创建一个新的单链表
func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{
		Head: nil,
		Size: 0,
	}
}

// InsertAtBeginning 在链表开头插入一个新节点
func (l *SinglyLinkedList) InsertAtBeginning(value int) {
	newNode := &Node{
		Value: value,
		Next:  l.Head,
	}
	l.Head = newNode
	l.Size++
}

// InsertAtEnd 在链表末尾插入一个新节点
func (l *SinglyLinkedList) InsertAtEnd(value int) {
	newNode := &Node{
		Value: value,
		Next:  nil,
	}

	if l.Head == nil {
		l.Head = newNode
		l.Size++
		return
	}

	current := l.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
	l.Size++
}

// InsertAt 在指定位置插入一个新节点
func (l *SinglyLinkedList) InsertAt(position int, value int) bool {
	// 如果位置无效，返回false
	if position < 0 || position > l.Size {
		return false
	}

	// 如果在开头插入
	if position == 0 {
		l.InsertAtBeginning(value)
		return true
	}

	// 如果在末尾插入
	if position == l.Size {
		l.InsertAtEnd(value)
		return true
	}

	// 在中间位置插入
	current := l.Head
	for i := 0; i < position-1; i++ {
		current = current.Next
	}

	newNode := &Node{
		Value: value,
		Next:  current.Next,
	}
	current.Next = newNode
	l.Size++
	return true
}

// DeleteFromBeginning 删除链表开头的节点
func (l *SinglyLinkedList) DeleteFromBeginning() bool {
	if l.Head == nil {
		return false
	}

	l.Head = l.Head.Next
	l.Size--
	return true
}

// DeleteFromEnd 删除链表末尾的节点
func (l *SinglyLinkedList) DeleteFromEnd() bool {
	if l.Head == nil {
		return false
	}

	// 如果只有一个节点
	if l.Head.Next == nil {
		l.Head = nil
		l.Size--
		return true
	}

	current := l.Head
	for current.Next.Next != nil {
		current = current.Next
	}

	current.Next = nil
	l.Size--
	return true
}

// DeleteAt 删除指定位置的节点
func (l *SinglyLinkedList) DeleteAt(position int) bool {
	// 如果位置无效，返回false
	if position < 0 || position >= l.Size || l.Head == nil {
		return false
	}

	// 如果删除开头的节点
	if position == 0 {
		return l.DeleteFromBeginning()
	}

	// 如果删除末尾的节点
	if position == l.Size-1 {
		return l.DeleteFromEnd()
	}

	// 删除中间位置的节点
	current := l.Head
	for i := 0; i < position-1; i++ {
		current = current.Next
	}

	current.Next = current.Next.Next
	l.Size--
	return true
}

// Search 在链表中查找指定值，返回其位置（如果找到）或-1（如果未找到）
func (l *SinglyLinkedList) Search(value int) int {
	if l.Head == nil {
		return -1
	}

	current := l.Head
	position := 0

	for current != nil {
		if current.Value == value {
			return position
		}
		current = current.Next
		position++
	}

	return -1
}

// GetSize 返回链表的大小
func (l *SinglyLinkedList) GetSize() int {
	return l.Size
}

// Display 打印链表的所有元素
func (l *SinglyLinkedList) Display() []int {
	var elements []int
	current := l.Head

	for current != nil {
		elements = append(elements, current.Value)
		current = current.Next
	}

	return elements
}
