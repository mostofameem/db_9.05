package handlers

import (
	"go_server_db/db"
	"net/http"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	db.ReadUser(w, r)
}
