package sq_queue

type SqQueue struct {
	Base    []interface{}
	Front   int
	Rear    int
	MaxSize int
}

func (sq *SqQueue) InitSqQueue(maxSize int) *SqQueue {
	return &SqQueue{
		Base:    make([]interface{}, maxSize),
		Front:   0,
		Rear:    0,
		MaxSize: maxSize,
	}
}

func (sq *SqQueue) QueueLength() int {
	if sq.Rear > sq.Front {
		return sq.Rear - sq.Front
	}
	if sq.Rear < sq.Front {
		return (sq.Rear) + (sq.MaxSize - sq.Front)
	}
	return 0
}

func (sq *SqQueue) EnQueue(data interface{}) *SqQueue {
	if (sq.Rear+1)%sq.MaxSize == sq.Front {
		return nil
	}
	sq.Base[sq.Rear] = data
	sq.Rear = (sq.Rear + 1) % sq.MaxSize
	return sq
}

func (sq *SqQueue) DeQueue() interface{} {
	if sq.Front == sq.Rear {
		return nil
	}
	e := sq.Base[sq.Front]
	sq.Front = (sq.Front + 1) % sq.MaxSize
	return e
}

func (sq *SqQueue) GetQueue() []interface{} {
	if sq.Rear > sq.Front {
		return sq.Base[sq.Front:sq.Rear]
	}
	if sq.Rear < sq.Front {
		partRear := sq.Base[0:sq.Rear]
		partFront := sq.Base[sq.Front:]
		if partRear != nil {
			return append(partFront, partRear...)
		} else {
			return partFront
		}
	}
	return nil
}
