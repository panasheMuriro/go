package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/order", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://localhost:8001/pizza/1")
		if err != nil {
			http.Error(w, "Error contacting menu service ", http.StatusInternalServerError)
			return
		}

		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		fmt.Fprintf(w, "Order placed for: %s", string(body))
	})

	fmt.Println("Order service running on 8002")
	http.ListenAndServe(":8002", nil)
}
