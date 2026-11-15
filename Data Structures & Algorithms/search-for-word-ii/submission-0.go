type TreeNode struct{
	Children [26]*TreeNode
	WordIdx int
}

func findWords(board [][]byte, words []string) []string {
	root := &TreeNode{WordIdx: -1}
    for i, word := range words{
		curr := root
		for j := 0; j < len(word); j++{
			idx := word[j] - 'a'
			if curr.Children[idx] == nil{
				curr.Children[idx] = &TreeNode{WordIdx: -1}
			}
			curr = curr.Children[idx]
		}
		curr.WordIdx = i
	}

	rows, cols := len(board), len(board[0])

	res := make([]string,0)

	var dfs func(r, c int, node *TreeNode)
	dfs = func(r, c int, node *TreeNode){
		if r < 0 || c < 0 || r >= rows || c >= cols || board[r][c] == '#'{
			return
		}

		char := board[r][c]
		nextNode := node.Children[char - 'a']
		if nextNode == nil{
			return
		}

		board[r][c] = '#'

		if nextNode.WordIdx != -1{
			res = append(res, words[nextNode.WordIdx])
			nextNode.WordIdx = -1
		}

		dfs(r+1, c, nextNode)
		dfs(r-1, c, nextNode)
		dfs(r, c+1, nextNode)
		dfs(r, c-1, nextNode)

		board[r][c] = char
	}

	for i := 0; i < len(board); i++{
		for j := 0; j < len(board[0]); j++{
			dfs(i, j, root)
		}
	}

	return res
}
