package account

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yohanalexander/desafio-banking-go/cmd/banking/models"
	"github.com/yohanalexander/desafio-banking-go/pkg/app"
)

// ListAccounts handler para listar accounts no DB
func ListAccounts(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		// capturando accounts no DB
		var a []models.Account
		if err := app.DB.Client.Find(&a); err.Error != nil {
			// caso tenha erro ao procurar no banco retorna 500
			http.Error(w, err.Error.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&a)
	}
}

// PostAccount handler para criar account no DB
func PostAccount(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		// capturando account no request
		a := &models.Account{}
		if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
			// caso tenha erro no decode do request retorna 400
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// validando json do struct account
		if err := app.Vld.Struct(a); err != nil {
			// caso o corpo do request seja inválido retorna 400
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// armazenando struct account no DB
		if result := a.CreateAccount(app); result != nil {
			// caso tenha erro ao armazenar no banco retorna 500
			http.Error(w, result.Error(), http.StatusInternalServerError)
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
		a := &models.Account{}
		if result := app.DB.Client.First(&a, &id); result.Error != nil {
			// caso tenha erro ao procurar no banco retorna 404
			http.Error(w, result.Error.Error(), http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]float64{"balance": a.Balance})
	}
}
