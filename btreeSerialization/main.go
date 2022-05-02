/*297. Serialize and Deserialize Binary Tree

Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed later in the same or another computer environment.

Design an algorithm to serialize and deserialize a binary tree. There is no restriction on how your serialization/deserialization algorithm should work. You just need to ensure that a binary tree can be serialized to a string and this string can be deserialized to the original tree structure.

Example:

You may serialize the following tree:

    1
   / \
  2   3
     / \
    4   5

as "[1,2,3,null,null,4,5]"

Clarification: The above format is the same as how LeetCode serializes a binary tree. You do not necessarily need to follow this format, so please be creative and come up with different approaches yourself.

Note: Do not use class member/global/static variables to store states. Your serialize and deserialize algorithms should be stateless.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

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

type Codec struct {
	vals []string
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	var result string
	if root == nil {
		result += "^,"
	} else {
		result += strconv.Itoa(root.Val) + "," + c.serialize(root.Left) + c.serialize(root.Right)
	}
	return result
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	c.vals = strings.Split(data, ",")
	return c.doDeserialize()
}

func (c *Codec) doDeserialize() *TreeNode {
	if len(c.vals) == 0 {
		return nil
	}
	if c.vals[0] == "^" {
		c.vals = c.vals[1:]
		return nil
	}

	val, _ := strconv.Atoi(c.vals[0])
	c.vals = c.vals[1:]

	return &TreeNode{Val: val,
		Left:  c.doDeserialize(),
		Right: c.doDeserialize()}
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */

func main() {
	obj := Constructor()
	node := obj.deserialize("1,2,^,^,3,4,5")
	ssli := obj.serialize(node)
	fmt.Println(ssli)
}
