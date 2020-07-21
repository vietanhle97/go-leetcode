package Serialize_And_Deserialize_Binary_Tree_297

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	root  *TreeNode
	input []string
}

func Constructor() Codec {
	x := Codec{root: nil, input: []string{}}
	return x
}

// Serializes a tree to a single string.
func (node *Codec) serialize(root *TreeNode) string {
	res := node.serializeHelper(root, "")
	return res
}

// Builds up the string
func (node *Codec) serializeHelper(root *TreeNode, res string) string {
	if root == nil {
		return res + "null,"
	} else {
		res += strconv.Itoa(root.Val) + ","
		res = node.serializeHelper(root.Left, res)
		res = node.serializeHelper(root.Right, res)
	}
	return res
}

// Deserializes your encoded data to tree.
func (node *Codec) deserialize(data string) *TreeNode {
	node.input = strings.Split(data, ",")
	return node.deserializeHelper()
}

func (node *Codec) deserializeHelper() *TreeNode {
	if node.input[0] == "null" {
		node.input = node.input[1:]
		return nil
	}
	val, _ := strconv.Atoi(node.input[0])
	node.input = node.input[1:]
	x := TreeNode{Val: val, Left: nil, Right: nil}
	x.Left = node.deserializeHelper()
	x.Right = node.deserializeHelper()
	return &x
}
