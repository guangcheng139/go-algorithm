package tree

// BST 二叉搜索树结构
type BST struct {
	Root *TreeNode
}

// NewBST 创建新的二叉搜索树
func NewBST() *BST {
	return &BST{
		Root: nil,
	}
}

// Insert 向二叉搜索树中插入值
func (bst *BST) Insert(val int) {
	if bst.Root == nil {
		bst.Root = NewTreeNode(val)
		return
	}

	insertRecursive(bst.Root, val)
}

// insertRecursive 递归插入节点
func insertRecursive(node *TreeNode, val int) {
	if val < node.Val {
		if node.Left == nil {
			node.Left = NewTreeNode(val)
		} else {
			insertRecursive(node.Left, val)
		}
	} else {
		if node.Right == nil {
			node.Right = NewTreeNode(val)
		} else {
			insertRecursive(node.Right, val)
		}
	}
}

// Search 在二叉搜索树中搜索值
func (bst *BST) Search(val int) bool {
	return searchRecursive(bst.Root, val)
}

// searchRecursive 递归搜索节点
func searchRecursive(node *TreeNode, val int) bool {
	if node == nil {
		return false
	}

	if val == node.Val {
		return true
	} else if val < node.Val {
		return searchRecursive(node.Left, val)
	} else {
		return searchRecursive(node.Right, val)
	}
}

// Delete 从二叉搜索树中删除值
func (bst *BST) Delete(val int) {
	bst.Root = deleteRecursive(bst.Root, val)
}

// deleteRecursive 递归删除节点
func deleteRecursive(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	// 查找要删除的节点
	if val < root.Val {
		root.Left = deleteRecursive(root.Left, val)
	} else if val > root.Val {
		root.Right = deleteRecursive(root.Right, val)
	} else {
		// 找到了要删除的节点

		// 情况1：叶子节点
		if root.Left == nil && root.Right == nil {
			return nil
		}

		// 情况2：只有一个子节点
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}

		// 情况3：有两个子节点
		// 找到右子树中的最小值，作为新的根节点
		successor := findMin(root.Right)
		root.Val = successor.Val
		// 从右子树中删除后继节点
		root.Right = deleteRecursive(root.Right, successor.Val)
	}

	return root
}

// findMin 查找树中的最小值节点
func findMin(node *TreeNode) *TreeNode {
	current := node

	// 一直向左走，直到找到最左边的节点
	for current != nil && current.Left != nil {
		current = current.Left
	}

	return current
}

// IsValidBST 验证二叉搜索树
// 时间复杂度: O(n)
// 空间复杂度: O(h)，h为树的高度
func IsValidBST(root *TreeNode) bool {
	return isValidBSTHelper(root, nil, nil)
}

// isValidBSTHelper 辅助函数，验证二叉搜索树
func isValidBSTHelper(node *TreeNode, min, max *int) bool {
	if node == nil {
		return true
	}

	// 检查当前节点的值是否在有效范围内
	if (min != nil && node.Val <= *min) || (max != nil && node.Val >= *max) {
		return false
	}

	// 递归验证左右子树
	return isValidBSTHelper(node.Left, min, &node.Val) && isValidBSTHelper(node.Right, &node.Val, max)
}
