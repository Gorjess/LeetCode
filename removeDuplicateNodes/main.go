/*面试题 02.01. Remove Duplicate Node LCCI
Write code to remove duplicates from an unsorted linked list.

Example1:

 Input: [1, 2, 3, 3, 2, 1]
 Output: [1, 2, 3]
Example2:

 Input: [1, 1, 1, 1, 2]
 Output: [1, 2]
Note:

The length of the list is within the range[0, 20000].
The values of the list elements are within the range [0, 20000].
Follow Up:

How would you solve this problem if a temporary buffer is not allowed?

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-duplicate-node-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

package main

import (
	"bytes"
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func generateList(in []int) *ListNode {
	head := &ListNode{
		Val:  in[0],
		Next: new(ListNode),
	}
	temp := head.Next
	for i := 1; i < len(in); i++ {
		temp.Val = in[i]
		temp.Next = new(ListNode)
		temp = temp.Next
	}
	return head
}

func removeDuplicateNodes(head *ListNode) *ListNode {
	rec := map[int]struct{}{}

	preN := head
	for n := head; n != nil; n = n.Next {
		if _, ok := rec[n.Val]; ok {
			preN.Next = n.Next
			continue
		}
		rec[n.Val] = struct{}{}
		preN = n
	}
	rec = map[int]struct{}{} // to reduce memory use

	var s bytes.Buffer
	for n := head; n != nil; n = n.Next {
		s.WriteString(fmt.Sprintf("%d,", n.Val))
	}
	fmt.Println(s.String())

	return head
}

func main() {
	inSrc := []int{1, 2, 3, 3, 2, 1}
	removeDuplicateNodes(generateList(inSrc))
}
