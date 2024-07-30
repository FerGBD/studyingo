package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Println("Salvando os dados do Usuário %s no banco de dados", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Gabigol", 20}
	usuario1.salvar()

	usuario2 := usuario{"Rodinei", 29}
	usuario2.salvar()

	maiorDeIdade := usuario2.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}
