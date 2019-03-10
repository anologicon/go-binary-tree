// Binary Tree in Golang
package main

import (
	"fmt"
	"io"
	"model"
	"os"
	"interfaces"
)

// BinaryNode is a model for node
type BinaryNode = model.BinaryNode

// BinaryNode is a model for node
type BinaryTree = model.BinaryTree

type iPrint = interfaces.IPrint

type Profundidade struct {
	iPrint
}

func (p Profundidade) Print(w io.Writer, node *BinaryNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.Data)
	p.Print(w, node.Right, ns+2, 'R')
	p.Print(w, node.Left, ns+2, 'L')
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

func print(t *BinaryTree, objectPrinter iPrint) {

	objectPrinter.Print(os.Stdout, t.Root, 0, 'M')
}

func main() {
	tree := new(BinaryTree)
	tree.Insert(5).
		Insert(3).
		Insert(4).
		Insert(7).
		Insert(6).
		Insert(17).
		Insert(12)

	print(tree, new(Profundidade))
}
