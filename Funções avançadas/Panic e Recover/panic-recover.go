package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("recuperando a execução")
	}
}

func alunosaAprovados(n1, n2 float32) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A média é exatamente 6!")

}

func main() {
	fmt.Println(alunosaAprovados(6, 6))
	fmt.Println("Aqui")
}
