package main

import "fmt"

// single return
func plus(a int, b int) int {
	return a + b
}

// multiple return
func vals() (int, int) {
	return 3, 7
}

// variadic functions like printLn
func sum(nums ...int) int {

	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}

	return total
}

//TODO: Anonymous functions, forming closures

func main() {
	s := plus(12, 13)
	fmt.Println("12 + 13 =", s)
	t, e := vals()
	fmt.Println(t, e)

	total := sum(12, 43, 65, 34, 11)
	fmt.Println(total)
}
