/*125. Valid Palindrome
Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

Note: For the purpose of this problem, we define empty string as valid palindrome.

Example 1:

Input: "A man, a plan, a canal: Panama"
Output: true
Example 2:

Input: "race a car"
Output: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-palindrome
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

package main

import (
	"fmt"
)

func isPalindrome(s string) bool {
	pureLetters := make([]uint8, 0, len(s))
	var c uint8
	var (
		a0 = uint8('0')
		a9 = uint8('9')
		aA = uint8('A')
		aZ = uint8('Z')
		aa = uint8('a')
		az = uint8('z')
	)
	for i := 0; i < len(s); i++ {
		c = s[i]
		if a0 <= c && c <= a9 || aA <= c && c <= aZ || aa <= c && c <= az {
			pureLetters = append(pureLetters, s[i])
		}
	}
	var (
		sz  = len(pureLetters)
		mid = sz / 2
	)

	var (
		asciiDiff = uint8(32)
		max       = func(a, b uint8) (uint8, uint8) {
			if a >= b {
				return a, b
			}
			return b, a
		}
		larger, smaller, diff uint8
	)

	for i := 0; i < mid; i++ {
		larger, smaller = max(pureLetters[i], pureLetters[sz-i-1])
		diff = larger - smaller
		if diff != 0 && diff != asciiDiff {
			return false
		}
		if diff == asciiDiff {
			if larger >= aA && smaller <= a9 {
				return false
			}
		}
	}

	return true
}

func main() {
	fmt.Println(isPalindrome("race a car"))
}
