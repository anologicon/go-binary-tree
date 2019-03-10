package search

import (
	"io"
	"fmt"
	"model"
)


type EmOrdem struct {
	IPrint interface{}
}

func (e EmOrdem) Print(w io.Writer, node *model.BinaryNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}

	e.Print(w, node.Right, ns+2, 'R')
	fmt.Fprintf(w, "%c:%v\n", ch, node.Data)
	e.Print(w, node.Left, ns+2, 'L')
}
