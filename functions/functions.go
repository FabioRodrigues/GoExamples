package main

import "fmt"

func main() {
	fmt.Println(somaValores(1, 2))
	fmt.Println()
	soma, ehPar := somaDoisValoresERetornaValorESeEhPar(2, 2)
	fmt.Println("A soma eh:", soma)
	if ehPar {
		fmt.Println("Eh par")
	} else {
		fmt.Println("Eh impar")
	}

}

func somaValores(a int, b int) int {
	return a + b
}

func somaDoisValoresERetornaValorESeEhPar(a int, b int) (int, bool) {
	result := somaValores(a, b)
	return result, (result%2 == 0)
}
