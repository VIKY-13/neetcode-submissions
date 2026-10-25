/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    stack := make([]*ListNode,0)

	dummy := head

	for dummy != nil{
		stack = append(stack, dummy)
		dummy = dummy.Next
	}

	if len(stack) == 0 || len(stack) <= 2{
		return
	}

	if len(stack) % 2 == 0{
		fmt.Println("even")
		stack = stack[len(stack)/2:]
	}else{
		fmt.Println("odd")
		stack = stack[(len(stack)/2)+1:]
	}

	dummy = head
	for len(stack) > 0{
		temp := dummy.Next
		sTop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		dummy.Next = sTop
		sTop.Next = temp
		dummy = temp
	}

	dummy.Next = nil
}
