package rbtree

import (
	"sync"
)

type RBNoder interface {
	Less(x RBNoder) bool
	Equal(x RBNoder) bool
}

type rbtree struct {
	root *node
	pool sync.Pool
}

type node struct {
	val    RBNoder
	parent *node
	left   *node
	right  *node
	red    bool
}

func New() *rbtree {
	rb := &rbtree{}
	rb.pool.New = func() interface{} {
		return &node{}
	}
	return rb
}

func (rb *rbtree) get() *node {
	n := rb.pool.Get().(*node)
	n.red = true
	return n
}

func (rb *rbtree) put(n *node) {
	n.val = nil
	n.parent = nil
	n.left = nil
	n.right = nil
	rb.pool.Put(n)
}

func (rb *rbtree) leftRotate(n *node) {
	right := n.right
	if right == nil {
		return
	}
	n.right = right.left
	if n.right != nil {
		n.right.parent = n
	}
	right.left = n
	right.parent = n.parent
	n.parent = right
	if right.parent != nil {
		if n == right.parent.right {
			right.parent.right = right
		} else {
			right.parent.left = right
		}
	}
	if rb.root == n {
		rb.root = right
	}
}

func (rb *rbtree) rightRotate(n *node) {
	left := n.left
	if left == nil {
		return
	}
	n.left = left.right
	if n.left != nil {
		n.left.parent = n
	}
	left.right = n
	left.parent = n.parent
	n.parent = left
	if left.parent != nil {
		if n == left.parent.left {
			left.parent.left = left
		} else {
			left.parent.right = left
		}
	}
	if rb.root == n {
		rb.root = left
	}
}

func (rb *rbtree) adjust(n *node) {
	p := n.parent
	if p == nil { // n is root node
		n.red = false
		return
	}
	if !p.red {
		return
	}
	gp := p.parent
	if gp == nil { // p is root node
		p.red = false
		return
	}
	var un *node
	if gp.left == p {
		un = gp.right
	} else {
		un = gp.left
	}
	if un == nil || !un.red {
		if p.left == n {
			p.red = false
			gp.red = true
			rb.rightRotate(gp)
		} else {
			rb.leftRotate(p)
			rb.adjust(p)
		}
	} else {
		p.red = false
		un.red = false
		gp.red = true
		rb.adjust(gp)
	}
}

func (rb *rbtree) Add(x RBNoder) bool {
	if rb.root == nil {
		rb.root = rb.get()
		rb.root.val = x
		rb.root.red = false
		return true
	}
	n := rb.root
	for n != nil {
		if n.val.Less(x) {
			if n.right == nil {
				nt := rb.get()
				nt.val = x
				nt.parent = n
				n.right = nt
				rb.adjust(nt)
				return true
			}
			n = n.right
		} else {
			if n.left == nil {
				nt := rb.get()
				nt.val = x
				nt.parent = n
				n.left = nt
				rb.adjust(nt)
				return true
			}
			n = n.left
		}
	}
	return false
}

func (rb *rbtree) del(n *node) {
	if n.left == nil || n.right == nil {
		nt := n.left
		if n.right != nil {
			nt = n.right
		}
		if n.parent != nil {
			if n.parent.left == n {
				n.parent.left = nt
			} else {
				n.parent.right = nt
			}
		} else {
			rb.root = nt
		}
	} else {

	}
	rb.put(n)
}

func (rb *rbtree) Del(x RBNoder) (ret bool) {
	n := rb.root
	for n != nil {
		if n.val.Equal(x) {
			ret = true
			break
		}
		if n.val.Less(x) {
			n = n.right
		} else {
			n = n.left
		}
	}
	if ret {
		rb.del(n)
	}
	return
}
