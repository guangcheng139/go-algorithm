package tree

import (
	"strconv"
	"strings"
)

// IsSameTree 判断两棵树是否相同
// 时间复杂度: O(n)
// 空间复杂度: O(h)，h为树的高度
func IsSameTree(p, q *TreeNode) bool {
	// 都为空，返回true
	if p == nil && q == nil {
		return true
	}

	// 一个为空，一个不为空，返回false
	if p == nil || q == nil {
		return false
	}

	// 值不相等，返回false
	if p.Val != q.Val {
		return false
	}

	// 递归判断左右子树是否相同
	return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}

// IsSymmetric 判断二叉树是否对称
// 时间复杂度: O(n)
// 空间复杂度: O(h)，h为树的高度
func IsSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSymmetricHelper(root.Left, root.Right)
}

// isSymmetricHelper 辅助函数
func isSymmetricHelper(left, right *TreeNode) bool {
	// 两个节点都为空，对称
	if left == nil && right == nil {
		return true
	}

	// 其中一个为空，不对称
	if left == nil || right == nil {
		return false
	}

	// 节点值不相等，不对称
	if left.Val != right.Val {
		return false
	}

	// 递归判断：左子树的左子节点和右子树的右子节点是否对称
	// 左子树的右子节点和右子树的左子节点是否对称
	return isSymmetricHelper(left.Left, right.Right) && isSymmetricHelper(left.Right, right.Left)
}

// LowestCommonAncestor 二叉树的最近公共祖先
// 时间复杂度: O(n)
// 空间复杂度: O(h)，h为树的高度
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 边界情况
	if root == nil || root == p || root == q {
		return root
	}

	// 在左子树中寻找
	left := LowestCommonAncestor(root.Left, p, q)
	// 在右子树中寻找
	right := LowestCommonAncestor(root.Right, p, q)

	// 如果左子树和右子树都找到了，说明根节点就是LCA
	if left != nil && right != nil {
		return root
	}

	// 如果只在左子树找到，说明LCA在左子树
	if left != nil {
		return left
	}

	// 如果只在右子树找到，说明LCA在右子树
	return right
}

// Serialize 二叉树序列化
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func Serialize(root *TreeNode) string {
	if root == nil {
		return "null"
	}

	// 前序遍历序列化
	return serializeHelper(root, "")
}

// serializeHelper 序列化辅助函数
func serializeHelper(node *TreeNode, s string) string {
	if node == nil {
		return s + "null,"
	}

	s += strconv.Itoa(node.Val) + ","
	s = serializeHelper(node.Left, s)
	s = serializeHelper(node.Right, s)

	return s
}

// Deserialize 二叉树反序列化
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func Deserialize(data string) *TreeNode {
	if data == "null" {
		return nil
	}

	nodes := strings.Split(data, ",")
	return deserializeHelper(&nodes)
}

// deserializeHelper 反序列化辅助函数
func deserializeHelper(nodes *[]string) *TreeNode {
	if len(*nodes) == 0 {
		return nil
	}

	val := (*nodes)[0]
	*nodes = (*nodes)[1:]

	if val == "null" {
		return nil
	}

	intVal, _ := strconv.Atoi(val)
	node := NewTreeNode(intVal)

	node.Left = deserializeHelper(nodes)
	node.Right = deserializeHelper(nodes)

	return node
}

// PathSum 路径总和 III
// 给定一个二叉树，它的每个结点都存放一个整数值。
// 找出路径和等于给定数值的路径总数。
// 注意: 路径不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的
// 时间复杂度: O(n^2) 最坏情况
// 空间复杂度: O(h)，h为树的高度
func PathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	// 从当前节点开始的路径总数
	pathsFromRoot := pathSumFromNode(root, sum)

	// 从左子节点开始的路径总数
	pathsFromLeft := PathSum(root.Left, sum)

	// 从右子节点开始的路径总数
	pathsFromRight := PathSum(root.Right, sum)

	return pathsFromRoot + pathsFromLeft + pathsFromRight
}

// pathSumFromNode 计算从某个节点开始的路径总数
func pathSumFromNode(node *TreeNode, sum int) int {
	if node == nil {
		return 0
	}

	count := 0

	// 当前节点值等于目标和，找到一条路径
	if node.Val == sum {
		count++
	}

	// 继续搜索左右子树
	count += pathSumFromNode(node.Left, sum-node.Val)
	count += pathSumFromNode(node.Right, sum-node.Val)

	return count
}
