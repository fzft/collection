// see http://staff.ustc.edu.cn/~csli/graduate/algorithms/book6/chap14.htm for details

package collection

type Color int

type Direction int

const (
	Black Color = iota
	Red
)

type t interface{}

type keyType interface {
	Less(interface{}) bool
}

type Rbt struct {
	root *RbtNode
	size int
}

type RbtNode struct {
	value               t
	key                 keyType
	parent, left, right *RbtNode
	color               Color
}

func NewRbt() *Rbt {
	return &Rbt{}
}

func (t *Rbt) Empty() bool {
	if t.root != nil {
		return false
	}
	return true
}

// Search for a node with given key in a binary tree
func (t *Rbt) Search(key keyType) *RbtNode {
	if t.Empty() || key == t.root.key {
		return t.root
	}
	if key.Less(t.root.key) {
		return t.search(t.root.left, key)
	} else {
		return t.search(t.root.right, key)
	}
}

func (t *Rbt) Clear() {
	t.root = nil
	t.size = 0
}

func (t *Rbt) Size() int {
	return t.size
}

// Insert implementation of the book intro to algorithms in chap 13.3 Tree-Insert procedure
func (t *Rbt) Insert(k keyType, val t) {
	var y *RbtNode
	x := t.root

	for x != nil {
		y = x
		if k.Less(x.key) {
			x = x.left
		} else {
			x = x.right
		}
	}

	z := &RbtNode{value: val, key: k, color: Red}
	t.size++
	z.parent = y
	if y == nil {
		z.color = Black
		t.root = z
		return
	} else if z.key.Less(y.key) {
		y.left = z
	} else {
		y.right = z
	}
	t.rbInsertFixup(z)
}

func (t *Rbt) Delete(k keyType) *RbtNode {
	z := t.Search(k)
	if z == nil {
		return nil
	}
	var x, y *RbtNode
	if z.left != nil && z.right != nil {
		y = t.successor(z)
	} else {
		y = z
	}
	if y.left != nil {
		x = y.left
	} else {
		x = y.right
	}
	if x != nil {
		x.parent = y.parent
	}

	if y.parent == nil {
		t.root = x
	} else if y == y.parent.left {
		y.parent.left = x
	} else {
		y.parent.right = x
	}
	if y != z {
		z.key = y.key
		z.value = y.value

	}
	if y.color == Black {
		t.deleteFixup(x, y.parent)
	}
	t.size--
	return y
}

// leftRotation assume right[x] != nil
func (t *Rbt) leftRotation(x *RbtNode) {
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

func (t *Rbt) rightRotation(x *RbtNode) {
	y := x.left
	x.left = y.right
	if y.right != nil {
		y.right.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.right {
		x.parent.right = y
	} else {
		x.parent.left = y
	}
	y.right = x
	x.parent = y
}

func (t *Rbt) insertFixup(z *RbtNode) {
	var y *RbtNode
	for z.parent != nil && z.parent.color == Red {
		if z.parent == z.parent.parent.left {
			y = z.parent.parent.right
			if y != nil && y.color == Red {
				z.parent.color = Black
				y.color = Black
				z.parent.parent.color = Red
				z = z.parent.parent
			} else {
				if z == z.parent.right {
					z = z.parent
					t.leftRotation(z)
				}
				z.parent.color = Black
				z.parent.parent.color = Red
				t.rightRotation(z.parent.parent)
			}
		} else {
			y = z.parent.parent.left
			if y != nil && y.color == Red {
				z.parent.color = Black
				y.color = Black
				z.parent.parent.color = Red
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					z = z.parent
					t.rightRotation(z)
				}
				z.parent.color = Black
				z.parent.parent.color = Red
				t.leftRotation(z.parent.parent)
			}
		}
	}
	t.root.color = Black
}

func (t *Rbt) rbInsertFixup(z *RbtNode) {
	var y *RbtNode
	for z.parent != nil && z.parent.color == Red {
		if z.parent == z.parent.parent.left {
			y = z.parent.parent.right
			if y != nil && y.color == Red {
				z.parent.color = Black
				y.color = Black
				z.parent.parent.color = Red
				z = z.parent.parent
			} else {
				if z == z.parent.right {
					z = z.parent
					t.leftRotation(z)
				}
				z.parent.color = Black
				z.parent.parent.color = Red
				t.rightRotation(z.parent.parent)
			}
		} else {
			y = z.parent.parent.left
			if y != nil && y.color == Red {
				z.parent.color = Black
				y.color = Black
				z.parent.parent.color = Red
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					z = z.parent
					t.rightRotation(z)
				}
				z.parent.color = Black
				z.parent.parent.color = Red
				t.leftRotation(z.parent.parent)
			}
		}
	}
	t.root.color = Black
}

// minimum an element in a binary tree whose key is minimum
func (t *Rbt) minimum(n *RbtNode) *RbtNode {
	for n.left != nil {
		n = n.left
	}
	return n
}

func (t *Rbt) search(x *RbtNode, k keyType) *RbtNode {
	for x != nil && k != x.key {
		if k.Less(x.key) {
			x = x.left
		} else {
			x = x.right
		}
	}
	return x
}

func (t *Rbt) successor(x *RbtNode) *RbtNode {
	var y *RbtNode
	if x.right != nil {
		return t.minimum(x.right)
	}
	y = x.parent
	for y != nil && x == y.right {
		x = y
		y = y.parent
	}
	return y
}

func (t *Rbt) deleteFixup(x, parent *RbtNode) {
	var w *RbtNode

	for x != t.root && getColor(x) == Black {
		if x != nil {
			parent = x.parent
		}
		if x == parent.left {
			w = parent.right
			if w.color == Red {
				w.color = Black
				parent.color = Red
				t.leftRotation(parent)
				w = parent.right
			}
			if getColor(w.left) == Black && getColor(w.right) == Black {
				w.color = Red
				x = parent
			} else {
				if getColor(w.right) == Black {
					if w.left != nil {
						w.left.color = Black
					}
					w.color = Red
					t.rightRotation(w)
					w = parent.right
				}
				w.color = parent.color
				parent.color = Black
				if w.right != nil {
					w.right.color = Black
				}
				t.leftRotation(parent)
				x = t.root
			}
		} else {
			w = parent.left
			if w.color == Red {
				w.color = Black
				parent.color = Red
				t.rightRotation(parent)
				w = parent.left
			}
			if getColor(w.left) == Black && getColor(w.right) == Black {
				w.color = Red
				x = parent
			} else {
				if getColor(w.left) == Black {
					if w.right != nil {
						w.right.color = Black
					}
					w.color = Red
					t.leftRotation(w)
					w = parent.left
				}
				w.color = parent.color
				parent.color = Black
				if w.left != nil {
					w.left.color = Black
				}
				t.rightRotation(parent)
				x = t.root
			}
		}
	}
	if x != nil {
		x.color = Black
	}
}

// getColor gets color of the node.
func getColor(n *RbtNode) Color {
	if n == nil {
		return Black
	}
	return n.color
}
