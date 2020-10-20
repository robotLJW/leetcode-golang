package _5_sort_colors

// 记数排序
func sortColors(nums []int) {
	var zeroCount, oneCount, twoCount, i int
	for i = 0; i < len(nums); i++ {
		switch nums[i] {
		case 0:
			zeroCount++
		case 1:
			oneCount++
		case 2:
			twoCount++
		}
	}

	for i = 0; i < zeroCount; i++ {
		nums[i] = 0
	}
	for i = zeroCount; i < zeroCount+oneCount; i++ {
		nums[i] = 1
	}
	for i = zeroCount + oneCount; i < len(nums); i++ {
		nums[i] = 2
	}
}

// 交换
func sortColorsTwo(nums []int) {
	countZero := swapColors(nums, 0)
	swapColors(nums[countZero:], 1)
}

func swapColors(colors []int, target int) (countTarget int) {
	for i, color := range colors {
		if color == target {
			colors[i], colors[countTarget] = colors[countTarget], colors[i]
			countTarget++
		}
	}
	return countTarget
}
