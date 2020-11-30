package hello

import (
	"fmt"
	"net/http"

	"github.com/yohanalexander/desafio-banking-go/pkg/app"
	"github.com/yohanalexander/desafio-banking-go/pkg/middleware"
)

func handlerHello(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	}
}

// HandlerHello um handler de hello world
func HandlerHello(app *app.App) http.HandlerFunc {
	mdw := []middleware.Middleware{
		middleware.LogRequest,
	}
	return middleware.Chain(handlerHello(app), mdw...)
}
