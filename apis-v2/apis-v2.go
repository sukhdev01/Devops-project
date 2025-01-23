package apisv2

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	var message = "Hello, Sukhdev! I hope you are doing well."

	fmt.Fprintf(w, "API-V2 msg: %s\n", message)
}

func ThanksHandler(w http.ResponseWriter, r *http.Request) {

	var message = "Hi Ginil, I'm doing great. Thank You! I hope you are doing well too."

	fmt.Fprintf(w, "API-V2 msg: %s\n", message)
}

func ByeHandler(w http.ResponseWriter, r *http.Request) {

	message := "bye bye... Have a nice day!"
	fmt.Fprintf(w, "API-V2 msg: %s", message)
}
