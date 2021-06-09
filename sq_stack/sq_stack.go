package sq_stack

type SqStack struct {
	Base      int
	Top       int
	StackSize int
	Data      []interface{}
}

func (sqStack *SqStack) InitSqStack(stackSize int) *SqStack {
	return &SqStack{
		Base:      0,
		Top:       0,
		StackSize: stackSize,
		Data:      make([]interface{}, stackSize),
	}
}

func (sqStack *SqStack) Push(data interface{}) *SqStack {
	if sqStack.Top-sqStack.Base == sqStack.StackSize {
		return nil
	}
	sqStack.Data[sqStack.Top] = data
	sqStack.Top++
	return sqStack
}

func (sqStack *SqStack) Pop() interface{} {
	if sqStack.Top == sqStack.Base {
		return nil
	}
	elem := sqStack.Data[sqStack.Top-1]
	sqStack.Top--
	return elem
}
