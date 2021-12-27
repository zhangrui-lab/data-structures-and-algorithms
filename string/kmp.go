package string

// KmpMatch KMP字符串匹配算法 (ASIIC)
func KmpMatch(pattern, text string) int {
	next := buildNext(pattern)
	i, n := 0, len(text)
	j, m := 0, len(pattern)
	for i < n && j < m {
		if j < 0 || text[i] == pattern[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	return i - j
}

// 构造 next 表
func buildNext(pattern string) []int {
	next := make([]int, len(pattern))
	next[0] = -1
	for i, t, size := 0, -1, len(pattern)-1; i < size; {
		if t < 0 || (pattern[i] == pattern[t]) {
			i++
			t++
			if pattern[i] == pattern[t] {
				next[i] = next[t]
			} else {
				next[i] = t
			}
		} else {
			t = next[t]
		}
	}
	return next
}
