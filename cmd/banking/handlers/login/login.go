package login

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/yohanalexander/desafio-banking-go/cmd/banking/models"
	"github.com/yohanalexander/desafio-banking-go/pkg/app"
	"github.com/yohanalexander/desafio-banking-go/pkg/secret"
)

// HandlerLogin handler para login na API e retorno do token JWT
func HandlerLogin(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		// criando a chave JWT usada para verificar a assinatura
		var jwtKey = []byte(app.Cfg.GetTokenKey())

		// capturando as credenciais no request
		creds := &models.Credentials{}
		if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
			// caso tenha erro no decode do request retorna 400
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// validando json das credenciais
		if err := app.Vld.Struct(creds); err != nil {
			// caso o corpo do request seja inválido retorna 400
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// capturando account no DB
		a := &models.Account{}
		if err := app.DB.Client.First(&a, "cpf = ?", creds.CPF); err.Error != nil {
			// caso tenha erro ao procurar no banco retorna 401
			http.Error(w, err.Error.Error(), http.StatusUnauthorized)
			return
		}

		// se a senha está incorreta
		if !secret.CheckPasswordHash(creds.Secret, a.Secret) {
			// caso tenha erro ao verificar o hash retorna 401
			http.Error(w, errors.New("Senha incorreta").Error(), http.StatusUnauthorized)
			return
		}

		// definindo o tempo de validade do token para 5 min
		expirationTime := time.Now().Add(5 * time.Minute)
		// criando o JWT claims que contém o CPF e tempo de validade
		claims := &models.Claims{
			CPF: creds.CPF,
			StandardClaims: jwt.StandardClaims{
				// no JWT o tempo de validade é dado em milisegundos unix
				ExpiresAt: expirationTime.Unix(),
			},
		}

		// declarando o token com o algoritmo usado para login
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		// criando a string do token JWT
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			// caso tenha erro ao criar o JWT retorna 500
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// adiciona cookie com o token JWT gerado
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})

		// retorna o token em formato JSON
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
	}
}
