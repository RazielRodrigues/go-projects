package main

import "fmt"

func main() {
	notas()
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
