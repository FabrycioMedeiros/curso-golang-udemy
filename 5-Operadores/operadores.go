package main

import "fmt"

func main() {
	//ARITMETICOS (+, -, /, *, %)
	soma := 10 + 5
	subtracao := 10 - 15
	divisao := 10 / 5
	multiplicacao := 10 * 5
	restoDeDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDeDivisao)

	/*
	Em Go não podemos realizar operações com variáveis de tipos diferentes, pois é extremamente tipado

	var numero1 int8 = 10
	var numero2 int16 = 25
	soma2 = numero1 + numero2
	fmt.Println(soma2)
	*/

	//A forma corretasempre deve seguir:
	/*
	var numero1 int16 = 10
	var numero2 int16 = 25
	var soma2 int16 = numero1 + numero2
	fmt.Println(soma2)

	//Com inferência
	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + numero2
	fmt.Println(soma2)
	
	// FIM DOS ARITMETICOS
	*/

	// ATRIBUIÇÃO - possui dois tipos(= , :=)
	//Primeiro utilizando a forma por extenso
	var variavel1 string = "String"

	//Segundo por inferência de tipos
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	//FIM DOS OPERADORES DE ATRIBUIÇÃO

	//OPERADORES RELACIONAIS - retornam TRUE ou FALSE
	fmt.Println(1 > 2) //MAIOR
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2) //COMPARAÇÃO
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2) //MENOR
	fmt.Println(1 != 2) //DIFERENTE 

	//FIM DOS RELACIONAIS

	//OPERADORES LÓGICOS
	fmt.Println(".............")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso) //Se todos forem verdadeiros, retornará verdadeiros
	fmt.Println(verdadeiro || falso) // Se UMA delas for verdadeira, vai retornar verdadeira
	fmt.Println(!falso) //Inverte o valor de uma variável Booleana

	//OPERADORES LÓGICOS

	//OPERADORES UNÁRIOS - para incrementar o valor de uma variável
	fmt.Println("....................")
	numero := 10
	numero++
	numero += 15
	fmt.Println(numero) //26

	numero1 := 3
	numero1 -= 12
	numero1--
	numero1 *= 2 //numero = numero * 2
	numero1 /= 2
	numero %= 2
	fmt.Println(numero1)

	//FIM DOS OPERADORES UNÁRIOS

	//OPERADOR TERNÁRIO - utilizado para adcionar um valor a variável, baseado em alguma condição

	//texto := numero > 5 ? "Maior que 5" : "Menor que 5" ISSO NÃO PODE SER FEITO NO GO, A FORMA CORRETA É USANDO O IF
fmt.Println("............")
	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}

	fmt.Println(texto)




}