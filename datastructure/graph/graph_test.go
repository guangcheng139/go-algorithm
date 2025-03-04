package graph

import (
	"reflect"
	"testing"
)

func TestGraphCreation(t *testing.T) {
	undirectedGraph := NewGraph(5, false)

	if undirectedGraph.Vertices != 5 {
		t.Errorf("NewGraph() Vertices = %v, want %v", undirectedGraph.Vertices, 5)
	}

	if undirectedGraph.IsDirected {
		t.Errorf("NewGraph() IsDirected = %v, want %v", undirectedGraph.IsDirected, false)
	}

	directedGraph := NewGraph(3, true)
	if !directedGraph.IsDirected {
		t.Errorf("NewGraph() IsDirected = %v, want %v", directedGraph.IsDirected, true)
	}
}

func TestAddEdge(t *testing.T) {
	// 测试无向图
	undirectedGraph := NewGraph(5, false)
	undirectedGraph.AddEdge(0, 1)
	undirectedGraph.AddEdge(0, 4)
	undirectedGraph.AddEdge(1, 2)

	// 验证边是否正确添加
	if !contains(undirectedGraph.AdjList[0], 1) {
		t.Errorf("Edge 0->1 not found in undirected graph")
	}
	if !contains(undirectedGraph.AdjList[1], 0) {
		t.Errorf("Edge 1->0 not found in undirected graph")
	}

	// 测试有向图
	directedGraph := NewGraph(5, true)
	directedGraph.AddEdge(0, 1)
	directedGraph.AddEdge(0, 4)
	directedGraph.AddEdge(1, 2)

	// 验证边是否正确添加
	if !contains(directedGraph.AdjList[0], 1) {
		t.Errorf("Edge 0->1 not found in directed graph")
	}
	if contains(directedGraph.AdjList[1], 0) {
		t.Errorf("Edge 1->0 should not exist in directed graph")
	}
}

func TestRemoveEdge(t *testing.T) {
	// 测试无向图
	undirectedGraph := NewGraph(5, false)
	undirectedGraph.AddEdge(0, 1)
	undirectedGraph.AddEdge(0, 4)

	// 移除边
	undirectedGraph.RemoveEdge(0, 1)

	// 验证边是否正确移除
	if contains(undirectedGraph.AdjList[0], 1) {
		t.Errorf("Edge 0->1 should be removed from undirected graph")
	}
	if contains(undirectedGraph.AdjList[1], 0) {
		t.Errorf("Edge 1->0 should be removed from undirected graph")
	}

	// 测试有向图
	directedGraph := NewGraph(5, true)
	directedGraph.AddEdge(0, 1)
	directedGraph.AddEdge(0, 4)

	// 移除边
	directedGraph.RemoveEdge(0, 1)

	// 验证边是否正确移除
	if contains(directedGraph.AdjList[0], 1) {
		t.Errorf("Edge 0->1 should be removed from directed graph")
	}
}

func TestBFS(t *testing.T) {
	graph := createTestGraph()

	result := graph.BFS(0)
	expected := []int{0, 1, 2, 3}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("BFS() = %v, want %v", result, expected)
	}
}

func TestDFS(t *testing.T) {
	graph := createTestGraph()

	result := graph.DFS(0)

	// DFS结果依赖于实现，但是每个顶点应该只访问一次
	if len(result) != graph.Vertices {
		t.Errorf("DFS() visited %v vertices, want %v", len(result), graph.Vertices)
	}

	// 检查是否有重复访问
	visited := make(map[int]bool)
	for _, v := range result {
		if visited[v] {
			t.Errorf("DFS() visited vertex %v multiple times", v)
		}
		visited[v] = true
	}
}

