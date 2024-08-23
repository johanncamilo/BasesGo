/*
strconv.Atoi == ParseInt
nil: forma que tiene go de determinar que una fn no tiene error
*/
package main

import (
	"fmt"
	"log"
	"strconv"
)

func parimpar(x int) (y bool) {
	if x%2 == 0 {
		y = true
	}
	if x%2 != 0 {
		y = false
	}

	return
}

func mainAtoi() {
	x := "69"
	y, e := strconv.Atoi(x)

	if e != nil {
		log.Fatal(e)
	}
	fmt.Printf("%T \n%d", y, y)

	fmt.Println()
	fmt.Println(parimpar(y))

}
