package tree

import (
	"reflect"
	"testing"
)

func TestNewTreeNode(t *testing.T) {
	node := NewTreeNode(5)
	if node.Val != 5 {
		t.Errorf("NewTreeNode() = %v, want %v", node.Val, 5)
	}
	if node.Left != nil || node.Right != nil {
		t.Errorf("NewTreeNode() created with non-nil children")
	}
}

func createTestTree() *TreeNode {
	/*
	       1
	      / \
	     2   3
	    / \   \
	   4   5   6
	*/
	root := NewTreeNode(1)
	root.Left = NewTreeNode(2)
	root.Right = NewTreeNode(3)
	root.Left.Left = NewTreeNode(4)
	root.Left.Right = NewTreeNode(5)
	root.Right.Right = NewTreeNode(6)

	return root
}

func TestPreorderTraversal(t *testing.T) {
	root := createTestTree()

	expected := []int{1, 2, 4, 5, 3, 6}
	result := PreorderTraversal(root)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("PreorderTraversal() = %v, want %v", result, expected)
	}
}

func TestInorderTraversal(t *testing.T) {
	root := createTestTree()

	expected := []int{4, 2, 5, 1, 3, 6}
	result := InorderTraversal(root)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("InorderTraversal() = %v, want %v", result, expected)
	}
}

func TestPostorderTraversal(t *testing.T) {
	root := createTestTree()

	expected := []int{4, 5, 2, 6, 3, 1}
	result := PostorderTraversal(root)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("PostorderTraversal() = %v, want %v", result, expected)
	}
}

func TestLevelOrderTraversal(t *testing.T) {
	root := createTestTree()

	expected := [][]int{
		{1},
		{2, 3},
		{4, 5, 6},
	}
	result := LevelOrderTraversal(root)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("LevelOrderTraversal() = %v, want %v", result, expected)
	}
}

func TestMaxDepth(t *testing.T) {
	root := createTestTree()

	expected := 3
	result := MaxDepth(root)

	if result != expected {
		t.Errorf("MaxDepth() = %v, want %v", result, expected)
	}
}

func TestIsSameTree(t *testing.T) {
	tree1 := createTestTree()
	tree2 := createTestTree()

	if !IsSameTree(tree1, tree2) {
		t.Errorf("IsSameTree() = false, want true for identical trees")
	}

	// 修改第二棵树
	tree2.Left.Right.Val = 10

	if IsSameTree(tree1, tree2) {
		t.Errorf("IsSameTree() = true, want false for different trees")
	}
}

func TestIsSymmetric(t *testing.T) {
	// 创建对称二叉树
	/*
	       1
	      / \
	     2   2
	    / \ / \
	   3  4 4  3
	*/
	root := NewTreeNode(1)
	root.Left = NewTreeNode(2)
	root.Right = NewTreeNode(2)
	root.Left.Left = NewTreeNode(3)
	root.Left.Right = NewTreeNode(4)
	root.Right.Left = NewTreeNode(4)
	root.Right.Right = NewTreeNode(3)

	if !IsSymmetric(root) {
		t.Errorf("IsSymmetric() = false, want true for symmetric tree")
	}

	// 修改为非对称
	root.Right.Right.Val = 5

	if IsSymmetric(root) {
		t.Errorf("IsSymmetric() = true, want false for asymmetric tree")
	}
}

func TestLowestCommonAncestor(t *testing.T) {
	root := createTestTree()

	// 节点4和节点5的LCA是节点2
	p := root.Left.Left   // 节点4
	q := root.Left.Right  // 节点5
	expected := root.Left // 节点2

	result := LowestCommonAncestor(root, p, q)

	if result != expected {
		t.Errorf("LowestCommonAncestor() = %v, want %v", result.Val, expected.Val)
	}

	// 节点4和节点6的LCA是节点1
	p = root.Left.Left   // 节点4
	q = root.Right.Right // 节点6
	expected = root      // 节点1

	result = LowestCommonAncestor(root, p, q)

	if result != expected {
		t.Errorf("LowestCommonAncestor() = %v, want %v", result.Val, expected.Val)
	}
}

func TestBSTOperations(t *testing.T) {
	bst := NewBST()
	bst.Insert(5)
	bst.Insert(3)
	bst.Insert(7)
	bst.Insert(2)
	bst.Insert(4)
	bst.Insert(6)
	bst.Insert(8)

	// 验证搜索功能
	if !bst.Search(5) {
		t.Errorf("BST.Search(5) = false, want true")
	}
	if !bst.Search(2) {
		t.Errorf("BST.Search(2) = false, want true")
	}
	if bst.Search(9) {
		t.Errorf("BST.Search(9) = true, want false")
	}

	// 验证删除功能
	bst.Delete(3)
	if bst.Search(3) {
		t.Errorf("BST.Search(3) = true after deletion, want false")
	}
	if !bst.Search(2) || !bst.Search(4) {
		t.Errorf("Children of deleted node not preserved")
	}
}

func TestIsValidBST(t *testing.T) {
	// 创建有效的BST
	/*
	       5
	      / \
	     3   7
	    / \ / \
	   2  4 6  8
	*/
	root := NewTreeNode(5)
	root.Left = NewTreeNode(3)
	root.Right = NewTreeNode(7)
	root.Left.Left = NewTreeNode(2)
	root.Left.Right = NewTreeNode(4)
	root.Right.Left = NewTreeNode(6)
	root.Right.Right = NewTreeNode(8)

	if !IsValidBST(root) {
		t.Errorf("IsValidBST() = false, want true for valid BST")
	}

	// 修改为无效的BST
	root.Right.Left.Val = 4 // 错误：右子树有小于根节点的值

	if IsValidBST(root) {
		t.Errorf("IsValidBST() = true, want false for invalid BST")
	}
}
