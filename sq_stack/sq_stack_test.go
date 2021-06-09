package sq_stack

import (
	"reflect"
	"testing"
)

func Test_Push(t *testing.T) {
	stack := &SqStack{}
	stack = stack.InitSqStack(2)
	stack = stack.Push(1)

	stack = stack.Push(2)
	expect := []interface{}{1, 2}
	if reflect.DeepEqual(expect, stack.Data) != true {
		t.Fatal()
	}
}

func Test_Pop(t *testing.T) {
	stack := &SqStack{}
	stack = stack.InitSqStack(3)
	stack = stack.Push(1)
	stack = stack.Push(2)
	stack = stack.Push(3)

	if stack.Pop() != 3 {
		t.Fatal()
	}

	if stack.Top != 2 {
		t.Fatal()
	}

}
