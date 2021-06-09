package link_queue

type QNode struct {
	Data interface{}
	Next *QNode
}

type LinkQueue struct {
	Length int
	Front  *QNode
	Rear   *QNode
}

func (lq *LinkQueue) InitSqQueue() *LinkQueue {
	node := &QNode{}
	return &LinkQueue{
		Length: 0,
		Front:  node,
		Rear:   node,
	}
}

func (lq *LinkQueue) enQueue(data interface{}) *LinkQueue {
	e := &QNode{
		Data: data,
		Next: nil,
	}
	if lq.Length == 0 {
		lq.Front = e
		lq.Rear = e
	} else {
		lq.Rear.Next = e
		lq.Rear = e
	}
	lq.Length++
	return lq
}

func (lq *LinkQueue) DeQueue() interface{} {
	if lq.Length == 0 {
		return nil
	}
	e := lq.Front
	data := e.Data
	if lq.Length == 1 {
		node := &QNode{}
		lq.Front = node
		lq.Rear = node
	} else {
		lq.Front = lq.Front.Next
	}
	lq.Length--
	e = nil
	return data
}

func (lq *LinkQueue) linkQueue2Slice() []interface{} {
	slice := []interface{}{}
	if lq.Length == 0 {
		return nil
	}
	cur := lq.Front
	for cur != nil {
		slice = append(slice, cur.Data)
		if cur.Next != nil {
			cur = cur.Next
		} else {
			break
		}
	}
	return slice
}
