package _25_implement_stack_using_queues

type MyStack struct {
	queue1, queue2 []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	stack := MyStack{}
	return stack
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.queue2 = append(this.queue2, x)
	for len(this.queue1) != 0 {
		this.queue2 = append(this.queue2, this.queue1[0])
		this.queue1 = this.queue1[1:]
	}
	this.queue1, this.queue2 = this.queue2, this.queue1
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	topValue := this.queue1[0]
	this.queue1 = this.queue1[1:]
	return topValue
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.queue1[0]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.queue1) == 0
}
