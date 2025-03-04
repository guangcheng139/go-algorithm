package stack

// 使用切片实现的栈
type Stack struct {
	items []interface{}
}

// 创建一个新的空栈
func NewStack() *Stack {
	return &Stack{
		items: []interface{}{},
	}
}

// 将元素推入栈顶
// 时间复杂度：O(1)
func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

// 从栈顶弹出元素并返回
// 如果栈为空，返回nil
// 时间复杂度：O(1)
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}

	n := len(s.items)
	item := s.items[n-1]
	s.items = s.items[:n-1]
	return item
}

// 查看栈顶元素但不移除
// 如果栈为空，返回nil
// 时间复杂度：O(1)
func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}

	return s.items[len(s.items)-1]
}

// 检查栈是否为空
// 时间复杂度：O(1)
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// 返回栈中元素数量
// 时间复杂度：O(1)
func (s *Stack) Size() int {
	return len(s.items)
}

// 清空栈
// 时间复杂度：O(1)
func (s *Stack) Clear() {
	s.items = []interface{}{}
}

// 将栈中的元素转换为切片
// 时间复杂度：O(n)
func (s *Stack) ToSlice() []interface{} {
	result := make([]interface{}, len(s.items))
	copy(result, s.items)
	return result
}

// 整数栈 - 特定类型的栈，仅用于整数
type IntStack struct {
	items []int
}

// 创建一个新的整数栈
func NewIntStack() *IntStack {
	return &IntStack{
		items: []int{},
	}
}

// 将整数推入栈顶
func (s *IntStack) Push(item int) {
	s.items = append(s.items, item)
}

// 从栈顶弹出整数并返回
// 如果栈为空，返回0和false
func (s *IntStack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	n := len(s.items)
	item := s.items[n-1]
	s.items = s.items[:n-1]
	return item, true
}

// 查看栈顶整数但不移除
// 如果栈为空，返回0和false
func (s *IntStack) Peek() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	return s.items[len(s.items)-1], true
}

// 检查整数栈是否为空
func (s *IntStack) IsEmpty() bool {
	return len(s.items) == 0
}

// 返回整数栈中元素数量
func (s *IntStack) Size() int {
	return len(s.items)
}
