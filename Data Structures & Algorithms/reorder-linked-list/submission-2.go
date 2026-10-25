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

	// 1. Find the middle of the list
	slow,fast := head,head
	for fast.Next !=nil && fast.Next.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 2. Reverse the second half of the list
	second := slow.Next
	slow.Next = nil	// Split the list into two halves
	var prev *ListNode
	for second != nil{
		temp := second.Next
		second.Next = prev
		prev = second
		second = temp
	}

	// 3. Merge the two halves (first and reversed second)
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
