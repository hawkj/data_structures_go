package link_queue

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_enQueue(t *testing.T) {
	lq := (&LinkQueue{}).InitSqQueue()
	lq.enQueue(1)
	lq.enQueue(2)
	lq.enQueue(3)
	lq.enQueue(4)
	if reflect.DeepEqual(lq.linkQueue2Slice(), []interface{}{1, 2, 3, 4}) != true {
		t.Fatal()
	}
	fmt.Println(lq.linkQueue2Slice())
}

func Test_DeQueue(t *testing.T) {
	lq := (&LinkQueue{}).InitSqQueue()
	lq.enQueue(1)
	lq.enQueue(2)
	lq.enQueue(3)
	lq.enQueue(4)
	data := lq.DeQueue()
	fmt.Println(data)
	fmt.Println(lq.linkQueue2Slice())
	if data != 1 {
		t.Fatal()
	}
	if reflect.DeepEqual(lq.linkQueue2Slice(), []interface{}{2, 3, 4}) != true {
		t.Fatal()
	}
}
