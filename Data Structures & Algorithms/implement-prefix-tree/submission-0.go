type PrefixTree struct {
	Children [26]*PrefixTree
	IsEnd bool
}

func Constructor() PrefixTree {
    return PrefixTree{}
}

func (this *PrefixTree) Insert(word string) {
	curr := this
	for _, val := range word{
		idx := val - 'a'
		if curr.Children[idx] == nil{
			curr.Children[idx] = &PrefixTree{}
		}
		curr = curr.Children[idx]
	}

	curr.IsEnd = true
}

func (this *PrefixTree) Search(word string) bool {
	curr := this

	for _, val := range word{
		idx := val - 'a'
		if curr.Children[idx] == nil{
			return false
		}

		curr = curr.Children[idx]
	}

	return curr.IsEnd
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	curr := this

	for _, val := range prefix{
		idx := val - 'a'

		if curr.Children[idx] == nil{
			return false
		}

		curr = curr.Children[idx]
	}
	return true
}
