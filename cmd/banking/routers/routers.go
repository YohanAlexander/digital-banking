package routers

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/yohanalexander/desafio-banking-go/cmd/banking/handlers/account"
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
	router.Path("/home").Handler(common.With(
		negroni.Wrap(homeRoutes),
	))
	home := homeRoutes.Path("/home").Subrouter()
	home.Methods("GET").HandlerFunc(hello.HandlerHello(app))

	// rota de login
	loginRoutes := mux.NewRouter()
	router.Path("/login").Handler(common.With(
		negroni.Wrap(loginRoutes),
	))
	login := loginRoutes.Path("/login").Subrouter()
	login.Methods("GET").HandlerFunc(hello.HandlerHello(app))

	// rota de accounts
	accountsRoutes := mux.NewRouter()
	router.Path("/accounts").Handler(common.With(
		negroni.Wrap(accountsRoutes),
	))
	accounts := accountsRoutes.Path("/accounts").Subrouter()
	accounts.Methods("GET").HandlerFunc(account.ListAccounts(app))
	accounts.Methods("POST").HandlerFunc(account.PostAccount(app))

	// rota de balance
	balanceRoutes := mux.NewRouter()
	router.Path("/accounts/{id}/balance").Handler(common.With(
		negroni.Wrap(balanceRoutes),
	))
	balance := balanceRoutes.Path("/accounts/{id}/balance").Subrouter()
	balance.Methods("GET").HandlerFunc(account.BalanceAccount(app))

	// rota de transfers
	transfersRoutes := mux.NewRouter()
	router.Path("/transfers").Handler(common.With(
		negroni.Wrap(transfersRoutes),
	))
	transfers := transfersRoutes.Path("/transfers").Subrouter()
	transfers.Methods("GET").HandlerFunc(hello.HandlerHello(app))

	return router
}
