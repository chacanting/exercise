package main

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	max := 0
	set := make(map[byte]int)
	var i, j int
	for j < n {
		if _, exist := set[s[j]]; exist &&  i<=set[s[j]] {
			i = set[s[j]] + 1
			set[s[j]] = j
			j += 1
		} else {
			set[s[j]] = j
			j += 1
			if j-i > max {
				max = j - i
			}
		}
	}
	return max
}