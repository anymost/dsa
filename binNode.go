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

func TraversePreOrder1(node *BinNode, callback func(val int)) {
	if node == nil {
		return
	}
	callback(node.data)
	TraversePreOrder1(node.leftChild, callback)
	TraversePreOrder1(node.rightChild, callback)
}


func TraverseMidOrder1(node *BinNode, callback func(val int)) {
	if node == nil {
		return
	}
	TraverseMidOrder1(node.leftChild, callback)
	callback(node.data)
	TraverseMidOrder1(node.rightChild, callback)
}


func TraversePreOrder2(node *BinNode, callback func(val int)) {
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


func TraverseMidOrder2(node *BinNode, callback func(val int)) {
	if node == nil {
		return
	}
	stack := NewStack()
	for {
		for node != nil {
			stack.Push(node)
			node = node.leftChild
		}
		if stack.Empty() {
			break
		}
		node = stack.Pop().(*BinNode)
		callback(node.data)
		node = node.rightChild
	}
}

func TraverseLevel(node *BinNode, callback func(val int)) {
	if node == nil {
		return
	}
	queue := NewQueue()
	queue.Enqueue(node)
	for !queue.Empty() {
		node := queue.Dequeue().(*BinNode)
		callback(node.data)
		if node.leftChild != nil {
			queue.Enqueue(node.leftChild)
		}
		if node.rightChild != nil {
			queue.Enqueue(node.rightChild)
		}
	}
}


