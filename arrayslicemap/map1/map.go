package main

import "fmt"

func main() {
	//var aprovados map[int]string
	//maps devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[12345678978] = "Maria"
	aprovados[78945612312] = "Pedro"
	aprovados[32165498745] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[32165498745])
	delete(aprovados, 32165498745)
	fmt.Println(aprovados[32165498745])
}
