package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 //tipo (float64) inferido pelo compilador

	//Sempre que declara uma variavel vc tem que usar ela
	//As variaveis são fortemente tipadas, se ela nasce com um tipo vai morrer com o mesmo
	//forma reduzida de criar uma var
	area := PI * math.Pow(raio, 2)
	fmt.Println("A área da circunferencia é", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false

	fmt.Println(e, f)

	g, h, i := 2, false, "epa!"

	fmt.Println(g, h, i)
}
