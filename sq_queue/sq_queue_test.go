package sq_queue

import (
	"reflect"
	"testing"
)

func Test_EnQueue(t *testing.T) {
	sqQueue := (&SqQueue{}).InitSqQueue(3)
	sqQueue.EnQueue(1)
	sqQueue.EnQueue(2)
	sqQueue.EnQueue(3)
	sqQueue.EnQueue(4)
	if reflect.DeepEqual([]interface{}{1, 2, nil}, sqQueue.Base) != true {
		t.Fatal()
	}
}

func Test_DeQueue(t *testing.T) {
	sqQueue := (&SqQueue{}).InitSqQueue(3)
	sqQueue.EnQueue(1)
	sqQueue.EnQueue(2)
	sqQueue.EnQueue(3)
	sqQueue.EnQueue(4)
	rs := sqQueue.DeQueue()
	if rs != 1 {
		t.Fatal()
	}
	if reflect.DeepEqual(sqQueue.GetQueue(), []interface{}{2}) != true {
		t.Fatal()
	}
}

func Test_QueueLength(t *testing.T) {
	sqQueue := (&SqQueue{}).InitSqQueue(5)
	sqQueue.EnQueue(1)
	sqQueue.EnQueue(2)
	sqQueue.EnQueue(3)
	sqQueue.EnQueue(4)
	sqQueue.DeQueue()
	sqQueue.EnQueue(5)
	sqQueue.DeQueue()
	if sqQueue.QueueLength() != 3 {
		t.Fatal()
	}
}

func Test_sqQueue(t *testing.T) {
	sqQueue := (&SqQueue{}).InitSqQueue(5)
	sqQueue.EnQueue(1)
	sqQueue.EnQueue(2)
	sqQueue.EnQueue(3)
	sqQueue.EnQueue(4)
	sqQueue.DeQueue()
	sqQueue.EnQueue(5)
	sqQueue.DeQueue()
	sqQueue.EnQueue(6)
	if reflect.DeepEqual(sqQueue.GetQueue(), []interface{}{3, 4, 5, 6}) != true {
		t.Fatal()
	}
}
