package graph

// Graph 表示基于邻接表的图
type Graph struct {
	Vertices   int            // 顶点数量
	IsDirected bool           // 是否为有向图
	AdjList    map[int][]Edge // 邻接表
}

// Edge 表示图中的一条边
type Edge struct {
	To     int // 目标顶点
	Weight int // 边的权重
}

// NewGraph 创建一个新的图
func NewGraph(vertices int, isDirected bool) *Graph {
	return &Graph{
		Vertices:   vertices,
		IsDirected: isDirected,
		AdjList:    make(map[int][]Edge),
	}
}

// AddEdge 向图中添加一条边
func (g *Graph) AddEdge(from, to, weight int) {
	// 检查顶点是否有效
	if from < 0 || from >= g.Vertices || to < 0 || to >= g.Vertices {
		return
	}

	// 添加从from到to的边
	g.AdjList[from] = append(g.AdjList[from], Edge{To: to, Weight: weight})

	// 如果是无向图，还需要添加从to到from的边
	if !g.IsDirected {
		g.AdjList[to] = append(g.AdjList[to], Edge{To: from, Weight: weight})
	}
}

// GetNeighbors 获取顶点的所有邻居
func (g *Graph) GetNeighbors(vertex int) []Edge {
	// 检查顶点是否有效
	if vertex < 0 || vertex >= g.Vertices {
		return nil
	}

	return g.AdjList[vertex]
}

// HasEdge 检查两个顶点之间是否存在边
func (g *Graph) HasEdge(from, to int) bool {
	// 检查顶点是否有效
	if from < 0 || from >= g.Vertices || to < 0 || to >= g.Vertices {
		return false
	}

	// 检查from到to的边是否存在
	for _, edge := range g.AdjList[from] {
		if edge.To == to {
			return true
		}
	}

	return false
}

// GetEdgeWeight 获取两个顶点之间边的权重
// 如果边不存在，返回-1
func (g *Graph) GetEdgeWeight(from, to int) int {
	// 检查顶点是否有效
	if from < 0 || from >= g.Vertices || to < 0 || to >= g.Vertices {
		return -1
	}

	// 查找from到to的边
	for _, edge := range g.AdjList[from] {
		if edge.To == to {
			return edge.Weight
		}
	}

	// 边不存在
	return -1
}

// RemoveEdge 从图中移除一条边
func (g *Graph) RemoveEdge(from, to int) {
	// 检查顶点是否有效
	if from < 0 || from >= g.Vertices || to < 0 || to >= g.Vertices {
		return
	}

	// 移除from到to的边
	var newEdges []Edge
	for _, edge := range g.AdjList[from] {
		if edge.To != to {
			newEdges = append(newEdges, edge)
		}
	}
	g.AdjList[from] = newEdges

	// 如果是无向图，还需要移除to到from的边
	if !g.IsDirected {
		newEdges = []Edge{}
		for _, edge := range g.AdjList[to] {
			if edge.To != from {
				newEdges = append(newEdges, edge)
			}
		}
		g.AdjList[to] = newEdges
	}
}

// GetAllVertices 获取图中的所有顶点
func (g *Graph) GetAllVertices() []int {
	vertices := make([]int, 0, g.Vertices)
	for i := 0; i < g.Vertices; i++ {
		vertices = append(vertices, i)
	}
	return vertices
}

// GetAllEdges 获取图中的所有边
func (g *Graph) GetAllEdges() [][3]int {
	var edges [][3]int

	for from, neighbors := range g.AdjList {
		for _, edge := range neighbors {
			// 如果是无向图，只添加from < to的边，避免重复
			if g.IsDirected || from < edge.To {
				edges = append(edges, [3]int{from, edge.To, edge.Weight})
			}
		}
	}

	return edges
}

// GetOutDegree 获取顶点的出度（从该顶点出发的边数）
func (g *Graph) GetOutDegree(vertex int) int {
	// 检查顶点是否有效
	if vertex < 0 || vertex >= g.Vertices {
		return 0
	}

	return len(g.AdjList[vertex])
}

// GetInDegree 获取顶点的入度（指向该顶点的边数）
func (g *Graph) GetInDegree(vertex int) int {
	// 检查顶点是否有效
	if vertex < 0 || vertex >= g.Vertices {
		return 0
	}

	count := 0
	for v, neighbors := range g.AdjList {
		if v == vertex && !g.IsDirected {
			// 对于无向图，自环算作入度
			for _, edge := range neighbors {
				if edge.To == vertex {
					count++
				}
			}
		} else {
			// 计算其他顶点指向该顶点的边
			for _, edge := range neighbors {
				if edge.To == vertex {
					count++
				}
			}
		}
	}

	return count
}
