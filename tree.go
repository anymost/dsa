package main
// 边数之和等于顶点数之和减一

type TreeNode struct {
	Parent *TreeNode
	FirstChild *TreeNode
	NextSibling *TreeNode
	val interface{}
}

type Tree struct {
	Root *TreeNode
}
