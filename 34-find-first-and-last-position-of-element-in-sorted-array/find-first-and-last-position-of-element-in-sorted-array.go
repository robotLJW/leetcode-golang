package _4_find_first_and_last_position_of_element_in_sorted_array

func searchRange(nums []int, target int) []int {
	ans := []int{-1, -1}
	if len(nums) == 0 || nums[0] > target {
		return ans
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			if ans[0] == -1 {
				ans[0] = i
				ans[1] = i
			} else {
				ans[1] = i
			}
		}
	}
	return ans
}
