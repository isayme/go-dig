package main

import (
	"fmt"
	"github.com/isayme/go-gdig"
)

type A struct {
}

func NewA() *A {
	fmt.Println("NewA called")
	return &A{}
}

type B struct {
	A *A
}

func NewB(a *A) *B {
	fmt.Println("NewB called")

	return &B{
		A: a,
	}
}

func main() {
	gdig.Provide(NewB)
	gdig.Provide(NewA)

	gdig.Invoke(func(b *B) {
		fmt.Println("gdig.String")
		fmt.Println(gdig.String())
	})
}
