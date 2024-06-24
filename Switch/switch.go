package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sábado"
	default:
		return "Digite um numero válido!"
	}
}

func time(numero int) string {
	var time string

	switch {
	case numero == 1:
		time = "Flamengo"
	case numero == 2:
		time = "São paulo"
	case numero == 3:
		time = "Palmeiras"
	case numero == 4:
		time = "Cruzeiro"
	default:
		time = "Numero invalido"
	}
	return time
}

func main() {
	fmt.Println("switch")
	dia := diaDaSemana(5)
	fmt.Println(dia)
	fmt.Println("|---------------|")
	time1 := time(1)
	fmt.Println(time1)
}
