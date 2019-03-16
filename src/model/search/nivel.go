package search

import (
	"fmt"
	"io"
	"model"
)

// Nivel - Estrutura de busca em largura
type Nivel struct {
	IPrint interface{}
}

/** Print - Metodo de busca por nivel
 * Recebe alem de varios dados que são meio desnecessarios. Um nodo raiz
 * 
 * var fila [] é um slice do tipo *BinaryNode;
 *
 * na fila é inserido o nodo root
 * 
 * fila[0 => nodoRaiz]
 *
 * no loop, que só termina se a fila não tiver mais nodos
 *
 * atual recebe o nodo que estava no topo da fila
 *
 * é exibido o valor dentro do nodo
 *
 * Existe um nodo esquerdo da raiz ?
 *
 * - Insere na fila
 * 
 * - fila[0 => nodoEsquerdo]
 *
 * Existe um nodo direito da raiz ?
 *
 * - Insere na fila
 * 
 * - fila[0 => nodoEsquerdo, 1 => nodoDireito]
 *
 * Fila ainda tem nodos la
 *
 * É retirado o primeiro
 *
 * - fila[0 => nodoDireito] retirado nodoEsquerdo
 *
 * atual = nodoEsquerdo
 *
 * imprime atual Valor
 *
 * Essa é a ideia, vamos inserindo na fila e exibindo pela ordem inserida.
 * 
*/
func (p Nivel) Print(w io.Writer, node *model.BinaryNode, ns int, ch rune) {
	
	var fila []*model.BinaryNode;

	fila = append(fila, node)
	
	var atual *model.BinaryNode;

	for len(fila) != 0 {

		// Pop no slice
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
