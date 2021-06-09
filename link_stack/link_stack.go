package link_stack

type StackNode struct {
	Data interface{}
	Next *StackNode
}

type LinkStack struct {
	Head *StackNode
	Size int
}

func (ls *LinkStack) InitStack() *LinkStack {
	HeadNode := &StackNode{}
	return &LinkStack{
		Head: HeadNode,
		Size: 0,
	}
}

func (ls *LinkStack) Push(e interface{}) *LinkStack {
	node := &StackNode{
		Data: e,
		Next: nil,
	}
	cur := ls.Head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = node
	ls.Size++
	return ls
}

func (ls *LinkStack) Pop() interface{} {
	if ls.Size == 0 {
		return nil
	}
	cur := ls.Head
	for cur.Next.Next != nil {
		cur = cur.Next
	}
	e := cur.Next
	cur.Next = nil
	ls.Size--
	return e.Data
}

func (ls *LinkStack) GetTop() interface{} {
	if ls.Size == 0 {
		return nil
	}
	cur := ls.Head
	for cur.Next != nil {
		cur = cur.Next
	}
	return cur.Data
}

func (ls *LinkStack) stack2slice() []interface{} {
	rs := make([]interface{}, 0)
	cur := ls.Head.Next
	if cur == nil {
		return nil
	}
	for cur != nil {
		rs = append(rs, cur.Data)
		cur = cur.Next
	}
	return rs
}
