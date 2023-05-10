package main

import (
	"fmt"
	"projetoTeste/calc"
)

func main() {
	var num1 int = 4
	num2 := 0

	result, err := calc.Dividir(num1, num2)

	if err != nil {
		fmt.Println("Erro: ", err)
		return
	}

	fmt.Println(result)
}