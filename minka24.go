
import (
  "math"
  "fmt"
)

type TreeNode struct {
  Val int
  Left *TreeNode
  Right *TreeNode
}

func insert(root *TreeNode, min, max int) bool {
  if root == nil {
    return true
  }

  if min <= root.Val && root.Val <= max {
    if insert(root.Left, min, root.Val) && insert(root.Right, root.Val, max) {
      return true
    }
  }

  return false
}

func isValidBST(root *TreeNode) bool {
  return insert(root, 0, 10000)
}
