package main

import "fmt"

// conseguimos definir um tipo
type dinheiro float64

type pessoa struct {
	name string
	dinheiro
}

type superpessoa struct {
	pessoa
	power string
}

func main() {
	sp := superpessoa{power: "raziel"}
	sp.name = "Jaiara"
	sp.dinheiro = 1000.0

	fmt.Println(sp.name, sp.dinheiro)
}
