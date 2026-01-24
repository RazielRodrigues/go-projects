package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"unicode"
	"unicode/utf8"
	"unsafe"
)

type ChavePrimaria int

type Response struct {
	ID ChavePrimaria
}

func runas() {
	var palavra string = "Açai 🚀"
	var r rune = '🚀'
	fmt.Println(palavra)
	fmt.Println(r)
	fmt.Println(len(palavra), utf8.RuneCount([]byte(palavra)))
	palavraCortada := "voumudar"
	runaPalavraCortada := []rune(palavraCortada)
	runaPalavraCortada[2] = 'A'
	novaPalavra := string(runaPalavraCortada)
	fmt.Println(novaPalavra)

	frase := "TESTE 123 NOVAMENTE ?"
	var fraseRuna []rune

	for _, v := range frase {
		if unicode.IsDigit(v) {
			fraseRuna = append(fraseRuna, v)
		}
		fmt.Println(string(fraseRuna))
	}

}

func valores() {
	//var x int32
	//var f1 float32
	//c1 := complex(1, 34)
	var bVal bool // false
	var troca bool
	trueS := "True"
	troca, _ = strconv.ParseBool(trueS)
	literal := "Normal"
	raw := `OLA 
	RAZIEL`
	// string to byte
	xx := "xx"
	var yy []byte
	yy = []byte(xx)
	// rune to string
	var runinha []rune
	runinha = []rune(xx)
	fmt.Println(bVal)
	fmt.Println(troca)
	fmt.Println(literal, raw, yy, runinha)

	// meu proprio tipo
	var pk ChavePrimaria
	pk = 1
	var response Response
	response = Response{
		ID: pk,
	}
	fmt.Println(response)

	// ponteiros
	var pkPoint *ChavePrimaria = &pk
	var defers ChavePrimaria = *pkPoint
	fmt.Println(defers)

	// interface
	var a interface{} = 12
	var b int = a.(int)
	fmt.Println(b)

	// switch with type
	switch a.(type) {
	case int:
		fmt.Println(1)
	case string:
		fmt.Println(2)
	}

	// unsafe bypass
	var i int = 42
	var p *int = &i
	var fp *float64 = (*float64)(unsafe.Pointer(p)) // *int to *float64
	fmt.Println(fp)

	// type casting em struct por causa do mesmos campos
	type Point struct {
		x int
	}
	type Angle struct {
		x int
	}
	var XX Point = Point{1}
	var AA Angle = Angle(XX)
	fmt.Println(AA)

	//array para slice
	var as []int = []int{123}
	var bd []int = as[1:]
	fmt.Print(bd)

	// iinterface vira nil
	var x interface{} = nil
	var y error = nil
	fmt.Print(x, y)

	// tipo de funcao se a assinatura for a mesma
	type FuncaoDiferente func(int) int
	var funcs FuncaoDiferente = FuncaoDiferente(testeFunc)
	fmt.Printf("funcs: %v\n", funcs)

	// array to slice conversion
	var arr [5]int = [5]int{1, 23, 4, 5}
	var sl []int = arr[:]
	fmt.Println(sl)
}

func testeFunc(n int) int {
	return 1
}

func masterFunction() string {
	valores()
	runas()
	response := "ok"
	return response
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		p := masterFunction()
		json.NewEncoder(w).Encode(p)
	})
	fmt.Println("listening: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
