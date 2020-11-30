package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yohanalexander/desafio-banking-go/cmd/banking/handlers/hello"
	"github.com/yohanalexander/desafio-banking-go/pkg/app"
)

// GetRouter retorna o roteador mux da API
func GetRouter(app *app.App) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", hello.HandlerHello(app)).Methods("GET")
	http.Handle("/", router)
	return router
}
