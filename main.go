package main

import (
	"fmt"
	"net/http"

	apisv2 "github.com/sukhdev01/Devops-project/apis-v2"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {

	message := `Welcome to Go...
		available apis:
			/api/v*/hello  
			/api/v*/thanks  
			/api/v*/bye 
		Try both versions: V1 and V2`
	fmt.Fprint(w, message)
}

func main() {

	http.HandleFunc("/", homeHandler)

	http.HandleFunc("/api/v1/hello", HelloHandler)
	http.HandleFunc("/api/v1/thanks", ThanksHandler)
	http.HandleFunc("/api/v1/bye", ByeHandler)

	http.HandleFunc("/api/v2/hello", apisv2.HelloHandler)
	http.HandleFunc("/api/v2/thanks", apisv2.ThanksHandler)
	http.HandleFunc("/api/v2/bye", apisv2.ByeHandler)

	fmt.Println("Starting server on :3456...")
	err := http.ListenAndServe(":3456", nil)

	if err != nil {
		fmt.Println("Failed to start server:", err)
	}

}
