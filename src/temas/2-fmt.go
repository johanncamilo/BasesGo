/*
fmt administra inputs y outputs en go
*/

package main

import "fmt"

func mainFmt() {

	numLengs := 5
	words := "Lenguajes de Programación"
	fmt.Printf("Conozco más de %d %s \n", numLengs, words)
	fmt.Printf("Conozco más de %v %v \n", numLengs, words)

	// Sprintf no imprime sólo genera un string
	message := fmt.Sprintf("Conozco más de %d %s", numLengs, words)
	fmt.Println("Ahora con Sprintf ⬇️")
	fmt.Println(message)

	fmt.Println("Tipos de datos:")
	fmt.Printf("%T", numLengs)
	fmt.Println()
	fmt.Printf("%T", words)

}
