package link_stack

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_Push(t *testing.T) {
	ls := &LinkStack{}
	stack := ls.InitStack()
	stack.Push(1)
	stack.Push(2)
	if reflect.DeepEqual(stack.stack2slice(), []interface{}{1, 2}) != true {
		t.Fatal()
	}
	if stack.Size != 2 {
		t.Fatal()
	}
}

func Test_Pop(t *testing.T) {
	ls := &LinkStack{}
	stack := ls.InitStack()
	stack.Push(1)
	stack.Push(2)
	data := stack.Pop()
	if data != 2 {
		t.Fatal()
	}
	if stack.Size != 1 {
		t.Fatal()
	}
}

func Test_GetTop(t *testing.T) {
	ls := &LinkStack{}
	stack := ls.InitStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println(stack.GetTop())
}
