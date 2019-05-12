package main

// 边数之和等于顶点数之和减一
// 满二叉树节点为2^(h+1)-1, h为高度
// 真二叉树：不存在子节点数为1的节点
// 对于有根有序树，通过长子兄弟表示法，将多叉树转化为二叉树

type BinNode struct {
	parent, leftChild, rightChild *BinNode
	data                          int
	height                        int
}

type BinTree struct {
	size int
	root *BinNode
}

func (binNode *BinNode) Size() int {
	size := 1
	if binNode.leftChild != nil {
		size += binNode.leftChild.Size()
	}
	if binNode.rightChild != nil {
		size += binNode.rightChild.Size()
	}
	return size
}

func (binNode *BinNode) InsertAsLeft(val int) *BinNode {
	binNode.leftChild = &BinNode{
		parent: binNode,
		data: val,
	}
	return binNode.leftChild
}

func (binNode *BinNode) InsertAsRight(val int) *BinNode {
	binNode.rightChild = &BinNode{
		parent: binNode,
		data: val,
	}
	return binNode.rightChild
}

func (binNode *BinNode) Succ() *BinNode {
	return &BinNode{}
}

func (binNode *BinNode) TravelLevel(callback func(node *BinNode)) {

}

func (binNode *BinNode) TravelPre(callback func(node *BinNode)) {

}

func (binNode *BinNode) TravelIn(callback func(node *BinNode)) {

}

func (binNode *BinNode) TravelPost(callback func(node *BinNode)) {

}

func (binTree *BinTree) Empty() bool {
	return binTree.size == 0
}
