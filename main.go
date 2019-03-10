// Binary Tree in Golang
package main

import (
	"model"
	"model/search"
	"os"
	"interfaces"
)

// BinaryNode is a model for node
type BinaryNode = model.BinaryNode

// BinaryTree is a model for Tree
type BinaryTree = model.BinaryTree

// Profundidade is a model to aply the 'deep search'
type Profundidade = search.Profundidade 

type iPrint = interfaces.IPrint

func print(t *BinaryTree, objectPrinter iPrint) {

	objectPrinter.Print(os.Stdout, t.Root, 0, 'M')
}

func main() {
	tree := new(BinaryTree)
	tree.Insert(5).
		Insert(3).
		Insert(4).
		Insert(7).
		Insert(6).
		Insert(17).
		Insert(12)

	print(tree, new(Profundidade))
}
