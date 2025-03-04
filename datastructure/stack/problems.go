package stack

// 有效的括号：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
// 有效字符串需满足：
// 1. 左括号必须用相同类型的右括号闭合。
// 2. 左括号必须以正确的顺序闭合。
// 时间复杂度：O(n)，空间复杂度：O(n)
func IsValid(s string) bool {
	// 创建一个字符栈
	stack := NewStack()

	// 遍历字符串中的每个字符
	for _, char := range s {
		switch char {
		case '(', '{', '[': // 如果是左括号，入栈
			stack.Push(char)
		case ')': // 如果是右括号，检查栈顶是否匹配
			if stack.IsEmpty() || stack.Pop() != '(' {
				return false
			}
		case '}':
			if stack.IsEmpty() || stack.Pop() != '{' {
				return false
			}
		case ']':
			if stack.IsEmpty() || stack.Pop() != '[' {
				return false
			}
		}
	}

	// 如果栈为空，说明所有括号都匹配
	return stack.IsEmpty()
}

// MinStack 设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。
// push(x) —— 将元素 x 推入栈中。
// pop() —— 删除栈顶的元素。
// top() —— 获取栈顶元素。
// getMin() —— 检索栈中的最小元素。
type MinStack struct {
	stack    []int // 主栈，存储所有元素
	minStack []int // 辅助栈，存储最小值
}

