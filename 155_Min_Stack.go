package leetcode

type MinStack struct {
	Stack    []int
	IsChange bool
	minVal   int
}

func Constructor155() MinStack {
	m := MinStack{}
	return m
}

func (this *MinStack) Push(val int) {
	this.IsChange = true
	this.Stack = append(this.Stack, val)
}

func (this *MinStack) Pop() {
	lenStack := len(this.Stack)
	if lenStack == 0 {
		return
	}
	this.IsChange = true
	this.Stack = this.Stack[:lenStack-1]
}

func (this *MinStack) Top() int {
	lenStack := len(this.Stack)
	if lenStack == 0 {
		return 0
	}
	return this.Stack[lenStack-1]
}

func (this *MinStack) GetMin() int {
	lenStack := len(this.Stack)

	if lenStack == 0 {
		return 0
	}
	min := this.Stack[0]
	if this.IsChange {
		for _, val := range this.Stack {
			if val < min {
				min = val
			}
		}
		this.IsChange = false
		return min
	}

	return this.minVal
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
