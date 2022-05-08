package main // importando biblioteca main para poder usar o metodo func

func main() {
	// Funcao principal

	var num int // Definição
	num = 100
	println(num)

	var num2 float32 = 25 // Definição e atribuição
	println(num2)

	palavra := "teste"
	println(palavra)

	myComplex := complex(2, 3)

	println(myComplex)

	println("Pegando os valores real e imaginario do numero complexo")
	println("Real: ", real(myComplex), "\nImaginario: ", imag(myComplex))
}
