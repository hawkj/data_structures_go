package link_list

import (
	"reflect"
	"testing"
)

func Test_sliceAndLinkList(t *testing.T) {
	slice4test := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	linkList := LinkList{}
	//先把切片转成list
	slice2LinkListResult := linkList.slice2LinkList(slice4test)
	//再把list转成切片
	linkList2SliceResult := slice2LinkListResult.linkList2Slice()
	if reflect.DeepEqual(slice4test, linkList2SliceResult) != true {
		t.Fatal()
	}
}

func Test_getElem(t *testing.T) {
	linkList := LinkList{}
	i := 11
	slice4test := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	slice2LinkListResult := linkList.slice2LinkList(slice4test)
	r := slice2LinkListResult.GetElem(i)
	if i > len(slice4test) {
		if r != nil {
			t.Fatal()
		}
	} else if r.Data != slice4test[i-1] {
		t.Fatal()
	}
}

func Test_localElem(t *testing.T) {
	linkList := LinkList{}
	slice4test := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	slice2LinkListResult := linkList.slice2LinkList(slice4test)
	r := slice2LinkListResult.LocalElem(4)
	expect := slice2LinkListResult.GetElem(4)
	if r != expect {
		t.Fatal()
	}
}

func Test_listInsert(t *testing.T) {
	linkList := LinkList{}
	slice4test := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	e := &LinkNode{
		Data: 99,
		Next: nil}
	slice2LinkListResult := linkList.slice2LinkList(slice4test)
	linkListInserted := slice2LinkListResult.ListInsert(4, e)
	r := linkListInserted.linkList2Slice()
	expect := []interface{}{1, 2, 3, 99, 4, 5, 6, 7, 8, 9, 0}
	if reflect.DeepEqual(r, expect) != true {
		t.Fatal()
	}
}

func Test_ListDelete(t *testing.T) {
	linkList := LinkList{}
	slice4test := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	i := 5
	slice2LinkListResult := linkList.slice2LinkList(slice4test)
	linkListDeleted := slice2LinkListResult.ListDelete(i)
	r := linkListDeleted.linkList2Slice()
	expect := []interface{}{1, 2, 3, 4, 6, 7, 8, 9, 0}
	if reflect.DeepEqual(r, expect) != true {
		t.Fatal()
	}
}

func Test_length(t *testing.T) {
	linkList := LinkList{}
	slice4test := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	slice2LinkListResult := linkList.slice2LinkList(slice4test)
	length := slice2LinkListResult.Length()
	if length != len(slice4test) {
		t.Fatal()
	}
}

func Test_add(t *testing.T) {
	linkList := LinkList{}
	slice4test := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	e := &LinkNode{
		Data: 99,
		Next: nil,
	}
	slice2LinkListResult := linkList.slice2LinkList(slice4test)
	r := slice2LinkListResult.Add(e)
	rSlice := r.linkList2Slice()
	expect := []interface{}{99, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	if reflect.DeepEqual(rSlice, expect) != true {
		t.Fatal()
	}
}

func Test_append(t *testing.T) {
	linkList := LinkList{}
	slice4test := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	e := &LinkNode{
		Data: 99,
		Next: nil,
	}
	slice2LinkListResult := linkList.slice2LinkList(slice4test)
	r := slice2LinkListResult.Append(e)
	rSlice := r.linkList2Slice()
	expect := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 99}
	if reflect.DeepEqual(rSlice, expect) != true {
		t.Fatal()
	}
}

func Test_ListDeleteByData(t *testing.T) {
	linkList := LinkList{}
	slice4test := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	slice2LinkListResult := linkList.slice2LinkList(slice4test)
	r := slice2LinkListResult.ListDeleteByData(4)
	rSlice := r.linkList2Slice()
	expect := []interface{}{1, 2, 3, 5, 6, 7, 8, 9, 0}
	if reflect.DeepEqual(rSlice, expect) != true {
		t.Fatal()
	}
}
