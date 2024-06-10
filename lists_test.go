package lists_test

import (
	"testing"

	"github.com/FrMnJ/lists"
)

func TestNode_String(t *testing.T) {
	node := lists.Node[int]{
		Value: 1,
		Next:  nil,
		Prev:  nil,
	}
	if node.String() != "(1)" {
		t.Error("The return value must be (1)")
	}
}

func TestNew(t *testing.T) {
	list := lists.New[int]()
	if list.Head() != nil || list.Tail() != nil {
		t.Error("Recently initialize list head and tail must be nil")
	}
	if list.Length() != 0 {
		t.Error("Recently initialize list length must be 0")
	}
}

func TestList_Add(t *testing.T) {
	list := lists.New[int]()
	list.Add(0)
	list.Add(1)
	if list.Tail().Value != 1 {
		t.Error("Tail value must be 1")
	}
}

func TestList_Insert(t *testing.T) {
	list := lists.New[int]()
	list.Add(0)
	list.Insert(1)
	if list.Head().Value != 1 {
		t.Error("Head value must be 1")
	}
	if list.Tail().Value != 0 {
		t.Error("Tail value must be 0")
	}
}

func TestList_Length(t *testing.T) {
	list := lists.New[int]()
	list.Add(0)
	list.Add(1)
	list.Add(2)
	list.Add(3)
	if list.Length() != 4 {
		t.Error("Length must be 4")
	}
}

func TestList_Find(t *testing.T) {
	list := lists.New[int]()
	list.Add(0)
	list.Add(1)
	list.Add(2)
	list.Add(3)
	node2 := list.Find(2)
	if node2 == nil || node2.Value != 2 {
		t.Error("Node must be 2")
	}
}

func TestList_Remove(t *testing.T) {
	list := lists.New[int]()
	list.Add(0)
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Remove(2)
	if list.Length() != 3 {
		t.Error("Length must be 4")
	}
	if list.Find(2) != nil {
		t.Error("Node 2 must be removed")
	}
}

func TestList_String(t *testing.T) {
	list := lists.New[int]()
	list.Add(0)
	list.Add(1)
	list.Add(2)
	list.Add(3)
	if list.String() != "(0)->(1)->(2)->(3)" {
		t.Error("String must be (0)->(1)->(2)->(3)")
	}
}
