package _99_bulls_and_cows

import "fmt"

func getHint(secret string, guess string) string {
	bulls, cows := 0, 0
	var cntS, cntG [10]int
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			bulls++
		} else {
			cntS[secret[i]-'0']++
			cntG[guess[i]-'0']++
		}
	}
	for i := 0; i < 10; i++ {
		cows += min(cntS[i], cntG[i])
	}
	return fmt.Sprintf("%dA%dB", bulls, cows)

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
