package _68_missing_number

func missingNumber(nums []int) int {
	length := len(nums)
	sum := length * (length + 1) / 2
	for i := 0; i < len(nums); i++ {
		sum = sum - nums[i]
	}
	return sum
}
