/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
    maxSum := -1001

	var dfs func(node *TreeNode)int

	dfs = func(node *TreeNode)int{
		if node == nil{
			return 0
		}

		leftSum := max(dfs(node.Left),0)
		rightSum := max(dfs(node.Right),0)

		currPathSum := node.Val + leftSum + rightSum
		if currPathSum > maxSum{
			maxSum = currPathSum
		}

		sum := node.Val + max(leftSum, rightSum)
		if sum > maxSum{
			maxSum = sum
		}

		return sum
	}

	dfs(root)
	return maxSum
}

func max(a, b int)int{
	if a > b{
		return a
	}
	return b
}
