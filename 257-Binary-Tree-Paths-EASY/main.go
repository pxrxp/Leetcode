/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var answer []string

func dfs(root *TreeNode, pathFollowed *[]int) {
    *pathFollowed = append(*pathFollowed, root.Val)

    if root.Left == nil && root.Right == nil {
        answer = append(answer, format(*pathFollowed))
        return
    }
    if root.Left != nil {
        dfs(root.Left, pathFollowed)
        *pathFollowed = pop(*pathFollowed)
    }
    if root.Right != nil {
        dfs(root.Right, pathFollowed)
        *pathFollowed = pop(*pathFollowed)
    }
}

func format(pathFollowed []int) string {
    l := len(pathFollowed)
    pathStr := make([]byte, 0)
    for i, v := range pathFollowed {
        pathStr = strconv.AppendInt(pathStr, int64(v), 10)
        if i != l-1 {
            pathStr = append(pathStr, '-')
            pathStr = append(pathStr, '>')
        }
    }
    return string(pathStr)
}

func pop(arr []int) []int {
    if l := len(arr); l > 0 {
        return (arr)[:l-1]
    }
    return arr
}

func binaryTreePaths(root *TreeNode) []string {
    answer = make([]string, 0)
    path := make([]int, 0)
    dfs(root, &path)
    return answer
}
