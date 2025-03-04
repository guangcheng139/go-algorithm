package classic

// NumIslands 计算二维网格中岛屿的数量
// 岛屿由相邻的陆地（'1'）组成，相邻指上下左右四个方向
// 网格被水（'0'）包围
func NumIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	count := 0

	// 遍历网格
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// 如果找到陆地，进行DFS并增加岛屿计数
			if grid[i][j] == '1' {
				count++
				dfs(grid, i, j, rows, cols)
			}
		}
	}

	return count
}

// dfs 深度优先搜索，将与当前陆地相连的所有陆地标记为已访问
func dfs(grid [][]byte, i, j, rows, cols int) {
	// 检查边界条件和是否为陆地
	if i < 0 || i >= rows || j < 0 || j >= cols || grid[i][j] != '1' {
		return
	}

	// 标记当前陆地为已访问（将其改为'0'）
	grid[i][j] = '0'

	// 递归访问上下左右四个方向
	dfs(grid, i+1, j, rows, cols) // 下
	dfs(grid, i-1, j, rows, cols) // 上
	dfs(grid, i, j+1, rows, cols) // 右
	dfs(grid, i, j-1, rows, cols) // 左
}

// NumIslandsBFS 使用BFS计算岛屿数量
func NumIslandsBFS(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	count := 0

	// 方向数组：上、右、下、左
	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	// 遍历网格
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// 如果找到陆地，进行BFS并增加岛屿计数
			if grid[i][j] == '1' {
				count++

				// 标记当前陆地为已访问
				grid[i][j] = '0'

				// 创建队列并将当前位置入队
				queue := [][]int{{i, j}}

				// BFS
				for len(queue) > 0 {
					// 出队
					cell := queue[0]
					queue = queue[1:]
					row, col := cell[0], cell[1]

					// 检查四个方向
					for _, dir := range directions {
						newRow, newCol := row+dir[0], col+dir[1]

						// 检查边界条件和是否为陆地
						if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols && grid[newRow][newCol] == '1' {
							// 标记为已访问并入队
							grid[newRow][newCol] = '0'
							queue = append(queue, []int{newRow, newCol})
						}
					}
				}
			}
		}
	}

	return count
}

// MaxAreaOfIsland 计算最大岛屿面积
func MaxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	maxArea := 0

	// 遍历网格
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// 如果找到陆地，计算岛屿面积
			if grid[i][j] == 1 {
				area := dfsArea(grid, i, j, rows, cols)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea
}

// dfsArea 深度优先搜索，计算岛屿面积
func dfsArea(grid [][]int, i, j, rows, cols int) int {
	// 检查边界条件和是否为陆地
	if i < 0 || i >= rows || j < 0 || j >= cols || grid[i][j] != 1 {
		return 0
	}

	// 标记当前陆地为已访问
	grid[i][j] = 0
	area := 1

	// 递归计算上下左右四个方向的面积
	area += dfsArea(grid, i+1, j, rows, cols) // 下
	area += dfsArea(grid, i-1, j, rows, cols) // 上
	area += dfsArea(grid, i, j+1, rows, cols) // 右
	area += dfsArea(grid, i, j-1, rows, cols) // 左

	return area
}

// NumDistinctIslands 计算不同形状的岛屿数量
func NumDistinctIslands(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])

	// 使用map存储不同形状的岛屿
	distinctIslands := make(map[string]bool)

	// 遍历网格
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// 如果找到陆地，进行DFS并记录形状
			if grid[i][j] == 1 {
				// 使用字符串表示岛屿形状
				path := ""
				dfsShape(grid, i, j, i, j, rows, cols, &path, "o") // 起点

				// 如果形状不为空，添加到集合中
				if path != "" {
					distinctIslands[path] = true
				}
			}
		}
	}

	// 返回不同形状的岛屿数量
	return len(distinctIslands)
}

// dfsShape 深度优先搜索，记录岛屿形状
// 使用相对坐标表示形状
func dfsShape(grid [][]int, i, j, baseI, baseJ, rows, cols int, path *string, dir string) {
	// 检查边界条件和是否为陆地
	if i < 0 || i >= rows || j < 0 || j >= cols || grid[i][j] != 1 {
		return
	}

	// 标记当前陆地为已访问
	grid[i][j] = 0

	// 记录当前方向
	*path += dir

	// 递归访问上下左右四个方向
	dfsShape(grid, i+1, j, baseI, baseJ, rows, cols, path, "d") // 下
	dfsShape(grid, i-1, j, baseI, baseJ, rows, cols, path, "u") // 上
	dfsShape(grid, i, j+1, baseI, baseJ, rows, cols, path, "r") // 右
	dfsShape(grid, i, j-1, baseI, baseJ, rows, cols, path, "l") // 左

	// 回溯标记
	*path += "b"
}
