/*10. Regular Expression Matching
Given an input string (s) and a pattern (p), implement regular expression matching with support for '.' and '*'.

'.' Matches any single character.
'*' Matches zero or more of the preceding element.
The matching should cover the entire input string (not partial).

Note:

s could be empty and contains only lowercase letters a-z.
p could be empty and contains only lowercase letters a-z, and characters like . or *.
Example 1:

Input:
s = "aa"
p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".
Example 2:

Input:
s = "aa"
p = "a*"
Output: true
Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
Example 3:

Input:
s = "ab"
p = ".*"
Output: true
Explanation: ".*" means "zero or more (*) of any character (.)".
Example 4:

Input:
s = "aab"
p = "c*a*b"
Output: true
Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore, it matches "aab".
Example 5:

Input:
s = "mississippi"
p = "mis*is*p*."
Output: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/regular-expression-matching
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

package main

import (
	"fmt"
	"time"
)

// todo: test space-complexity between pointer and struct
type subPattern struct {
	letter  uint8
	pattern uint8
}

/*
'.' Matches any single character.
'*' Matches zero or more of the preceding element.
The matching should cover the entire input string (not partial).
*/
func isMatch(s string, p string) bool {
	start := time.Now()
	patternStack := make([]*subPattern, 0, len(p)) // pattern for each letter
	sp := new(subPattern)

	ca := uint8(97)  // ascii for 'a'
	cz := uint8(122) // ascii for 'z'

	// isLowerLtr returns if in is a lowercase letter
	isLowerLtr := func(in uint8) bool {
		return ca <= in && in <= cz
	}
	// isStar returns if in is a *
	isStar := func(in uint8) bool {
		return in == uint8('*')
	}
	// isDot returns if in is a .
	isDot := func(in uint8) bool {
		return in == uint8('.')
	}

	for i := 0; i < len(p); i++ {
		char := p[i]
		if !isStar(char) && !isDot(char) && !isLowerLtr(char) {
			return false
		}

		if isStar(char) {
			sp.pattern = char
		} else {
			if sp.letter != 0 { // we got a new letter
				sp = new(subPattern)
			}
			sp.letter = char
			patternStack = append(patternStack, sp)
		}
	}

	var (
		cursor = 0
		head   *subPattern
	)

	for len(patternStack) > 0 {
		head = patternStack[0]
		patternStack = patternStack[1:]

		if cursor >= len(s) {
			if isStar(head.pattern) {
				continue
			} else {
				return false
			}
		}

		if !isDot(head.letter) {
			if head.letter != s[cursor] {
				if !isStar(head.pattern) {
					return false
				}
				continue
			}
			cursor++
		}

		if isStar(head.pattern) {
			if isDot(head.letter) && len(patternStack) == 0 {
				return true
			}
			for cursor < len(s) {
				if s[cursor] == head.letter {
					cursor++
				} else if isDot(head.letter) {
					if s[cursor] == patternStack[1].letter {
						break
					}
				} else {
					break
				}
			}
		}
	}

	// pattern must cover the entire input string
	if cursor >= len(s) && len(patternStack) == 0 {
		fmt.Println(time.Now().Sub(start))
		return true
	}
	fmt.Println(time.Now().Sub(start))
	return false
}

func main() {
	fmt.Println(isMatch("a", ".*..a*"))
}
