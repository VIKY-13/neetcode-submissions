/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    if head == nil || head.Next == nil{
		return
	}

	slow,fast := head,head
	for fast.Next !=nil && fast.Next.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next
	}

	second := slow.Next
	slow.Next = nil
	var prev *ListNode
	for second != nil{
		temp := second.Next
		second.Next = prev
		prev = second
		second = temp
	}

	dummy := head
	second = prev
	for second != nil{
		firstNext := dummy.Next
		secondNext := second.Next

		dummy.Next = second
		second.Next = firstNext

		dummy = firstNext
		second = secondNext
	}
}
