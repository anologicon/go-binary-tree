package model

type BinaryNode struct {
	Left  *BinaryNode
	Right *BinaryNode
	Data  int64
}

// Insert é um metodo de inserção de dados
func (n *BinaryNode) Insert(data int64) {
	if n == nil {
		return
	} else if data <= n.Data {
		if n.Left == nil {
			n.Left = &BinaryNode{Data: data, Left: nil, Right: nil}
		} else {
			n.Left.Insert(data)
		}
	} else {
		if n.Right == nil {
			n.Right = &BinaryNode{Data: data, Left: nil, Right: nil}
		} else {
			n.Right.Insert(data)
		}
	}
}