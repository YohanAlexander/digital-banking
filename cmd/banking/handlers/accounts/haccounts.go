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
		// caso tenha erro ao procurar no banco retorna 500
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
		// caso tenha erro no decode do request retorna 400
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// validando json do struct account
		invalid := app.Vld.Struct(a)
		// caso o corpo do request seja inválido retorna 400
		if invalid != nil {
			http.Error(w, invalid.Error(), http.StatusBadRequest)
			return
		}
		// armazenando struct account no DB
		result := a.Create(app)
		// caso tenha erro ao armazenar no banco retorna 500
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"Resposta": "Conta criada"})
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
		// caso tenha erro ao procurar no banco retorna 404
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]float64{"balance": a.Balance})
	}
}
