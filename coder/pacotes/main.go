package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(Soma(2, 3))

	fmt.Println(runtime.NumCPU())
}
