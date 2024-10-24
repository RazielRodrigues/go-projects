package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

type fornecedor struct {
	nome string
	pais string
}

func (p pessoa) toString() string {
	return fmt.Sprintf("%v %v", p.nome, p.sobrenome)
}

func (p produto) toString() string {
	return fmt.Sprintf("%v - R$ %.2f", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

/* parte dois
 */

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo      string
	turboLigado bool
	velocidade  int
}

func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

/* parte 3 */
type luxooso interface {
	ligarTurbo()
}

type esportivoBom interface {
	baliza()
}

type esportivoLuxo interface {
	esportivoBom
	luxooso

	/* 	mais metodos
	 */
}

type BM2 struct{}

func (b BM2) ligarTurbo() {
	fmt.Println("Turbo ligado")
}

func (b BM2) baliza() {
	fmt.Println("Baliza ligada")
}

func main() {
	var coisa imprimivel = pessoa{"Raziel", "Rodrigues"}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	coisa = produto{"Notebook", 1899.90}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	carro1 := ferrari{"VIPER", false, 0}
	carro1.ligarTurbo()

	var carro2 esportivo = &ferrari{"VIPER 2", false, 0}
	carro2.ligarTurbo()

	fmt.Println(carro1, carro2)

	var BM2 esportivoLuxo = BM2{}
	BM2.ligarTurbo()
	BM2.baliza()

}
