package main

import "fmt"

func print(message string) {
	for _, item := range []int{1, 2, 3} {
		fmt.Println(message, ":", item)
	}
}

func main() {
	go print("async 1")
	print("sincrono")
	go print("async 2")

	fmt.Scanln()
	fmt.Println("pronto")
}
