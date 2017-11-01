package longestPalindrome

func longestPalindrome(s string) string {
	max := ""
	for i := 0; i < len(s); i++ {
		j := 1
		for i-j >= 0 && i+j < len(s) && s[i-j] == s[i+j] {
			j += 1
		}
		if len(max) < len(s[i-j+1:i+j]) {
			max = s[i-j+1 : i+j]
		}
		j = 1
		for i-j+1 >= 0 && i+j < len(s) && s[i-j+1] == s[i+j] {
			j += 1
		}
		if i+j == len(s) {
			if len(max) < len(s[i-j+1+1:]) {
				max = s[i-j+1+1:]
			}
		} else if len(max) < len(s[i-j+1+1:i+j]) {
			max = s[i-j+1+1 : i+j]
		}

	}
	return max
}