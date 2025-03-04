package queue

import (
	"reflect"
	"testing"
)

// 测试通用队列的基本操作
func TestQueue(t *testing.T) {
	// 创建一个新队列
	q := NewQueue()

	// 测试初始状态
	if !q.IsEmpty() {
		t.Error("新创建的队列应该为空")
	}

	if q.Size() != 0 {
		t.Errorf("新创建的队列大小应该为0，得到 %d", q.Size())
	}

	// 测试入队和查看操作
	q.Enqueue(1)
	q.Enqueue("test")
	q.Enqueue(3.14)

	if q.Size() != 3 {
		t.Errorf("队列大小应该为3，得到 %d", q.Size())
	}

	if q.Peek() != 1 {
		t.Errorf("队列头部元素应该为1，得到 %v", q.Peek())
	}

	// 测试出队操作
	item := q.Dequeue()
	if item != 1 {
		t.Errorf("出队的元素应该为1，得到 %v", item)
	}

	if q.Size() != 2 {
		t.Errorf("出队后队列大小应该为2，得到 %d", q.Size())
	}

	// 测试ToSlice操作
	slice := q.ToSlice()
	expected := []interface{}{"test", 3.14}
	if !reflect.DeepEqual(slice, expected) {
		t.Errorf("ToSlice应返回 %v，得到 %v", expected, slice)
	}

	// 测试Clear操作
	q.Clear()
	if !q.IsEmpty() {
		t.Error("Clear后队列应该为空")
	}

	// 测试空队列操作
	if q.Dequeue() != nil {
		t.Error("从空队列Dequeue应该返回nil")
	}

	if q.Peek() != nil {
		t.Error("空队列Peek应该返回nil")
	}
}

// 测试整数队列的基本操作
func TestIntQueue(t *testing.T) {
	// 创建一个新的整数队列
	q := NewIntQueue()

	// 测试初始状态
	if !q.IsEmpty() {
		t.Error("新创建的整数队列应该为空")
	}

	if q.Size() != 0 {
		t.Errorf("新创建的整数队列大小应该为0，得到 %d", q.Size())
	}

	// 测试入队和查看操作
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	if q.Size() != 3 {
		t.Errorf("队列大小应该为3，得到 %d", q.Size())
	}

	val, ok := q.Peek()
	if !ok || val != 10 {
		t.Errorf("队列头部元素应该为10，得到 %v (ok=%v)", val, ok)
	}

	// 测试出队操作
	val, ok = q.Dequeue()
	if !ok || val != 10 {
		t.Errorf("出队的元素应该为10，得到 %v (ok=%v)", val, ok)
	}

	if q.Size() != 2 {
		t.Errorf("出队后队列大小应该为2，得到 %d", q.Size())
	}

	// 测试空队列操作
	q = NewIntQueue()
	val, ok = q.Dequeue()
	if ok || val != 0 {
		t.Errorf("从空队列Dequeue应该返回(0,false)，得到 (%v,%v)", val, ok)
	}

	val, ok = q.Peek()
	if ok || val != 0 {
		t.Errorf("空队列Peek应该返回(0,false)，得到 (%v,%v)", val, ok)
	}
}

// 测试环形队列的基本操作
func TestCircularQueue(t *testing.T) {
	// 创建一个容量为3的环形队列
	q := NewCircularQueue(3)

	// 测试初始状态
	if !q.IsEmpty() {
		t.Error("新创建的环形队列应该为空")
	}

	if q.IsFull() {
		t.Error("新创建的环形队列不应该满")
	}

	// 测试Enqueue和Peek操作
	if !q.Enqueue(1) {
		t.Error("入队应该成功")
	}

	if !q.Enqueue(2) {
		t.Error("入队应该成功")
	}

	val, ok := q.Peek()
	if !ok || val != 1 {
		t.Errorf("队列前端元素应该是1，得到 %v (ok=%v)", val, ok)
	}

	// 测试队列已满的情况
	if !q.Enqueue(3) {
		t.Error("入队应该成功")
	}

	if !q.IsFull() {
		t.Error("队列应该已满")
	}

	if q.Enqueue(4) {
		t.Error("队列已满，入队应该失败")
	}

	// 测试Dequeue操作
	val, ok = q.Dequeue()
	if !ok || val != 1 {
		t.Errorf("出队的元素应该为1，得到 %v (ok=%v)", val, ok)
	}

	// 测试环形特性
	if !q.Enqueue(4) {
		t.Error("出队后应该能够再次入队")
	}

	// 验证队列内容
	expected := []interface{}{2, 3, 4}
	for i, exp := range expected {
		val, ok := q.Dequeue()
		if !ok || val != exp {
			t.Errorf("第%d次出队的元素应该为%v，得到 %v (ok=%v)", i, exp, val, ok)
		}
	}

	// 测试空队列操作
	if !q.IsEmpty() {
		t.Error("所有元素出队后，队列应该为空")
	}

	_, ok = q.Dequeue()
	if ok {
		t.Error("空队列出队应该失败")
	}
}

