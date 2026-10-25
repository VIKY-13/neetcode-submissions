/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
    dummy := &ListNode{
		Next: head,
	}
	groupPrev := dummy

	for{
		kth := getKth(groupPrev, k)
		if kth == nil{
			break
		}

		groupNext := kth.Next
		prev := kth.Next
		curr := groupPrev.Next
		for {
			if curr == groupNext{
				break
			}
			temp := curr.Next
			curr.Next = prev
			prev = curr
			curr = temp
		}

		tmp := groupPrev.Next
		groupPrev.Next = kth
		groupPrev = tmp
	}

	return dummy.Next
}


func getKth(node *ListNode, k int) *ListNode{
	for node != nil && k > 0{
		k--
		node = node.Next
	}

	return node
}