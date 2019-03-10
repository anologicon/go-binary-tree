package search

import (
	"fmt"
	"io"
	"model"
)

type Nivel struct {
	IPrint interface{}
}


// Print - Metodo para ordernar a busca por profundidade
func (p Nivel) Print(w io.Writer, node *model.BinaryNode, ns int, ch rune) {
	var	fila []*model.BinaryNode
	
	fila = append(fila, node)
	
	for len(fila) != 0{
		atual := fila[0]

		fmt.Println(atual.Data, "\n")

		if (atual.Left != nil) {
			fila = append(fila, atual.Left)
		}

		if (atual.Right != nil) {
			fila = append(fila, atual.Right)
		}
	}
}
