/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
    var isValid func(node *TreeNode, minVal, maxVal *int) bool

	isValid = func(node *TreeNode, minVal, maxVal *int) bool{
		if node == nil{
			return true
		}

		if minVal != nil && node.Val <= *minVal{
			return false
		}
		if maxVal != nil && node.Val >= *maxVal{
			return false
		}

		return isValid(node.Left, minVal, &node.Val) &&
			isValid(node.Right, &node.Val, maxVal)
	}

	return isValid(root, nil, nil)
}
