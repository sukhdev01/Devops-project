package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {

	message := "Welcome to Go..."
	fmt.Fprint(w, message)
}

func main() {

	http.HandleFunc("/", homeHandler)

	fmt.Println("Starting server on :3456...")
	err := http.ListenAndServe(":3456", nil)

	if err != nil {
		fmt.Println("Failed to start server:", err)
	}

}
