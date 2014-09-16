package main

import (
	"gosample/httphandler"
	"net/http"
)

func main() {
	http.HandleFunc("/", httphandler.IndexHandler)
	http.ListenAndServe(":3000", nil)
}
