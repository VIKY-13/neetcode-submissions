/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil{
		return nil
	}

	var prev *ListNode = nil
	curr := head

	for curr != nil{
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}

	return prev
}
