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

    @annotation.tailrec
    def traverse(current: List[TreeNode], acc: List[List[Int]] = Nil): List[List[Int]] = {
      if (current.isEmpty) acc.reverse
      else {
        val values = current.map(_.value)
        val next = 
            for {
                node <- current
                children <- List(node.left, node.right).flatMap(Option(_))
            } yield children
        traverse(next, values :: acc)
      }
    }

    Option(root).map(r => traverse(List(r))).getOrElse(Nil)
  }
}
