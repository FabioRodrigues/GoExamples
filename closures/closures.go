package main

import "fmt"

func main() {
	function := IntSeq()

	fmt.Println(function())
	fmt.Println(function())
	fmt.Println(function())
	fmt.Println(function())
	fmt.Println(function())
}

func IntSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
