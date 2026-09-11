/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head, next *ListNode) *ListNode {
	if next != nil {
		nextNext := next.Next
		next.Next = head
		return reverseList(next, nextNext)
	}
	return head
}

func findMiddle(slow, fast *ListNode) (prev, middle *ListNode) {
	newFast := fast.Next
	if newFast != nil {
		newFast = newFast.Next
	}
	if newFast == nil {
		return slow, slow.Next
	}

	return findMiddle(slow.Next, newFast)
}

func mergeLists(head1, head2 *ListNode) *ListNode {
	if head1 == nil {
		return head2
	}

	if head2 == nil {
		return head1
	}

	next1 := head1.Next
	next2 := head2.Next

	head1.Next = head2
	head2.Next = mergeLists(next1, next2)

	return head1
}

func reorderList(head *ListNode) {
	prev, middle := findMiddle(head, head)
	prev.Next = nil
	reversedHalf := reverseList(nil, middle)
	head = mergeLists(head, reversedHalf)
}
