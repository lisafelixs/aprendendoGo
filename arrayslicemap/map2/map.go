package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José João":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.0,
	}

	funcsESalarios["Rafael Filho"] = 1350.0
	delete(funcsESalarios, "Inexistente")

	//se colocar só 1 ele vai considerar a chave, para que considere só o valor pode usar o _
	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
