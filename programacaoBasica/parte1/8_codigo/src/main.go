/*
	Estudando estrutura de repetição
*/

package main

func main() {

	myMap := make(map[string]string)

	myMap["primeiro"] = "Aprendendo"
	myMap["segundo"] = "linguagem"
	myMap["terceiro"] = "Go"

	for k, v := range myMap {
		println(k, v)
	}

	/*

		array := [...]string{"lucas", "raquel", "silvana"}
		i := len(array)
		for j := 0; j < i; j++ {

			println(array[j])

		}

		for id, v := range array {

			println(id, v)
		}

		for idx, v := range array{

			println(idx, v)

		}
	*/

	/*

		j := 0

		for {
			j++
			println(j)
			if j >= 5 {
			break
			}
		}
		for i := 0; i < 5; i++ {

			println(i)

		}
	*/

}
