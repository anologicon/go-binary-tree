// Binary Tree in Golang
package main

import (
	"bufio"
	"fmt"
	"interfaces"
	"model"
	"model/search"
	"os"
	"strconv"
)

// Print - metodo que executa a exibição dos dados
func Print(t *model.BinaryTree, objectPrinter interfaces.IPrint) {
	objectPrinter.Print(os.Stdout, t.Root, 0, 'M')
}


func main() {
	tree := new(model.BinaryTree)

	snr := bufio.NewScanner(os.Stdin)
	enter := "Arvore Binaria em GOlang\n2 - Remover dados\n3 - Imprimir Largura\n6 - Imprimir Profundidade\n4 - Sair\n"

	tree.Insert(5);
	tree.Insert(3);
	tree.Insert(7);
	tree.Insert(4);
	tree.Insert(6);
	tree.Insert(17);
	tree.Insert(12);

	for fmt.Println(enter); snr.Scan(); fmt.Println(enter) {
		
		line := snr.Text()

		switch line {
		case "2":
			fmt.Println("\n Digite o numero que deseja remover\n")

			snr.Scan()

			numero := snr.Text()

			numeroInteiro, err := strconv.ParseInt(numero, 10, 64)

			if err != nil {
				fmt.Fprintln(os.Stderr, "reading standard input:", err)
			}

			tree.Remover(tree.Root, numeroInteiro);

			break
		case "3":
			fmt.Println("\n\nImprimindo Nivel:\n")

			Print(tree, search.Nivel{})

 			fmt.Println("\n\n\n")

			break
		case "6":
			fmt.Println("\n\nImprimindo Profundidade:\n")

			Print(tree, search.Profundidade{})

			fmt.Println("\n\n\n")

			break
		case "4":
			fmt.Println("Fim")
			os.Exit(0)
			break
		default:
			fmt.Println(enter)
		}

	}

	Print(tree, new(search.Profundidade))
}
