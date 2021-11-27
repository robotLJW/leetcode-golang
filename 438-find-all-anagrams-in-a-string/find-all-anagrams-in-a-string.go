package _38_find_all_anagrams_in_a_string

import "reflect"

func findAnagrams(s string, p string) []int {
	ans := make([]int, 0)
	if len(s) < len(p) {
		return ans
	}
	pLength := len(p)
	pLetterCount := calLetters(p)
	for i := 0; i <= len(s)-pLength; i++ {
		tmpStr := s[i : i+pLength]
		if reflect.DeepEqual(calLetters(tmpStr), pLetterCount) {
			ans = append(ans, i)
		}
	}
	return ans
}

func calLetters(str string) [26]int {
	ans := [26]int{}
	for _, ch := range str {
		ans[ch-'a']++
	}
	return ans
}
