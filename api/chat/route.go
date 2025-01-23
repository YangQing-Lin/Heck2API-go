package handler

import (
	"net/http"
)

func EntryPoint(w http.ResponseWriter, r *http.Request) {
	Handler(w, r)
}
