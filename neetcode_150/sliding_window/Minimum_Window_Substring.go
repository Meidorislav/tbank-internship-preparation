func isSubstring(freq [128]int, t string) bool {
	for i := 0; i < len(t); i++ {
		if freq[t[i]] > 0 {
			return false
		}
	}
	return true
}

func minWindow(s string, t string) string {
	if len(t) > len(s) || len(t) == 0 {
		return ""
	}

	var freq [128]int
	for i := 0; i < len(t); i++ {
		freq[t[i]]++
	}

	count := len(t)
	shortest_left := -1
	shortest_right := -1
	left := 0
	for right := 0; right < len(s); right++ {
		if freq[s[right]] > 0 {
			count--
		}
		freq[s[right]]--

		for count == 0 {
			if shortest_left == -1 || (right-left) < (shortest_right-shortest_left) {
				shortest_left = left
				shortest_right = right
			}

			freq[s[left]]++
			if freq[s[left]] > 0 {
				count++
			}
			left++
		}
	}

	if shortest_left != -1 {
		return s[shortest_left : shortest_right+1]
	}
	return ""
}

