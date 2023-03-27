package main

import "fmt"

func main() {

	numeros := []float64{1.5, 2.3, 3.7, 5.2, 6.8, 9.1}
	fmt.Println(numeros)
	n := 6
	soma := 0.0

	for i := 0; i < n; i++ {
		soma += numeros[i]
	}

	media := (float64(soma)) / (float64(n))
	fmt.Println("A média é igual a,", media)

}
