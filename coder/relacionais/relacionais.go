package main

import (
	"fmt"
	"time"
)

func main() {

	i := 20

	i += 8
	i -= 12
	i *= 54
	i /= 2
	i %= 1

	fmt.Println(i)

	x, y := 3, 32

	x, y = y, x

	fmt.Println(x, y)

	fmt.Println(3 != 3)
	fmt.Println(3 < 3)
	fmt.Println(3 > 3)
	fmt.Println(3 >= 3)
	fmt.Println(3 <= 3)

	dt1 := time.Unix(0, 0)
	dt2 := time.Unix(0, 0)

	fmt.Println(dt1 == dt2)

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"Raziel"}
	p2 := Pessoa{"Raziel"}

	fmt.Println(p1 == p2)
	i++
	fmt.Println("aa", i)

	fmt.Println(varios(true, true))

	// ponteiros :o

	num := 0

	var ponteiro *int = nil

	ponteiro = &num

	*ponteiro++
	*ponteiro++
	*ponteiro++

	fmt.Println(ponteiro, *ponteiro, num, &num)

	if *ponteiro == 3 {
		fmt.Println("arroz")
	} else if *ponteiro < 3 {
		fmt.Println("feijao")

	} else {
		fmt.Println("feijao")
	}

	if agora := true; agora != false { // funciona no siwtch
		fmt.Println("feijao e carne")
	}

	// nao tem while usa esse for para subistutiur
	cont := 0
	for cont <= 10 {
		fmt.Println(cont)
		cont++
	}

	for aaa := 0; aaa <= 10; aaa++ {
		fmt.Println(aaa)
	}

	/* 	for {
		fmt.Println("feijao e carne")
		time.Sleep(5)
	} */

	var nota int = 10
	switch nota {
	case 10:
		fmt.Println("QUE QUE QUE")
	case 3, 1:
		fmt.Println("aushuashuashau")
	case 22, 9:
		fallthrough
	default:
		fmt.Println("aaaaaaaaaa beeeeeeeeeee ceeeeeeeeee deeeeeeeeeeeee")
	}

	var nota2 int = 5
	// acessa a varuiavel de fora
	switch {
	case nota2 >= 12:
		fmt.Println("QUE QUE QUE")
	case nota2 >= 12:
		fmt.Println("aushuashuashau")
	case nota2 >= 12:
		fallthrough
	default:
		fmt.Println("aaaaaaaaaa beeeeeeeeeee ceeeeeeeeee deeeeeeeeeeeee")
	}

	fmt.Println(tipo("eu sou o ultrin"))
	fmt.Println(tipo(11111111111))

}

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	default:
		return "fooooooo"
	}
}

func varios(x, y bool) (bool, bool) {
	um := x && y
	dois := x || y

	return um, dois
}

// go nao tem operador ternario
