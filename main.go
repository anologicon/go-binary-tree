package main

import	"fmt"


type Tree struct {
	Left *Tree
	Value int
	Right *Tree
}

func Walk(t *Tree, ch chan int)  {
	if t == nil {
		return
	}

	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

func Walker(t *Tree) <- chan int {
	ch := make(chan int)
	go func(){
		Walk(t, ch)
		close(ch)
	}()

	return ch
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
		return t
	}
	t.Right = insert(t.Right, v)
	return t
}

func Reader(t1 *Tree) {
	c1 := Walker(t1)
	for {
		v1 := <-c1
		
		if v1 == 0 { 
			fmt.Println("Fim");
			
			break
		}

		fmt.Println("Valor", v1);
	}
}

func main() {

	fmt.Println("Comecou")

	var t *Tree

	t = insert(t, 1)
	t = insert(t, 2)
	t = insert(t, 3)
	t = insert(t, 4)
	t = insert(t, 5)

	Reader(t)
}