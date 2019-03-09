// Binary Tree in Golang
package main

import (
	"fmt"
	"io"
	"model"
	"os"
)

type BinaryNode = model.BinaryNode

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
	fmt.Fprintf(w, "%c:%v\n", ch, node.Data)
	p.print(w, node.Right, ns+2, 'R')
	p.print(w, node.Left, ns+2, 'L')
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

	e.print(w, node.Right, ns+2, 'R')
	fmt.Fprintf(w, "%c:%v\n", ch, node.Data)
	e.print(w, node.Left, ns+2, 'L')
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
		t.root = &BinaryNode{Data: data, Left: nil, Right: nil}
	} else {
		t.root.Insert(data)
	}
	return t
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
