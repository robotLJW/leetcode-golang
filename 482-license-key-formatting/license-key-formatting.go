package _82_license_key_formatting

import "unicode"

func licenseKeyFormatting(s string, k int) string {
	ans, count := []byte{}, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != '-' {
			ans = append(ans, byte(unicode.ToUpper(rune(s[i]))))
			count++
			if count%k == 0 {
				ans = append(ans, '-')
			}
		}
	}
	if len(ans) > 0 && ans[len(ans)-1] == '-' {
		ans = ans[:len(ans)-1]
	}
	for i, j := 0, len(ans)-1; i < len(ans)/2; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return string(ans)
}
