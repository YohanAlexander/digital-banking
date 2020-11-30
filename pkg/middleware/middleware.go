package middleware

import (
	"net/http"
)

// Middleware tipo para encadeiar os middlewares
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Chain encadeia os middlewares da direita para esquerda
func Chain(f http.HandlerFunc, m ...Middleware) http.HandlerFunc {
	// se a cadeia esta vazia use o handler original
	if len(m) == 0 {
		return f
	}
	// senão encadeia os middlewares recursivamente
	return m[0](Chain(f, m[1:]...))
}
