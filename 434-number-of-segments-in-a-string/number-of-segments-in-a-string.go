package _34_number_of_segments_in_a_string

func countSegments(s string) int {
	count := 0
	isBlank, isWord := false, false
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if isBlank == true {
				continue
			} else {
				if isWord == true {
					count++
					isWord = false
				}
				isBlank = true
			}
		} else {
			isBlank = false
			if isWord == true {
				continue
			} else {
				isWord = true
			}
		}
	}
	if isWord == true {
		count++
	}
	return count
}
