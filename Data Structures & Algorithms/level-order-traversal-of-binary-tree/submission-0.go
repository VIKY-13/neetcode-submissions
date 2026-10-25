/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int,0)
	if root == nil{
		return res
	}

    q := make([]*TreeNode,0)
	q = append(q,root)

	for len(q) > 0{
		nodes := make([]int,0)
		qLen := len(q)

		for qLen > 0{
			node := q[0]
			if node.Left != nil{
				q = append(q, node.Left)
			}

			if node.Right != nil{
				q = append(q, node.Right)
			}

			nodes = append(nodes, node.Val)

			q = q[1:]
			qLen--
		}

		res = append(res, nodes)
	}

	return res
}
