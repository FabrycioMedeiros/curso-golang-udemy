package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1

	variavel1++

	fmt.Println(variavel1, variavel2) //11 e 10

	//Ponteiro
	var variavel3 int //0
	var ponteiro *int //nil

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, ponteiro) //desreferenciação - desfazendo a referência, quando coloca o *, ele vai no endereço de memória e le o valor da variável

	variavel3 = 150
	fmt.Println(variavel3, *ponteiro)

	//ARRAYS INTERNOS
	fmt.Println("..........")
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)

}

//PONTEIRO é uma referência de memória, nào jogamos o valor da variaveel dentro dela, estamos jogando o endereço de memória onde está variavel está salva!
