package graph

// Graph 图的邻接表表示
type Graph struct {
	Vertices   int
	AdjList    map[int][]int
	IsDirected bool
}

// NewGraph 创建一个新的图
// vertices: 顶点数量，从0到vertices-1编号
// isDirected: 是否为有向图
func NewGraph(vertices int, isDirected bool) *Graph {
	adjList := make(map[int][]int)
	for i := 0; i < vertices; i++ {
		adjList[i] = []int{}
	}

	return &Graph{
		Vertices:   vertices,
		AdjList:    adjList,
		IsDirected: isDirected,
	}
}

// AddEdge 添加边
// 时间复杂度: O(1)
func (g *Graph) AddEdge(src, dest int) {
	// 检查顶点是否有效
	if src < 0 || src >= g.Vertices || dest < 0 || dest >= g.Vertices {
		return
	}

	// 添加边 src -> dest
	g.AdjList[src] = append(g.AdjList[src], dest)

	// 如果是无向图，添加边 dest -> src
	if !g.IsDirected {
		g.AdjList[dest] = append(g.AdjList[dest], src)
	}
}

// RemoveEdge 移除边
// 时间复杂度: O(E)，E为目标顶点的边数
func (g *Graph) RemoveEdge(src, dest int) {
	// 移除边 src -> dest
	for i, v := range g.AdjList[src] {
		if v == dest {
			g.AdjList[src] = append(g.AdjList[src][:i], g.AdjList[src][i+1:]...)
			break
		}
	}

	// 如果是无向图，移除边 dest -> src
	if !g.IsDirected {
		for i, v := range g.AdjList[dest] {
			if v == src {
				g.AdjList[dest] = append(g.AdjList[dest][:i], g.AdjList[dest][i+1:]...)
				break
			}
		}
	}
}

// BFS 广度优先搜索
// 时间复杂度: O(V+E)，V为顶点数，E为边数
// 空间复杂度: O(V)
func (g *Graph) BFS(start int) []int {
	if start < 0 || start >= g.Vertices {
		return []int{}
	}

	visited := make([]bool, g.Vertices)
	result := []int{}

	// 创建队列并将起始顶点入队
	queue := []int{start}
	visited[start] = true

	for len(queue) > 0 {
		// 出队
		vertex := queue[0]
		queue = queue[1:]

		// 处理当前顶点
		result = append(result, vertex)

		// 访问所有相邻顶点
		for _, adjacent := range g.AdjList[vertex] {
			if !visited[adjacent] {
				visited[adjacent] = true
				queue = append(queue, adjacent)
			}
		}
	}

	return result
}

// DFS 深度优先搜索
// 时间复杂度: O(V+E)，V为顶点数，E为边数
// 空间复杂度: O(V)
func (g *Graph) DFS(start int) []int {
	if start < 0 || start >= g.Vertices {
		return []int{}
	}

	visited := make([]bool, g.Vertices)
	result := []int{}

	g.dfsUtil(start, visited, &result)

	return result
}

// dfsUtil DFS辅助函数
func (g *Graph) dfsUtil(vertex int, visited []bool, result *[]int) {
	// 标记当前顶点为已访问
	visited[vertex] = true

	// 处理当前顶点
	*result = append(*result, vertex)

	// 访问所有相邻顶点
	for _, adjacent := range g.AdjList[vertex] {
		if !visited[adjacent] {
			g.dfsUtil(adjacent, visited, result)
		}
	}
}

// HasCycle 检测无向图是否有环
// 时间复杂度: O(V+E)，V为顶点数，E为边数
// 空间复杂度: O(V)
func (g *Graph) HasCycle() bool {
	if g.IsDirected {
		return g.hasDirectedCycle()
	}
	return g.hasUndirectedCycle()
}

// hasUndirectedCycle 检测无向图是否有环
func (g *Graph) hasUndirectedCycle() bool {
	visited := make([]bool, g.Vertices)

	for i := 0; i < g.Vertices; i++ {
		if !visited[i] {
			if g.hasUndirectedCycleUtil(i, -1, visited) {
				return true
			}
		}
	}

	return false
}

// hasUndirectedCycleUtil 辅助函数
func (g *Graph) hasUndirectedCycleUtil(vertex, parent int, visited []bool) bool {
	visited[vertex] = true

	for _, adjacent := range g.AdjList[vertex] {
		// 如果相邻顶点未被访问
		if !visited[adjacent] {
			if g.hasUndirectedCycleUtil(adjacent, vertex, visited) {
				return true
			}
		} else if adjacent != parent {
			// 如果相邻顶点已被访问，且不是当前顶点的父节点，则存在环
			return true
		}
	}

	return false
}

// hasDirectedCycle 检测有向图是否有环
func (g *Graph) hasDirectedCycle() bool {
	visited := make([]bool, g.Vertices)
	recStack := make([]bool, g.Vertices) // 递归栈

	for i := 0; i < g.Vertices; i++ {
		if !visited[i] {
			if g.hasDirectedCycleUtil(i, visited, recStack) {
				return true
			}
		}
	}

	return false
}

// hasDirectedCycleUtil 辅助函数
func (g *Graph) hasDirectedCycleUtil(vertex int, visited, recStack []bool) bool {
	visited[vertex] = true
	recStack[vertex] = true

	for _, adjacent := range g.AdjList[vertex] {
		if !visited[adjacent] {
			if g.hasDirectedCycleUtil(adjacent, visited, recStack) {
				return true
			}
		} else if recStack[adjacent] {
			// 如果相邻顶点在当前递归栈中，则存在环
			return true
		}
	}

	recStack[vertex] = false
	return false
}

// TopologicalSort 拓扑排序（仅适用于有向无环图）
// 时间复杂度: O(V+E)，V为顶点数，E为边数
// 空间复杂度: O(V)
func (g *Graph) TopologicalSort() []int {
	if !g.IsDirected || g.hasDirectedCycle() {
		return []int{} // 无法对无向图或有环图进行拓扑排序
	}

	visited := make([]bool, g.Vertices)
	stack := []int{} // 用于存储结果

	for i := 0; i < g.Vertices; i++ {
		if !visited[i] {
			g.topologicalSortUtil(i, visited, &stack)
		}
	}

	// 反转栈以获得正确的顺序
	for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
		stack[i], stack[j] = stack[j], stack[i]
	}

	return stack
}

// topologicalSortUtil 拓扑排序辅助函数
func (g *Graph) topologicalSortUtil(vertex int, visited []bool, stack *[]int) {
	visited[vertex] = true

	for _, adjacent := range g.AdjList[vertex] {
		if !visited[adjacent] {
			g.topologicalSortUtil(adjacent, visited, stack)
		}
	}

	// 当前顶点的所有相邻顶点都已经处理完成
	*stack = append(*stack, vertex)
}
