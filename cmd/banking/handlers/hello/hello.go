package hello

import (
	"fmt"
	"net/http"

	"github.com/yohanalexander/desafio-banking-go/pkg/app"
)

// HandlerHello um handler de hello world
func HandlerHello(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	}
}
