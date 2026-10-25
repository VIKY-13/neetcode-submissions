/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	if head == nil{
		return false
	}
    slow, fast := head,head

	for fast.Next != nil && fast.Next.Next != nil{
		fast = fast.Next.Next
		if slow == fast{
			return true
		}
		slow = slow.Next
	}

	return false
}
