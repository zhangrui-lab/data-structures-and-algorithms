// Package string 字符串回文串算法
package string

import (
	"math"
	"strings"
)

// LongestPalindromeSubstr 获取字符串str的最长回文串
func LongestPalindromeSubstr(str string) string {
	os := str
	str = "#" + strings.Join(strings.Split(str, ""), "#") + "#"
	c, r, mx, s := 0, 0, 0, 0
	p := make([]int, len(str))
	extend := func(i int) {
		for lo, hi := i-p[i]-1, i+p[i]+1; lo > -1 && hi < len(str) && str[lo] == str[hi]; {
			lo--
			hi++
			p[i]++
		}
	}
	for i := 0; i < len(str); i++ {
		if i < r {
			p[i] = int(math.Min(float64(p[c<<1-i]), float64(r-i)))
		}
		extend(i)
		if r < i+p[i] {
			c = i
			r = i + p[i]
		}
		if mx < p[i] {
			mx = p[i]
			s = (i - mx) >> 1
		}
	}
	return os[s : s+mx]
}
