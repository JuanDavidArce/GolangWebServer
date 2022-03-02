package main

import (
	"fmt"
	"net/http"
)

func HadleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the API Endpoint")
}
