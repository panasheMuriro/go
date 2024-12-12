package main

import (
	"fmt"
	"math/rand"
)

func main() {

	some_number := rand.Intn(10)

	switch {
	case some_number < 5:
		fmt.Println("number less than 5", some_number)
	case some_number >= 5:
		fmt.Println("number at least 5", some_number)
	}

	whichAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Ths is a bool", t)
		}
	}

	whichAmI(true)

}
