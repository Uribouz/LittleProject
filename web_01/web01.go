package main

import (
	"fmt"
	"net/http"
)

func myFuntionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey its kinda working")
}

func myFuntionHandlerAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You dont know mee !!")
}

func main() {
	http.HandleFunc("/", myFuntionHandler)
	http.HandleFunc("/about", myFuntionHandlerAbout)
	http.ListenAndServe(":8000", nil)
}
