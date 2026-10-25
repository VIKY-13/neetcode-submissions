/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
	res := make([]int,0)
	if root == nil{
		return res
	}

	q := make([]*TreeNode,0)
	q = append(q, root)

	for len(q) > 0{
		res = append(res,q[len(q)-1].Val)

		qLen := len(q)
		for qLen > 0{
			node := q[0]
			if node.Left != nil{
				q = append(q, node.Left)
			}
			if node.Right != nil{
				q = append(q, node.Right)
			}
			qLen--
			q = q[1:]
		}
	}

	return res 
}
