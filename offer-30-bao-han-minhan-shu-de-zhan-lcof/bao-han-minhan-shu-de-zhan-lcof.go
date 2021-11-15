package offer_30_bao_han_minhan_shu_de_zhan_lcof

import "math"

type MinStack struct {
	Data       []int
	MinData    []int
	Length     int
	CurrentMin int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		Data:       make([]int, 0),
		MinData:    make([]int, 0),
		Length:     0,
		CurrentMin: math.MaxInt64,
	}
}

func (this *MinStack) Push(x int) {
	if len(this.Data) == this.Length {
		this.Data = append(this.Data, x)
		if x < this.CurrentMin {
			this.MinData = append(this.MinData, x)
			this.CurrentMin = x
		} else {
			this.MinData = append(this.MinData, this.CurrentMin)
		}
	} else {
		this.Data[this.Length] = x
		if x < this.CurrentMin {
			this.MinData[this.Length] = x
			this.CurrentMin = x
		} else {
			this.MinData[this.Length] = this.CurrentMin
		}
	}
	this.Length++
}

func (this *MinStack) Pop() {
	if this.Length < 1 {
		return
	}
	this.Length--
	if this.Length == 0 {
		this.CurrentMin = math.MaxInt64
	}else{
		this.CurrentMin = this.MinData[this.Length-1]
	}
}

func (this *MinStack) Top() int {
	if this.Length > 0 {
		return this.Data[this.Length-1]
	}
	return 0
}

func (this *MinStack) Min() int {
	if this.Length > 0 {
		return this.MinData[this.Length-1]
	}
	return 0
}
