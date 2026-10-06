/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorder(root *TreeNode, res *[]int) {
    if root == nil {
        return
    }
    inorder(root.Left, res)
    *res = append(*res, root.Val)
    inorder(root.Right, res)
}

func build(traversal []int, left, right int) *TreeNode {
    if left > right {
        return nil
    }
    mid := left+(right-left)/2
    return &TreeNode {
        traversal[mid],
        build(traversal, left, mid-1),
        build(traversal, mid+1, right),
    }
}

func balanceBST(root *TreeNode) *TreeNode {
    traversal := make([]int, 0)
    inorder(root, &traversal)
    return build(traversal, 0, len(traversal)-1)
}
