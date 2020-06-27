/*41. First Missing Positive
Given an unsorted integer array, find the smallest missing positive integer.

Example 1:

Input: [1,2,0]
Output: 3
Example 2:

Input: [3,4,-1,1]
Output: 2
Example 3:

Input: [7,8,9,11,12]
Output: 1
Note:

Your algorithm should run in O(n) time and uses constant extra space.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/first-missing-positive
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateIntArray() []int {
	rand.Seed(time.Now().UnixNano())
	sz := rand.Intn(10)
	offset := 20
	out := make([]int, sz)
	for i := 0; i < sz; i++ {
		out[i] = rand.Intn(100) - offset
	}
	return out
}

func firstMissingPositive(nums []int) int {
	var min int
	return min
}

func main() {
	fmt.Println(firstMissingPositive(generateIntArray()))
}
