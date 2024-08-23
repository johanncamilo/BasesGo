package main

import (
	"fmt"
	"math"
)

// múltiples parametros, múltiples return
func complexFn(a, b int) (c, d int) {
	return a * b, a + b
}

func calculosCirculo(r int) (area, perimetro float64) {

	const pi = math.Pi
	radio := float64(r)

	area = pi * math.Pow(radio, 2)
	perimetro = 2 * pi * radio

	return area, perimetro

}

func mainFunciones() {
	value1, _ := complexFn(45, 58)

	fmt.Println(value1)

	radio := 74
	area, perimetro := calculosCirculo(radio)

	fmt.Printf("el área de círculo de radio %d es: %f \n", radio, area)
	fmt.Printf("el perímetro del círculo de radio %d es: %f \n", radio, perimetro)
}
