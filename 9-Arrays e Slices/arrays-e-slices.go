package main

import (
	"fmt"
	"reflect"
)

func main() {

	fmt.Println("Arrays e Slices")

	//Em GO temos duas formas de criar Arrays
	//1 Forma - VARIAVEL/NOME/[TAMANHO DO ARRAY]/TIPO
	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	//2 forma - Utilizando inferência de tipos :=
	array2 := [5]string{"valor0", "valor1", "valor2", "valor3", "valor4"}
	fmt.Println(array2)

	//Forma de utilizar array de forma flexivel [...] esta forma vai fixar o array, de acordo com o numero de itens que passarmos dento do {valores}, mas não deixa com valor DINAMICO, será relacionado ao numero de valores passados.
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)
	// EX array3[7] = 10 , ele não irá aceitar, pois não é dinamico, e dessa forma pula um valor do campo!

	//SLICE(fatia) - não tem tamanho fixo, o tamanho é de acordo com a necessidade, seu tamanho É DINAMICO
	//possui limitação de tipo
	//IMPORTANTE, o slice não é um array, só parece, é como se fosse um ponteiro que aponta para um array!!!
	slice := []int{10, 20, 30, 40, 50, 60, 70, 80}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 18)
	fmt.Println(slice)

	//REFERENCIAR UM ARRAY - seria uma fatia, pedaço do array2, nisso ele irá pegar os valores 1 e 2 do array, deixando o 3 valor de fora
	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posição alterada"
	fmt.Println(slice2)



}

//Arrays são listas de valores(strings, inteiros, etc) temos uma variável que guardará uma serie de valores
//Lembrar que o GO não possibilita deixar criar algo de qualquer tipo(Extremamente tipado) e todos os itens tem que ser obrigatóriamente do mesmo tipo, boa prática - especificar o tamanho do ARRAY, quando não o fazemos, ele se torna um SLICE
//O array tem tamanho fixo e os tipos de dados também fixos 
// FUNÇÃO TypeOf - trás o tipo de uma variável
// FUNÇÃO Append - ADD um item no SLICE e retornar um slice novo com este item adicionado