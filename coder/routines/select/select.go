package main

import (
	"fmt"
	"net/http"
)

// quando se coloca <-chan no parametro é apenas de leitura quando nao se coloca e de escrita/leitura
func Titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			for _, url := range urls {
				resp, _ := http.Get(url)
				c <- resp.Status
			}
			close(c)
		}(url)
	}
	return c
}

func maisRapido() string {
	c1 := Titulo("https://siga.cps.sp.gov.br/aluno/login.aspx")
	c2 := Titulo("https://youtube.com")
	c3 := Titulo("https://www.githubuniverse.com")

	// retorna o que for mais rapido
	select {
	case t1 := <-c1:
		return t1 + " siga"
	case t2 := <-c2:
		return t2 + " facebook"
	case t3 := <-c3:
		return t3 + " githubuniverse"
		/* 	default: fmt.Println("nenhum") */
	}
}

func main() {
	fmt.Println(maisRapido())

}
