package dfs_bfs

import (
	"reflect"
	"testing"
)

// 创建测试用的二叉树
func createTestTree() *TreeNode {
	/*
	       1
	      / \
	     2   3
	    / \   \
	   4   5   6
	*/
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 6}

	return root
}

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected int
	}{
		{
			name:     "空树",
			root:     nil,
			expected: 0,
		},
		{
			name:     "只有根节点",
			root:     &TreeNode{Val: 1},
			expected: 1,
		},
		{
			name:     "标准树",
			root:     createTestTree(),
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxDepth(tt.root)
			if result != tt.expected {
				t.Errorf("MaxDepth() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestMaxDepthBFS(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected int
	}{
		{
			name:     "空树",
			root:     nil,
			expected: 0,
		},
		{
			name:     "只有根节点",
			root:     &TreeNode{Val: 1},
			expected: 1,
		},
		{
			name:     "标准树",
			root:     createTestTree(),
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxDepthBFS(tt.root)
			if result != tt.expected {
				t.Errorf("MaxDepthBFS() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestLevelOrder(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected [][]int
	}{
		{
			name:     "空树",
			root:     nil,
			expected: [][]int{},
		},
		{
			name:     "只有根节点",
			root:     &TreeNode{Val: 1},
			expected: [][]int{{1}},
		},
		{
			name:     "标准树",
			root:     createTestTree(),
			expected: [][]int{{1}, {2, 3}, {4, 5, 6}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LevelOrder(tt.root)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("LevelOrder() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestZigzagLevelOrder(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected [][]int
	}{
		{
			name:     "空树",
			root:     nil,
			expected: [][]int{},
		},
		{
			name:     "只有根节点",
			root:     &TreeNode{Val: 1},
			expected: [][]int{{1}},
		},
		{
			name:     "标准树",
			root:     createTestTree(),
			expected: [][]int{{1}, {3, 2}, {4, 5, 6}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ZigzagLevelOrder(tt.root)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ZigzagLevelOrder() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestNumIslands(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]byte
		expected int
	}{
		{
			name:     "空网格",
			grid:     [][]byte{},
			expected: 0,
		},
		{
			name: "标准示例1",
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			expected: 1,
		},
		{
			name: "标准示例2",
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建网格副本，因为函数会修改原网格
			grid := make([][]byte, len(tt.grid))
			for i := range tt.grid {
				grid[i] = make([]byte, len(tt.grid[i]))
				copy(grid[i], tt.grid[i])
			}

			result := NumIslands(grid)
			if result != tt.expected {
				t.Errorf("NumIslands() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestSurroundedRegions(t *testing.T) {
	tests := []struct {
		name     string
		board    [][]byte
		expected [][]byte
	}{
		{
			name:     "空矩阵",
			board:    [][]byte{},
			expected: [][]byte{},
		},
		{
			name: "标准示例",
			board: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'X'},
			},
			expected: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
			},
		},
		{
			name: "边界上的O不变",
			board: [][]byte{
				{'X', 'O', 'X', 'X'},
				{'O', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'O'},
				{'X', 'X', 'O', 'X'},
			},
			expected: [][]byte{
				{'X', 'O', 'X', 'X'},
				{'O', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'O'},
				{'X', 'X', 'O', 'X'},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建矩阵副本，因为函数会修改原矩阵
			board := make([][]byte, len(tt.board))
			for i := range tt.board {
				board[i] = make([]byte, len(tt.board[i]))
				copy(board[i], tt.board[i])
			}

			SurroundedRegions(board)

			if !reflect.DeepEqual(board, tt.expected) {
				t.Errorf("SurroundedRegions() = %v, want %v", board, tt.expected)
			}
		})
	}
}

func TestWordSearch(t *testing.T) {
	tests := []struct {
		name     string
		board    [][]byte
		word     string
		expected bool
	}{
		{
			name:     "空网格",
			board:    [][]byte{},
			word:     "ABCCED",
			expected: false,
		},
		{
			name: "标准示例-存在",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:     "ABCCED",
			expected: true,
		},
		{
			name: "标准示例-不存在",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:     "ABCB",
			expected: false,
		},
		{
			name: "单个字符",
			board: [][]byte{
				{'A'},
			},
			word:     "A",
			expected: true,
		},
		{
			name: "长路径",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'E', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:     "ABCESEEEFS",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建网格副本，因为函数会修改原网格
			board := make([][]byte, len(tt.board))
			for i := range tt.board {
				board[i] = make([]byte, len(tt.board[i]))
				copy(board[i], tt.board[i])
			}

			result := WordSearch(board, tt.word)
			if result != tt.expected {
				t.Errorf("WordSearch() = %v, want %v", result, tt.expected)
			}
		})
	}
}
