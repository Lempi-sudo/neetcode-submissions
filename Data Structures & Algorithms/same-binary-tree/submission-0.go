/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
   	if q == nil && p == nil {
		return true
	}

	flag := true
	if p != nil && q != nil && p.Val == q.Val {
		flag = flag && isSameTree(p.Left, q.Left)
		flag = flag && isSameTree(p.Right, q.Right)
	} else {
		return false
	}
	return flag
} 

