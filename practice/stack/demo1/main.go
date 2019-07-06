package main

func main() {

}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

type MinStack struct {
	data interface{}
	next *MinStack
}

/** initialize your data structure here. */
func Constructor() MinStack {
	var minStack MinStack
	return minStack
}

func (this *MinStack) Push(x int) {
	if this.data == nil {
		this.data = x
	} else {
		this.next.data = this.data
		this.data = x
	}
}

func (this *MinStack) Pop() {
	return
}

func (this *MinStack) Top() int {
	return 9
}

func (this *MinStack) GetMin() int {
	return 1
}
