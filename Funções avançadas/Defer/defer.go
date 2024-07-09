package main

import "fmt"

func flamengo() {
	fmt.Println("Maior time do mundo")
}

func bills() {
	fmt.Println("The biggest team in nfl")
}

func campeao(p1, p2 float32) bool {
	defer fmt.Println("Pontos somados, o flamengo foi campeão ?")
	fmt.Println("Vamos verificar se o flamengo foi campeao ou nao da série a do brasileirão")

	pontos := (p1 + p2)

	if pontos >= 70 {
		return true
	}

	return false
}

func main() {
	flamengo()
	//defer == adiar
	bills()

	fmt.Println(campeao(35, 35))
}
