/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	var candidate *TreeNode

	min, max := getMinMax(p.Val, q.Val)
    for root != nil{
		rootVal := root.Val
		if min <= rootVal && max >= rootVal{
			candidate = root
			break
		}

		if min <= rootVal && max <= rootVal{
			root = root.Left
		}else{
			root = root.Right
		}
	}

	return candidate
}

func getMinMax(a , b int)(int,int){
	if a < b{
		return a, b
	}
	return b, a
}
