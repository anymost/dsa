package main

type BinNode struct {
	Parent, LeftChild, RightChild *BinNode
	Data                          int
	Height                        int
}

func (binNode *BinNode) Size() int {
	size := 1
	if binNode.LeftChild != nil {
		size += binNode.LeftChild.Size()
	}
	if binNode.RightChild != nil {
		size += binNode.RightChild.Size()
	}
	return size
}

func (binNode *BinNode) InsertAsLeft(val int) *BinNode {
	binNode.LeftChild = &BinNode{
		Parent: binNode,
		Data:   val,
	}
	return binNode.LeftChild
}

func (binNode *BinNode) InsertAsRight(val int) *BinNode {
	binNode.RightChild = &BinNode{
		Parent: binNode,
		Data:   val,
	}
	return binNode.RightChild
}

func (binNode *BinNode) Succ() *BinNode {
	var succ *BinNode
	if binNode.RightChild != nil {
		node := binNode.RightChild
		for node != nil {
			succ = node
			node = node.LeftChild
		}
	}
	return succ
}


func TraversePreOrder1(node *BinNode, callback func(val int)) {
	if node == nil {
		return
	}
	callback(node.Data)
	TraversePreOrder1(node.LeftChild, callback)
	TraversePreOrder1(node.RightChild, callback)
}


func TraverseMidOrder1(node *BinNode, callback func(val int)) {
	if node == nil {
		return
	}
	TraverseMidOrder1(node.LeftChild, callback)
	callback(node.Data)
	TraverseMidOrder1(node.RightChild, callback)
}

func TraversePreOrder(node *BinNode, visitor func(val int)) {
	if node == nil {
		return
	}
	stack := NewStack()
	for {
		for node != nil {
			visitor(node.Data)
			stack.Push(node.RightChild)
			node = node.LeftChild
		}
		if stack.Empty() {
			break
		}
		node = stack.Pop().(*BinNode)
	}
}

func TraverseInOrder(node *BinNode, visitor func(val int)) {
	if node == nil {
		return
	}
	stack := NewStack()
	for {
		for node != nil {
			stack.Push(node)
			node = node.LeftChild
		}
		if stack.Empty() {
			break
		}
		node = stack.Pop().(*BinNode)
		visitor(node.Data)
		node = node.RightChild
	}
}

func TraverseLevel(node *BinNode, visitor func(val int)) {
	if node == nil {
		return
	}
	queue := NewQueue()
	queue.Enqueue(node)
	for !queue.Empty() {
		node = queue.Dequeue().(*BinNode)
		visitor(node.Data)
		if node.LeftChild != nil {
			queue.Enqueue(node.LeftChild)
		}
		if node.RightChild != nil {
			queue.Enqueue(node.RightChild)
		}
	}
}


