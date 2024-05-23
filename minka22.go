package main

import "strconv"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func i2s(x int) string {
	return strconv.Itoa(x)
}

func s2i(str string) int {
	x, err := strconv.Atoi(str)
	if err == nil {
		return x
	}
	return 0
}

type Codec struct {
	s string
}

func Constructor() Codec {
	return Codec{s: ""}
}

func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	this.s = this.s + "(" + i2s(root.Val)

	if root.Left != nil {
		this.s = this.serialize(root.Left)
	} else {
		this.s = this.s + "x"
	}
	if root.Right != nil {
		this.s = this.serialize(root.Right)
	} else {
		this.s = this.s + "x"
	}
	return this.s + ")"
}

func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	var (
		tree = &TreeNode{Val: 0, Left: nil, Right: nil}
		spl1 = 0
		spl2 = 0
		spl3 = 0
	)
	for i := 1; ; i++ {
		if data[i] == 'x' || data[i] == '(' {
			spl1 = i
			break
		}
	}
	tree.Val = s2i(data[1:spl1])

	leftbr := 0
	for i := spl1; ; i++ {
		if data[i] == '(' {
			leftbr++
		}
		if data[i] == ')' {
			leftbr--
		}
		if leftbr == 0 {
			spl2 = i
			break
		}
	}
	if spl2 == spl1 {
		tree.Left = nil
	} else {
		tree.Left = this.deserialize(data[spl1 : spl2+1])
	}

	for i := spl2 + 1; ; i++ {
		if data[i] == '(' {
			leftbr++
		}
		if data[i] == ')' {
			leftbr--
		}
		if leftbr == 0 {
			spl3 = i
			break
		}
	}
	if spl2 == spl3-1 {
		tree.Right = nil
	} else {
		tree.Right = this.deserialize(data[spl2+1 : spl3+1])
	}

	return tree
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

func main() {
	ser := Constructor()
	deser := Constructor()
	data := ser.serialize(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
	})
	tree := deser.deserialize(data)
	println(data, tree.Val)
}
