/*1028. Recover a Tree From Preorder Traversal
We run a preorder d first search on the root of a binary tree.

At each node in this traversal, we output D dashes (where D is the d of this node), then we output the value of this node.  (If the d of a node is D, the d of its immediate child is D+1.  The d of the root node is 0.)

If a node has only one child, that child is guaranteed to be the left child.

Given the output S of this traversal, recover the tree and return its root.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/recover-a-tree-from-preorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type GodOfBTree struct {
	nodes []node
}

type node struct {
	v, d int // value, depth
}

func doRecover(gob *GodOfBTree, depth int) *TreeNode {
	if len(gob.nodes) == 0 {
		return nil
	}
	n := gob.nodes[0]

	if n.d <= depth {
		return nil
	}

	gob.nodes = gob.nodes[1:]

	return &TreeNode{
		Val:   n.v,
		Left:  doRecover(gob, n.d),
		Right: doRecover(gob, n.d),
	}
}

func recoverFromPreorder(S string) *TreeNode {
	var (
		ss  = strings.Split(S, "-")
		sz  = len(ss)
		n   = node{}
		gob = &GodOfBTree{nodes: make([]node, 0, sz)}
	)
	for i := 0; i < sz; i++ {
		s := ss[i]
		if s != "" {
			n.v, _ = strconv.Atoi(s)
			if i != 0 { // exclude root
				n.d++
			}
			gob.nodes = append(gob.nodes, n)
			n = node{}
		} else {
			n.d++
		}
	}

	return doRecover(gob, -1)
}

func main() {
	n := recoverFromPreorder("1-2--3--4-5--6--7")
	fmt.Println(n)
}
