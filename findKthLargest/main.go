/*215. Kth Largest Element in an Array
Find the kth largest element in an unsorted array. Note that it is the kth largest element in the sorted order, not the kth distinct element.

Example 1:

Input: [3,2,1,5,6,4] and k = 2
Output: 5
Example 2:

Input: [3,2,3,1,2,4,5,5,6] and k = 4
Output: 4
Note:
You may assume k is always valid, 1 ≤ k ≤ array's length.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/kth-largest-element-in-an-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

package main

import (
	"fmt"
)

// qSort recursivly calls partion to sort in in a decreasing order.
// Note that tail is excluded in partion.
func qSort(in []int, head, tail int) {
	if head < tail {
		flag := partion(in, head, tail)
		qSort(in, head, flag)   // elements before flag
		qSort(in, flag+1, tail) // elements behind flag
	}
}

// partion reorganizes elements in decreasing order
// and returns the index of a new pivot.
func partion(array []int, head, tail int) int {
	// pick the first elem as pivot
	pivot, flag := array[head], tail

	swap := func(idxI, idxJ int) {
		if idxI == idxJ {
			return
		}
		array[idxI], array[idxJ] = array[idxJ], array[idxI]
	}

	for i := tail - 1; i > head; i-- {
		if array[i] < pivot {
			flag--
			swap(i, flag)
		}
	}
	// place pivot at the right position
	swap(flag-1, head)

	return flag - 1
}

func findKthLargest(nums []int, k int) int {
	sz := len(nums)

	// perform quick-sort on nums
	return sz
}

func main() {
	a := []int{4, 1, 8, 5, 8}
	fmt.Println(a)
	qSort(a, 0, len(a))
	fmt.Println(a)
}
