package leetcode

func lengthOfLongestSubstring(s string) int {
	for i := 0; i < len(s); i++ {
		maxLen := 0
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				maxLen += i
			}
		}
	}
	return len(s)
}