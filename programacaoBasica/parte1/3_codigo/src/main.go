package main

const (
	primeiro = iota // iota é uma função geradora de indeces; ela gera um indice para a constante
	segundo  = iota
	terceiro = iota
)

/*
const (
	primeiro = "primeiro"
	segundo  = "segundo"
	terceiro = "terceiro"
)
*/

func main() {

	println(primeiro)
	println(segundo)
	println(terceiro)

}
