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
	enter := "Arvore Binaria em GOlang\n1 - Inserir dados\n2 - Remover dados\n3 - Imprimir\n4 - Sair\n"

	for fmt.Println(enter); snr.Scan(); fmt.Println(enter) {

		line := snr.Text()

		switch line {
		case "1":
			fmt.Println("\n Digite o numero que deseja inserir\n")

			snr.Scan()

			numero := snr.Text()

			numeroInteiro, err := strconv.ParseInt(numero, 10, 64)

			if err != nil {
				fmt.Fprintln(os.Stderr, "reading standard input:", err)
			}

			tree.Insert(numeroInteiro)
			break
		case "3":
			fmt.Println("\n\nImprimindo:\n")

			Print(tree, search.Nivel{})

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
