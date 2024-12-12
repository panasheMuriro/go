package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define route handlers
	http.HandleFunc("/greet/", greetUser)
	http.HandleFunc("/search", handleQueryString)
	http.HandleFunc("/notfound", handleNotFound)
	http.HandleFunc("/error", handleInternalServerError)

	// Start the server on port 8080
	fmt.Println("Server started on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

// Handler to greet users based on URL path
func greetUser(w http.ResponseWriter, r *http.Request) {
	// Extract URL parameter
	parts := r.URL.Path[len("/greet/"):]
	if parts == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Hello, %s!\n", parts)
}

// Handler for search query parameters
func handleQueryString(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	name := queryParams.Get("name")
	age := queryParams.Get("age")

	if name == "" || age == "" {
		http.Error(w, "Both name and age are required", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Hello, %s! You are %s years old.\n", name, age)
}

// Handler for 404 Not Found
func handleNotFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

// Handler for Internal Server Error
func handleInternalServerError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}
