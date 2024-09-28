package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type lay struct {
	allnil  bool
	arr     []*TreeNode
	lasttre *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	otv := make([]int, 0)

	arr1 := make([]*TreeNode, 0)
	lastlay := lay{arr: arr1, allnil: false, lasttre: nil}
	arr2 := make([]*TreeNode, 0)
	thislay := lay{arr: arr2, allnil: false, lasttre: nil}

	lastlay.arr = append(lastlay.arr, root)
	lastlay.allnil = false
	lastlay.lasttre = root

	for lastlay.allnil == false {
		otv = append(otv, lastlay.lasttre.Val)
		for _, tree := range lastlay.arr {
			if tree != nil {
				if tree.Left != nil {
					thislay.lasttre = tree.Left
				}

				if tree.Right != nil {
					thislay.lasttre = tree.Right
				}

				thislay.arr = append(thislay.arr, tree.Left, tree.Right)
			}
		}
		if thislay.lasttre == nil {
			thislay.allnil = true
		}
		lastlay = thislay
		thislay = lay{arr: make([]*TreeNode, 0), allnil: false, lasttre: nil}
	}
	return otv
}

func main() {
	fmt.Println(rightSideView(&TreeNode{
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
	}))
}
