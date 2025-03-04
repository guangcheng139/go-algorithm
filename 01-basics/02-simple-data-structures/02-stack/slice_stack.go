package stack

// SliceStack 表示基于切片的栈
type SliceStack struct {
	items []int
}

// NewSliceStack 创建一个新的基于切片的栈
func NewSliceStack() *SliceStack {
	return &SliceStack{
		items: make([]int, 0),
	}
}

// Push 将元素推入栈顶
func (s *SliceStack) Push(value int) {
	s.items = append(s.items, value)
}

// Pop 从栈顶弹出元素
func (s *SliceStack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	index := len(s.items) - 1
	value := s.items[index]
	s.items = s.items[:index]
	return value, true
}

// Peek 查看栈顶元素但不移除
func (s *SliceStack) Peek() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	return s.items[len(s.items)-1], true
}

// IsEmpty 检查栈是否为空
func (s *SliceStack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size 返回栈中元素的数量
func (s *SliceStack) Size() int {
	return len(s.items)
}

// Clear 清空栈
func (s *SliceStack) Clear() {
	s.items = make([]int, 0)
}

// ToSlice 将栈转换为切片（从栈底到栈顶）
func (s *SliceStack) ToSlice() []int {
	result := make([]int, len(s.items))
	copy(result, s.items)
	return result
}
