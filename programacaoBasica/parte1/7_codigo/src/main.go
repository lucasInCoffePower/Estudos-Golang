/*
	Estruturas condicionais
*/
package main

func main() {

	var num int
	num = 5

	/*
		if num:=5; num == 1 { // podemos criar a variavel e testar de imediato
		println("num é igual a 1")
		} else {

		println("num é diferente de 1")
		}

	*/
	if num == 1 {
		println("num é igual a 1")
	} else {

		println("num é diferente de 1")
	}

	switch {

	case num == 1:
		println("num é igual a 1")
	case num > 1:
		println("num é maior que 1")

	}

}
