package main

import (
	"fmt"
	"log"
	"net/http"
)

// helloHandler handles requests to the "/hello" path.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the requested path is "/hello".
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// Check if the HTTP method is GET.
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	// Respond with "hello!" for successful GET requests.
	fmt.Fprintf(w, "hello!")
}

// formHandler handles POST requests to the "/form" path.
func formHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the requested path is "/form".
	if r.URL.Path != "/form" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// Check if the HTTP method is POST.
	if r.Method != "POST" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	// Parse the form data from the request.
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() error: %v", err)
		return
	}

	// Respond with a success message for POST requests.
	fmt.Fprint(w, "POST request successful\n")

	// Retrieve form values for "name" and "address".
	name := r.FormValue("name")
	address := r.FormValue("address")

	// Print the form values.
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func main() {
	// Serve static files from the "./static" directory.
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)

	// Register the helloHandler for the "/hello" path.
	http.HandleFunc("/hello", helloHandler)

	// Register the formHandler for the "/form" path.
	http.HandleFunc("/form", formHandler)

	// Start the server on port 8080.
	fmt.Printf("Starting server at http://localhost:8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
