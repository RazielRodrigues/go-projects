package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	/*
		_ m "math"
		atalho e forma d eperder o importe

	*/)

func main() {
	const PI float64 = 5.6
	var raio = 3.2

	area := PI * math.Pow(raio, 2)
	fmt.Println("aaa", area)

	fmt.Print("Primeiro")
	fmt.Print("Programa")

	const (
		a = 1
		b = 2
	)

	var (
		r = 3
		c = 5
	)

	// ~multipla declaracao
	var e, f bool = true, false

	g, h, i := 2, "aaaa", false
	fmt.Print(a, b, r, c, e, f, g, h, i)

	/* 		concatena
	 */
	str := "true"

	parsed, _ := strconv.ParseBool(str)
	fmt.Println("ggggg", reflect.TypeOf(g))
	fmt.Println("ggggg", len(h))
	fmt.Println("ggggg" + h)
	fmt.Println("e o que", parsed)
}
