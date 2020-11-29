package banking

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

// InitHTTP cria o roteamento das rotas REST
func InitHTTP() {
	router := mux.NewRouter()
	router.HandleFunc("/", helloHandler).Methods("GET")
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
