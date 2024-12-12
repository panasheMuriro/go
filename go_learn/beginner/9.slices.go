package main

import "fmt"

func main() {
	// var s = make([]string, 0)
	// var s = []string{}
	// s = append(s, "Panashe", "Muriro", "Paula")
	// fmt.Println(s[1:2])

	s := [][]int{}

	for i := 0; i < 10; i++ {
		inner := []int{}
		for j := 0; j < 10; j++ {
			inner = append(inner, j)
		}
		s = append(s, inner)
	}
	fmt.Println(s)
}
