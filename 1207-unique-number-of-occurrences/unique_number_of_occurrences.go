package _207_unique_number_of_occurrences

import "sort"

func uniqueOccurrences(arr []int) bool {
	if len(arr) == 0 {
		return false
	}
	sort.Ints(arr)
	dataMap := make(map[int]bool)
	count := 1
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			count++
		} else {
			if _, ok := dataMap[count]; ok {
				return false
			} else {
				dataMap[count] = true
			}
			count = 1
		}
	}
	if _, ok := dataMap[count]; ok {
		return false
	} else {
		dataMap[count] = true
	}
	return true
}
