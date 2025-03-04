package dfs_bfs

// TreeNode 二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// MaxDepth 二叉树的最大深度 - DFS实现
// 给定一个二叉树，找出其最大深度。最大深度是指从根节点到最远叶子节点的最长路径上的节点数量。
// 时间复杂度: O(n)
// 空间复杂度: O(h)，h为树的高度
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := MaxDepth(root.Left)
	rightDepth := MaxDepth(root.Right)

	// 取较大值 + 1（当前节点）
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

// MaxDepthBFS 二叉树的最大深度 - BFS实现
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func MaxDepthBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	depth := 0

	for len(queue) > 0 {
		size := len(queue)
		depth++

		// 处理当前层的所有节点
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:] // 出队

			// 将子节点加入队列
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return depth
}

// LevelOrder 二叉树的层序遍历 - BFS实现
// 给你一个二叉树，请你返回其按层序遍历得到的节点值（即逐层地，从左到右访问所有节点）
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)
		level := []int{}

		// 处理当前层的所有节点
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:] // 出队

			level = append(level, node.Val)

			// 将子节点加入队列
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, level)
	}

	return result
}

// ZigzagLevelOrder 二叉树的锯齿形层序遍历
// 给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推）
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func ZigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := []*TreeNode{root}
	leftToRight := true

	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, size)

		// 处理当前层的所有节点
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:] // 出队

			// 根据方向决定元素放置位置
			idx := i
			if !leftToRight {
				idx = size - 1 - i
			}
			level[idx] = node.Val

			// 将子节点加入队列
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, level)
		leftToRight = !leftToRight // 切换方向
	}

	return result
}

// NumIslands 岛屿数量
// 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
// 时间复杂度: O(m*n)
// 空间复杂度: O(m*n)，最坏情况下递归深度
func NumIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	count := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				count++
				dfs(grid, i, j, rows, cols)
			}
		}
	}

	return count
}

// dfs 深度优先搜索，将相连的陆地标记为已访问
func dfs(grid [][]byte, i, j, rows, cols int) {
	// 边界检查和水域检查
	if i < 0 || i >= rows || j < 0 || j >= cols || grid[i][j] != '1' {
		return
	}

	// 标记为已访问
	grid[i][j] = '0'

	// 访问上下左右四个方向
	dfs(grid, i+1, j, rows, cols)
	dfs(grid, i-1, j, rows, cols)
	dfs(grid, i, j+1, rows, cols)
	dfs(grid, i, j-1, rows, cols)
}

// SurroundedRegions 被围绕的区域
// 给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' 组成，找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。
// 被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。
// 时间复杂度: O(m*n)
// 空间复杂度: O(m*n)
func SurroundedRegions(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	rows, cols := len(board), len(board[0])

	// 从边界开始DFS，标记所有与边界连通的'O'
	// 第一列和最后一列
	for i := 0; i < rows; i++ {
		if board[i][0] == 'O' {
			markBoundaryConnected(board, i, 0, rows, cols)
		}
		if board[i][cols-1] == 'O' {
			markBoundaryConnected(board, i, cols-1, rows, cols)
		}
	}

	// 第一行和最后一行
	for j := 0; j < cols; j++ {
		if board[0][j] == 'O' {
			markBoundaryConnected(board, 0, j, rows, cols)
		}
		if board[rows-1][j] == 'O' {
			markBoundaryConnected(board, rows-1, j, rows, cols)
		}
	}

	// 处理整个矩阵：
	// - 'T'表示与边界连通的'O'，恢复为'O'
	// - 其他'O'被'X'围绕，替换为'X'
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X' // 被围绕的'O'替换为'X'
			} else if board[i][j] == 'T' {
				board[i][j] = 'O' // 恢复与边界连通的'O'
			}
		}
	}
}

// markBoundaryConnected 标记与边界连通的'O'
func markBoundaryConnected(board [][]byte, i, j, rows, cols int) {
	if i < 0 || i >= rows || j < 0 || j >= cols || board[i][j] != 'O' {
		return
	}

	// 标记为临时字符'T'
	board[i][j] = 'T'

	// 递归标记上下左右
	markBoundaryConnected(board, i+1, j, rows, cols)
	markBoundaryConnected(board, i-1, j, rows, cols)
	markBoundaryConnected(board, i, j+1, rows, cols)
	markBoundaryConnected(board, i, j-1, rows, cols)
}

// WordSearch 单词搜索
// 给定一个二维网格和一个单词，找出该单词是否存在于网格中。
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中"相邻"单元格是那些水平相邻或垂直相邻的单元格。
// 同一个单元格内的字母不允许被重复使用。
// 时间复杂度: O(m*n*4^L)，其中L是单词长度
// 空间复杂度: O(m*n)
func WordSearch(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 || len(word) == 0 {
		return false
	}

	rows, cols := len(board), len(board[0])

	// 从每个格子开始尝试
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == word[0] && dfsWordSearch(board, word, i, j, 0, rows, cols) {
				return true
			}
		}
	}

	return false
}

// dfsWordSearch 深度优先搜索单词
func dfsWordSearch(board [][]byte, word string, i, j, index, rows, cols int) bool {
	// 已找到单词
	if index == len(word) {
		return true
	}

	// 边界检查和字符匹配检查
	if i < 0 || i >= rows || j < 0 || j >= cols || board[i][j] != word[index] {
		return false
	}

	// 标记当前格子为已访问（避免重复使用）
	temp := board[i][j]
	board[i][j] = '#'

	// 尝试四个方向
	found := dfsWordSearch(board, word, i+1, j, index+1, rows, cols) ||
		dfsWordSearch(board, word, i-1, j, index+1, rows, cols) ||
		dfsWordSearch(board, word, i, j+1, index+1, rows, cols) ||
		dfsWordSearch(board, word, i, j-1, index+1, rows, cols)

	// 恢复当前格子
	board[i][j] = temp

	return found
}
