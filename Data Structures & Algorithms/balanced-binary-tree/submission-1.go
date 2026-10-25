/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    var dfs func(node *TreeNode) int
	res := true

	dfs = func(node *TreeNode) int{
		if node == nil{
			return 0
		}

		leftHeight := dfs(node.Left)
		rightHeight := dfs(node.Right)

		if abs(leftHeight - rightHeight) > 1{
			res = false 
		}

		return max(leftHeight, rightHeight) + 1
	}

	dfs(root)
	return res
}

func abs(a int) int{
	if a < 0{
		return a * -1
	}

	return a
}

func max(a, b int)int{
	if a > b{
		return a
	}
	return b
}
