package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {

	message := `Welcome to Go...
		available apis:
			/api/v1/hello  
			/api/v1/thanks  
			/api/v1/bye`
	fmt.Fprint(w, message)
}

func main() {

	http.HandleFunc("/", homeHandler)

	http.HandleFunc("/api/v1/hello", HelloHandler)
	http.HandleFunc("/api/v1/thanks", ThanksHandler)
	http.HandleFunc("/api/v1/bye", ByeHandler)

	fmt.Println("Starting server on :3456...")
	err := http.ListenAndServe(":3456", nil)

	if err != nil {
		fmt.Println("Failed to start server:", err)
	}

}
