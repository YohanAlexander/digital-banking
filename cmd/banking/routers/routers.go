package routers

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/yohanalexander/desafio-banking-go/cmd/banking/handlers/hello"
	"github.com/yohanalexander/desafio-banking-go/pkg/app"
)

// GetRouter retorna o roteador mux da API
func GetRouter(app *app.App) *mux.Router {

	// middleware compartilhado em todas as rotas da API
	common := negroni.New(
		negroni.NewLogger(),
	)

	router := mux.NewRouter()

	// rota de home
	homeRoutes := mux.NewRouter()
	router.PathPrefix("/home").Handler(common.With(
		negroni.Wrap(homeRoutes),
	))
	home := homeRoutes.Path("/home").Subrouter()
	home.Methods("GET").HandlerFunc(hello.HandlerHello(app))

	// rota de login
	loginRoutes := mux.NewRouter()
	router.PathPrefix("/login").Handler(common.With(
		negroni.Wrap(loginRoutes),
	))
	login := loginRoutes.Path("/login").Subrouter()
	login.Methods("GET").HandlerFunc(hello.HandlerHello(app))

	// rota de accounts
	accountsRoutes := mux.NewRouter()
	router.PathPrefix("/accounts").Handler(common.With(
		negroni.Wrap(accountsRoutes),
	))
	accounts := accountsRoutes.Path("/accounts").Subrouter()
	accounts.Methods("GET").HandlerFunc(hello.HandlerHello(app))

	// rota de transfers
	transfersRoutes := mux.NewRouter()
	router.PathPrefix("/transfers").Handler(common.With(
		negroni.Wrap(transfersRoutes),
	))
	transfers := accountsRoutes.Path("/transfers").Subrouter()
	transfers.Methods("GET").HandlerFunc(hello.HandlerHello(app))

	return router
}
