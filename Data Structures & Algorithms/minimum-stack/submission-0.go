type MinStack struct {
	currentMin int
	stack []node
}

type node struct{
	val int
	currentMin int
}

func Constructor() MinStack {
	return MinStack{
		stack: make([]node,0),
	}
}

func (this *MinStack) Push(val int) {
	if val < this.currentMin || len(this.stack) == 0{
		this.currentMin = val
	}
	node := node{
		val: val,
		currentMin: this.currentMin,
	}

	this.stack = append(this.stack, node)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]

	if len(this.stack) == 0{
		this.currentMin = 0
		return
	}

	currentTop := this.stack[len(this.stack)-1]
	this.currentMin = currentTop.currentMin
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1].val
}

func (this *MinStack) GetMin() int {
	return this.currentMin
}
