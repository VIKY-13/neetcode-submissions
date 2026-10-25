/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	maxD := 0

	var dfs func(node *TreeNode) int

	dfs = func(node *TreeNode)int{
		if node == nil{
			return 0
		}

		leftHeight := dfs(node.Left)
		rightHeight := dfs(node.Right)

		if leftHeight + rightHeight > maxD{
			maxD = leftHeight + rightHeight
		}

		return 1+max(leftHeight, rightHeight)
	}
    
	dfs(root)
	return maxD
}

func max(a, b int)int{
	if a > b{
		return a
	}
	return b
}
