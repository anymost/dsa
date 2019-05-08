package main

// 边数之和等于顶点数之和减一

type BinTreeNode struct {
	Parent     *BinTreeNode
	LeftChild  *BinTreeNode
	RightChild *BinTreeNode
	val        int
}

func (binTreeNode *BinTreeNode) InsertAsLeftChild(node *BinTreeNode) {
	left := binTreeNode.LeftChild
	binTreeNode.LeftChild = node
	binTreeNode.LeftChild.LeftChild = left
}

func (binTreeNode *BinTreeNode) InsertAsRightChild(node *BinTreeNode) {
	right := binTreeNode.RightChild
	binTreeNode.RightChild = node
	binTreeNode.RightChild.RightChild = right
}

type BinTree struct {
	Root *BinTreeNode
}

func NewBinTree() *BinTree  {
	return &BinTree{
		Root: &BinTreeNode{
			Parent:     nil,
			LeftChild:  nil,
			RightChild: nil,
			val:        -1,
		},
	}
}


