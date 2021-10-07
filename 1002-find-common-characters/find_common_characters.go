package _002_find_common_characters

import "math"

func commonChars(a []string) (ans []string) {
	minFreq := [26]int{}
	for i := 0; i < len(minFreq); i++ {
		minFreq[i] = math.MaxInt64
	}
	for _, word := range a {
		tempFreq := [26]int{}
		for _, b := range word {
			tempFreq[b-'a']++
		}
		for i := 0; i < len(tempFreq); i++ {
			if tempFreq[i] < minFreq[i] {
				minFreq[i] = tempFreq[i]
			}
		}
	}
	for i := byte(0); i < 26; i++ {
		for j := 0; j < minFreq[i]; j++ {
			ans = append(ans, string('a'+i))
		}
	}
	return ans
}
