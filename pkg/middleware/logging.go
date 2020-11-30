package middleware

import (
	"net/http"

	"github.com/yohanalexander/desafio-banking-go/pkg/logger"
)

// LogRequest é um middleware para log das requests na API
func LogRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
		next(w, r)
	}
}
