/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
        return nil
    }

	interval :=1
	for interval < len(lists){
		for i := 0; i< len(lists) - interval; i += 2 * interval{
			lists[i] = merge(lists[i], lists[i+interval])
		}

		interval *= 2
	}

	return lists[0]
}

func merge(l1, l2 *ListNode) *ListNode{
	head := &ListNode{}
	dummy := head

	for l1 != nil && l2 != nil{
		if l1.Val < l2.Val{
			dummy.Next = l1
			l1 = l1.Next
		} else {
			dummy.Next = l2
			l2 = l2.Next
		}
		dummy = dummy.Next
	}

	if l1 == nil{
		dummy.Next = l2
	}else {
		dummy.Next = l1
	}

	return head.Next
}