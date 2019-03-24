package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func (p pessoa) getNome() string {
	return p.nome
}

func (p *pessoa) getIdade(adicional int) int {
	return p.idade + adicional
}

func main() {
	p1 := pessoa{"fabio", 10}
	fmt.Println(p1)

	p2 := &p1

	p2.nome = "Jose"

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println("Nome via metodo:", p2.getNome())
	fmt.Println("idade via metodo:", p2.getIdade(2))
}
