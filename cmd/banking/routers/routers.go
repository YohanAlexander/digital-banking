package routers

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/yohanalexander/desafio-banking-go/cmd/banking/handlers/account"
	"github.com/yohanalexander/desafio-banking-go/cmd/banking/handlers/login"
	"github.com/yohanalexander/desafio-banking-go/cmd/banking/handlers/transfer"
	"github.com/yohanalexander/desafio-banking-go/pkg/app"
)

// GetRouter retorna o roteador mux da API
func GetRouter(app *app.App) *mux.Router {

	// middleware compartilhado em todas as rotas da API
	common := negroni.New(
		negroni.NewLogger(),
	)

	// criando roteador base
	router := mux.NewRouter()

	// rota de login
	loginRoutes := mux.NewRouter()
	router.Path("/login").Handler(common.With(
		negroni.Wrap(loginRoutes),
	))
	logins := loginRoutes.Path("/login").Subrouter()
	logins.Methods("POST").HandlerFunc(login.HandlerLogin(app))

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
	balances := balanceRoutes.Path("/accounts/{id}/balance").Subrouter()
	balances.Methods("GET").HandlerFunc(account.BalanceAccount(app))

	// rota de transfers
	transfersRoutes := mux.NewRouter()
	router.Path("/transfers").Handler(common.With(
		negroni.Wrap(transfersRoutes),
	))
	transfers := transfersRoutes.Path("/transfers").Subrouter()
	transfers.Methods("GET").HandlerFunc(transfer.ListTransfers(app))
	transfers.Methods("POST").HandlerFunc(transfer.PostTransfer(app))

	return router
}
