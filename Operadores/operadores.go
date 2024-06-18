package main

import "fmt"

func main() {

	soma := 1 + 2
	subtracao := 2 - 1
	multplicacao := 10 * 10
	divisao := 10 / 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, multplicacao, divisao, restoDaDivisao)

	var n1 int16 = 10
	var n2 int16 = 25
	somim := n1 + n2
	fmt.Println(somim)

	//Atribuição
	var variavel1 string = "String"
	variavel2 := "String 2"
	fmt.Println(variavel1, variavel2)

	//operadores relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	//operadores logicos
	fmt.Println("--------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	//operadores unirios
	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)

	numero *= 3
	fmt.Println(numero)

	//ternarios

	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}

	fmt.Println(texto)

}
