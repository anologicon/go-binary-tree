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
	
	fila := make([]*model.BinaryNode, 0)

	fila = append(fila, node)
	
	var atual *model.BinaryNode;
	
	for len(fila) != 0 {

		atual, fila = fila[len(fila)-1], fila[:len(fila)-1]
		
		fmt.Println(atual.Data)
		
		if (atual.Right != nil) {
			fila = append(fila, atual.Right)
		}

		if (atual.Left != nil) {
			fila = append(fila, atual.Left)
		}
	}
}
