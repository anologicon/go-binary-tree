package model

import (
	"fmt"
)


// BinaryTree - Estrutura da arvore
type BinaryTree struct {
	Root *BinaryNode

	IPrint interface{}
}

// Insert - Insere os nodos na arvore
func (t *BinaryTree) Insert(data int64) *BinaryTree {
	if t.Root == nil {
		t.Root = &BinaryNode{Data: data, Left: nil, Right: nil}
	} else {
		t.Root.Insert(data)
	}
	return t
}

/** Remover - Remove um valor da arvore
  * Metodo recebe um nodo, e um valor a ser removido da arvore
  *
  * O valor procurado, é menor que a raiz ?
  *
  * - Recursivamente chama o nodo a esquerda da raiz
  * - Se não, recursivamente a direita da raiz
  * 
  * nodo.Data = 5; data 7
  *
  * 7 < 5 false
  * 
  * 7 > 5 true; O nodo direito da raiz ira ser setado recursivamente
  *
  * Quando o 7 encontra o 7 o nodo que tem o valor
  *
  * Procura o menor valor da arvore
  *
  * 12
  *
  * e seta o valor do nodo como o menor
  *
  * E recursivamente ira organizando os outros nodos
  *
  * 
*/
func (t *BinaryTree) Remover(n *BinaryNode, data int64) *BinaryNode {
	if(n == nil) {
		return n;
	}

	if (data < n.Data) {	
		
		fmt.Println(data, "<", n.Data);

		fmt.Println("Agora esquerdo de", n.Data, "sera realocado");

		n.Left = t.Remover(n.Left, data);
	} else if (data > n.Data) {

		fmt.Println(data, ">", n.Data);

		fmt.Println("Agora direito de", n.Data, "sera realocado")
		n.Right = t.Remover(n.Right, data);
	} else {
		fmt.Println(data, "==", n.Data);

		if(n.Left == nil) {
			fmt.Println(n.Data, "removido");
			return n.Right;

		} else if (n.Right == nil) {

			fmt.Println(n.Data, "removido");
			return n.Left;
		}

		fmt.Println("Nodo: ", n.Data);
		n.Data = menorValor(n.Right);
		fmt.Println("Nodo agora  é : ", n.Data);
		n.Right = t.Remover(n.Right, n.Data);
	}

	return n;
}

func menorValor(n *BinaryNode) int64  {
	var minimo int64;

	minimo = n.Data;

	for n.Left != nil {
		minimo = n.Left.Data;
		n = n.Left;
	}

	return minimo;
}