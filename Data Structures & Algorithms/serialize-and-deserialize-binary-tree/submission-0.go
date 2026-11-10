/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    finalStr := ""

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode){
		if node == nil{
			finalStr += "N,"
			return
		}

		finalStr += strconv.Itoa(node.Val)+","

		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)

	// fmt.Println(finalStr)

	return finalStr
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	preorder := strings.Split(data, ",")

	// fmt.Println(preorder)

	var dfs func() *TreeNode

	idx := -1
	dfs = func() *TreeNode{
		idx++
		if preorder[idx] == "N"{
			return nil
		}

		val, _:= strconv.Atoi(preorder[idx])
		root := &TreeNode{
			Val: val,
		}

		root.Left = dfs()
		root.Right = dfs()

		return root
	}

    return dfs()
}
