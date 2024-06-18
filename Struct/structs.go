package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro  string
	numero      uint8
	complemento string
}

func main() {
	fmt.Println("Arquivo Structs")

	var u usuario
	u.nome = "Fernando"
	u.idade = 19
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua das flores", 137, "Proximo ao instagrill"}

	usuario2 := usuario{"davi", 23, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "ronaldo"}
	fmt.Println(usuario3)
}
