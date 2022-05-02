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
)

func firstMissingPositive(nums []int) int {
	sz := len(nums)
	// rewrite array
	for i := 0; i < sz; i++ {
		if nums[i] > sz || nums[i] <= 0 {
			nums[i] = sz + 1
		}
	}

	abs := func(in int) int {
		if in < 0 {
			return -in
		}
		return in
	}

	// mark elements
	var e int
	for i := 0; i < sz; i++ {
		e = abs(nums[i])
		if e > sz {
			continue
		}
		nums[e-1] = -abs(nums[e-1])
	}
	e = 0
	// find the idx whose corresponding elements is positive
	for i := 0; i < sz; i++ {
		if nums[i] > 0 {
			break
		}
		e++
	}

	return e + 1
}

func firstMissingPositiveSwap(nums []int) int {
	sz := len(nums)
	for i := 0; i < sz; i++ {
		for nums[i] > 0 && nums[i] <= sz && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < sz; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return sz + 1
}

func main() {
	//fmt.Println(firstMissingPositive(generateIntArray()))
	fmt.Println(firstMissingPositiveSwap([]int{1, 1}))
}
