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

func hablar(letra string, c chan string) {

	c <- letra
	time.Sleep(time.Second)
	time.Sleep(time.Second)

	c <- letra + " 2"
	time.Sleep(time.Second)
	time.Sleep(time.Second)

	c <- letra + " 3"
	time.Sleep(time.Second)
	time.Sleep(time.Second)
	time.Sleep(time.Second)
	time.Sleep(time.Second)
	time.Sleep(time.Second)

	c <- letra + " 31111"
	time.Sleep(time.Second)
	time.Sleep(time.Second)
	time.Sleep(time.Second)
	time.Sleep(time.Second)
	time.Sleep(time.Second)
	time.Sleep(time.Second)
	time.Sleep(time.Second)
	time.Sleep(time.Second)
	time.Sleep(time.Second)
	time.Sleep(time.Second)

}

func falando() {

	fale("João", "Pq vc não fala comigo?", 1) // funcao sincrona
	go fale("Maria", "Aaah", 2)                 // funcao assincrona
	fale("João", "Pq vc não fala comigo?", 3) // funcao sincrona

	ch1 := make(chan string, 1) // canal de string

	ch1 <- "Ola"
	fmt.Println(<-ch1)
}

func main() {

	c := make(chan string)

	go hablar("FALA AI JAO", c)
	go fale("Maria", "Aaah", 30) // funcao assincrona

	a, b := <-c, <-c
	fmt.Println(a, b)

	fmt.Println("Fim!")
	fmt.Println(<-c)

	fmt.Println("QUASE LA e o Fim!")

	fmt.Println(<-c)

	fmt.Println("agora e o Fim!")

}
