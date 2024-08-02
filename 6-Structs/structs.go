package main

import "fmt"

type usuario struct{
	nome string
	idade uint8
	endereco endereco
}

//Structs dentro de Structs
type endereco struct {
	logradouro string
	numero uint8

}

func main() {
	fmt.Println("Arquivos Structs")

	//1 forma de escrita
	var u usuario
	u.nome = "Davi"
	u.idade = 21
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos Bobos", 0}

	//2 forma de escrita
	usuario2 := usuario{"Davi", 21, enderecoExemplo}
	fmt.Println(usuario2)

	//3 forma - quando não se deseja utilizar todas as variavéis
	usuario3 := usuario{idade: 22}
	fmt.Println(usuario3)
}

/*
STRUCTS são coleções de campos, cada campo tem um nome e tipo

*/