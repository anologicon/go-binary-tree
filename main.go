// Binary Tree in Golang
package main

import (
	"fmt"
	"io"
	"os"
)

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int64
}

type Profundidade struct {
	iPrint
}

func (p Profundidade) print(w io.Writer, node *BinaryNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	p.print(w, node.right, ns+2, 'R')
	p.print(w, node.left, ns+2, 'L')
}

type EmOrdem struct {
	iPrint
}

func (e EmOrdem) print(w io.Writer, node *BinaryNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}

	e.print(w, node.right, ns+2, 'R')
	fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	e.print(w, node.left, ns+2, 'L')
}

type BinaryTree struct {
	root *BinaryNode
	iPrint
}

type iPrint interface {
	print(w io.Writer, node *BinaryNode, ns int, ch rune)
}

func (t *BinaryTree) insert(data int64) *BinaryTree {
	if t.root == nil {
		t.root = &BinaryNode{data: data, left: nil, right: nil}
	} else {
		t.root.insert(data)
	}
	return t
}

func (n *BinaryNode) insert(data int64) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.right.insert(data)
		}
	}
}

func print(t *BinaryTree, objectPrinter iPrint) {

	objectPrinter.print(os.Stdout, t.root, 0, 'M')
}

func main() {
	tree := new(BinaryTree)
	tree.insert(5).
		insert(3).
		insert(4).
		insert(7).
		insert(6).
		insert(17).
		insert(12)

	print(tree, new(Profundidade))
}
