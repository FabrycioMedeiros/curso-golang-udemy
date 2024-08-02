package main

import "fmt"

// Para declarar funções
//Palavra chave da função / nome da função / (parametros da função) Se tiver retorno, especificamos depois do fechamento do parametro {aplicamos os métodos que desejamos da função: return }
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

//Particularidade que encontramos em GO, as funções podem ter mais de um retorno Exemplo (int8, int8)
//Se parametros forem do mesmo tipo, só necesitamos nomear o ultimo parametro.
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao

}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	//Declarando variável do tipo função e jogar uma função dentro da mesma
	var f = func (txt string) string  {
		fmt.Println(txt + " Impressão")
		return txt
		
	}

	resultado := f("Texto da função 1")
	fmt.Println(resultado)

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	/* Quando nào queremos chamar um dos valores utilizamos _ no lugar do nome da variável

	resultadoSoma, _ := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma) //Resposta: 25 a segunda resposta será gerada, mas não apresentada
	*/

}

