package main

import "fmt"

func main() {
	lista1 := [2]int{1, 2}
	for _, item := range lista1 {
		fmt.Println(item)
	}

	fmt.Println()
	slice1 := []string{"a", "b"}
	for i := 0; i < len(slice1); i++ {
		fmt.Println(slice1[i])
	}

	fmt.Println()
	slice2 := make([]int, 2)
	fmt.Println(len(slice2))

	fmt.Println()
	var slice3 []int
	fmt.Println(len(slice3))

	for i := 0; i < 10; i++ {
		slice3 = append(slice3, i)
	}
	fmt.Println("slice3:", slice3)
}
