package main

import "fmt"

func main() {
	notas()
	pedacoes()
}

func notas() {
	var notas [3]float64
	notas[0], notas[1], notas[2] = 1.1, 4.5, 9.0

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64((len(notas)))

	fmt.Println("Media: ", media)
}

func pedacoes() {

	var numeros = [...]int{1, 2, 3, 4, 5}

	// EQUIVALENTE FOR EACH
	for index, valor := range numeros {
		println(index, valor)
	}
	for index, _ := range numeros {
		println(index)
	}
	for sopegaoindex := range numeros {
		println(sopegaoindex)
	}
}
