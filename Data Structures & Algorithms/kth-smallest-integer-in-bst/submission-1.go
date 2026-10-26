/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    count := 0
    res := 0
    
    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil || count >= k {
            return
        }
        
        // 1. Traverse left subtree
        dfs(node.Left)
        
        // 2. Visit current node
        count++
        if count == k {
            res = node.Val
            return
        }
        
        // 3. Traverse right subtree
        dfs(node.Right)
    }
    
    dfs(root)
    return res
}