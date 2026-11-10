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

	var dfs func(left, right int)*TreeNode
	dfs = func(left, right int)*TreeNode{
		if left > right{
			return nil
		}

		mid := inorderMap[preorder[preorderIdx]]
		root := &TreeNode{
			Val : preorder[preorderIdx],
		}
		preorderIdx++

		root.Left = dfs(left, mid-1)
		root.Right = dfs(mid+1, right)

		return root
	}

	return dfs(0,len(preorder)-1)
}
