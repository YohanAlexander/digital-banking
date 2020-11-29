package hello

import (
	"fmt"
	"net/http"
)

// HandlerHello um handler de hello world
func HandlerHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}
