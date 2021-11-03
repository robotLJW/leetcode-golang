package offer_09_yong_liang_ge_zhan_shi_xian_dui_lie_lcof

import "container/list"

type CQueue struct {
	stackOne, stackTwo *list.List
}

func Constructor() CQueue {
	return CQueue{
		stackOne: list.New(),
		stackTwo: list.New(),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.stackOne.PushBack(value)
}

func (this *CQueue) DeleteHead() int {
	if this.stackTwo.Len() == 0 {
		for this.stackOne.Len() != 0 {
			this.stackTwo.PushBack(this.stackOne.Remove(this.stackOne.Back()))
		}
	}
	if this.stackTwo.Len() != 0 {
		ans := this.stackTwo.Back()
		this.stackTwo.Remove(ans)
		return ans.Value.(int)
	}
	return -1
}
