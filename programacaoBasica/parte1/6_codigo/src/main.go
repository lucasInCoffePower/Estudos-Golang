package main

func main() {
	var num int = 0
	var num2 int = 5
	var produto int = 7
	var divisao float32
	divisao = 24

	num++  // incrementando valor
	num2-- // decremento
	soma := 3 + 6
	subtracao := 4 - 8
	produto *= 5
	divisao /= 8
	modulo := 5 % 2

	println(
		"Soma: ", soma,
		"\nSubtracao: ", subtracao,
		"\nProduto: ", produto,
		"\nDivisao: ", divisao,
		"\nModulo: ", modulo,
	)
}
