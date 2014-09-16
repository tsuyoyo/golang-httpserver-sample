package httphandler

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
	fmt.Printf("Request has come : %s", r.Host)
}
