package model

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

// Remover - Remove um valor da arvore
func (t *BinaryTree) Remover(n *BinaryNode, data int64) *BinaryNode {
	if(n == nil) {
		return n;
	}

	if (data < n.Data) {	
		n.Left = t.Remover(n.Left, data);
	} else if (data > n.Data) {
		n.Right = t.Remover(n.Right, data);
	} else {
		if(n.Left == nil) {
			
			return n.Right;

		} else if (n.Right == nil) {
			return n.Left;
		}

		n.Data = menorValor(n.Right);

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