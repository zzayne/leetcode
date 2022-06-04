package design

type MinStack struct {
	nums []int
	mins []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if len(this.mins) == 0 || val <= this.GetMin() {
		this.mins = append(this.mins, val)
	}
	this.nums = append(this.nums, val)
}

func (this *MinStack) Pop() {
	if len(this.nums) == 0 {
		return
	}
	if this.Top() == this.mins[len(this.mins)-1] {
		this.mins = this.mins[:len(this.mins)-1]
	}
	this.nums = this.nums[:len(this.nums)-1]
}

func (this *MinStack) Top() int {
	if len(this.nums) == 0 {
		return 0
	}
	return this.nums[len(this.nums)-1]
}

func (this *MinStack) GetMin() int {
	return this.mins[len(this.mins)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
