package main

import "fmt"

func main() {
	dict := make(map[string]string)

	fmt.Println("dicionario criado:", dict)

	fmt.Println()
	dict["first"] = "a"
	dict["second"] = "b"
	fmt.Println("Dicionario após adicionar dois elementos", dict)

	fmt.Println()
	delete(dict, "first")
	fmt.Println("dicionario após remover o primeiro elemento", dict)
}
