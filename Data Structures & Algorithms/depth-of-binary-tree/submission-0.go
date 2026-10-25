/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
	// maxDepth := 0

	var dfs func(node *TreeNode)int

	dfs = func(node *TreeNode) int{
		if node == nil{
			return 0
		}

		leftHeight := dfs(node.Left)
		rightHeight := dfs(node.Right)

		return 1+max(leftHeight, rightHeight)
	}
    
	return dfs(root)
}

func max(a, b int) int{
	if a > b{
		return a
	}
	return b
}