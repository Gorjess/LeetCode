package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumberRecursivly(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		res  = ListNode{}
		tmp  = &res
		add1 = tmp.Val
		add2 = 0
	)

	if l1 == nil && l2 == nil {
		return &res
	} else {
		if l1 != nil {
			add1 += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			add2 = l2.Val
			l2 = l2.Next
		}

		var (
			digit = (add1 + add2) % 10
			ten   = (add1 + add2) / 10
		)

		tmp.Val = digit
		if ten > 0 {
			tmp.Next = new(ListNode)
			tmp.Next.Val = ten
		}
		addTwoNumberRecursivly(l1, l2)
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	add := func(v1, v2 int) (int, int) {
		return (v1 + v2) % 10, (v1 + v2) / 10
	}

	res := ListNode{}
	tmp := &res
	for l1 != nil || l2 != nil {
		var (
			add1 = tmp.Val
			add2 = 0
		)
		if l1 != nil {
			add1 += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			add2 = l2.Val
			l2 = l2.Next
		}
		digit, ten := add(add1, add2)
		tmp.Val = digit
		if ten > 0 || l1 != nil || l2 != nil {
			tmp.Next = new(ListNode)
			tmp.Next.Val = ten
			tmp = tmp.Next
		}
	}

	return &res
}

func main() {
	testL1 := []int{1}
	testL2 := []int{9, 9}

	toList := func(in []int) *ListNode {
		out := ListNode{}
		tmp := &out
		for i := 0; i < len(in); i++ {
			tmp.Val = in[i]
			if i < len(in)-1 {
				tmp.Next = new(ListNode)
				tmp = tmp.Next
			}
		}
		return &out
	}

	result := addTwoNumbers(toList(testL1), toList(testL2))
	fmt.Print(result)
}
