/*
	Given an array of integers, return indices of the two numbers such that they add up to a specific target.
	You may assume that each input would have exactly one solution, and you may not use the same element twice.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

func twoSum(nums []int, target int) []int {
	record := map[int]int{} // num -> index
	l := len(nums)
	for i := 0; i < l; i++ {
		if idx, ok := record[target-nums[i]]; ok {
			return []int{i, idx}
		}
		record[nums[i]] = i
	}
	return nil
}
