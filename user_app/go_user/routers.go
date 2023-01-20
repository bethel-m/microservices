package main

import (
	"net/http"
)

func loadRoutes(h *BaseHandler) {
	http.HandleFunc("/v1/register", h.registerHandler)
	http.HandleFunc("/v1/get-user", h.getUserHandler)
}
