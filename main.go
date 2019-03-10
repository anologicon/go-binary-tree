// Binary Tree in Golang
package main

import (
	"model"
	"model/search"
	"os"
	"interfaces"
)

// Print - metodo que executa a exibição dos dados
func Print(t *model.BinaryTree, objectPrinter interfaces.IPrint) {
	objectPrinter.Print(os.Stdout, t.Root, 0, 'M')
}

func main() {
	tree := new(model.BinaryTree)
	tree.Insert(5).
		Insert(3).
		Insert(4).
		Insert(7).
		Insert(6).
		Insert(17).
		Insert(12)

	Print(tree, new(search.Profundidade))
}
