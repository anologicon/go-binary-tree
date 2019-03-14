package search

import (
	"fmt"
	"io"
	"model"
)

type Nivel struct {
	IPrint interface{}
}

// Print - Metodo para ordernar a busca por nivel
func (p Nivel) Print(w io.Writer, node *model.BinaryNode, ns int, ch rune) {
	
	var fila []*model.BinaryNode;

	fila = append(fila, node)
	
	var atual *model.BinaryNode;
	
	for len(fila) != 0 {

		// This is a pop slice
		atual, fila = fila[0], fila[1:]
		
		fmt.Println(atual.Data);
		
		if (atual.Left != nil) {
			fila = append(fila, atual.Left)
		}

		if (atual.Right != nil) {
			fila = append(fila, atual.Right)
		}
		
	}
}
