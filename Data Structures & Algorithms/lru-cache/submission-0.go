type LRUCache struct {
    cap int
	cache map[int]*Node
	left *Node
	right *Node
}

type Node struct{
	key, val int
	prev, next *Node
}

func Constructor(capacity int) LRUCache {
    left, right := &Node{}, &Node{}
	left.next = right
	right.prev = left

	return LRUCache{
		cap: capacity,
		cache: make(map[int]*Node),
		left: left,
		right: right,
	}
}

func (this *LRUCache) insert(node *Node){
	prev := this.right.prev
	next := this.right

	node.prev = prev
	prev.next = node
	node.next = next
	next.prev = node
}

func (this *LRUCache) remove(node *Node){
	prev := node.prev
	next := node.next

	prev.next = next
	next.prev = prev
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok{
		return -1
	}

	this.remove(node)
	this.insert(node)
	return node.val
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok{
		this.remove(node)
		node.val = value
		this.insert(node)
		return
	}

	if len(this.cache) == this.cap{
		lru := this.left.next
		this.remove(lru)
		delete(this.cache, lru.key)
	}

    newNode := &Node{
		key: key,
		val: value,
	}
	this.cache[key] = newNode
	this.insert((newNode))
}
