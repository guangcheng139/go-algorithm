package queue

// 滑动窗口最大值：给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。
// 时间复杂度：O(n)，空间复杂度：O(k)
func MaxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k <= 0 {
		return []int{}
	}

	n := len(nums)
	result := make([]int, n-k+1)
	deque := make([]int, 0) // 存储元素索引的双端队列

	for i := 0; i < n; i++ {
		// 移除所有小于当前元素的索引，因为它们不可能是最大值
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}

		// 添加当前元素索引
		deque = append(deque, i)

		// 移除超出窗口范围的元素
		if deque[0] <= i-k {
			deque = deque[1:]
		}

		// 当窗口大小满足k时，记录当前窗口的最大值
		if i >= k-1 {
			result[i-k+1] = nums[deque[0]]
		}
	}

	return result
}

// 设计双端队列：实现一个双端队列，支持以下操作
// insertFront()、insertLast()、deleteFront()、deleteLast()、getFront()、getRear()、isEmpty()、isFull()
type Deque struct {
	items    []interface{}
	front    int
	rear     int
	size     int
	capacity int
}

// 创建一个新的双端队列，指定容量
func NewDeque(capacity int) *Deque {
	return &Deque{
		items:    make([]interface{}, capacity),
		front:    0,
		rear:     capacity - 1, // 从尾部开始，方便从前面插入
		size:     0,
		capacity: capacity,
	}
}

// 从队列前端添加元素
// 如果队列已满，返回false
// 时间复杂度：O(1)
func (d *Deque) InsertFront(item interface{}) bool {
	if d.IsFull() {
		return false
	}

	// 更新front指针，需要考虑环形结构
	d.front = (d.front - 1 + d.capacity) % d.capacity
	d.items[d.front] = item
	d.size++
	return true
}

// 从队列后端添加元素
// 如果队列已满，返回false
// 时间复杂度：O(1)
func (d *Deque) InsertLast(item interface{}) bool {
	if d.IsFull() {
		return false
	}

	// 更新rear指针，考虑环形结构
	d.rear = (d.rear + 1) % d.capacity
	d.items[d.rear] = item
	d.size++
	return true
}

// 从队列前端删除元素
// 如果队列为空，返回nil和false
// 时间复杂度：O(1)
func (d *Deque) DeleteFront() (interface{}, bool) {
	if d.IsEmpty() {
		return nil, false
	}

	item := d.items[d.front]
	d.front = (d.front + 1) % d.capacity
	d.size--
	return item, true
}

// 从队列后端删除元素
// 如果队列为空，返回nil和false
// 时间复杂度：O(1)
func (d *Deque) DeleteLast() (interface{}, bool) {
	if d.IsEmpty() {
		return nil, false
	}

	item := d.items[d.rear]
	d.rear = (d.rear - 1 + d.capacity) % d.capacity
	d.size--
	return item, true
}

// 获取队列前端元素
// 如果队列为空，返回nil和false
// 时间复杂度：O(1)
func (d *Deque) GetFront() (interface{}, bool) {
	if d.IsEmpty() {
		return nil, false
	}

	return d.items[d.front], true
}

// 获取队列后端元素
// 如果队列为空，返回nil和false
// 时间复杂度：O(1)
func (d *Deque) GetRear() (interface{}, bool) {
	if d.IsEmpty() {
		return nil, false
	}

	return d.items[d.rear], true
}

// 检查队列是否为空
// 时间复杂度：O(1)
func (d *Deque) IsEmpty() bool {
	return d.size == 0
}

// 检查队列是否已满
// 时间复杂度：O(1)
func (d *Deque) IsFull() bool {
	return d.size == d.capacity
}

// 队列中元素的数量
// 时间复杂度：O(1)
func (d *Deque) Size() int {
	return d.size
}

// 获取双端队列的容量
// 时间复杂度：O(1)
func (d *Deque) Capacity() int {
	return d.capacity
}

// 使用两个栈实现队列
type MyQueueWithStacks struct {
	stackPush []int // 用于入队操作的栈
	stackPop  []int // 用于出队操作的栈
}

// 创建一个新的队列
func NewMyQueueWithStacks() *MyQueueWithStacks {
	return &MyQueueWithStacks{
		stackPush: []int{},
		stackPop:  []int{},
	}
}

// 将元素推入队列
// 时间复杂度：O(1)
func (q *MyQueueWithStacks) Push(x int) {
	q.stackPush = append(q.stackPush, x)
}

// 从队列中弹出元素
// 时间复杂度：均摊O(1)
func (q *MyQueueWithStacks) Pop() int {
	q.moveIfNeeded()

	if len(q.stackPop) == 0 {
		return -1 // 或者错误处理
	}

	val := q.stackPop[len(q.stackPop)-1]
	q.stackPop = q.stackPop[:len(q.stackPop)-1]
	return val
}

// 查看队列头部元素
// 时间复杂度：均摊O(1)
func (q *MyQueueWithStacks) Peek() int {
	q.moveIfNeeded()

	if len(q.stackPop) == 0 {
		return -1 // 或者错误处理
	}

	return q.stackPop[len(q.stackPop)-1]
}

// 检查队列是否为空
// 时间复杂度：O(1)
func (q *MyQueueWithStacks) Empty() bool {
	return len(q.stackPush) == 0 && len(q.stackPop) == 0
}

// 将元素从Push栈移动到Pop栈
// 只有当Pop栈为空时才进行移动
func (q *MyQueueWithStacks) moveIfNeeded() {
	// 只有当Pop栈为空时，才将Push栈中的所有元素倒入Pop栈
	if len(q.stackPop) == 0 {
		for len(q.stackPush) > 0 {
			n := len(q.stackPush)
			q.stackPop = append(q.stackPop, q.stackPush[n-1])
			q.stackPush = q.stackPush[:n-1]
		}
	}
}

// 任务调度器：给定一个用字符数组表示的 CPU 需要执行的任务列表，以及一个正整数 n，表示两个相同种类的任务之间必须间隔至少 n 个时间单位，
// 计算完成所有任务所需要的最短时间。
// 时间复杂度：O(nlogn)，其中n是任务数量，主要耗时在排序
// 空间复杂度：O(1)，因为字母种类最多只有26种
func LeastInterval(tasks []byte, n int) int {
	// 统计每种任务的数量
	count := make([]int, 26)
	for _, task := range tasks {
		count[task-'A']++
	}

	// 找出最大任务数量
	maxCount := 0
	for _, c := range count {
		if c > maxCount {
			maxCount = c
		}
	}

	// 找出有多少种任务达到了最大数量
	maxCountTasks := 0
	for _, c := range count {
		if c == maxCount {
			maxCountTasks++
		}
	}

	// 计算最少需要的时间单位
	// 公式：(maxCount - 1) * (n + 1) + maxCountTasks
	// 解释：(最大任务数量 - 1) 组，每组需要 (n + 1) 个时间单位，最后一组需要 maxCountTasks 个时间单位
	result := (maxCount-1)*(n+1) + maxCountTasks

	// 如果计算结果小于任务总数，那么返回任务总数
	// 这种情况发生在间隔n很小，任务种类很多时
	if result < len(tasks) {
		return len(tasks)
	}

	return result
}
