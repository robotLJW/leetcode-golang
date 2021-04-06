package _74_longest_continuous_increasing_subsequence

func findLengthOfLCIS(nums []int) int {
	maxLength, currentLength := 0, 1
	if len(nums) == 0 {
		return 0
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			currentLength++
		} else {
			if maxLength < currentLength {
				maxLength = currentLength
			}
			currentLength = 1
		}
	}
	if maxLength < currentLength {
		maxLength = currentLength
	}
	return maxLength
}
