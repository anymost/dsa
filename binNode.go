package main

type BinNode struct {
	parent, leftChild, rightChild *BinNode
	data                          int
	height                        int
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
		data:   val,
	}
	return binNode.leftChild
}

func (binNode *BinNode) InsertAsRight(val int) *BinNode {
	binNode.rightChild = &BinNode{
		parent: binNode,
		data:   val,
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

func Traverse(node *BinNode, callback func(val int)) {
	if node == nil {
		return
	}
	callback(node.data)
	Traverse(node.leftChild, callback)
	Traverse(node.rightChild, callback)
}


func TravelT1(node *BinNode, callback func(val int)) {
	if node == nil {
		return
	}
	stack := NewStack()
	stack.Push(node)
	for !stack.Empty() {
		node = stack.Pop().(*BinNode)
		callback(node.data)
		if node.rightChild != nil {
			stack.Push(node.rightChild)
		}
		if node.leftChild != nil {
			stack.Push(node.leftChild)
		}
	}
}

func TravelT2(node *BinNode, callback func(val int)) {
	if node == nil {
		return
	}
	stack := NewStack()
	for {
		for node != nil {
			callback(node.data)
			stack.Push(node.rightChild)
			node = node.leftChild
		}
		if stack.Empty() {
			break
		}
		node = stack.Pop().(*BinNode)
	}
}

