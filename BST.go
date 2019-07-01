package main


type BST struct {
	tree *BinTree
	hot *BinNode
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
	bst.hot = node
	if target < node.Data {
		node = node.LeftChild
	} else {
		node = node.RightChild
	}
	return bst.searchIn(node, target, hot)
}


func (bst *BST) Insert(target int) *BinNode {
	node := bst.Search(target)
	if node == nil {
		node = &BinNode{
			Data: target,
			Parent: bst.hot,
		}
		bst.tree.size++
		bst.tree.UpdateHeightAbove(node)
	}
	return node
}

func (bst *BST) Remove(target int) bool {
	node := bst.Search(target)
	if node == nil {
		return false
	}
	bst.tree.RemoveAt(node, bst.hot)
	bst.tree.size--
	bst.tree.UpdateHeightAbove(node)
	return true
}

