package array

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func checkTree(root *TreeNode) bool {
	if root.Left.Val+root.Right.Val == root.Val {
		return true
	} else {
		fmt.Println(root.Val)
		fmt.Println(root.Left.Val)
		return false
	}
}
