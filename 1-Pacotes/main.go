package main

import (
	"fmt"
	"modulo/auxiliar"
	"github.com/badoux/checkmail"
)

//Quando se exporta uma função, é boa prática colocar um comentário acima dele, falando o que esta sendo feito
// Escrever registra uma mensagem na tela
func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("devbook@gmail.com")
	fmt.Println(erro)
}
