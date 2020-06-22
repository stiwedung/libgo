package rbtree

import (
	"fmt"
	"testing"
)

type testDataStruct int

func (t testDataStruct) Less(x RBNoder) bool {
	ux := x.(testDataStruct)
	return t < ux
}

func (t testDataStruct) Equal(x RBNoder) bool {
	ux := x.(testDataStruct)
	return t == ux
}

func testLayerPrint(rb *rbtree) {
	if rb.root == nil {
		return
	}
	s := []*node{rb.root}
	for {
		var hasNode bool
		ts := []*node{}
		for _, n := range s {
			if n == nil {
				fmt.Printf("nil\t")
				ts = append(ts, nil, nil)
			} else {
				if n.red {
					fmt.Printf("\033[31m%d\033[0m\t", n.val)
				} else {
					fmt.Printf("%d\t", n.val)
				}
				ts = append(ts, n.left, n.right)
				if n.left != nil || n.right != nil {
					hasNode = true
				}
			}
		}
		fmt.Println()
		if !hasNode {
			break
		}
		s = ts
	}
	fmt.Println()
}

func testAdd(rb *rbtree, val int) {
	rb.Add(testDataStruct(val))
	testLayerPrint(rb)
}

func TestAdd(t *testing.T) {
	rb := New()
	testAdd(rb, 80)
	testAdd(rb, 120)
	testAdd(rb, 40)
	testAdd(rb, 140)
	testAdd(rb, 20)
	testAdd(rb, 60)
	testAdd(rb, 50)
	testAdd(rb, 70)
	testAdd(rb, 35)
}
