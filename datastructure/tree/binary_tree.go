package tree

// TreeNode 二叉树节点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NewTreeNode 创建新的二叉树节点
func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

// PreorderTraversal 前序遍历（递归）
// 时间复杂度: O(n)
// 空间复杂度: O(h)，h为树的高度
func PreorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}

	// 根-左-右
	result = append(result, root.Val)
	result = append(result, PreorderTraversal(root.Left)...)
	result = append(result, PreorderTraversal(root.Right)...)

	return result
}

// InorderTraversal 中序遍历（递归）
// 时间复杂度: O(n)
// 空间复杂度: O(h)，h为树的高度
func InorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}

	// 左-根-右
	result = append(result, InorderTraversal(root.Left)...)
	result = append(result, root.Val)
	result = append(result, InorderTraversal(root.Right)...)

	return result
}

// PostorderTraversal 后序遍历（递归）
// 时间复杂度: O(n)
// 空间复杂度: O(h)，h为树的高度
func PostorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}

	// 左-右-根
	result = append(result, PostorderTraversal(root.Left)...)
	result = append(result, PostorderTraversal(root.Right)...)
	result = append(result, root.Val)

	return result
}

// LevelOrderTraversal 层序遍历
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func LevelOrderTraversal(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		levelValues := []int{}

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:] // 出队

			levelValues = append(levelValues, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, levelValues)
	}

	return result
}

// MaxDepth 二叉树的最大深度
// 时间复杂度: O(n)
// 空间复杂度: O(h)，h为树的高度
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := MaxDepth(root.Left)
	rightDepth := MaxDepth(root.Right)

	// 返回左右子树深度的较大值 + 1（当前节点）
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}
