package main

import "fmt"

type pessoa struct {
	name string
}

type superpessoa struct {
	pessoa
	power string
}

func main() {
	sp := superpessoa{power: "raziel"}
	sp.name = "Jaiara"
	fmt.Println(sp.name)
}
