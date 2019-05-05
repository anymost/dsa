package main

type Queue struct {
	val []interface{}
}

func NewQueue() *Queue {
	return &Queue{
		val: make([]interface{}, 0),
	}
}

func (queue *Queue) Enqueue(v interface{}) {
	queue.val = append(queue.val, v)
}

func (queue *Queue) Dequeue() interface{} {
	v := queue.val[0]
	queue.val = queue.val[1:]
	return v
}

func (queue *Queue) Front() interface{} {
	return queue.val[0]
}

func (queue *Queue) Empty() bool {
	return len(queue.val) == 0
}
