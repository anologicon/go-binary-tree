package interfaces

import (
	"io"
	"model"
)

type BinaryNode = model.BinaryNode

type IPrint interface {
	Print(w io.Writer, node *BinaryNode, ns int, ch rune)
}
