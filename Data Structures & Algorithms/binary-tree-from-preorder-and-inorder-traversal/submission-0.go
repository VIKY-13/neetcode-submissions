/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
    inorderMap := make(map[int]int)
	for i, val := range inorder{
		inorderMap[val] = i
	}

	preorderIdx := 0

	var dfs func(left, right int) *TreeNode

	dfs = func(left, right int) *TreeNode{
		if left > right{
			return nil
		}

		rootVal := preorder[preorderIdx]
		rootNode := &TreeNode{
			Val: rootVal,
		}

		mid := inorderMap[rootVal]

		preorderIdx++

		rootNode.Left = dfs(left, mid-1)
		rootNode.Right = dfs(mid+1, right)

		return rootNode
	}

	return dfs(0,len(preorder)-1)
}
