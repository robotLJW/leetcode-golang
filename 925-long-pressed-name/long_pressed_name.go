package _25_long_pressed_name

// 计算每个单词中不重复的，并统计其个数
func isLongPressedName(name string, typed string) bool {
	if len(typed) < len(name) {
		return false
	}
	data1, count1 := removeSameWords(name)
	data2, count2 := removeSameWords(typed)
	if data1 != data2 {
		return false
	}
	for i := 0; i < len(count1); i++ {
		if count2[i] < count1[i] {
			return false
		}
	}
	return true
}

func removeSameWords(str string) (string, []int) {
	data := make([]byte, 0)
	count := make([]int, 0)
	tempCount := 1
	i := 1
	for ; i < len(str); i++ {
		if str[i] == str[i-1] {
			tempCount++
		} else {
			data = append(data, str[i-1])
			count = append(count, tempCount)
			tempCount = 1
		}
	}
	data = append(data, str[i-1])
	count = append(count, tempCount)
	return string(data), count
}
