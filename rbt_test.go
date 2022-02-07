package collection

import (
	"fmt"
	"testing"
)

type key int

func (k key) Less(b interface{}) bool {
	val := b.(key)
	return k < val
}

func (t *Rbt) Preorder() {
	fmt.Println("preorder begin")
	if t.root != nil {
		t.root.preorder()
	}
	fmt.Println("preorder end")
}

func (n *RbtNode) preorder() {
	fmt.Printf("(%v %v) ", n.key, n.value)
	if n.parent == nil {
		fmt.Printf("nil")
	} else {
		fmt.Printf("whose parent is %v", n.parent.key)
	}
	if n.color == Red {
		fmt.Println(" and color RED")
	} else {
		fmt.Println(" and color BLACK")
	}
	if n.left != nil {
		fmt.Printf("%v's left child is ", n.key)
		n.left.preorder()
	}
	if n.right != nil {
		fmt.Printf("%v's right child is ", n.key)
		n.right.preorder()
	}
}

func TestPreorder(t *testing.T) {
	tree := NewRbt()
	if !tree.Empty() {
		t.Error("tree not empty")
	}

	tree.Insert(key(1), "123")
	tree.Insert(key(3), "234")
	tree.Insert(key(4), "dfa3")
	tree.Insert(key(6), "sd4")
	tree.Insert(key(5), "jcd4")
	tree.Insert(key(2), "bcd4")
	if tree.Size() != 6 {
		t.Error("Error size")
	}
	if tree.Empty() {
		t.Error("tree empty")
	}
	tree.Preorder()
}

func TestFind(t *testing.T) {

	tree := NewRbt()
	tree.Insert(key(1), "123")
	tree.Insert(key(3), "234")
	tree.Insert(key(4), "dfa3")
	tree.Insert(key(6), "sd4")
	tree.Insert(key(5), "jcd4")
	tree.Insert(key(2), "bcd4")

	n := tree.Search(key(4))
	if n.value != "dfa3" {
		t.Error("Error value")
	}
	n.value = "bdsf"
	if n.value != "bdsf" {
		t.Error("Error value modify")
	}
	value := tree.Search(key(5)).value.(string)
	if value != "jcd4" {
		t.Error("Error value after modifyed other node")
	}
}

func TestDelete(t *testing.T) {
	tree := NewRbt()

	tree.Insert(key(1), "123")
	tree.Insert(key(3), "234")
	tree.Insert(key(4), "dfa3")
	tree.Insert(key(6), "sd4")
	tree.Insert(key(5), "jcd4")
	tree.Insert(key(2), "bcd4")
	for i := 1; i <= 6; i++ {
		tree.Delete(key(i))
		if tree.Size() != 6-i {
			t.Error("Delete Error")
		}
	}
	tree.Insert(key(1), "bcd4")
	tree.Clear()
	tree.Preorder()
	if tree.Search(key(1)) != nil {
		t.Error("Can't clear")
	}
}
