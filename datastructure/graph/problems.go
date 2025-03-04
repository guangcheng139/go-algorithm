package graph

import (
	"container/heap"
	"math"
)

// NumIslands 岛屿数量
// 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
// 时间复杂度: O(M*N)，M和N是网格的行数和列数
// 空间复杂度: O(M*N)
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

// dfs 深度优先搜索，将相连的陆地都标记为已访问
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

// CanFinish 课程表（拓扑排序）
// 你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses-1 。
// 在选修某些课程之前需要一些先修课程。先修课程按数组 prerequisites 给出，
// 其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则必须先学习课程 bi 。
// 检查是否可能完成所有课程的学习。
// 时间复杂度: O(V+E)，V为顶点数，E为边数
// 空间复杂度: O(V+E)
func CanFinish(numCourses int, prerequisites [][]int) bool {
	// 构建有向图
	graph := NewGraph(numCourses, true)

	// 添加边
	for _, prereq := range prerequisites {
		course, prerequisite := prereq[0], prereq[1]
		graph.AddEdge(prerequisite, course) // 先修课程指向后修课程
	}

	// 检查图中是否有环，如果有环则无法完成所有课程
	return !graph.HasCycle()
}

// Dijkstra 迪杰斯特拉算法，求解单源最短路径
// 时间复杂度: O((V+E)logV)，V为顶点数，E为边数
// 空间复杂度: O(V)
func Dijkstra(graph *WeightedGraph, start int) []int {
	// 初始化距离数组，所有距离设为无穷大
	dist := make([]int, graph.Vertices)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[start] = 0 // 起点到自己的距离为0

	// 创建优先队列
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{value: start, priority: 0})

	// 记录已处理的顶点
	processed := make([]bool, graph.Vertices)

	for pq.Len() > 0 {
		// 取出当前距离最小的顶点
		item := heap.Pop(pq).(*Item)
		u := item.value

		// 如果已经处理过，跳过
		if processed[u] {
			continue
		}
		processed[u] = true

		// 更新所有相邻顶点的距离
		for _, edge := range graph.AdjList[u] {
			v, weight := edge.To, edge.Weight

			// 如果通过当前顶点可以得到更短的路径
			if !processed[v] && dist[u] != math.MaxInt32 && dist[u]+weight < dist[v] {
				dist[v] = dist[u] + weight
				heap.Push(pq, &Item{value: v, priority: dist[v]})
			}
		}
	}

	return dist
}

// WeightedGraph 带权图的定义
type WeightedGraph struct {
	Vertices   int
	AdjList    map[int][]Edge
	IsDirected bool
}

// Edge 边的定义
type Edge struct {
	To     int
	Weight int
}

// NewWeightedGraph 创建一个新的带权图
func NewWeightedGraph(vertices int, isDirected bool) *WeightedGraph {
	adjList := make(map[int][]Edge)
	for i := 0; i < vertices; i++ {
		adjList[i] = []Edge{}
	}

	return &WeightedGraph{
		Vertices:   vertices,
		AdjList:    adjList,
		IsDirected: isDirected,
	}
}

// AddEdge 添加带权边
func (g *WeightedGraph) AddEdge(src, dest, weight int) {
	// 检查顶点是否有效
	if src < 0 || src >= g.Vertices || dest < 0 || dest >= g.Vertices {
		return
	}

	// 添加边 src -> dest
	g.AdjList[src] = append(g.AdjList[src], Edge{To: dest, Weight: weight})

	// 如果是无向图，添加边 dest -> src
	if !g.IsDirected {
		g.AdjList[dest] = append(g.AdjList[dest], Edge{To: src, Weight: weight})
	}
}

// Item 优先队列中的元素
type Item struct {
	value    int // 顶点
	priority int // 距离（优先级）
	index    int // 在堆中的索引
}

// PriorityQueue 优先队列实现
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // 避免内存泄漏
	item.index = -1 // 标记为已删除
	*pq = old[0 : n-1]
	return item
}

// CloneGraph 克隆无向图
// 时间复杂度: O(V+E)，V为顶点数，E为边数
// 空间复杂度: O(V)
func CloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	// 使用map记录已经克隆的节点
	visited := make(map[*Node]*Node)

	return cloneGraphDFS(node, visited)
}

// Node 无向图节点定义
type Node struct {
	Val       int
	Neighbors []*Node
}

// cloneGraphDFS 深度优先搜索克隆图
func cloneGraphDFS(node *Node, visited map[*Node]*Node) *Node {
	// 如果节点已经被访问过，直接返回克隆的节点
	if clone, ok := visited[node]; ok {
		return clone
	}

	// 创建当前节点的克隆
	clone := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, 0, len(node.Neighbors)),
	}

	// 标记为已访问
	visited[node] = clone

	// 递归克隆所有邻居节点
	for _, neighbor := range node.Neighbors {
		clone.Neighbors = append(clone.Neighbors, cloneGraphDFS(neighbor, visited))
	}

	return clone
}
