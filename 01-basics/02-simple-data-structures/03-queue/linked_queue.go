package queue

// Node 表示队列中的一个节点
type Node struct {
	Value int
	Next  *Node
}

// LinkedQueue 表示基于链表的队列
type LinkedQueue struct {
	Front *Node // 队列前端，用于出队
	Rear  *Node // 队列后端，用于入队
	Size  int   // 队列大小
}

// NewLinkedQueue 创建一个新的基于链表的队列
func NewLinkedQueue() *LinkedQueue {
	return &LinkedQueue{
		Front: nil,
		Rear:  nil,
		Size:  0,
	}
}

// Enqueue 将元素添加到队列末尾
func (q *LinkedQueue) Enqueue(value int) {
	newNode := &Node{
		Value: value,
		Next:  nil,
	}

	// 如果队列为空
	if q.IsEmpty() {
		q.Front = newNode
		q.Rear = newNode
	} else {
		// 将新节点添加到队列末尾
		q.Rear.Next = newNode
		q.Rear = newNode
	}

	q.Size++
}

// Dequeue 从队列前端移除元素
func (q *LinkedQueue) Dequeue() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}

	value := q.Front.Value
	q.Front = q.Front.Next
	q.Size--

	// 如果队列变为空
	if q.Front == nil {
		q.Rear = nil
	}

	return value, true
}

// Peek 查看队列前端元素但不移除
func (q *LinkedQueue) Peek() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}

	return q.Front.Value, true
}

// IsEmpty 检查队列是否为空
func (q *LinkedQueue) IsEmpty() bool {
	return q.Front == nil
}

// GetSize 返回队列中元素的数量
func (q *LinkedQueue) GetSize() int {
	return q.Size
}

// Clear 清空队列
func (q *LinkedQueue) Clear() {
	q.Front = nil
	q.Rear = nil
	q.Size = 0
}

// ToSlice 将队列转换为切片（从队列前端到后端）
func (q *LinkedQueue) ToSlice() []int {
	var result []int
	current := q.Front

	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}

	return result
}
