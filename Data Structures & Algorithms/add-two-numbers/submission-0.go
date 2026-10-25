/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    l1Q, l2Q := make([]*ListNode,0), make([]*ListNode,0)

	for l1 != nil{
		l1Q = append(l1Q, l1)
		l1 = l1.Next
	}

	for l2 != nil{
		l2Q = append(l2Q, l2)
		l2 = l2.Next
	}

	carry := 0
	var curr *ListNode
	for len(l1Q) > 0 || len(l2Q) > 0 || carry != 0{
		l1Val, l2Val := 0, 0
		if len(l1Q) > 0{
			l1Val = l1Q[0].Val
			l1Q = l1Q[1:]
		}
		if len(l2Q) > 0{
			l2Val = l2Q[0].Val
			l2Q = l2Q[1:]
		}

		sum := l1Val + l2Val + carry
		val := sum % 10
		carry = sum / 10

		node := &ListNode{
			Val: val,
			Next: curr,
		}

		curr = node
	}

	var prev *ListNode
	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}

	return prev
}
