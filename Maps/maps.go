package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Andres",
		"sobrenome": "Aquino",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"segundo":  "pedro",
		},
		"curso": {
			"nome":   "TADS",
			"campus": "UERJ-ZO",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["Time"] = map[string]string{
		"nome": "Chelsea",
	}
	fmt.Println(usuario2)
}
