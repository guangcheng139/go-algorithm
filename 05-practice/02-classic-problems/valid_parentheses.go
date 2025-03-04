package classic

// IsValid 检查字符串中的括号是否有效
// 有效的括号必须以正确的顺序闭合
// 例如："{[]}" 是有效的，但 "{[}]" 是无效的
func IsValid(s string) bool {
	// 创建一个栈用于存储左括号
	var stack []rune

	// 定义括号对应关系
	brackets := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	// 遍历字符串中的每个字符
	for _, char := range s {
		// 如果是左括号，入栈
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else if char == ')' || char == '}' || char == ']' {
			// 如果是右括号，检查栈是否为空
			if len(stack) == 0 {
				return false
			}

			// 弹出栈顶元素，检查是否匹配
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// 如果不匹配，返回false
			if top != brackets[char] {
				return false
			}
		}
	}

	// 如果栈为空，说明所有括号都匹配
	return len(stack) == 0
}

// LongestValidParentheses 计算最长有效括号子串的长度
// 例如：")()())" 中最长有效括号子串是 "()()"，长度为4
func LongestValidParentheses(s string) int {
	maxLen := 0
	stack := []int{-1} // 初始化栈，-1作为基准

	// 遍历字符串
	for i, char := range s {
		if char == '(' {
			// 左括号，将索引入栈
			stack = append(stack, i)
		} else {
			// 右括号，弹出栈顶元素
			stack = stack[:len(stack)-1]

			if len(stack) == 0 {
				// 栈为空，将当前索引作为新的基准
				stack = append(stack, i)
			} else {
				// 计算当前有效括号子串的长度
				currLen := i - stack[len(stack)-1]
				if currLen > maxLen {
					maxLen = currLen
				}
			}
		}
	}

	return maxLen
}

// GenerateParenthesis 生成所有可能的有效括号组合
// n表示括号对数
func GenerateParenthesis(n int) []string {
	var result []string
	backtrack(&result, "", 0, 0, n)
	return result
}

// backtrack 回溯生成有效括号
// open: 已使用的左括号数量
// close: 已使用的右括号数量
// n: 括号对数
func backtrack(result *[]string, current string, open, close, n int) {
	// 如果当前字符串长度等于2n，说明已生成完整的括号组合
	if len(current) == 2*n {
		*result = append(*result, current)
		return
	}

	// 如果左括号数量小于n，可以添加左括号
	if open < n {
		backtrack(result, current+"(", open+1, close, n)
	}

	// 如果右括号数量小于左括号数量，可以添加右括号
	if close < open {
		backtrack(result, current+")", open, close+1, n)
	}
}

// MinAddToMakeValid 计算使括号有效的最少添加数
// 例如："())" 需要添加1个左括号变成 "(())"
func MinAddToMakeValid(s string) int {
	// 未匹配的左括号数量
	open := 0
	// 需要添加的括号数量
	add := 0

	for _, char := range s {
		if char == '(' {
			// 左括号，增加未匹配计数
			open++
		} else if char == ')' {
			// 右括号，检查是否有未匹配的左括号
			if open > 0 {
				// 有未匹配的左括号，减少未匹配计数
				open--
			} else {
				// 没有未匹配的左括号，需要添加一个左括号
				add++
			}
		}
	}

	// 最终结果是需要添加的括号数量加上未匹配的左括号数量
	return add + open
}

// ScoreOfParentheses 计算括号字符串的分数
// 规则：
// - "()" 的分数是1
// - "AB" 的分数是A的分数加上B的分数，其中A和B是有效括号字符串
// - "(A)" 的分数是A的分数的2倍，其中A是有效括号字符串
func ScoreOfParentheses(s string) int {
	// 使用栈记录分数
	stack := []int{0} // 初始化栈，0表示当前层的分数

	for _, char := range s {
		if char == '(' {
			// 左括号，开始新的一层，初始分数为0
			stack = append(stack, 0)
		} else {
			// 右括号，计算当前层的分数
			v := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// 计算新的分数
			// 如果当前层的分数为0，则为"()"，分数为1
			// 否则为"(A)"，分数为A的2倍
			score := 1
			if v > 0 {
				score = 2 * v
			}

			// 将分数添加到上一层
			stack[len(stack)-1] += score
		}
	}

	// 返回最外层的分数
	return stack[0]
}
