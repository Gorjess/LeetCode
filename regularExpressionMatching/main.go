package main

import "fmt"

// the char 'x' among any of the following comments means
// the value of the character indexed by j in pattern p.
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	matches := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	f := make([][]bool, m+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]bool, n+1)
	}
	f[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				f[i][j] = f[i][j-2]  // situation-1: assume that "x*" matches no character in s
				if matches(i, j-1) { // situation-2: check if 'x' matches s[i]
					// f[i][j] is true as long as one of the above two situations returns true.
					// so the operator "||" is used here
					f[i][j] = f[i][j] || f[i-1][j]
				}
			} else if matches(i, j) {
				f[i][j] = f[i-1][j-1]
			}
		}
	}
	return f[m][n]
}

func main() {
	fmt.Println(isMatch("aaa", "ab*a*c*a"))
	//fmt.Println(isMatch("hello", "hel*o"))
}
