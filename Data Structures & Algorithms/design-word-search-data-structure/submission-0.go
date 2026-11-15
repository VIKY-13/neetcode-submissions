type WordDictionary struct {
    children [26]*WordDictionary
	isEnd bool
}

func Constructor() WordDictionary {
    return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string)  {
    curr := this
	
	for _, val := range word{
		idx := val - 'a'

		if curr.children[idx] == nil{
			curr.children[idx] = &WordDictionary{}
		}

		curr = curr.children[idx]
	}

	curr.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	return this.search(word, 0)
}

func (this *WordDictionary) search(word string, wordIdx int) bool {
	curr := this
	
	for i := wordIdx; i < len(word); i++{
		if word[i] == '.'{
			for _, child := range curr.children{
				if child != nil && child.search(word, i+1){
					return true
				}
			}
			return false
		}
	
		idx := word[i] - 'a'
		if curr.children[idx] == nil{
			return false
		}

		curr = curr.children[idx]
	}

	return curr.isEnd
}
