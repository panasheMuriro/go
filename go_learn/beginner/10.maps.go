package main

import "fmt"

func main() {
	snack_count := make(map[string]int)
	snack_count["lays"] = 3
	snack_count["oreo"] = 10
	delete(snack_count, "oreo")
	fmt.Println((snack_count))
}
