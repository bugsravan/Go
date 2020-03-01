package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func WalkTree(t *tree.Tree, c chan int) {
	Walk(t, c)
}

func Walk(t *tree.Tree, c chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, c)
	fmt.Println("Sending:", t.Value)
	c <- t.Value
	Walk(t.Right, c)
}

func Same(t1 *tree.Tree, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)
	go WalkTree(t1, c1)
	go WalkTree(t2, c2)
	for i := range c1 {
		if i != <-c2 {
			return false
		}
		fmt.Println(i)
	}
	return true
}

func main() {
	t1 := tree.New(3)
	t2 := tree.New(5)
	isSame := Same(t1, t2)
	fmt.Println(isSame)
}
