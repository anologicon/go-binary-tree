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