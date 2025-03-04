package graph

import (
	"container/list"
)

// BFS 广度优先遍历图
// 从指定的起始顶点开始，返回访问顶点的顺序
func BFS(g *Graph, startVertex int) []int {
	// 检查起始顶点是否有效
	if startVertex < 0 || startVertex >= g.Vertices {
		return nil
	}

	// 初始化访问结果和已访问标记
	var result []int
	visited := make([]bool, g.Vertices)

	// 创建队列并将起始顶点入队
	queue := list.New()
	queue.PushBack(startVertex)
	visited[startVertex] = true

	// BFS主循环
	for queue.Len() > 0 {
		// 出队一个顶点
		vertex := queue.Remove(queue.Front()).(int)
		result = append(result, vertex)

		// 访问所有未访问的邻居
		for _, edge := range g.GetNeighbors(vertex) {
			neighbor := edge.To
			if !visited[neighbor] {
				visited[neighbor] = true
				queue.PushBack(neighbor)
			}
		}
	}

	return result
}

// DFS 深度优先遍历图
// 从指定的起始顶点开始，返回访问顶点的顺序
func DFS(g *Graph, startVertex int) []int {
	// 检查起始顶点是否有效
	if startVertex < 0 || startVertex >= g.Vertices {
		return nil
	}

	// 初始化访问结果和已访问标记
	var result []int
	visited := make([]bool, g.Vertices)

	// 调用递归DFS函数
	dfsRecursive(g, startVertex, visited, &result)

	return result
}

// dfsRecursive 递归实现DFS
func dfsRecursive(g *Graph, vertex int, visited []bool, result *[]int) {
	// 标记当前顶点为已访问
	visited[vertex] = true
	*result = append(*result, vertex)

	// 递归访问所有未访问的邻居
	for _, edge := range g.GetNeighbors(vertex) {
		neighbor := edge.To
		if !visited[neighbor] {
			dfsRecursive(g, neighbor, visited, result)
		}
	}
}

// DFSIterative 使用迭代方式实现DFS
// 从指定的起始顶点开始，返回访问顶点的顺序
func DFSIterative(g *Graph, startVertex int) []int {
	// 检查起始顶点是否有效
	if startVertex < 0 || startVertex >= g.Vertices {
		return nil
	}

	// 初始化访问结果和已访问标记
	var result []int
	visited := make([]bool, g.Vertices)

	// 创建栈并将起始顶点入栈
	stack := list.New()
	stack.PushBack(startVertex)

	// DFS主循环
	for stack.Len() > 0 {
		// 出栈一个顶点
		vertex := stack.Remove(stack.Back()).(int)

		// 如果顶点未访问，则标记为已访问并添加到结果中
		if !visited[vertex] {
			visited[vertex] = true
			result = append(result, vertex)

			// 将所有未访问的邻居入栈（注意：倒序入栈以保持与递归DFS相同的访问顺序）
			neighbors := g.GetNeighbors(vertex)
			for i := len(neighbors) - 1; i >= 0; i-- {
				neighbor := neighbors[i].To
				if !visited[neighbor] {
					stack.PushBack(neighbor)
				}
			}
		}
	}

	return result
}

// FindPath 使用BFS查找从起点到终点的最短路径
// 返回路径上的顶点序列，如果不存在路径则返回nil
func FindPath(g *Graph, startVertex, endVertex int) []int {
	// 检查顶点是否有效
	if startVertex < 0 || startVertex >= g.Vertices ||
		endVertex < 0 || endVertex >= g.Vertices {
		return nil
	}

	// 如果起点和终点相同
	if startVertex == endVertex {
		return []int{startVertex}
	}

	// 初始化已访问标记和前驱数组
	visited := make([]bool, g.Vertices)
	predecessor := make([]int, g.Vertices)
	for i := range predecessor {
		predecessor[i] = -1 // -1表示没有前驱
	}

	// 创建队列并将起始顶点入队
	queue := list.New()
	queue.PushBack(startVertex)
	visited[startVertex] = true

	// BFS主循环
	foundPath := false
	for queue.Len() > 0 && !foundPath {
		// 出队一个顶点
		vertex := queue.Remove(queue.Front()).(int)

		// 访问所有未访问的邻居
		for _, edge := range g.GetNeighbors(vertex) {
			neighbor := edge.To
			if !visited[neighbor] {
				visited[neighbor] = true
				predecessor[neighbor] = vertex
				queue.PushBack(neighbor)

				// 如果找到终点，结束搜索
				if neighbor == endVertex {
					foundPath = true
					break
				}
			}
		}
	}

	// 如果没有找到路径
	if !foundPath {
		return nil
	}

	// 重建路径（从终点回溯到起点）
	var path []int
	for at := endVertex; at != -1; at = predecessor[at] {
		path = append([]int{at}, path...)
	}

	return path
}

// HasCycle 检测图中是否存在环
func HasCycle(g *Graph) bool {
	// 对于无向图和有向图使用不同的检测方法
	if g.IsDirected {
		return hasDirectedCycle(g)
	}
	return hasUndirectedCycle(g)
}

// hasDirectedCycle 检测有向图中是否存在环
func hasDirectedCycle(g *Graph) bool {
	// 初始化访问状态数组
	// 0: 未访问, 1: 正在访问, 2: 已完成访问
	state := make([]int, g.Vertices)

	// 对每个未访问的顶点进行DFS
	for i := 0; i < g.Vertices; i++ {
		if state[i] == 0 && dfsCheckDirectedCycle(g, i, state) {
			return true
		}
	}

	return false
}

// dfsCheckDirectedCycle 使用DFS检测有向图中的环
func dfsCheckDirectedCycle(g *Graph, vertex int, state []int) bool {
	// 标记当前顶点为正在访问
	state[vertex] = 1

	// 检查所有邻居
	for _, edge := range g.GetNeighbors(vertex) {
		neighbor := edge.To

		// 如果邻居正在被访问，则存在环
		if state[neighbor] == 1 {
			return true
		}

		// 如果邻居未访问，且从邻居出发存在环
		if state[neighbor] == 0 && dfsCheckDirectedCycle(g, neighbor, state) {
			return true
		}
	}

	// 标记当前顶点为已完成访问
	state[vertex] = 2
	return false
}

// hasUndirectedCycle 检测无向图中是否存在环
func hasUndirectedCycle(g *Graph) bool {
	// 初始化已访问数组
	visited := make([]bool, g.Vertices)

	// 对每个未访问的顶点进行DFS
	for i := 0; i < g.Vertices; i++ {
		if !visited[i] && dfsCheckUndirectedCycle(g, i, -1, visited) {
			return true
		}
	}

	return false
}

// dfsCheckUndirectedCycle 使用DFS检测无向图中的环
func dfsCheckUndirectedCycle(g *Graph, vertex, parent int, visited []bool) bool {
	// 标记当前顶点为已访问
	visited[vertex] = true

	// 检查所有邻居
	for _, edge := range g.GetNeighbors(vertex) {
		neighbor := edge.To

		// 如果邻居未访问
		if !visited[neighbor] {
			// 从邻居出发继续DFS，如果发现环则返回true
			if dfsCheckUndirectedCycle(g, neighbor, vertex, visited) {
				return true
			}
		} else if neighbor != parent {
			// 如果邻居已访问且不是父节点，则存在环
			return true
		}
	}

	return false
}
