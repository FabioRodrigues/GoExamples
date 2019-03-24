package main

import "fmt"

func main() {
	var i int = 0
	fmt.Println("valor inicial de i:", i)
	fmt.Println()

	val(i)
	fmt.Println("valor do i na passagem por valor:", i)
	fmt.Println()

	ref(&i)
	fmt.Println("valor do i na passagem por referência (ponteiro):", i)

}

func val(i int) {
	i = 10
}

func ref(i *int) {
	*i = 10
}
