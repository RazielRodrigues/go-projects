package main

import (
	"fmt"
	"net/http"
)

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

func main() {
	c := Titulo("https://www.google.com.br", "https://www.facebook.com.br", "https://www.instagram.com.br")
	c1 := Titulo("https://www.google.com", "https://www.facebook.com", "https://www.instagram.com")
	fmt.Println("c " + <-c)
	fmt.Println("c1 " + <-c1)

}
