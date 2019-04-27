package main

type ListNode struct {
	val int
	prev *ListNode
	next *ListNode
}

func (node *ListNode) Prev() *ListNode{
	return node.prev
}

func (node *ListNode) Next() *ListNode{
	return node.next
}

func (node *ListNode) Data() int {
	return node.val
}

func (node *ListNode) InsertAsPrev (temp *ListNode) {
	temp.prev = node.prev
	temp.next = node
	node.prev = temp
}

func (node *ListNode) InsertAsNext(temp *ListNode) {
	temp.prev = node
	temp.next = node.next
	node.next = temp
}

type List struct {
	head *ListNode
	tail *ListNode
	size int
}

func NewEmptyList() *List {
	head := &ListNode{}
	tail := &ListNode{}
	head.next = tail
	tail.prev = head
	return &List{
		head,
		tail,
		0,
	}
}

func (list *List) Head() *ListNode {
	return list.head
}

func (list *List) Tail() *ListNode {
	return list.tail
}

func (list *List) Get(index int) int {
	head := list.Head()
	for index > -1 {
		head = head.Next()
		index--
	}
	return head.Data()
}

func (list *List) Find(target int, length int, start *ListNode) *ListNode {
	for length > 0 {
		if start.Data() == target{
			return start
		} else {
			start = start.Prev()
		}
		length--
	}
	return nil
}


func (list *List) InsertBefore(target int, start *ListNode) {
	newNode := &ListNode{
		val: target,
	}
	start.InsertAsPrev(newNode)
	list.size++
}

func (list *List) InsertAfter(start *ListNode, target int ) {
	newNode := &ListNode{
		val: target,
	}
	start.InsertAsNext(newNode)
	list.size++
}


func (list *List) InsertAsLast(node *ListNode) {
	list.tail.InsertAsPrev(node)
	list.size++
}

func (list *List) Copy(node *ListNode, length int) {
	for length > 0 {
		list.InsertAsLast(node)
		node = node.Next()
		list.size++
		length--
	}
}
