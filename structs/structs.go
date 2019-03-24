package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func main() {
	p1 := pessoa{"fabio", 10}
	fmt.Println(p1)

	p2 := &p1

	p2.nome = "Jose"

	fmt.Println(p1)
	fmt.Println(p2)
}
