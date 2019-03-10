package search

import (
	"fmt"
	"io"
	"model"
)

// BinaryNode is a model for node
type BinaryNode = model.BinaryNode

// Profundidade modelo para ordenação
type Profundidade struct {
	IPrint interface{}
}

// Print - Metodo para ordernar a busca por profundidade
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
