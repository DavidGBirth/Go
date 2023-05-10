package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	greeting := fmt.Sprintf("Hello, %s", "James")
	fmt.Println(greeting)

	r := mux.NewRouter()

	http.ListenAndServe(":9000", r)
}