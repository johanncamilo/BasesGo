package main

import "fmt"

func mainZero() {

	const pi float64 = 3.14

	// los : indican que la varible no ha sido declarada anteriormente
	base := 332

	fmt.Println("base: ", base)

	// ZERO VALUES: valores x default
	var a int
	var b float64
	var c bool
	var d string

	fmt.Println(a, b, c, d)

}
