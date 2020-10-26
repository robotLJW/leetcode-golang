package _365_how_many_numbers_are_smaller_than_the_current_number

import "sort"

func smallerNumbersThanCurrent(nums []int) []int {
	data := make([]int, len(nums))
	dataMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		data[i] = nums[i]
	}
	sort.Ints(data)
	for i := 0; i < len(data); i++ {
		if _, ok := dataMap[data[i]]; ok {
		} else {
			dataMap[data[i]] = i
		}
	}
	for i := 0; i < len(nums); i++ {
		data[i] = dataMap[nums[i]]
	}
	return data
}
