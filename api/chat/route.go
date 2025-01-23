package handler

import (
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	handleRequest(w, r)
}
