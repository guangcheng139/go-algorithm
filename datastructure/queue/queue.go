package queue

// 使用切片实现的通用队列
type Queue struct {
	items []interface{}
}

// 创建一个新的空队列
func NewQueue() *Queue {
	return &Queue{
		items: []interface{}{},
	}
}

// 将元素入队（添加到队列末尾）
// 时间复杂度：O(1)
func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

// 元素出队（从队列头部移除）并返回
// 如果队列为空，返回nil
// 时间复杂度：最坏O(n)，因为需要移动元素，但可以用环形缓冲区优化
func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// 查看队列头部元素但不移除
// 如果队列为空，返回nil
// 时间复杂度：O(1)
func (q *Queue) Peek() interface{} {
	if q.IsEmpty() {
		return nil
	}

	return q.items[0]
}

// 检查队列是否为空
// 时间复杂度：O(1)
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// 返回队列中元素数量
// 时间复杂度：O(1)
func (q *Queue) Size() int {
	return len(q.items)
}

// 清空队列
// 时间复杂度：O(1)
func (q *Queue) Clear() {
	q.items = []interface{}{}
}

// 将队列转换为切片
// 时间复杂度：O(n)
func (q *Queue) ToSlice() []interface{} {
	result := make([]interface{}, len(q.items))
	copy(result, q.items)
	return result
}

// 整数队列 - 特定类型的队列，仅用于整数
type IntQueue struct {
	items []int
}

// 创建一个新的整数队列
func NewIntQueue() *IntQueue {
	return &IntQueue{
		items: []int{},
	}
}

// 将整数入队
func (q *IntQueue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// 整数出队并返回
// 如果队列为空，返回0和false
func (q *IntQueue) Dequeue() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

// 查看队列头部整数但不移除
// 如果队列为空，返回0和false
func (q *IntQueue) Peek() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}

	return q.items[0], true
}

// 检查整数队列是否为空
func (q *IntQueue) IsEmpty() bool {
	return len(q.items) == 0
}

// 返回整数队列中元素数量
func (q *IntQueue) Size() int {
	return len(q.items)
}

// 环形队列的实现
// 环形队列是一个定长队列，可以充分利用空间
type CircularQueue struct {
	items    []interface{} // 存储元素的数组
	front    int           // 队列头部索引
	rear     int           // 队列尾部索引
	size     int           // 当前元素数量
	capacity int           // 队列容量
}

// 创建一个新的环形队列，指定容量
func NewCircularQueue(capacity int) *CircularQueue {
	return &CircularQueue{
		items:    make([]interface{}, capacity),
		front:    0,
		rear:     -1,
		size:     0,
		capacity: capacity,
	}
}

// 将元素入队到环形队列
// 如果队列已满，返回false
// 时间复杂度：O(1)
func (q *CircularQueue) Enqueue(item interface{}) bool {
	if q.IsFull() {
		return false
	}

	q.rear = (q.rear + 1) % q.capacity
	q.items[q.rear] = item
	q.size++
	return true
}

// 从环形队列中出队元素
// 如果队列为空，返回nil和false
// 时间复杂度：O(1)
func (q *CircularQueue) Dequeue() (interface{}, bool) {
	if q.IsEmpty() {
		return nil, false
	}

	item := q.items[q.front]
	q.front = (q.front + 1) % q.capacity
	q.size--
	return item, true
}

// 查看环形队列头部元素
// 如果队列为空，返回nil和false
// 时间复杂度：O(1)
func (q *CircularQueue) Peek() (interface{}, bool) {
	if q.IsEmpty() {
		return nil, false
	}

	return q.items[q.front], true
}

// 检查环形队列是否为空
// 时间复杂度：O(1)
func (q *CircularQueue) IsEmpty() bool {
	return q.size == 0
}

// 检查环形队列是否已满
// 时间复杂度：O(1)
func (q *CircularQueue) IsFull() bool {
	return q.size == q.capacity
}

// 返回环形队列中的元素数量
// 时间复杂度：O(1)
func (q *CircularQueue) Size() int {
	return q.size
}

// 获取环形队列的容量
// 时间复杂度：O(1)
func (q *CircularQueue) Capacity() int {
	return q.capacity
}