// 创建一个新的最小栈
func NewMinStack() *MinStack {
	return &MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

// 将元素推入最小栈
// 时间复杂度：O(1)
func (ms *MinStack) Push(x int) {
	// 元素入主栈
	ms.stack = append(ms.stack, x)

	// 如果辅助栈为空或x小于等于当前最小值，则x也入辅助栈
	if len(ms.minStack) == 0 || x <= ms.minStack[len(ms.minStack)-1] {
		ms.minStack = append(ms.minStack, x)
	}
}

// 从最小栈中弹出元素
// 时间复杂度：O(1)
func (ms *MinStack) Pop() {
	// 如果栈为空，不执行操作
	if len(ms.stack) == 0 {
		return
	}

	// 如果弹出的元素是当前最小值，辅助栈也要弹出
	if ms.stack[len(ms.stack)-1] == ms.minStack[len(ms.minStack)-1] {
		ms.minStack = ms.minStack[:len(ms.minStack)-1]
	}

	// 弹出主栈顶元素
	ms.stack = ms.stack[:len(ms.stack)-1]
}

// 获取最小栈的栈顶元素
// 时间复杂度：O(1)
func (ms *MinStack) Top() int {
	// 如果栈为空，返回0（实际应用中可能需要错误处理）
	if len(ms.stack) == 0 {
		return 0
	}

	return ms.stack[len(ms.stack)-1]
}

// 获取最小栈中的最小元素
// 时间复杂度：O(1)
func (ms *MinStack) GetMin() int {
	// 如果辅助栈为空，返回0（实际应用中可能需要错误处理）
	if len(ms.minStack) == 0 {
		return 0
	}

	return ms.minStack[len(ms.minStack)-1]
}

// 逆波兰表达式求值：根据逆波兰表示法，求表达式的值。
// 有效的运算符包括 +, -, *, / 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。
// 时间复杂度：O(n)，空间复杂度：O(n)
func EvalRPN(tokens []string) int {
	stack := NewIntStack()

	for _, token := range tokens {
		// 如果是操作符，弹出两个数进行计算，并将结果压入栈中
		if token == "+" || token == "-" || token == "*" || token == "/" {
			// 弹出两个操作数
			num2, _ := stack.Pop()
			num1, _ := stack.Pop()

			// 根据操作符进行计算
			var result int
			switch token {
			case "+":
				result = num1 + num2
			case "-":
				result = num1 - num2
			case "*":
				result = num1 * num2
			case "/":
				result = num1 / num2
			}

			// 将结果压入栈中
			stack.Push(result)
		} else {
			// 如果是数字，转换为整数后压入栈中
			num := 0
			sign := 1

			// 处理负数
			if token[0] == '-' {
				sign = -1
				token = token[1:]
			}

			// 将字符串转换为整数
			for i := 0; i < len(token); i++ {
				num = num*10 + int(token[i]-'0')
			}

			stack.Push(num * sign)
		}
	}

	// 最终栈中应该只剩下一个元素，即表达式的结果
	result, _ := stack.Pop()
	return result
}

// 用栈实现队列：使用栈实现队列的下列操作
// push(x) -- 将一个元素放入队列的尾部
// pop() -- 从队列首部移除元素
// peek() -- 返回队列首部的元素
// empty() -- 返回队列是否为空
type MyQueue struct {
	inStack  []int // 用于入队操作的栈
	outStack []int // 用于出队操作的栈
}

// 创建一个新的队列
func NewMyQueue() *MyQueue {
	return &MyQueue{
		inStack:  []int{},
		outStack: []int{},
	}
}

// 将元素推入队列
// 时间复杂度：O(1)
func (q *MyQueue) Push(x int) {
	q.inStack = append(q.inStack, x)
}

// 从队列中弹出元素
// 时间复杂度：平均O(1)，最坏O(n)
func (q *MyQueue) Pop() int {
	// 确保outStack非空
	q.ensureOutStack()

	if len(q.outStack) == 0 {
		return -1 // 实际应用中可能需要错误处理
	}

	// 从outStack弹出元素
	val := q.outStack[len(q.outStack)-1]
	q.outStack = q.outStack[:len(q.outStack)-1]
	return val
}

// 查看队列顶部元素
// 时间复杂度：平均O(1)，最坏O(n)
func (q *MyQueue) Peek() int {
	// 确保outStack非空
	q.ensureOutStack()

	if len(q.outStack) == 0 {
		return -1 // 实际应用中可能需要错误处理
	}

	return q.outStack[len(q.outStack)-1]
}

// 检查队列是否为空
// 时间复杂度：O(1)
func (q *MyQueue) Empty() bool {
	return len(q.inStack) == 0 && len(q.outStack) == 0
}

// 确保outStack非空（如果可能）
// 时间复杂度：平均O(1)，最坏O(n)
func (q *MyQueue) ensureOutStack() {
	// 如果outStack为空，将inStack中的所有元素倒入outStack
	if len(q.outStack) == 0 {
		for len(q.inStack) > 0 {
			val := q.inStack[len(q.inStack)-1]
			q.inStack = q.inStack[:len(q.inStack)-1]
			q.outStack = append(q.outStack, val)
		}
	}
}

// 每日温度：给定一个整数数组 temperatures 表示每天的温度，返回一个数组，
// 对于每一天，返回需要等待多少天才会有更高的温度。如果之后没有更高的温度，则返回0。
// 时间复杂度：O(n)，空间复杂度：O(n)
func DailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	result := make([]int, n)
	stack := []int{} // 存储索引的栈

	for i := 0; i < n; i++ {
		// 当栈不为空，且当前温度大于栈顶索引对应的温度时
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			// 弹出栈顶索引
			prevIdx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// 计算等待天数
			result[prevIdx] = i - prevIdx
		}

		// 将当前索引入栈
		stack = append(stack, i)
	}

	// 栈中剩余的索引，对应的结果为0（已经初始化为0）
	return result
}

// 柱状图中最大的矩形：给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
// 求在该柱状图中，能够勾勒出来的矩形的最大面积。
// 时间复杂度：O(n)，空间复杂度：O(n)
func LargestRectangleArea(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}

	// 在heights数组前后增加0高度的柱子，方便处理边界情况
	newHeights := make([]int, n+2)
	copy(newHeights[1:n+1], heights)

	stack := []int{0} // 存储索引的栈，初始放入0（即左边界）
	maxArea := 0

	for i := 1; i < len(newHeights); i++ {
		// 当栈不为空，且当前高度小于栈顶索引对应的高度时
		// 说明找到了右边界，可以计算面积
		for len(stack) > 0 && newHeights[i] < newHeights[stack[len(stack)-1]] {
			// 弹出栈顶索引
			height := newHeights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]

			// 计算宽度：当前索引 - 新的栈顶索引 - 1
			width := i - stack[len(stack)-1] - 1

			// 更新最大面积
			area := height * width
			if area > maxArea {
				maxArea = area
			}
		}

		// 将当前索引入栈
		stack = append(stack, i)
	}

	return maxArea
}
