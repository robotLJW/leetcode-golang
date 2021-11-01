package _75_distribute_candies

func distributeCandies(candyType []int) int {
	isRepeat := make(map[int]int)
	notRepeatCount := 0
	for i := 0; i < len(candyType); i++ {
		if _, ok := isRepeat[candyType[i]]; ok {
			continue
		} else {
			isRepeat[candyType[i]] = 1
			notRepeatCount++
		}
	}
	if notRepeatCount > len(candyType)/2 {
		return len(candyType) / 2
	}
	return notRepeatCount
}

func distributeCandiesTwo(candyType []int) int {
	set := make(map[int]struct{})
	for _, value := range candyType {
		set[value] = struct{}{}
	}
	ans := len(candyType) / 2
	if len(set) < ans {
		ans = len(set)
	}
	return ans
}
