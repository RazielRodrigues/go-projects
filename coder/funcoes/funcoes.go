package main

import "fmt"

func somar(a int, b int) int {
	return a + b
}

func imprimir(valor int) {
	fmt.Println(valor)
}

func main() {
	resultado := somar(5, 4)
	imprimir(resultado)
}
