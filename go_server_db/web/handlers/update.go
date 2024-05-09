package handlers

import (
	"go_server_db/db"
	"log"
	"net/http"
)

func Update(w http.ResponseWriter, r *http.Request) {
	err := db.UpdateUser(w, r)
	if err != nil {
		log.Fatal("Update Failed")
	}
}
