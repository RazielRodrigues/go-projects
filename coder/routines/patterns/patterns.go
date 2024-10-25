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

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

func juntar(c1, c2 <-chan string) <-chan string {
	junto := make(chan string)
	go encaminhar(c1, junto)
	go encaminhar(c2, junto)
	return junto
}

func main() {
	// c := Titulo("https://www.google.com.br", "https://www.facebook.com.br", "https://www.instagram.com.br")
	// c1 := Titulo("https://www.google.com", "https://www.facebook.com", "https://www.instagram.com")
	// fmt.Println("c " + <-c)
	// fmt.Println("c1 " + <-c1)

	juntasso := juntar(
		Titulo("https://www.google.com.br", "https://www.facebook.com.br", "https://www.instagram.com.br"),
		Titulo("https://www.google.com", "https://www.facebook.com", "https://www.instagram.com"),
	)

	fmt.Println("c " + <-juntasso)
	fmt.Println("c " + <-juntasso)
	fmt.Println("c " + <-juntasso)
	fmt.Println("c " + <-juntasso)

}
