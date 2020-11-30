package handleraccounts

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yohanalexander/desafio-banking-go/cmd/banking/models/accounts"
	"github.com/yohanalexander/desafio-banking-go/pkg/app"
)

// ListAccounts handler para listar accounts no DB
func ListAccounts(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		// capturando accounts no DB
		var accounts []accounts.Account
		result := app.DB.Client.Find(&accounts)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&accounts)
	}
}

// PostAccount handler para criar account no DB
func PostAccount(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		// capturando account no request
		a := &accounts.Account{}
		err := json.NewDecoder(r.Body).Decode(&a)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// validando struct account
		invalid := app.Vld.Struct(a)
		if invalid != nil {
			http.Error(w, invalid.Error(), http.StatusBadRequest)
			return
		}
		// armazenando account no DB
		result := a.Create(app)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(&a)
	}
}

// BalanceAccount handler para retornar o saldo da account no DB
func BalanceAccount(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		// capturando id na url
		id := mux.Vars(r)["id"]
		// capturando account no DB
		a := &accounts.Account{}
		result := app.DB.Client.First(&a, &id)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&a.Balance)
	}
}
