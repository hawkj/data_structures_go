package link_list

type LinkNode struct {
	Data interface{}
	Next *LinkNode
}

type LinkList struct {
	Head *LinkNode
}

func (linklist *LinkList) InitLinkList() LinkList {
	list := LinkList{
		&LinkNode{nil, nil},
	}
	return list
}

//单链表查找第i个元素
func (linklist *LinkList) GetElem(i int) *LinkNode {
	cur := linklist.Head
	for j := 1; j < i; j++ {
		if cur.Next == nil {
			return nil
		}
		cur = cur.Next
	}
	return cur
}

//单链表的按值查找
func (linklist *LinkList) LocalElem(data interface{}) *LinkNode {
	cur := linklist.Head
	if cur == nil && cur.Data != data {
		return nil
	}
	for {
		if cur.Data == data {
			return cur
		}
		if cur.Next == nil {
			break
		}
		cur = cur.Next
	}
	return nil
}

//将值为e的新结点插入到表的第i个结点的位置上，
func (linklist *LinkList) ListInsert(i int, e *LinkNode) *LinkList {
	preNode := linklist.GetElem(i - 1)
	if preNode == nil {
		return nil
	}
	temp := preNode.Next
	preNode.Next = e
	e.Next = temp
	return linklist
}

//单链表的删除
func (linklist *LinkList) ListDelete(i int) *LinkList {
	preNode := linklist.GetElem(i - 1)
	if preNode == nil {
		return nil
	}
	delElem := preNode.Next
	preNode.Next = preNode.Next.Next
	delElem.Next = nil
	return linklist
}

func (linklist *LinkList) Length() int {
	cur := linklist.Head
	count := 0
	for cur != nil {
		count++
		cur = cur.Next
	}
	return count
}

func (linklist *LinkList) Add(e *LinkNode) *LinkList {
	head := linklist.Head
	if head == nil {
		return nil
	}
	e.Next = head
	linklist.Head = e
	return linklist
}

func (linklist *LinkList) Append(e *LinkNode) *LinkList {
	cur := linklist.Head
	if cur == nil {
		return nil
	}

	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = e
	return linklist
}

func (linklist *LinkList) ListDeleteByData(data interface{}) *LinkList {
	cur := linklist.Head
	//单独处理头结点
	if cur.Data == data {
		linklist.Head = cur.Next
		return linklist
	}
	for cur.Next != nil {
		if cur.Next.Data == data {
			temp := cur.Next
			cur.Next = cur.Next.Next
			temp.Next = nil
			return linklist
		} else {
			cur = cur.Next
		}
	}
	return nil
}

//tools
func (List LinkList) slice2LinkList(s []interface{}) *LinkList {
	if len(s) == 0 {
		return nil
	}
	cur := &LinkNode{}
	sentinel := cur
	for _, data := range s {
		linknode := &LinkNode{
			Data: data,
			Next: nil,
		}
		cur.Next = linknode
		cur = linknode
	}
	linklist := LinkList{
		Head: sentinel.Next,
	}
	return &linklist
}

func (linklist *LinkList) linkList2Slice() []interface{} {
	rs := make([]interface{}, 0)
	cur := linklist.Head
	if cur == nil {
		return nil
	}
	for cur != nil {
		rs = append(rs, cur.Data)
		cur = cur.Next
	}
	return rs
}
