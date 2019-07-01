package main

import "math"

// 边数之和等于顶点数之和减一
// 满二叉树节点为2^(h+1)-1, h为高度
// 真二叉树：不存在子节点数为1的节点
// 对于有根有序树，通过长子兄弟表示法，将多叉树转化为二叉树


type BinTree struct {
	size int
	root *BinNode
}

func NewTree(val int) *BinTree {
	return &BinTree{
		size: 1,
		root: &BinNode{
			Data: val,
		},
	}
}

func (binTree *BinTree) Empty() bool {
	return binTree.size == 0
}

func (binTree *BinTree) UpdateHeight(node *BinNode) int {
	leftHeight := float64(structure(node.LeftChild))
	rightHeight := float64(structure(node.RightChild))
	node.Height = int(math.Max(leftHeight, rightHeight)) + 1
	return node.Height
}

func (BinTree *BinTree) UpdateHeightAbove(node *BinNode) {
	for node != nil {
		BinTree.UpdateHeight(node)
		node = node.Parent
	}
}

func (binTree *BinTree) InsertAsLeft(node *BinNode, val int) *BinNode {
	binTree.size++
	node.InsertAsLeft(val)
	binTree.UpdateHeightAbove(node)
	return node.LeftChild
}

func (binTree *BinTree) InsertAsRight(node *BinNode, val int) *BinNode {
	binTree.size++
	node.InsertAsRight(val)
	binTree.UpdateHeightAbove(node)
	return node.RightChild
}


func (binTree *BinTree) RemoveAt(node *BinNode, hot *BinNode) *BinNode {
	w := node
	var succ *BinNode
	if node.LeftChild == nil {
		node = node.RightChild
		succ = node
	} else if node.RightChild == nil {
		node = node.LeftChild
		succ = node
	} else {
		succ := w.Succ()
		succ.Data, w.Data = w.Data, succ.Data
		u := w.Parent
		_ = u
	}
	hot = w.Parent
	if succ != nil {
		succ.Parent = hot
	}
	return succ
}

func structure(p *BinNode) int {
	if p == nil {
		return -1
	} else {
		return p.Height
	}
}
