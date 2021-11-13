package _20_detect_capital

func detectCapitalUse(word string) bool {
	if isUppercase(word[0]) {
		upperNum, lowerNum := 0, 0
		for i := 1; i < len(word); i++ {
			if isUppercase(word[i]) {
				upperNum++
			} else {
				lowerNum++
			}
			if upperNum*lowerNum != 0 {
				return false
			}
		}
	} else {
		for i := 1; i < len(word); i++ {
			if isUppercase(word[i]) {
				return false
			}
		}
	}
	return true
}

func isUppercase(b byte) bool {
	if b >= 'A' && b <= 'Z' {
		return true
	}
	return false
}
