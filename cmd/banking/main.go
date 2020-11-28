package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handler).Methods("GET")
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
