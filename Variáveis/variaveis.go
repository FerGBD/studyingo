package main

import "fmt"

func main() {
	// go nao deixa que nao utilizemos variaveis inutilizaveis
	var pai string = "Flamengo"
	nome := "Fluminense"
	fmt.Println(pai)
	fmt.Println(nome)

	var (
		variavel  string = "botafogo"
		variavel1 string = " vasco"
	)
	fmt.Print(variavel, variavel1)

	variavel2, variavel3 := "Palmeiras", "Corinthians"
	fmt.Println(variavel2, variavel3)

	const constante1 string = "constante 1"
	fmt.Println(constante1)

	pai, nome = nome, pai
	fmt.Println(pai, nome)
}
