package main

import "fmt"

func main() {
	arrays()
	ranges()
	pedacos()
	makerzao()
	mapeando()

}

func arrays() {
	var notas [3]float64
	notas[0], notas[1], notas[2] = 1.1, 4.5, 9.0

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64((len(notas)))

	fmt.Println("Media: ", media)
}

func ranges() {

	var numeros = [...]int{1, 2, 3, 4, 5, 7, 8, 9}

	// EQUIVALENTE FOR EACH
	for index, valor := range numeros {
		fmt.Println(index, valor)
	}
	for index, _ := range numeros {
		fmt.Println(index)
	}
	for sopegaoindex := range numeros {
		fmt.Println(sopegaoindex)
	}
}

func pedacos() {
	var numeros = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var fixos = [...]string{"a", "b", "c"}

	var pedacoes = []int{4, 5, 6}

	fmt.Println(numeros, fixos, pedacoes)

	pedacoFixo := numeros[5:9]
	pedacoDoInicio := numeros[:4]
	pedacoDoPedaco := pedacoFixo[1:2]

	fmt.Println(pedacoFixo, pedacoDoInicio, pedacoDoPedaco)
}

func makerzao() {
	pedaco := make([]int, 10) // o array criado é de 10 assim como o tamanho do slice

	pedacoDoPedaco := make([]int, 5, 30) // o slice criado tem o tamanho de cinco e o tamanho do array é 30

	pedacoMenor := make([]int, 0, 10)
	fmt.Println(pedacoMenor, len(pedacoMenor), cap(pedacoMenor))

	pedacoMenor = append(pedacoMenor, 10, 10, 10)

	fmt.Println(pedaco, pedacoDoPedaco, pedacoMenor)

	var copia = make([]int, 2)

	copy(copia, pedacoMenor)
	fmt.Println(copia)
}

func mapeando() {
	aprovados := make(map[int]string)

	aprovados[1] = "Raziel"
	aprovados[2] = "Jaiara"
	aprovados[3] = "Sergio"
	aprovados[4] = "Marcos"
	// mesma chave sobrescreve
	aprovados[4] = "Lisa"
	fmt.Println(aprovados)
	for chave, nome := range aprovados {
		fmt.Println(chave, nome)
	}

	delete(aprovados, 1)
	fmt.Println(aprovados)

	dados := map[string]string{
		"nome":    "ai sim",
		"entendi": "eu tambem",
	}

	fmt.Println(dados)

	muitosDados := map[int]map[string]int{
		0: {
			"OLA": 1000,
		},
		1: {
			"que coisa feia!": 2000,
		},
	}

	// da pra deletar e fazer um for dentro do for com essa estrutura
	fmt.Println(muitosDados)

}
