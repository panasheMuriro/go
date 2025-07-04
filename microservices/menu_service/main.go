package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/pizza/1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"name": "Cheese", "price" : 5.99}`)
	})

	fmt.Println("running on port 8001")
	http.ListenAndServe(":8001", nil)
}
