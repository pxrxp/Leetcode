/**
 * Definition for a binary tree node.
 * class TreeNode(_value: Int = 0, _left: TreeNode = null, _right: TreeNode = null) {
 *   var value: Int = _value
 *   var left: TreeNode = _left
 *   var right: TreeNode = _right
 * }
 */

object Solution {
    def levelOrder(root: TreeNode): List[List[Int]] = {

        def traverseLevel(nodes: List[TreeNode]): List[List[Int]] = {
            if (nodes.isEmpty) Nil
            else {
                val currentLevelValues = nodes.map(_.value)
                val nextLevelNodes = nodes.flatMap(node => List(node.left, node.right).filter(_ != null))

                currentLevelValues :: traverseLevel(nextLevelNodes)
            }
        }

        if (root == null) Nil
        else traverseLevel(List(root))
    }
}
