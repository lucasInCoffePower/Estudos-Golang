package main

import "fmt" // importando biblioteca para usar a função Println que permite iterar sobre as estruturas de dados

func main() {

	myMap := make(map[int]string) // criando dicionario com chaves int e valores strings

	fmt.Println(myMap) // imprimindo dicionario vazio

	myMap[0] = "ola"     // atribuindo a chave 0 o valor "ola"
	myMap[100] = "tchau" // atribuindo a chave 100 o valor "tchau"

	fmt.Println(myMap) // imprimindo tudo

}
