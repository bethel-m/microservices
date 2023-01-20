package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type BaseHandler struct {
	usermodel userModel
}

func (h *BaseHandler) getUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "user returned")
}

// register a user
func (h *BaseHandler) registerHandler(w http.ResponseWriter, r *http.Request) {
	new_user := user{}
	body_err := json.NewDecoder(r.Body).Decode(&new_user)
	if body_err != nil {
		log.Println("erroy reading body", body_err)
	}
	h.usermodel.create_user(new_user)
}
