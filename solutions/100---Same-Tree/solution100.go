package solution100

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == q {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	a := isSameTree(p.Left, q.Left)
	if !a {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Right, q.Right)
}
