package transfer

import (
	"encoding/json"
	"net/http"

	"github.com/yohanalexander/desafio-banking-go/cmd/banking/models"
	"github.com/yohanalexander/desafio-banking-go/pkg/app"
)

// ListTransfers handler para listar transfers da account no DB
func ListTransfers(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		// capturando transfers no DB
		var t []models.Transfer
		if err := app.DB.Client.Find(&t); err.Error != nil {
			// caso tenha erro ao procurar no banco retorna 500
			http.Error(w, err.Error.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&t)
	}
}

// PostTransfer handler para criar transfer no DB
func PostTransfer(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		// capturando transfer no request
		t := &models.Transfer{}
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			// caso tenha erro no decode do request retorna 400
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// validando json do struct transfer
		if err := app.Vld.Struct(t); err != nil {
			// caso o corpo do request seja inválido retorna 400
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// armazenando struct transfer no DB
		if err := t.CreateTransfer(app); err != nil {
			// caso tenha erro ao armazenar no banco retorna 500
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"Resposta": "Transferência realizada"})
	}
}
