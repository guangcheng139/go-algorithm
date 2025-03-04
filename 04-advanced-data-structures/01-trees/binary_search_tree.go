package trees

// TreeNode 表示二叉树中的一个节点
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// BST 表示二叉搜索树
type BST struct {
	Root *TreeNode
}

// NewBST 创建一个新的二叉搜索树
func NewBST() *BST {
	return &BST{
		Root: nil,
	}
}

// Insert 向二叉搜索树中插入一个值
func (bst *BST) Insert(value int) {
	bst.Root = insertRecursive(bst.Root, value)
}

// insertRecursive 递归插入值到二叉搜索树
func insertRecursive(root *TreeNode, value int) *TreeNode {
	// 如果树为空，创建一个新节点
	if root == nil {
		return &TreeNode{
			Value: value,
			Left:  nil,
			Right: nil,
		}
	}

	// 如果值小于当前节点的值，插入到左子树
	if value < root.Value {
		root.Left = insertRecursive(root.Left, value)
	} else if value > root.Value { // 如果值大于当前节点的值，插入到右子树
		root.Right = insertRecursive(root.Right, value)
	}
	// 如果值等于当前节点的值，不做任何操作（BST通常不允许重复值）

	return root
}

// Search 在二叉搜索树中查找一个值
// 返回包含该值的节点，如果未找到则返回nil
func (bst *BST) Search(value int) *TreeNode {
	return searchRecursive(bst.Root, value)
}

// searchRecursive 递归查找值
func searchRecursive(root *TreeNode, value int) *TreeNode {
	// 如果树为空或找到了值
	if root == nil || root.Value == value {
		return root
	}

	// 如果值小于当前节点的值，在左子树中查找
	if value < root.Value {
		return searchRecursive(root.Left, value)
	}

	// 如果值大于当前节点的值，在右子树中查找
	return searchRecursive(root.Right, value)
}

// Delete 从二叉搜索树中删除一个值
func (bst *BST) Delete(value int) {
	bst.Root = deleteRecursive(bst.Root, value)
}

// deleteRecursive 递归删除值
func deleteRecursive(root *TreeNode, value int) *TreeNode {
	// 如果树为空，返回nil
	if root == nil {
		return nil
	}

	// 如果值小于当前节点的值，在左子树中删除
	if value < root.Value {
		root.Left = deleteRecursive(root.Left, value)
	} else if value > root.Value { // 如果值大于当前节点的值，在右子树中删除
		root.Right = deleteRecursive(root.Right, value)
	} else { // 找到了要删除的节点
		// 情况1：叶子节点（没有子节点）
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
		// 找到右子树中的最小值（后继节点）
		root.Value = findMin(root.Right).Value
		// 从右子树中删除后继节点
		root.Right = deleteRecursive(root.Right, root.Value)
	}

	return root
}

// findMin 找到树中的最小值节点
func findMin(root *TreeNode) *TreeNode {
	current := root
	// 一直向左遍历，直到找到最左边的节点
	for current != nil && current.Left != nil {
		current = current.Left
	}
	return current
}

// InOrderTraversal 中序遍历二叉搜索树
// 返回按升序排列的值
func (bst *BST) InOrderTraversal() []int {
	var result []int
	inOrderTraversalRecursive(bst.Root, &result)
	return result
}

// inOrderTraversalRecursive 递归中序遍历
func inOrderTraversalRecursive(root *TreeNode, result *[]int) {
	if root != nil {
		inOrderTraversalRecursive(root.Left, result)
		*result = append(*result, root.Value)
		inOrderTraversalRecursive(root.Right, result)
	}
}

// PreOrderTraversal 前序遍历二叉搜索树
func (bst *BST) PreOrderTraversal() []int {
	var result []int
	preOrderTraversalRecursive(bst.Root, &result)
	return result
}

// preOrderTraversalRecursive 递归前序遍历
func preOrderTraversalRecursive(root *TreeNode, result *[]int) {
	if root != nil {
		*result = append(*result, root.Value)
		preOrderTraversalRecursive(root.Left, result)
		preOrderTraversalRecursive(root.Right, result)
	}
}

// PostOrderTraversal 后序遍历二叉搜索树
func (bst *BST) PostOrderTraversal() []int {
	var result []int
	postOrderTraversalRecursive(bst.Root, &result)
	return result
}

// postOrderTraversalRecursive 递归后序遍历
func postOrderTraversalRecursive(root *TreeNode, result *[]int) {
	if root != nil {
		postOrderTraversalRecursive(root.Left, result)
		postOrderTraversalRecursive(root.Right, result)
		*result = append(*result, root.Value)
	}
}

// Height 计算二叉树的高度
func (bst *BST) Height() int {
	return heightRecursive(bst.Root)
}

// heightRecursive 递归计算高度
func heightRecursive(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := heightRecursive(root.Left)
	rightHeight := heightRecursive(root.Right)

	// 返回左右子树中较大的高度加1
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}