func TestHasCycle(t *testing.T) {
	// 创建无环无向图
	acyclicUndirected := NewGraph(3, false)
	acyclicUndirected.AddEdge(0, 1)
	acyclicUndirected.AddEdge(1, 2)

	if acyclicUndirected.HasCycle() {
		t.Errorf("HasCycle() = true, want false for acyclic undirected graph")
	}

	// 创建有环无向图
	cyclicUndirected := NewGraph(3, false)
	cyclicUndirected.AddEdge(0, 1)
	cyclicUndirected.AddEdge(1, 2)
	cyclicUndirected.AddEdge(2, 0)

	if !cyclicUndirected.HasCycle() {
		t.Errorf("HasCycle() = false, want true for cyclic undirected graph")
	}

	// 创建无环有向图
	acyclicDirected := NewGraph(3, true)
	acyclicDirected.AddEdge(0, 1)
	acyclicDirected.AddEdge(1, 2)

	if acyclicDirected.HasCycle() {
		t.Errorf("HasCycle() = true, want false for acyclic directed graph")
	}

	// 创建有环有向图
	cyclicDirected := NewGraph(3, true)
	cyclicDirected.AddEdge(0, 1)
	cyclicDirected.AddEdge(1, 2)
	cyclicDirected.AddEdge(2, 0)

	if !cyclicDirected.HasCycle() {
		t.Errorf("HasCycle() = false, want true for cyclic directed graph")
	}
}

func TestTopologicalSort(t *testing.T) {
	// 创建无环有向图
	dag := NewGraph(6, true)
	dag.AddEdge(5, 2)
	dag.AddEdge(5, 0)
	dag.AddEdge(4, 0)
	dag.AddEdge(4, 1)
	dag.AddEdge(2, 3)
	dag.AddEdge(3, 1)

	result := dag.TopologicalSort()

	// 检查结果是否为有效的拓扑排序
	visited := make([]bool, dag.Vertices)
	for _, v := range result {
		visited[v] = true
		for _, adj := range dag.AdjList[v] {
			if visited[adj] {
				t.Errorf("TopologicalSort() invalid ordering: %v should come before %v", v, adj)
			}
		}
	}

	// 创建有环有向图
	cyclicGraph := NewGraph(3, true)
	cyclicGraph.AddEdge(0, 1)
	cyclicGraph.AddEdge(1, 2)
	cyclicGraph.AddEdge(2, 0)

	result = cyclicGraph.TopologicalSort()
	if len(result) != 0 {
		t.Errorf("TopologicalSort() = %v, want empty slice for cyclic graph", result)
	}
}

func TestNumIslands(t *testing.T) {
	grid1 := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}

	if count := NumIslands(grid1); count != 1 {
		t.Errorf("NumIslands() = %v, want %v", count, 1)
	}

	grid2 := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	if count := NumIslands(grid2); count != 3 {
		t.Errorf("NumIslands() = %v, want %v", count, 3)
	}
}

func TestCanFinish(t *testing.T) {
	// 可以完成的课程
	prerequisites1 := [][]int{{1, 0}}
	if !CanFinish(2, prerequisites1) {
		t.Errorf("CanFinish() = false, want true")
	}

	// 不可以完成的课程（有环）
	prerequisites2 := [][]int{{1, 0}, {0, 1}}
	if CanFinish(2, prerequisites2) {
		t.Errorf("CanFinish() = true, want false")
	}
}

func TestWeightedGraph(t *testing.T) {
	graph := NewWeightedGraph(4, true)
	graph.AddEdge(0, 1, 1)
	graph.AddEdge(0, 2, 4)
	graph.AddEdge(1, 2, 2)
	graph.AddEdge(1, 3, 5)
	graph.AddEdge(2, 3, 1)

	// 测试单源最短路径
	dist := Dijkstra(graph, 0)
	expectedDist := []int{0, 1, 3, 4}

	if !reflect.DeepEqual(dist, expectedDist) {
		t.Errorf("Dijkstra() = %v, want %v", dist, expectedDist)
	}
}

// 辅助函数

// createTestGraph 创建测试用的图
func createTestGraph() *Graph {
	/*
	   0 -- 1
	   |    |
	   2 -- 3
	*/
	graph := NewGraph(4, false)
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 3)

	return graph
}

// contains 检查切片中是否包含某个元素
func contains(slice []int, item int) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
