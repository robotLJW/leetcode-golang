package _81_insert_delete_getrandom_o1_duplicates_allowed

import (
	"math/rand"
)

type RandomizedCollection struct {
	idx  map[int]map[int]struct{}
	nums []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	return RandomizedCollection{
		idx: map[int]map[int]struct{}{},
	}
}

/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	ids, ok := this.idx[val]
	if !ok {
		ids = map[int]struct{}{}
		this.idx[val] = ids
	}
	ids[len(this.nums)] = struct{}{}
	this.nums = append(this.nums, val)
	return !ok
}

/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
	ids, ok := this.idx[val]
	if !ok {
		return false
	}
	var i int
	for k := range ids {
		i = k
		break
	}
	n := len(this.nums)
	this.nums[i] = this.nums[n-1]
	delete(ids, i)
	delete(this.idx[this.nums[i]], n-1)
	if i < n-1 {
		this.idx[this.nums[i]][i] = struct{}{}
	}
	if len(ids) == 0 {
		delete(this.idx, val)
	}
	this.nums = this.nums[:n-1]
	return true
}

/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}
