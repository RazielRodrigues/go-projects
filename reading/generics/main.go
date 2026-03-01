package main

import (
	"fmt"
)

type Numbers interface { // constraint interface a equipe do go fez um pacote com diversas constraints
	float64 | int64 // basicamente faz um union types
}

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func SumNumbers[K comparable, V Numbers](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

type StructGenerica[T int64 | string | any] struct {
	items []T
	nome  T
}

// T is type parameter here, with any constraint
type MyStruct[T any] struct {
	inner T
}

// No new type parameter is allowed in struct methods
func (m *MyStruct[T]) Get() T {
	return m.inner
}
func (m *MyStruct[T]) Set(v T) {
	m.inner = v
}

type Caixa[T any] struct{ Conteudo T }

func Processar[T any](lista []T) { // obriga que todos os dados sejam do tipo T

}

func ProcessarValorGenerico[T any](valorgenerico T) T { // obriga que todos os dados sejam do tipo T
	var zero T
	return zero // retornar o valor zero mesmo que o tipo seja generic
}

func main() {

	caixa := Caixa[int]{ // para fazer instancia de struct generica
		Conteudo: 1,
	}
	fmt.Print(caixa)

	ints := map[string]int64{
		"first":  12,
		"second": 15,
	}

	floats := map[string]float64{
		"first":  12,
		"second": 15,
	}

	// var interfaceGenericaFazParteDosGenerics any

	fmt.Printf("valor sem generic %v , floats %v. Generics %v e %v e com numbers %v",
		SumInts(ints),
		SumFloats(floats),
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats),
		SumNumbers(floats), // ele entende pois faz a inferencia dos dados type inference
		//SumIntsOrFloats[string, int64](ints),
		//SumIntsOrFloats[string, float64](floats))
	)
}
