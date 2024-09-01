package main

import "fmt"

func somar(a int, b int) int {
	return a + b
}

func imprimir(valor int) {
	fmt.Println(valor)
}

// pode usar retorno limpo pq ja declarou em cima na parte do retorno
// as variaveis do retorno tambem sao acessiveis dentro dele
// cada variavel pode ter um tipo diferente
// podemos tamvem fazer como a forma minimizada
func retornoNomeado(p1, p2 int) (segundo int, primeiro int, nome string) {
	segundo = p2
	primeiro = p1
	nome = "aaa"

	return // retorno limpo
}

var soma = func(a, b int) {
	fmt.Println("funcao dentro de variavel")
}

func vai(a string) {
	fmt.Println(a)
}

func agora(funcao func(string)) {
	funcao("TA VENDO, FOI!")
}

func comoassim(chamandobb func(string), teste string) {
	chamandobb("IDENTIFICOU QUE A FUNÇÃO QUE EU TINHA PASSADO NA TINHA RETORNO HUUUUUUUUUUUUUUM TA INTELIGENTE EM!")
	fmt.Println(teste)
}

func variariandoTipoBrutos(varios ...int) {
	for _, v := range varios {
		fmt.Println(v)
	}
}

func variandoUmSlicePoisEcomoDesctructing(pedaco ...string) {
	for _, v := range pedaco {
		fmt.Println(v)
	}
}

func main() {

	// basico funcao e metodos com e sem retorno
	resultado := somar(5, 4)
	imprimir(resultado)

	// multiplo retorno
	segundo, primeiro, _ := retornoNomeado(1, 2)
	fmt.Println(segundo, primeiro)

	// funcao dentro de variavel
	soma(1, 2)

	// funcao dentro de variavel em escopo local
	queijo := func() {
		fmt.Println("queijo!")
	}
	queijo()

	// passando funcao pra funcao
	agora(vai)
	comoassim(vai, "chama")

	// argumentos variaveis
	variariandoTipoBrutos(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	pedacao := []string{"Raziel", "Jaiara"}
	variandoUmSlicePoisEcomoDesctructing(pedacao...)

	// multiplica o resultado da multiplicaçºao atual com o resultado da proxima menos o numero atual
	fmt.Println(recursividadeFatorial(3))
	fmt.Println(recursividadeFatorial(4))
	fmt.Println(recursividadeFatorial(5))
	fmt.Println(recursividadeFatorial(6))

	atrasandoOutraFuncao()
}

func closure() func() {

	mostra := func() {
		fmt.Println("fechado!")
	}
	return mostra
}

func recursividadeFatorial(numero uint) uint {
	switch {
	case numero == 0:
		return 1
	default:
		resultado := numero * recursividadeFatorial(numero-1)
		return resultado
	}
}

func atrasandoOutraFuncao() {
	defer fmt.Println("Chamando atrasado")

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
}

func init() {
	fmt.Println("Aqui a gente consegue iniciar algo para o pacote sera sempre a primeira funcao buscada, esse voce pode ter varios no projeto mais de um ja a funcao main é uma por projeto")
}
