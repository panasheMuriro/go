package main

import "fmt"

func main() {
	for i := 0; i < 20; i++ {
		fmt.Println("Value is ", i)
	}

	for i := range 5 {
		fmt.Println("val is ", i)
	}

}
