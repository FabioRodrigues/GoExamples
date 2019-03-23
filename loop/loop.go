package main

import "fmt"

func main() {
	for i := 1; i <= 2; i++ {
		fmt.Println("classico", i)
	}

	fmt.Println()
	i := 1
	for i <= 2 {
		fmt.Println("simples", i)
		i++
	}

	fmt.Println()
	var j int
	j = 1
	for {
		fmt.Println("infinito", j)
		if j == 2 {
			break
		}

		j++
	}

	fmt.Println()
	lista := [2]int{1, 2}
	for _, item := range lista {
		fmt.Println(item)
	}

}
