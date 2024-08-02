package main

import (
	"errors"
	"fmt"
)

func main() {
	//int8, int16, int32, int64

	var numero int8 = 12
	var numero1 int16 = 12000
	var numero2 int32 = 1200000000
	var numero3 int64 = 1200000000000000000

	//Ao utilizar i int sem especificação, ele usa a arquitetura do seu pc como base

	var numero4 int = 1111111111111111111 //base 64
	// usando inferencia de tipos 
	numero5 := -2222222222222222222

	fmt.Println(numero, numero1, numero2, numero3, numero4, numero5)

	// Dentro do GO temos o uint = unsygned interger que é um int sem sinal

	var numero6 uint = 1111111111111111111 //base 64
	// não se pode colocar sinal,  dará erro pois não permite sinais

	fmt.Println(numero6)

	//alias = apelido
	//INT32 = RUNE
	var numero7 rune = 1111111111 //base 32

	fmt.Println(numero7)

	//UINT8 = BYTE
	var numero8 byte = 111 //base 8

	fmt.Println(numero8)

	//FIM NUMEROS INTEIROS

	//Numeros reais float32, float64 a diferença e que suportam ponto flutuane(" . ")Sempre usa ponto!!!
	var numeroReal7 float32 = 111111111111111111111111111111111111111.11 //base 32

	var numeroReal8 float64 = 1111111111111111111111111.11 //base 64

	fmt.Println(numeroReal7, numeroReal8)

	//Inferência de tipos e le usara sua arquitetura como parametro
	numeroReal9 := 1234.45
	fmt.Println(numeroReal9)

	//FIM NUMEROS REAIS

	//STRINGS
	// Em GO não temos CHAR, sempre utilizar aspas duplas

	var str string = "Texto"
	fmt.Println(str)

	//Inferência de tipos
	str2 := "Texto2"
	fmt.Println(str2)

	//exemplo CHAR, ele irá buscar o numero correspondente a tabela ASCII, e se colocar mais de um dará erro

	char := 'B'
	fmt.Println(char) //R. 66

	//FIM STRING

	//Conceito de VALOR ZERO - Valor que será atribuido a uma variável quando não se inicializa a mesma com valor
	// No GO todo tipo de dado tem seu valor inicial, STRING é vazio e INT, FLOAT será 0, para ERRO é <NIL> ou nulo.

	var texto1 string
	fmt.Println(texto1) //Resposta sairá em branco, sem nada apresentável

	var texto2 int16
	fmt.Println(texto1, texto2) //Resposta sairá 0

	texto := 5
	fmt.Println(texto) //vai ser lido como inteiro

	//BOOLEANO

	var booleano1 bool = true
	var booleano2 bool = false // O valor zero é false
	var booleano3 bool // O valor zero é false
	fmt.Println(booleano1, booleano2, booleano3)

	//ERRO <nil> ou nulo

	//Para usarmos um valor, temos que utilizar um PACOTE NATIVO chamado "errors"
	// variável/nome da variaável/ tipo da variável
	var erro error
	fmt.Println(erro) //<nil>

	//ERRORS = nome do pacote nativo
	var erro1 error = errors.New("Erro interno criado")
	fmt.Println(erro1) //<nil>

}