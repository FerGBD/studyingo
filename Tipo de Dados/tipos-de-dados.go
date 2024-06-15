package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int = 200
	fmt.Println(numero)

	teste := 100000000
	fmt.Println(teste)

	var numero2 uint32 = 1000
	fmt.Println(numero2)

	//alias
	//int 32 = rune
	var numero3 rune = 123456
	fmt.Println(numero3)

	//byte = uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	var nume float32 = 21.3
	fmt.Println(nume)

	//string
	var time string = "Flamengo"
	fmt.Println(time)

	//possivel char
	char := 'S'
	fmt.Println(char)

	//boolean
	var boolean1 bool = true
	fmt.Println(boolean1)

	var erro error = errors.New("Erros interno")
	fmt.Println(erro)

}
