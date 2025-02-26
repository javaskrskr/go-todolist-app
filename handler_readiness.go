package main

import (
	"net/http"
)

// func you have to use, if you want to define HTTP handler
// in the way that the go standard library
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}
