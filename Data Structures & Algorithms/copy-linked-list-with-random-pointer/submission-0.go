/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    dummy := head

	hash := make(map[*Node]*Node)
	for dummy != nil{
		node := &Node{}
		hash[dummy] = node
		dummy = dummy.Next
	}

	// return deepCopy.Next

	dummy = head
	for dummy != nil{
		node := hash[dummy]
		node.Val = dummy.Val
		node.Next = hash[dummy.Next]
		node.Random = hash[dummy.Random]

		dummy = dummy.Next
	}

	return hash[head]
}
