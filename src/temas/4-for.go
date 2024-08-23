package main

import (
	"fmt"
)

func mainFor() {
	count := 22
	for i := 0; i < count; i++ {
		fmt.Println(i)
	}

	count = 0
	fmt.Println("ahora con For while")
	for count < 17 {
		fmt.Println(count)
		count++
	}

	wakandaForever := 0
	for {
		fmt.Println(wakandaForever)
		wakandaForever++
	}
}
