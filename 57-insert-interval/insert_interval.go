package _7_insert_interval

func insert(intervals [][]int, newInterval []int) [][]int {
	ans := make([][]int, 0)
	start, end := newInterval[0], newInterval[1]
	i, indexStart, indexEnd := 0, -1, -1
	for ; i < len(intervals); i++ {
		temp := intervals[i]
		if temp[0] > end {
			break
		}
		if temp[1] < start {
			ans = append(ans, temp)
			continue
		}

		if indexStart == -1 && start >= temp[0] && start <= temp[1] {
			indexStart = i
		}
		if end >= temp[0] && end <= temp[1] {
			indexEnd = i
		}
	}
	if indexStart == -1 {
		if indexEnd == -1 {
			ans = append(ans, newInterval)
		} else {
			temp := make([]int, 2)
			temp[0] = start
			temp[1] = intervals[indexEnd][1]
			ans = append(ans, temp)
		}
	} else {
		if indexEnd == -1 {
			temp := intervals[indexStart]
			if end > temp[1] {
				temp[1] = end
			}
			ans = append(ans, temp)
		} else {
			temp := make([]int, 2)
			temp[0] = intervals[indexStart][0]
			temp[1] = intervals[indexEnd][1]
			ans = append(ans, temp)
		}
	}
	for ; i < len(intervals); i++ {
		temp := intervals[i]
		ans = append(ans, temp)
	}
	return ans
}
