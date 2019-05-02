package main

import "fmt"

type ListNode struct {
	val  int
	prev *ListNode
	next *ListNode
}

func (node *ListNode) Destroy() {
	node.prev = nil
	node.next = nil
}

func (node *ListNode) Prev() *ListNode {
	return node.prev
}

func (node *ListNode) Next() *ListNode {
	return node.next
}

func (node *ListNode) Data() int {
	return node.val
}

func (node *ListNode) InsertAsPrev(temp *ListNode) {
	prev := node.Prev()
	temp.prev = prev
	temp.next = node
	prev.next = temp
	node.prev = temp
}

func (node *ListNode) InsertAsNext(temp *ListNode) {
	next := node.Next()
	temp.prev = node
	temp.next = next
	next.prev = temp
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

func NewListWithArray(array []int) *List {
	list := NewEmptyList()
	for _, v := range array {
		node := &ListNode{
			val: v,
		}
		list.Tail().InsertAsPrev(node)
	}
	list.size = len(array)
	return list
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

func (list *List) Find(target, length int, start *ListNode) *ListNode {
	start = start.Prev()
	for length > 0 {
		if start.Data() == target {
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

func (list *List) InsertAfter(start *ListNode, target int) {
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
		length--
	}
}

func (list *List) Remove(node *ListNode) int {
	next := node.next
	node.prev.next = next
	node.next.prev = node.prev
	list.size--
	node.Destroy()
	return node.Data()
}

func (list *List) Clear() {
	head, n := list.Head(), list.size
	for n > 0 {
		head = head.Next()
		list.Remove(head)
	}
	list.size = 0
}

func (list *List) Traverse(fn func(*ListNode)) {
	head := list.Head().Next()
	for head != list.Tail() {
		fn(head)
		head = head.next
	}
}

func (list *List) ArrayData() []int {
	array, i := make([]int, list.size), 0
	list.Traverse(func(node *ListNode) {
		array[i] = node.Data()
		i++
	})
	return array
}

func (list *List) Deduplicate() {
	head, tail, r, size := list.Head().Next(), list.Tail(), 0, list.size
	for head != tail {
		node := list.Find(head.Data(), r, head)
		if node != nil {
			list.Remove(node)
			size--
		} else {
			r++
		}
		head = head.Next()
	}
	list.size = size
}

func (list *List) Uniquify() {
	head := list.Head().next
	for head != list.Tail() {
		if head.Next() != list.Tail() && head.Data() == head.Next().Data() {
			list.Remove(head.Next())
		} else {
			head = head.Next()
		}
	}
}

func (list *List) Search(target int, length int, node *ListNode) *ListNode {
	node = node.Prev()
	for length > 0 && node != list.Head() {
		if node.Data() == target {
			return node
		}
		node = node.Prev()
	}
	return nil
}

func (list *List) SelectMax(node *ListNode, length int) *ListNode {
	max := node
	for length > 1 {
		node = node.Next()
		if node.Data() > max.Data() {
			max = node
		}
		length--
	}
	return max
}

func (list *List) SelectionSort() {
	head, last, size := list.Head().Next(), list.Tail(), list.size
	for size > 1 {
		max := list.SelectMax(head, size)
		fmt.Println(max)
		list.Remove(max)
		fmt.Println(max)
		last.InsertAsPrev(max)
		fmt.Println(max)
		last = max
		fmt.Println(max)
		size--
	}
}
