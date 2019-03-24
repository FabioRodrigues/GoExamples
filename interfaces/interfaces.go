package main

import "fmt"

type veiculo interface {
	andar() int
	parar() int
}

type carro struct {
	limiteMaximo   int
	tempoParaParar int
}

type moto struct {
	limiteMaximo   int
	tempoParaParar int
}

func (c carro) andar() int {
	return c.limiteMaximo
}

func (c carro) parar() int {
	return c.tempoParaParar
}

func (m moto) andar() int {
	return m.limiteMaximo
}

func (m moto) parar() int {
	return m.tempoParaParar
}

func testarVeiculo(v veiculo) {
	fmt.Println("Velocidade maxima ao andar:", v.andar())
	fmt.Println("Tempo para parar:", v.parar())
}

func main() {
	testarVeiculo(carro{100, 15})
	testarVeiculo(moto{limiteMaximo: 50, tempoParaParar: 100})
}
