/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    // A dummy node to simplify managing the new head and group connections
    dummy := &ListNode{Next: head}
    groupPrev := dummy

    for {
        // 1. Find the k-th node from groupPrev
        kth := getKth(groupPrev, k)
        if kth == nil {
            break
        }
        groupNext := kth.Next

        // 2. Reverse the k group
        prev := kth.Next
        curr := groupPrev.Next
        for curr != groupNext {
            temp := curr.Next
            curr.Next = prev
            prev = curr
            curr = temp
        }

        // 3. Connect the reversed group with the rest of the list
        tmp := groupPrev.Next
        groupPrev.Next = kth
        groupPrev = tmp
    }

    return dummy.Next
}

// Helper function to find the k-th node starting from curr
func getKth(curr *ListNode, k int) *ListNode {
    for curr != nil && k > 0 {
        curr = curr.Next
        k--
    }
    return curr
}