// 测试双端队列的基本操作
func TestDeque(t *testing.T) {
	// 创建一个容量为5的双端队列
	d := NewDeque(5)

	// 测试初始状态
	if !d.IsEmpty() {
		t.Error("新创建的双端队列应该为空")
	}

	// 测试从前端插入
	if !d.InsertFront(1) {
		t.Error("前端插入应该成功")
	}

	// 测试从后端插入
	if !d.InsertLast(2) {
		t.Error("后端插入应该成功")
	}

	if d.Size() != 2 {
		t.Errorf("队列大小应该为2，得到 %d", d.Size())
	}

	// 测试获取前端和后端元素
	frontVal, ok := d.GetFront()
	if !ok || frontVal != 1 {
		t.Errorf("前端元素应该为1，得到 %v (ok=%v)", frontVal, ok)
	}

	rearVal, ok := d.GetRear()
	if !ok || rearVal != 2 {
		t.Errorf("后端元素应该为2，得到 %v (ok=%v)", rearVal, ok)
	}

	// 测试从前端删除
	frontVal, ok = d.DeleteFront()
	if !ok || frontVal != 1 {
		t.Errorf("删除的前端元素应该为1，得到 %v (ok=%v)", frontVal, ok)
	}

	// 测试从后端删除
	rearVal, ok = d.DeleteLast()
	if !ok || rearVal != 2 {
		t.Errorf("删除的后端元素应该为2，得到 %v (ok=%v)", rearVal, ok)
	}

	// 测试空队列
	if !d.IsEmpty() {
		t.Error("所有元素删除后，队列应该为空")
	}

	// 测试队列满
	for i := 0; i < d.Capacity(); i++ {
		if !d.InsertLast(i) {
			t.Errorf("第%d次插入应该成功", i)
		}
	}

	if !d.IsFull() {
		t.Error("队列应该已满")
	}

	if d.InsertFront(100) || d.InsertLast(100) {
		t.Error("队列已满，插入应该失败")
	}
}

// 测试滑动窗口最大值
func TestMaxSlidingWindow(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
		{[]int{1}, 1, []int{1}},
		{[]int{1, -1}, 1, []int{1, -1}},
		{[]int{9, 11}, 2, []int{11}},
		{[]int{4, 3, 2, 1}, 3, []int{4, 3}},
		{[]int{}, 0, []int{}},
	}

	for i, test := range tests {
		result := MaxSlidingWindow(test.nums, test.k)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("测试%d失败: MaxSlidingWindow(%v, %d) = %v，期望值 %v",
				i, test.nums, test.k, result, test.expected)
		}
	}
}

// 测试用栈实现的队列
func TestMyQueueWithStacks(t *testing.T) {
	q := NewMyQueueWithStacks()

	// 测试初始状态
	if !q.Empty() {
		t.Error("新创建的队列应该为空")
	}

	// 测试Push和Peek操作
	q.Push(1)
	q.Push(2)

	if q.Peek() != 1 {
		t.Errorf("队列头部元素应该为1，得到 %d", q.Peek())
	}

	// 测试Pop操作
	val := q.Pop()
	if val != 1 {
		t.Errorf("出队的元素应该为1，得到 %d", val)
	}

	// 测试Pop后的状态
	if q.Peek() != 2 {
		t.Errorf("出队后队列头部元素应该为2，得到 %d", q.Peek())
	}

	// 再次Push和测试
	q.Push(3)
	q.Push(4)

	// 连续Pop并验证
	expected := []int{2, 3, 4}
	for i, exp := range expected {
		val := q.Pop()
		if val != exp {
			t.Errorf("第%d次出队的元素应该为%d，得到 %d", i, exp, val)
		}
	}

	// 测试空队列状态
	if !q.Empty() {
		t.Error("所有元素出队后，队列应该为空")
	}
}

// 测试任务调度器
func TestLeastInterval(t *testing.T) {
	tests := []struct {
		tasks    []byte
		n        int
		expected int
	}{
		{[]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2, 8},
		{[]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 0, 6},
		{[]byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}, 2, 16},
		{[]byte{'A', 'B', 'C', 'D', 'E', 'A', 'B', 'C', 'D', 'E'}, 4, 10},
	}

	for i, test := range tests {
		result := LeastInterval(test.tasks, test.n)
		if result != test.expected {
			t.Errorf("测试%d失败: LeastInterval(%v, %d) = %d，期望值 %d",
				i, test.tasks, test.n, result, test.expected)
		}
	}
}
