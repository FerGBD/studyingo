package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for i < 10 {
		i++
		fmt.Println("Intermentando i")
		time.Sleep(time.Second)
	}
	fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando j", j)
		time.Sleep(time.Second)
	}

	nomes := [3]string{"Lopes", "Joao", "Yuri"}

	for i, nome := range nomes {
		fmt.Println(i, nome)
	}

	for i, letra := range "ORIGAMI" {
		fmt.Println(i, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Davi",
		"sobrenome": "augusto",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

}
