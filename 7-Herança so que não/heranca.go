package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

// para utilizar a HERANÇA do Struct anterior, utilizamos o nome do outro struct, e não atribuimos um tipo
type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"João", "Pedro", 20, 178}
	fmt.Println(p1)

	// Instânciando o estudante
	e1 := estudante{p1, "Engenharia", "USP"}
	fmt.Println(e1.nome) //imprime só o JOAO
	fmt.Println(e1) // imprime {{João Pedro 20 178} Engenharia USP}
}
