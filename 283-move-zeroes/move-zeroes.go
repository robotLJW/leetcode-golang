package _83_move_zeroes

func moveZeroes(nums []int) {
	zeroCount := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroCount++
		} else {
			nums[i-zeroCount] = nums[i]
		}
	}
	for zeroCount > 0 {
		nums[len(nums)-zeroCount] = 0
		zeroCount--
	}
}
