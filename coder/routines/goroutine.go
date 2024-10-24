package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		fmt.Println(pessoa, texto, i)
		time.Sleep(time.Second)
	}
}

func main() {

	/* 	Adicionando a palavra go voce cria uma funcao que roda
	   	indepente ou seja de forma assincrona, entretanto ela
	   	depende do fluxo principal do programa, se o fluxo termnia
	   	o programa nao vai esperar a execucação completa da go function */

	fale("João", "Pq vc não fala comigo?", 1) // funcao sincrona
	go fale("Maria", "Aaah", 2)                 // funcao assincrona
	fale("João", "Pq vc não fala comigo?", 3) // funcao sincrona

	/* 	os canais sao tipos de uma variavel em go assim como o tipo inteiro ou string
	   	os canais recebem dados e criam ligações que voce consegue buscar o resultado
	   	de algo
	*/

	ch1 := make(chan string, 1) // canal de string

	ch1 <- "Ola"
	fmt.Println(<-ch1)

}
