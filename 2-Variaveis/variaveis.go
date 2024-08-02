package main

import "fmt"

func main() {
	//Inferência de tipo de variáveis, baseado em seu valor
	var variavel1 string = "variavel 1" //declaração explicita
	variavel2 :=  "Variavel 2" //declaração implicita

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var(
		variavel3 string = "valor1"
		variavel4 string = "valor2"
	)

	fmt.Println(variavel3,variavel4)

	//Utilizando inferência de tipos, LEMBRAR :=
	variavel5, variavel6 := "Valor3", "Valor4"

	fmt.Println(variavel5, variavel6)

	//A maioria das ações em GO só possuem uma forma, mas para declarar variáveis e constantes temos várias formas como as acima

	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	//inverter o valor de duas variáveis, em outras linguagens precisariamos de uma variável auxiliar, no GO não necessitamos, desta forma podemos trocar valores rapidamente assim:

	variavel5, variavel6 = variavel6, variavel5

	fmt.Println(variavel5, variavel6)
}
