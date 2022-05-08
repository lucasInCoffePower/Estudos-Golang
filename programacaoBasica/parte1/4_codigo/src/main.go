/*

	Trabalhando com estrutura de dados

*/

package main

import "fmt" // // importando biblioteca para usar a função Println que permite iterar sobre as estruturas de dados

func main() {

	lista := [...]int{28, 26, 27} // array sem definir numero de elementos

	lista2 := [...]int{28, 34, 78}

	lista3 := []float32{32., 78, 93.}

	lista4 := make([]int, 5) // gerando array de 5 elementos do tipo int

	fatia := lista2[0:2] // Fazendo slice

	fatia = append(fatia, 4) // Fazendo fatia receber a lista fatia atualizada com o valor 4

	fmt.Println(len(lista2)) // Verificando o tamanho da lista através da função len
	fmt.Println(fatia)
	fmt.Println(lista3)
	fmt.Println(lista4)

	/*
		lista := [3]int{28, 26, 27} // array de tres elementos inicializado
		lista[0] = 28
		lista[1] = 26
		lista[2] = 27
	*/

	/*
		lista := [3]int{} // definicao do array
		lista[0] = 28
		lista[1] = 26
		lista[2] = 27
	*/
	fmt.Println(lista)
}
