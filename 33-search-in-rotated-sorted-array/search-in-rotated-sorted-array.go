package _3_search_in_rotated_sorted_array

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	minIndex := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			minIndex = i
		}
	}
	if nums[minIndex] > target {
		return -1
	}
	var start, end int
	if minIndex == 0 {
		start, end = 0, len(nums)-1
	} else {
		if nums[len(nums)-1] < target {
			start, end = 0, minIndex-1
		} else {
			start, end = minIndex, len(nums)-1
		}
	}
	for start <= end {
		middle := (start + end) / 2
		if nums[middle] == target {
			return middle
		} else if nums[middle] < target {
			start = middle + 1
		} else {
			end = middle - 1
		}
	}
	return -1
}
