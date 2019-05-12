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

func (binTree *BinTree) Empty() bool {
	return binTree.size == 0
}

func (binTree *BinTree) UpdateHeight(node *BinNode) int {
	leftHeight := float64(structure(node.leftChild))
	rightHeight := float64(structure(node.rightChild))
	node.height = int(math.Max(leftHeight, rightHeight)) + 1
	return node.height
}

func (binTree *BinTree) UpdateHeightAbove(node *BinNode) {
	for node != nil {
		binTree.UpdateHeight(node)
		node = node.parent
	}
}

func (binTree *BinTree) InsertAsLeft(node *BinNode, val int) *BinNode {
	binTree.size++
	node.InsertAsLeft(val)
	binTree.UpdateHeightAbove(node)
	return node.leftChild
}

func (binTree *BinTree) InsertAsRight(node *BinNode, val int) *BinNode {
	binTree.size++
	node.InsertAsRight(val)
	binTree.UpdateHeightAbove(node)
	return node.rightChild
}

func structure(p *BinNode) int {
	if p == nil {
		return -1
	} else {
		return p.height
	}
}
