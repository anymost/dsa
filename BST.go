package main


type BST struct {
	tree *BinTree
}

func NewBST(val int) *BST {
	return &BST{
		tree: NewTree(val),
	}
}

func (bst *BST) Search(target int) *BinNode {
	var hot BinNode
	return bst.searchIn(bst.tree.root, target, &hot)
}

func (bst *BST) searchIn(node *BinNode, target int, hot *BinNode) *BinNode {
	if node == nil || node.Data == target {
		return node
	}
	// hot节点等效的认为是查找命中节点(可能为空，可能命中)的父节点
	hot = node
	if target < node.Data {
		node = node.LeftChild
	} else {
		node = node.RightChild
	}
	return bst.searchIn(node, target, hot)
}



func (bst *BST) Insert(target int) {

}

func (bst *BST) Remove(target int) {

}
