package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    float32
}
type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")
	p1 := pessoa{"Andres", "Aquino", 20, 1.70}
	fmt.Println(p1)

	e1 := estudante{p1, "TADS", "UERJ"}
	fmt.Println(e1)
}
