package _5_jump_game_ii

func jump(nums []int) int {
	position, steps := len(nums)-1, 0
	for position > 0 {
		for i := 0; i < len(nums); i++ {
			if i+nums[i] >= position {
				position = i
				steps++
				break
			}
		}
	}
	return steps
}

func jump2(nums []int) int {
	maxPosition, steps, end := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			steps++
		}
	}
	return steps
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
