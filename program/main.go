package main

import (
	"fmt"
	"main/calc"
)

func main() {
	var num1 int = 4
	num2 := 2

	result, err := calc.Dividir(num1, num2)

	if err != nil {
		fmt.Println("Erro: ", err)
		return
	}

	fmt.Println(result)
}