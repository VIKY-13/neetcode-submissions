/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	res := 0
	count := 0
    var dfs func(node *TreeNode)

	dfs = func(node *TreeNode){
		if node == nil{
			return
		}

		if count == k{
			return
		}

		dfs(node.Left)
		count++
		// fmt.Printf("node - %d, count - %d \n", node.Val, count)
		if count == k{
			res = node.Val
			return
		}

		dfs(node.Right)
	}

	dfs(root)
	return res
}
