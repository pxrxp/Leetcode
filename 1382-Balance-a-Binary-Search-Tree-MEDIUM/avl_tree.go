/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var heights map[*TreeNode]int

func inorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, res)
	*res = append(*res, root.Val)
	inorder(root.Right, res)
}

// Difference in height between left and right side of tree
func balance(root *TreeNode) int {
	return height(root.Left, false) - height(root.Right, false)
}

func height(root *TreeNode, recompute bool) int {
	if root == nil {
		return -1
	}

	h := -1

	h_memoized, exists := heights[root]
	if exists && !recompute {
		return h_memoized
	}

	l, r := height(root.Left, false), height(root.Right, false)
	if l < r {
		h = r + 1
	} else {
		h = l + 1
	}
	heights[root] = h
	return h
}

func leftRotate(root *TreeNode) *TreeNode {
	// x < ((x2) < x1 < (x3))
	// becomes
	// (x < (x2)) < x1 < (x3)

	x := root
	x1 := x.Right
	x2 := x1.Left

	x.Right = x2
	x1.Left = x

	_ = height(x, true)
	_ = height(x1, true)

	return x1
}

func rightRotate(root *TreeNode) *TreeNode {
	// ((x2) < x1 < (x3)) < x
	// becomes
	// (x2) < x1 < ((x3) < x)

	x := root
	x1 := x.Left
	x3 := x1.Right

	x.Left = x3
	x1.Right = x

	_ = height(x, true)
	_ = height(x1, true)

	return x1
}

func insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		node := &TreeNode{val, nil, nil}
		_ = height(node, true)
		return node
	}

	if val < root.Val {
		root.Left = insert(root.Left, val)
	} else if val > root.Val {
		root.Right = insert(root.Right, val)
	} else {
		return root
	}

	_ = height(root, true)

	if balance(root) > 1 {
		// Left heavy
		if balance(root.Left) >= 0 {
			// Left-Left heavy
			root = rightRotate(root)

			// (x2 < x1) < x <
			// becomes
			// (x2) < x < (x1)
		} else {
			// Left-Right heavy
			root.Left = leftRotate(root.Left)
			root = rightRotate(root)

			// (x2 > x1) < x <
			// becomes
			// (x1 < x2) < x <
			// becomes
			// (x1) < x2 < (x)
		}
	} else if balance(root) < -1 {
		// Right heavy
		if balance(root.Right) > 0 {
			// Right-Left heavy
			root.Right = rightRotate(root.Right)
			root = leftRotate(root)

			// < x < (x1 > x2)
			// becomes
			// < x < (x2 < x1)
			// becomes
			// (x) < x2 < (x1)
		} else {
			// Right-Right heavy
			root = leftRotate(root)

			// < x < (x1 < x2)
			// becomes
			// (x) < x1 < (x2)
		}
	}

	// Balanced
	return root
}

func balanceBST(root *TreeNode) *TreeNode {
	heights = make(map[*TreeNode]int)
	traversal := make([]int, 0)
	inorder(root, &traversal)
	var newRoot *TreeNode
	for _, v := range traversal {
		newRoot = insert(newRoot, v)
	}
	return newRoot
}
