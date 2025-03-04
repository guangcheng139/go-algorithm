package stack

import (
	"reflect"
	"testing"
)

// 测试通用栈的基本操作
func TestStack(t *testing.T) {
	// 创建新栈
	stack := NewStack()

	// 测试初始状态
	if !stack.IsEmpty() {
		t.Error("新创建的栈应该为空")
	}

	if stack.Size() != 0 {
		t.Errorf("新创建的栈大小应该为0，得到 %d", stack.Size())
	}

	// 测试Push和Peek操作
	stack.Push(1)
	stack.Push("test")
	stack.Push(3.14)

	if stack.Size() != 3 {
		t.Errorf("栈大小应该为3，得到 %d", stack.Size())
	}

	if stack.Peek() != 3.14 {
		t.Errorf("栈顶元素应该为3.14，得到 %v", stack.Peek())
	}

	// 测试Pop操作
	item := stack.Pop()
	if item != 3.14 {
		t.Errorf("弹出的元素应该为3.14，得到 %v", item)
	}

	if stack.Size() != 2 {
		t.Errorf("Pop后栈大小应该为2，得到 %d", stack.Size())
	}

	// 测试Clear操作
	stack.Clear()
	if !stack.IsEmpty() {
		t.Error("Clear后栈应该为空")
	}

	// 测试空栈操作
	if stack.Pop() != nil {
		t.Error("从空栈Pop应该返回nil")
	}

	if stack.Peek() != nil {
		t.Error("空栈Peek应该返回nil")
	}
}

// 测试整数栈的基本操作
func TestIntStack(t *testing.T) {
	// 创建新的整数栈
	stack := NewIntStack()

	// 测试初始状态
	if !stack.IsEmpty() {
		t.Error("新创建的整数栈应该为空")
	}

	if stack.Size() != 0 {
		t.Errorf("新创建的整数栈大小应该为0，得到 %d", stack.Size())
	}

	// 测试Push和Peek操作
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	if stack.Size() != 3 {
		t.Errorf("栈大小应该为3，得到 %d", stack.Size())
	}

	val, ok := stack.Peek()
	if !ok || val != 30 {
		t.Errorf("栈顶元素应该为30，得到 %v (ok=%v)", val, ok)
	}

	// 测试Pop操作
	val, ok = stack.Pop()
	if !ok || val != 30 {
		t.Errorf("弹出的元素应该为30，得到 %v (ok=%v)", val, ok)
	}

	if stack.Size() != 2 {
		t.Errorf("Pop后栈大小应该为2，得到 %d", stack.Size())
	}

	// 测试空栈操作
	stack = NewIntStack()
	val, ok = stack.Pop()
	if ok || val != 0 {
		t.Errorf("从空栈Pop应该返回(0,false)，得到 (%v,%v)", val, ok)
	}

	val, ok = stack.Peek()
	if ok || val != 0 {
		t.Errorf("空栈Peek应该返回(0,false)，得到 (%v,%v)", val, ok)
	}
}

// 测试有效括号函数
func TestIsValid(t *testing.T) {
	tests := []struct {
		s        string
		expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"", true},
		{"((", false},
		{"))", false},
		{"([]){", false},
	}

	for i, test := range tests {
		result := IsValid(test.s)
		if result != test.expected {
			t.Errorf("测试%d失败: IsValid(%q) = %v，期望值 %v",
				i, test.s, result, test.expected)
		}
	}
}

// 测试最小栈
func TestMinStack(t *testing.T) {
	// 创建一个最小栈
	minStack := NewMinStack()

	// 测试基本操作
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)

	// 检查最小值是否为-3
	if minStack.GetMin() != -3 {
		t.Errorf("最小值应该为-3，得到 %d", minStack.GetMin())
	}

	// 弹出栈顶元素(-3)
	minStack.Pop()

	// 检查栈顶元素是否为0
	if minStack.Top() != 0 {
		t.Errorf("栈顶元素应该为0，得到 %d", minStack.Top())
	}

	// 检查最小值是否为-2
	if minStack.GetMin() != -2 {
		t.Errorf("最小值应该为-2，得到 %d", minStack.GetMin())
	}

	// 弹出所有元素
	minStack.Pop()
	minStack.Pop()

	// 添加新元素
	minStack.Push(5)
	minStack.Push(3)

	// 检查最小值是否为3
	if minStack.GetMin() != 3 {
		t.Errorf("最小值应该为3，得到 %d", minStack.GetMin())
	}
}

// 测试逆波兰表达式求值
func TestEvalRPN(t *testing.T) {
	tests := []struct {
		tokens   []string
		expected int
	}{
		{[]string{"2", "1", "+", "3", "*"}, 9},
		{[]string{"4", "13", "5", "/", "+"}, 6},
		{[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
		{[]string{"3", "-4", "+"}, -1},
		{[]string{"3", "4", "-"}, -1},
	}

	for i, test := range tests {
		result := EvalRPN(test.tokens)
		if result != test.expected {
			t.Errorf("测试%d失败: EvalRPN(%v) = %d，期望值 %d",
				i, test.tokens, result, test.expected)
		}
	}
}

// 测试用栈实现的队列
func TestMyQueue(t *testing.T) {
	// 创建一个队列
	queue := NewMyQueue()

	// 测试空队列
	if !queue.Empty() {
		t.Error("新创建的队列应该为空")
	}

	// 添加元素
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)

	// 检查非空状态
	if queue.Empty() {
		t.Error("添加元素后队列不应该为空")
	}

	// 测试Peek操作
	if queue.Peek() != 1 {
		t.Errorf("队列头部元素应该为1，得到 %d", queue.Peek())
	}

	// 测试Pop操作
	if queue.Pop() != 1 {
		t.Errorf("弹出的元素应该为1，得到 %d", queue.Pop())
	}

	// 再次检查头部元素
	if queue.Peek() != 2 {
		t.Errorf("队列头部元素应该为2，得到 %d", queue.Peek())
	}

	// 继续添加元素
	queue.Push(4)

	// 按顺序弹出所有元素
	expected := []int{2, 3, 4}
	for i, exp := range expected {
		val := queue.Pop()
		if val != exp {
			t.Errorf("第%d次弹出的元素应该为%d，得到 %d", i, exp, val)
		}
	}

	// 检查队列是否为空
	if !queue.Empty() {
		t.Error("弹出所有元素后队列应该为空")
	}
}

// 测试每日温度函数
func TestDailyTemperatures(t *testing.T) {
	tests := []struct {
		temperatures []int
		expected     []int
	}{
		{[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{[]int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
		{[]int{30, 60, 90}, []int{1, 1, 0}},
		{[]int{}, []int{}},
		{[]int{30}, []int{0}},
	}

	for i, test := range tests {
		result := DailyTemperatures(test.temperatures)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("测试%d失败: DailyTemperatures(%v) = %v，期望值 %v",
				i, test.temperatures, result, test.expected)
		}
	}
}

// 测试柱状图中最大的矩形
func TestLargestRectangleArea(t *testing.T) {
	tests := []struct {
		heights  []int
		expected int
	}{
		{[]int{2, 1, 5, 6, 2, 3}, 10},
		{[]int{2, 4}, 4},
		{[]int{1, 1}, 2},
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{0}, 0},
		{[]int{1, 2, 3, 4, 5}, 9},
	}

	for i, test := range tests {
		result := LargestRectangleArea(test.heights)
		if result != test.expected {
			t.Errorf("测试%d失败: LargestRectangleArea(%v) = %d，期望值 %d",
				i, test.heights, result, test.expected)
		}
	}
}
