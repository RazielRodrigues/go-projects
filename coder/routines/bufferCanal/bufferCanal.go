package main

import (
	"fmt"
	"time"
)

func rotina(c chan int, v int) {
	time.Sleep(time.Second)
	time.Sleep(time.Second)

	c <- v + 10
	c <- v + 1
	time.Sleep(time.Second)
	time.Sleep(time.Second)
	c <- v + 2
	time.Sleep(time.Second)
	time.Sleep(time.Second)
	c <- v + 3
	time.Sleep(time.Second)
	time.Sleep(time.Second)
	time.Sleep(time.Second)
	c <- v + 4

}

func roda(chances int, c chan string) {
	for i := 2; i < chances; i++ {
		time.Sleep(time.Second)
		if i%2 != 0 {
			c <- "Ganho!"
			continue
		}
		c <- "Perdeu!"
	}
	close(c)
}

func comBuffer() {
	// Quando tem o tamanho do buffer o deadlock nao acontece se voce chamar a quantidade  do buffer
	// voce pode chamar do que tem no buffer mas pode ociaosar um deadlock que é chamar um canal
	// quando ja nao tem mais dados a ser enviados
	c := make(chan int, 3)
	go rotina(c, 1)

	fmt.Println(<-c)
	fmt.Println("Fim!")
	fmt.Println(<-c)
	fmt.Println("Fim! 2")

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func main() {

	c := make(chan string, 10)
	go roda(cap(c), c)

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("Fim!")
}
