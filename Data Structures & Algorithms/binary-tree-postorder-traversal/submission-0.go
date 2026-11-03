/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func postorderTraversal(root *TreeNode) []int {
    res := make([]int, 0)

	var post func(node *TreeNode)

	post = func(node *TreeNode){
		if node == nil{
			return
		}

		post(node.Left)
		post(node.Right)

		res = append(res, node.Val)
	}

	post(root)
	return res
}
