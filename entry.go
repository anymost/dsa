package main

// 二叉搜索树 BST 特点是任一几点不小于左后代、不大于右后代，满足顺序性
// 循关键码访问
// 每个BST的节点包含一个词条，Entry结构即为词条，其中的Key为关键码

type Entry struct {
	Key int
	Value interface{}
}

func NewEntry(key int, value interface{}) *Entry {
	return &Entry{
		Key: key,
		Value: value,
	}
}

func (entry *Entry) Less(other *Entry) bool {
	return entry.Key < other.Key
}
func (entry *Entry) Larger(other *Entry) bool {
	return entry.Key > other.Key
}

func (entry *Entry) Equal (other *Entry) bool {
	return entry.Key == other.Key
}

