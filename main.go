package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Message struct to handle JSON responses
type Message struct {
	Text string `json:"message"`
}

// Handler for the root route
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "🚀 Welcome to Go Server!")
}

// Handler for /api/hello endpoint
func apiHelloHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")
	
	// Create response message
	response := Message{Text: "Hello World from Go API!"}
	
	// Convert to JSON and send response
	json.NewEncoder(w).Encode(response)
}

// Handler for /api/greet endpoint with query parameter
func greetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Get name from query parameter, default to "Developer"
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Developer"
	}
	
	response := Message{Text: fmt.Sprintf("Hello, %s! Welcome to Go!", name)}
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Define routes
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/api/hello", apiHelloHandler)
	http.HandleFunc("/api/greet", greetHandler)
	
	// Server configuration
	port := ":8080"
	fmt.Printf("🌟 Server starting on http://localhost%s\n", port)
	fmt.Println("Available endpoints:")
	fmt.Println("  GET / - Welcome message")
	fmt.Println("  GET /api/hello - JSON Hello World")
	fmt.Println("  GET /api/greet?name=YourName - Personalized greeting")
	
	// Start server
	log.Fatal(http.ListenAndServe(port, nil))
}